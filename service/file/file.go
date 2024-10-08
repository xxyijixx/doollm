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
	"fmt"
	"os"
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
	MaxTxtLen    = 100000
	FileTypePDF  = "pdf"
	FileTypeDoc  = "document"
	FileTypeTxt  = "txt"
	FileTypeWord = "word"
	FileTypePPT  = "ppt"

	FOLDER = "folder"
)

type FileService interface {
	Traversal()
	UploadWorkspace()
	Update()
	UpdateByFileUser()
	ClearNotExistFile()
	Delete()
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
	log.Infof("Start of file processing")
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

	// 更新工作区
	f.UploadWorkspace()
	// 检测用户权限变化
	f.UpdateByFileUser()
	// 清除不存在文件
	f.ClearNotExistFile()

	log.Infof("End of file processing")
}

func (f *FileServiceImpl) UploadWorkspace() {
	log.Debugf("Uploading user workspace ...")
	documents, err := repo.LlmDocument.WithContext(context.Background()).Where(repo.LlmDocument.LinkType.Eq(linktype.FILE)).Find()
	if err != nil {
		return
	}
	for _, document := range documents {
		handleFileAuth(document.LinkId)
	}

}

// Delete 删除文件，移除知识库文档并更新用户工作区
func (f *FileServiceImpl) Delete(fileId int64) {
	ctx := context.Background()
	document, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.FILE), repo.LlmDocument.LinkId.Eq(fileId)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Debugf("未找到相关文档信息")
		}
		return
	}
	documentService.Remove(document.ID)
}

// Update 文件访问权限变更，更新用户工作区
func (f *FileServiceImpl) Update(fileId int64) {
	log.Debugf("正在处理用户文件更新 fileid=%d", fileId)
	ctx := context.Background()
	file, err := repo.File.WithContext(ctx).Where(repo.File.ID.Eq(fileId)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Debugf("未找到相关文件信息")
		}
		return
	}
	// 获取工作区用户
	workspaceList, userWorkspaceMap, err := workspaceService.GetWorkspaceUser()
	if err != nil {
		log.Infof("Error query user workspace: %v", err)
		return
	}

	// 查找文件可见用户，按用户ID进行升序排序，userid为0表示所有人
	fileUsers, err := repo.FileUser.WithContext(ctx).Where(repo.FileUser.FileID.Eq(file.ID)).Order(repo.FileUser.Userid.Asc()).Find()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Debugf("error query file user %v", err)
			return
		}
	}
	fileUserIds := make([]int64, len(fileUsers))
	fileUserMap := make(map[int64]*model.FileUser)

	for i, fileUser := range fileUsers {
		fileUserIds[i] = fileUser.Userid
		fileUserMap[fileUser.Userid] = fileUser
	}
	// 在文件可见用户中，哪些用户没有工作区权限，需要移除
	permissionDeniedUser := make([]int64, 0)
	if len(fileUsers) != 0 {
		if fileUsers[0].Userid != 0 {
			for _, userWorkspace := range workspaceList {
				if _, exists := fileUserMap[userWorkspace.Userid]; !exists {
					permissionDeniedUser = append(permissionDeniedUser, userWorkspace.Userid)
				}
			}
		}
	}
	// 文件夹
	if file.Type == FOLDER {
		files := make([]*model.File, 0)
		GetAllFile(ctx, file.ID, &files)
		// 处理文件
		for _, file := range files {
			updateFile(ctx, file, permissionDeniedUser, fileUsers, userWorkspaceMap)
		}
	} else {
		updateFile(ctx, file, permissionDeniedUser, fileUsers, userWorkspaceMap)
	}
}

// updateFile 更新文件
func updateFile(ctx context.Context, file *model.File, permissionDeniedUser []int64, fileUsers []*model.FileUser, userWorkspaceMap map[int64]*model.LlmWorkspace) {
	// 判断文件类型是否支持
	if !isSupport(file) {
		log.Debugf("文件[#%d]不支持的文件类型: %s", file.ID, file.Type)
		return
	}

	document, err := repo.LlmDocument.WithContext(ctx).
		Where(repo.LlmDocument.LinkType.Eq(linktype.FILE), repo.LlmDocument.LinkId.Eq(file.ID)).
		First()
	if err != nil {
		log.Debugf("错误查询文件[#%d]对应的文档信息: %v", file.ID, err)
		return
	}

	for _, fileUser := range fileUsers {
		// 如果共享所有人，对所有用户工作区上传，否则单个上传
		if fileUser.Userid == 0 {
			for _, userWorkspace := range userWorkspaceMap {
				workspaceService.Upload(userWorkspace.Userid, document.ID)
			}
			break
		} else {
			workspaceService.Upload(fileUser.Userid, document.ID)
		}
	}
	// 处理哪些用户需要将文件移除
	for _, userid := range permissionDeniedUser {
		if userid == file.Userid {
			continue
		}
		workspaceService.RemoveDocument(userid, document.ID)
	}
	// 文件所有者上传
	workspaceService.Upload(file.Userid, document.ID)
}

