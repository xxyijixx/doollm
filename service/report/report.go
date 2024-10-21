package report

import (
	"context"
	"doollm/clients/anythingllm"
	"doollm/clients/anythingllm/documents"
	"doollm/repo"
	"doollm/repo/model"
	"doollm/service/common"
	"doollm/service/document"
	linktype "doollm/service/document/type"
	"doollm/service/workspace"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type ReportService interface {
	Traversal()
	UploadWorkspace()
	Update()
}

type ReportServiceImpl struct{}

type ReportJsonData struct {
	Title       string
	Type        string
	Owner       string
	ReceiveUser string
	Content     string
}

var anythingllmClient = anythingllm.NewClient()

var documentService = &document.DocumentServiceImpl{}
var workspaceService = &workspace.WorkspaceServiceImpl{}

// Traversal 对用户的工作报告进行处理
func (r *ReportServiceImpl) Traversal() {

	ctx := context.Background()
	_, userMap, err := common.GetUserAndBiuldMap(ctx)
	if err != nil {
		return
	}
	var reports []*model.Report
	// 分批处理
	// 排除汇报的内容，该字段类型为longtext，在后续使用的地方再查询内容
	err = repo.Report.WithContext(ctx).Omit(repo.Report.Content).FindInBatches(&reports, 100, func(tx gen.Dao, batch int) error {
		for _, report := range reports {
			document, err := repo.LlmDocument.WithContext(ctx).
				Where(repo.LlmDocument.LinkType.Eq(linktype.REPORT)).
				Where(repo.LlmDocument.LinkId.Eq(report.ID)).
				Where(repo.LlmDocument.LinkParantId.Eq(0)).
				First()
			if err != nil && err != gorm.ErrRecordNotFound {
				log.Debugf("Error query document by report id: %v, err: %v", report.ID, err)
				return nil
			}
			if err := r.updateOrInsertDocument(ctx, report, document, userMap); err != nil {
				log.Error("Error update or insert document", err)
				return nil
			}
		}
		return nil
	})
	if err != nil {
		log.Error("分批处理汇报失败", err)
		return
	}

	// 上传用户工作区
	r.UploadWorkspace()
}

// UploadWorkspace 上传至用户工作区，目前至上传到工作报告所有者的工作区
func (r *ReportServiceImpl) UploadWorkspace() {
	ctx := context.Background()
	documents, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.REPORT)).Find()
	if err != nil {
		log.Debugf("Error query documents %v", err)
		return
	}
	for _, document := range documents {
		workspaceService.Upload(document.Userid, document.ID)
	}
}

// Update 维护状态
// TODO
// 1. 查询拥有工作区权限的用户
// 2. 查询这些用户拥有的全部报告ID
// 3. 将这些ID与已经上传的文档ID进行比较，如报告被删除，删除上传文档并更新工作区文档
func (r *ReportServiceImpl) Update() {
	panic("not implemented") // TODO: Implement
}

func (fr *ReportServiceImpl) updateOrInsertDocument(ctx context.Context, report *model.Report, document *model.LlmDocument, userMap map[int64]*model.User) error {
	// 更新文档
	if document != nil && document.LastModifiedAt.Equal(report.UpdatedAt) {
		log.WithField("reportId", report.ID).Debug("report内容没有更新")
		return nil
	}

	user, exists := userMap[report.Userid]
	if !exists {
		return fmt.Errorf("查询不到用户信息 userid: %v", report.Userid)
	}
	receiveUserNames := handleReceive(ctx, *report, userMap)

	// 构建上传的文本内容
	reportContent, err := repo.Report.WithContext(ctx).Select(repo.Report.Content).Where(repo.Report.ID.Eq(report.ID)).First()
	if err != nil {
		return fmt.Errorf("查询工作汇报内容错误: %v", err)
	}
	report.Content = reportContent.Content
	reportJson := ReportJsonData{
		Title:       report.Title,
		Type:        report.Type,
		Owner:       user.Nickname,
		ReceiveUser: receiveUserNames,
		Content:     report.Content,
	}

	text, err := json.Marshal(reportJson)
	if err != nil {
		return err
	}
	rowTitle := fmt.Sprintf("report-%d-%d-%d", report.Userid, report.ID, time.Now().Unix())
	params := documents.RawTextParams{
		TextContent: string(text),
		Metadata: documents.RawTextMetadata{
			Title: rowTitle,
		},
	}

	res, err := anythingllmClient.DocumentUploadRowText(params)
	if err != nil || !res.Success {
		return err
	}

	if len(res.Documents) == 0 {
		log.Debug("上传文档失败, res.Document length is 0")
		return nil
	}
	doc := res.Documents[0]
	if document == nil {
		// 插入新文档
		log.WithField("reportId", report.ID).Debug("report内容没有上传")
		newDocument := &model.LlmDocument{
			LinkType:           linktype.REPORT,
			LinkId:             report.ID,
			LinkParantId:       0,
			DocID:              doc.ID,
			Location:           doc.Location,
			Title:              doc.Title,
			Userid:             report.Userid,
			TokenCountEstimate: int64(doc.TokenCountEstimate),
			LastModifiedAt:     report.UpdatedAt,
			CreatedAt:          time.Now(),
		}
		return repo.LlmDocument.WithContext(ctx).Create(newDocument)
	}

	log.WithField("reportId", report.ID).Debug("report内容存在更新")
	result, err := repo.LlmDocument.WithContext(ctx).
		Where(repo.LlmDocument.ID.Eq(document.ID)).
		Updates(&model.LlmDocument{
			LastModifiedAt:     report.UpdatedAt,
			Location:           doc.Location,
			Title:              doc.Title,
			DocID:              doc.ID,
			TokenCountEstimate: int64(doc.TokenCountEstimate),
		})
	if err != nil || result.RowsAffected == 0 {
		return err
	}
	// 移除旧文档并更新工作区
	return documentService.RemoveAndUpdateWorkspace(document.ID, doc.Location, document.Location)

}

// handleReceive 处理汇报对象,返回汇报对象的名称集合，多个使用逗号进行分隔
func handleReceive(ctx context.Context, report model.Report, userMap map[int64]*model.User) string {
	reportReceive, err := repo.ReportReceive.WithContext(ctx).Where(repo.ReportReceive.Rid.Eq(report.ID)).Find()
	if err != nil {
		return ""
	}
	receiveUserIds := make([]int64, len(reportReceive))
	for i, receive := range reportReceive {
		receiveUserIds[i] = receive.Userid
	}
	receiveUserNames := common.GetUserNames(receiveUserIds, &userMap)

	return receiveUserNames
}
