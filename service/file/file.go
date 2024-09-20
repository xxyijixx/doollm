package file

import (
	"context"
	"doollm/clients/anythingllm"
	"doollm/config"
	"doollm/repo"
	"doollm/repo/model"
	"doollm/service/document"
	linktype "doollm/service/document/type"
	"doollm/service/workspace"
	"strconv"
	"strings"

	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	MaxPPTSize   = 1024 * 1024 * 100 // 100MB
	MaxPDFSize   = 1024 * 1024 * 50  // 50MB
	FileTypePDF  = "pdf"
	FileTypeDoc  = "document"
	FileTypeTxt  = "txt"
	FileTypeWord = "word"
	FileTypePPT  = "ppt"
)

type FileService interface {
	Traversal()
	UploadWorkspace()
	Update()
}

type FileServiceImpl struct{}

var documentService = &document.DocumentServiceImpl{}
var workspaceService = &workspace.WorkspaceServiceImpl{}

// const ALLOW_TYPE = [],

var anythingllmClient = anythingllm.NewClient()

type Content struct {
	From   string `json:"from"`
	Type   string `json:"type"`
	Ext    string `json:"ext"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (f *FileServiceImpl) Traversal() {
	ctx := context.Background()
	// 将所有符合条件的文件查询出来
	conditions := []string{
		FileTypePDF,
		FileTypeDoc,
		FileTypeTxt,
		FileTypeWord,
		FileTypePPT}

	files, err := repo.File.WithContext(ctx).Where(repo.File.Type.In(conditions...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	fileIDs := make([]int64, len(files))
	for i, file := range files {
		fileIDs[i] = file.ID
	}

	for _, file := range files {
		fileContent, err := repo.FileContent.WithContext(ctx).
			Where(repo.FileContent.Fid.Eq(file.ID)).
			Order(repo.FileContent.ID.Desc()).
			First()
		if err != nil || !verifyFile(file, fileContent) {
			continue
		}

		var content Content
		err = json.Unmarshal([]byte(fileContent.Content), &content)
		if err != nil {
			continue
		}

		document, err := repo.LlmDocument.WithContext(ctx).
			Where(repo.LlmDocument.LinkType.Eq(linktype.FILE), repo.LlmDocument.LinkId.Eq(file.ID)).
			First()
		if err != nil && err != gorm.ErrRecordNotFound {
			continue
		}

		if err := f.updateOrInsertDocument(ctx, file, content, document); err != nil {
			log.Error(err)
		}

	}

}

func (f *FileServiceImpl) UploadWorkspace() {
	documents, err := repo.LlmDocument.WithContext(context.Background()).Where(repo.LlmDocument.LinkType.Eq(linktype.FILE)).Find()
	if err != nil {
		return
	}
	for _, document := range documents {
		handleFileAuth(document.LinkId)
	}

}

func (f *FileServiceImpl) Update() {

}

func verifyFile(file *model.File, fileContent *model.FileContent) bool {
	switch file.Type {
	case FileTypePPT:
		return fileContent.Size <= MaxPPTSize
	case FileTypePDF:
		return fileContent.Size <= MaxPDFSize
	case FileTypeTxt, FileTypeDoc:
		return true
	default:
		return true
	}
}

func (f *FileServiceImpl) updateOrInsertDocument(ctx context.Context, file *model.File, content Content, document *model.LlmDocument) error {
	if document != nil && document.LastModifiedAt.Equal(file.UpdatedAt) {
		log.Debugf("File[#%d]没有更新", file.ID)
		return nil
	}
	filePath := config.PublicPath(content.URL)
	log.Info("正在处理", filePath)
	res, err := anythingllmClient.UploadFile(filePath, file.Ext)
	if err != nil || !res.Success {
		return err
	}
	if len(res.Documents) == 0 {
		return nil
	}
	doc := res.Documents[0]
	if document == nil {
		// 插入新文档
		log.Debugf("File[#%d]没有上传", file.ID)
		newDocument := &model.LlmDocument{
			LinkType:           linktype.FILE,
			LinkId:             file.ID,
			LinkParantId:       0,
			DocID:              doc.ID,
			Location:           doc.Location,
			Title:              doc.Title,
			Userid:             file.Userid,
			TokenCountEstimate: int64(doc.TokenCountEstimate),
			LastModifiedAt:     file.UpdatedAt,
			CreatedAt:          time.Now(),
		}
		return repo.LlmDocument.WithContext(ctx).Create(newDocument)
	}
	// 更新文档
	log.Debugf("File[#%d]存在更新", file.ID)
	result, err := repo.LlmDocument.WithContext(ctx).
		Where(repo.LlmDocument.ID.Eq(document.ID)).
		Updates(&model.LlmDocument{
			LastModifiedAt:     file.UpdatedAt,
			Location:           doc.Location,
			Title:              doc.Title,
			DocID:              doc.ID,
			TokenCountEstimate: int64(doc.TokenCountEstimate),
		})
	if err != nil || result.RowsAffected == 0 {
		return err
	}
	// 移除旧文档
	return documentService.RemoveAndUpdateWorkspace(document.ID, doc.Location, document.Location)
}

// 处理文件可以被哪些用户查看， 后续通过参数传入工作区信息，默认ID对应的文件非文件夹
func handleFileAuth(fileId int64) {
	ctx := context.Background()
	file, err := repo.File.WithContext(ctx).Where(repo.File.ID.Eq(fileId)).First()
	log.Infof("正在处理文件[$%v]共享情况", fileId)
	if err != nil {
		log.Debugf("Error query file %v", err)
		return
	}
	workspaceList, err := repo.LlmWorkspace.WithContext(ctx).Find()
	if err != nil {
		log.Debugf("Error query workspace %v", err)
		return
	}
	workspaceUserIds := make([]int64, len(workspaceList))
	for i, work := range workspaceList {
		workspaceUserIds[i] = work.Userid
	}
	if file.Pid == 0 {
		fileUsers, err := repo.FileUser.WithContext(ctx).Order(repo.FileUser.Userid.Asc()).Find()
		document, err2 := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.FILE), repo.LlmDocument.LinkId.Eq(file.ID)).First()
		if err2 != nil {
			log.Debugf("Error query document %v", err)
			return
		}
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return
			}
			log.Debugf("文件[#%v]不存在共享", file.ID)
		} else {
			if fileUsers[0].Userid == 0 {
				// 共享所有人
				for _, userid := range workspaceUserIds {
					workspaceService.Upload(userid, document.ID)
				}
			} else {
				for _, fileUser := range fileUsers {
					workspaceService.Upload(fileUser.Userid, document.ID)
				}
			}
		}
		workspaceService.Upload(file.Userid, document.ID)
	} else {
		// 对于非位于顶级目录的文件
		shareUserIds := getFileShareUsers(file, workspaceUserIds)
		document, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.FILE), repo.LlmDocument.LinkId.Eq(file.ID)).First()
		if err != nil {
			return
		}
		log.Debugf("文件[#%v]共享用户：%v", file.ID, shareUserIds)
		for _, userid := range shareUserIds {
			workspaceService.Upload(userid, document.ID)
		}
	}

}

func getFileShareUsers(file *model.File, workspaceUserIds []int64) []int64 {
	pids := file.Pids
	parts := strings.Split(strings.Trim(pids, ","), ",")
	ctx := context.Background()
	shareUsers := make([]int64, 0)
	for _, part := range parts {
		pid, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return shareUsers
		}
		fileUsers, err := repo.FileUser.WithContext(ctx).Where(repo.FileUser.FileID.Eq(pid)).Order(repo.FileUser.Userid.Asc()).Find()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return []int64{}
		}
		for _, fileUser := range fileUsers {
			if fileUser.Userid == 0 {
				return workspaceUserIds
			}
			shareUsers = append(shareUsers, fileUser.Userid)
		}
	}
	return shareUsers
}