// UpdateByFileUser 更新工作区，根据文件共享情况进行更新
func (f *FileServiceImpl) UpdateByFileUser() {
	log.Debugf("processing user share ...")
	fileUsers, err := repo.FileUser.WithContext(context.Background()).Distinct(repo.FileUser.FileID).Find()
	if err != nil {
		return
	}
	for _, fileUser := range fileUsers {
		f.Update(fileUser.FileID)
	}
}

// ClearNotExistFile 清除文件不存在的文档信息
func (f *FileServiceImpl) ClearNotExistFile() {
	log.Info("")
	ctx := context.Background()
	// 文件类型限制
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

	documents, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.FILE)).Find()
	if err != nil {
		return
	}

	fileIdsMap := make(map[int64]struct{})
	for _, file := range files {
		fileIdsMap[file.ID] = struct{}{}
	}

	for _, document := range documents {
		if _, exist := fileIdsMap[document.LinkId]; !exist {
			log.Debugf("文档[#%d]对应的文件[#%d]不存在，移除文档", document.ID, document.LinkId)
			// 在文档信息中找不到对应的文件ID
			documentService.Remove(document.ID)
		}
	}
}

// GetAllFile 获取文件的所有子文件
func GetAllFile(ctx context.Context, fileId int64, fileListPtr *[]*model.File) {
	fileList := *fileListPtr

	subFiles, err := repo.File.WithContext(ctx).Where(repo.File.Pid.Eq(fileId)).Find()
	if err != nil {
		return
	}

	*fileListPtr = append(fileList, subFiles...)

	for _, file := range subFiles {
		if file.Type == FOLDER {
			GetAllFile(ctx, file.ID, fileListPtr)
		}
	}
}

// verifyFile 验证文件是否符合上传规则
func verifyFile(file *model.File, fileContent *model.FileContent) bool {
	switch file.Type {
	case FileTypePPT:
		return fileContent.Size <= MaxPPTSize
	case FileTypePDF:
		return fileContent.Size <= MaxPDFSize
	case FileTypeTxt:
		return VerifyTxtFile(fileContent)
	case FileTypeDoc:
		return true
	default:
		return true
	}
}

// isSupport 是否支持该文件
func isSupport(file *model.File) bool {
	switch file.Type {
	case FileTypePPT, FileTypeWord, FileTypePDF, FileTypeTxt, FileTypeDoc:
		return true
	default:
		return false
	}
}

// verifyTxtFile 验证txt文件是否符合上传规则，文件字符长度不能超过10万字符
func VerifyTxtFile(fileContent *model.FileContent) bool {
	var content Content
	err := json.Unmarshal([]byte(fileContent.Content), &content)
	if err != nil {
		log.Debugf("Error unmarshal json %v", err)
		return false
	}
	data, err := os.ReadFile(config.PublicPath(content.URL)) // 替换为你的文件路径
	if err != nil {
		log.Printf("无法读取文件: %v", err)
		return false
	}

	// 计算字符长度
	charCount := len(data)

	// 判断是否超过10万字符
	if charCount > MaxTxtLen {
		fmt.Printf("文件包含 %d 个字符，超过10万字符。\n", charCount)
		return false
	}
	return true
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
	log.Infof("正在处理文件[#%v]共享情况", fileId)
	if err != nil {
		log.Debugf("Error query file %v", err)
		return
	}
	// 获取工作区列表
	workspaceList, _, err := workspaceService.GetWorkspaceUser()
	if err != nil {
		log.Debugf("Error query workspace %v", err)
		return
	}
	workspaceUserIds := make([]int64, len(workspaceList))
	for i, work := range workspaceList {
		workspaceUserIds[i] = work.Userid
	}
	document, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.FILE),
		repo.LlmDocument.LinkId.Eq(file.ID)).First()
	if err != nil {
		log.Debugf("Error query document %v", err)
		return
	}
	if file.Pid == 0 {

		fileUsers, err := repo.FileUser.WithContext(ctx).Order(repo.FileUser.Userid.Asc()).Find()
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
		log.Debugf("文件[#%v]共享用户：%v", file.ID, shareUserIds)
		for _, userid := range shareUserIds {
			workspaceService.Upload(userid, document.ID)
		}
	}
	workspaceService.Upload(file.Userid, document.ID)
}

// getFileShareUsers 获取文件可见用户列表
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
