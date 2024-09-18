package report

import (
	"context"
	"doollm/clients/anythingllm"
	"doollm/clients/anythingllm/documents"
	"doollm/clients/anythingllm/system"
	"doollm/repo"
	"doollm/repo/model"
	linktype "doollm/service/document/type"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReportService interface {
	Traversal()
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

// Traversal 对用户的工作报告进行处理，暂时不分批处理
func (r *ReportServiceImpl) Traversal() {

	ctx := context.Background()
	users, err := repo.User.WithContext(ctx).
		Where(repo.User.Bot.Eq(0)).
		Find()
	if err != nil {
		log.Error("select user fail :", err)
	}

	userIds := make([]int64, len(users))
	for i, user := range users {
		userIds[i] = user.Userid
	}
	if len(userIds) == 0 {
		return
	}
	reports, err := repo.Report.WithContext(ctx).Where(repo.Report.Userid.In(userIds...)).Find()
	if err != nil {
		return
	}
	// 构建 UserId 到 User 的映射
	userMap := make(map[int64]*model.User)
	for _, user := range users {
		userMap[user.Userid] = user
	}

	for _, report := range reports {
		document, err := repo.LlmDocument.WithContext(ctx).
			Where(repo.LlmDocument.LinkType.Eq(linktype.REPORT)).
			Where(repo.LlmDocument.LinkId.Eq(report.ID)).
			Where(repo.LlmDocument.LinkParantId.Eq(0)).
			First()
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Info("错误：", err)
			return
		}

		if err := r.updateOrInsertDocument(ctx, report, document, userMap); err != nil {
			log.Error(err)
		}
	}
}

func (r *ReportServiceImpl) Update() {
	panic("not implemented") // TODO: Implement
}

func (fr *ReportServiceImpl) updateOrInsertDocument(ctx context.Context, report *model.Report, document *model.LlmDocument, userMap map[int64]*model.User) error {
	// 更新文档
	if document.LastModifiedAt.Equal(report.UpdatedAt) {
		log.Debugf("Report[#%d]内容没有更新", report.ID)
		return nil
	}

	user, exists := userMap[report.Userid]
	if !exists {
		log.Warn("查询不到用户信息")
	}
	receiveUserNames := handleReceive(ctx, *report, userMap)
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
	rowTitle := "report-" + strconv.FormatInt(report.Userid, 10) + "-" + strconv.FormatInt(report.ID, 10) + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	params := documents.RawTextParams{
		TextContent: string(text),
		Metadata: documents.RawTextMetadata{
			Title: rowTitle,
		},
	}

	res, err := anythingllmClient.UploadRowText(params)
	if err != nil || !res.Success {
		return err
	}

	if len(res.Documents) == 0 {
		return nil
	}
	doc := res.Documents[0]
	if document == nil {
		// 插入新文档
		log.Debugf("Report[#%d]内容没有上传", report.ID)
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

	log.Debugf("Report[#%d]内容存在更新", report.ID)
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
	// 移除旧文档
	return anythingllmClient.RemoveDocument(system.RemoveDocumentParams{
		Names: []string{document.Location},
	})
}

// handleReceive 处理汇报对象，暂时汇报对象不进行上传
func handleReceive(ctx context.Context, report model.Report, userMap map[int64]*model.User) string {
	reportReceive, err := repo.ReportReceive.WithContext(ctx).Where(repo.ReportReceive.Rid.Eq(report.ID)).Find()
	if err != nil {
		return ""
	}
	receiveUserIds := make([]int64, len(reportReceive))
	for i, receive := range reportReceive {
		receiveUserIds[i] = receive.Userid
	}
	receiveUsers := findUserByIDs(userMap, receiveUserIds)
	receiveUserNames := ""
	if len(receiveUsers) != 0 {
		names := make([]string, len(receiveUsers))
		for i, receiveUser := range receiveUsers {
			names[i] = receiveUser.Nickname
		}
		receiveUserNames = strings.Join(names, ",")
	}
	return receiveUserNames
}

func findUserByIDs(items map[int64]*model.User, ids []int64) []*model.User {

	var result []*model.User

	for _, id := range ids {
		if item, exists := items[id]; exists {
			result = append(result, item)
		}
	}

	return result
}
