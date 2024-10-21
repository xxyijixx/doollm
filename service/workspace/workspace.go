package workspace

import (
	"context"
	"doollm/clients/anythingllm"
	wk "doollm/clients/anythingllm/workspace"
	"doollm/repo"
	"doollm/repo/model"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var anythingllmClient = anythingllm.NewClient()

type WorkspaceService interface {
	Verify(userid int64) bool
	Upload(userid int64, documentId int64) error
	Create(userid int64)
	SelectByUserId(userid int64) (*model.LlmWorkspace, error)
}

type WorkspaceServiceImpl struct {
}

func (w *WorkspaceServiceImpl) Verify(userid int64) bool {

	_, err := repo.WorkspacePermission.WithContext(context.Background()).Where(repo.WorkspacePermission.WorkspaceID.IsNotNull(), repo.WorkspacePermission.IsCreate.Is(true)).First()
	if err != nil {
		log.Debugf("查询用户工作区错误: %v", err)
		return false
	}
	return true
}

// GetWorkspaceUser 获取工作区用户
// 返回工作区用户列表和一个key为userid 值为工作区信息的Map
func (w *WorkspaceServiceImpl) GetWorkspaceUser() ([]*model.LlmWorkspace, map[int64]*model.LlmWorkspace, error) {
	ctx := context.Background()
	userWorkspacePermission, err := repo.WorkspacePermission.WithContext(ctx).Where(repo.WorkspacePermission.IsCreate.Is(true),
		repo.WorkspacePermission.WorkspaceID.IsNotNull()).Find()
	if err != nil {
		return nil, nil, err
	}
	// 进行转换
	userWorkspaces := make([]*model.LlmWorkspace, len(userWorkspacePermission))
	for i, workspacePermission := range userWorkspacePermission {
		userWorkspaces[i] = &model.LlmWorkspace{
			ID:        workspacePermission.ID,
			Userid:    workspacePermission.UserID,
			Name:      workspacePermission.WorkspaceID,
			Slug:      workspacePermission.WorkspaceID,
			CreatedAt: workspacePermission.CreateTime,
		}
	}
	userWorkspaceMapByUserid := make(map[int64]*model.LlmWorkspace)
	for _, userWorkspace := range userWorkspaces {
		userWorkspaceMapByUserid[userWorkspace.Userid] = userWorkspace
	}
	return userWorkspaces, userWorkspaceMapByUserid, nil
}

func (w *WorkspaceServiceImpl) Upload(userid int64, documentId int64) error {
	log.WithFields(log.Fields{
		"userId":     userid,
		"documentId": documentId,
	}).Debug("正在上传至工作区")
	var err error
	if !w.Verify(userid) {
		log.WithField("userId", userid).Debugf("用户没有工作区权限")
		return fmt.Errorf("%v", "用户没有工作区权限")
	}
	ctx := context.Background()
	document, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.ID.Eq(documentId)).First()
	if err != nil {
		return err
	}
	// 获取用户工作区
	workspace, err := w.SelectByUserId(userid)
	if err != nil {
		return err
	}

	workspaceDocument, err := repo.LlmWorkspaceDocument.WithContext(ctx).
		Where(repo.LlmWorkspaceDocument.WorkspaceSlug.Eq(workspace.Slug), repo.LlmWorkspaceDocument.DocumentID.Eq(documentId)).
		First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if workspaceDocument != nil {
		log.Debugf("文档[#%d]已经存在于用户[#%d]的工作区", documentId, userid)
		return fmt.Errorf("the document already exists in the workspace")
	}

	resp, err := anythingllmClient.UpdateEmbeddings(workspace.Slug, wk.UpdateEmbeddingsParams{
		Adds: []string{document.Location},
	})
	if err != nil {
		return err
	}
	// jsonData, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Printf("更新工作区响应:\n %v", string(jsonData))
	flag := false
	for _, workspaceDocument := range resp.Workspace.Documents {
		if workspaceDocument.Docpath == document.Location {
			flag = true
			break
		}
	}
	if !flag {
		log.Debugf("文档[#%d]移动到工作区失败:", document.ID)
		return fmt.Errorf("failed to move document to the workspace")
	}
	err = repo.LlmWorkspaceDocument.WithContext(ctx).Create(
		&model.LlmWorkspaceDocument{
			WorkspaceID:   workspace.ID,
			WorkspaceSlug: workspace.Slug,
			DocumentID:    documentId,
			CreatedAt:     time.Now(),
		},
	)

	return err
}

func (w *WorkspaceServiceImpl) RemoveDocument(userid int64, documentId int64) error {
	log.WithFields(log.Fields{
		"userId":     userid,
		"documentId": documentId,
	}).Debug("移除文档")
	ctx := context.Background()
	workspace, err := w.SelectByUserId(userid)
	if err != nil {
		return err
	}
	document, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.ID.Eq(documentId)).First()
	if err != nil {
		return err
	}
	// workspaceDocument, err := repo.LlmWorkspaceDocument.WithContext(ctx).
	// 	Where(repo.LlmWorkspaceDocument.WorkspaceID.Eq(workspace.ID), repo.LlmWorkspaceDocument.DocumentID.Eq(documentId)).
	// 	First()
	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return err
	// }
	resultInfo, err := repo.LlmWorkspaceDocument.WithContext(ctx).
		Where(repo.LlmWorkspaceDocument.WorkspaceSlug.Eq(workspace.Slug),
			repo.LlmWorkspaceDocument.DocumentID.Eq(document.ID),
		).Delete()
	_ = resultInfo
	if err != nil {
		return err
	}
	resp, err := anythingllmClient.UpdateEmbeddings(workspace.Slug, wk.UpdateEmbeddingsParams{
		Deletes: []string{document.Location},
	})
	_ = resp
	return err
}

// deprecated 不由该方法进行创建
func (w *WorkspaceServiceImpl) Create(userid int64) {
	workspace, err := w.SelectByUserId(userid)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
	}
	if workspace != nil {
		fmt.Println("不要重复创建工作区")
		return
	}
	resp, err := anythingllmClient.CreateWorkspace(wk.CreateParams{
		Name: "tttttttttttttttt",
	})
	if err != nil {
		return
	}
	repo.LlmWorkspace.WithContext(context.Background()).Create(&model.LlmWorkspace{
		Userid:    userid,
		Name:      resp.Workspace.Name,
		Slug:      resp.Workspace.Slug,
		CreatedAt: time.Now(),
	})
}

// SelectByUserId 获取并转换
func (w *WorkspaceServiceImpl) SelectByUserId(userid int64) (*model.LlmWorkspace, error) {
	workspacePermission, err := repo.WorkspacePermission.WithContext(context.Background()).Where(repo.WorkspacePermission.WorkspaceID.IsNotNull(),
		repo.WorkspacePermission.IsCreate.Is(true),
		repo.WorkspacePermission.UserID.Eq(userid)).First()
	if err != nil {
		return nil, err
	}

	workspace := &model.LlmWorkspace{
		ID:        workspacePermission.ID,
		Userid:    workspacePermission.UserID,
		Name:      workspacePermission.WorkspaceID,
		Slug:      workspacePermission.WorkspaceID,
		CreatedAt: workspacePermission.CreateTime,
	}
	return workspace, nil
}

// Delete 删除工作区,删除数据库记录
func (w *WorkspaceServiceImpl) Delete(workspaceSlug string) error {

	_, err := repo.LlmWorkspaceDocument.WithContext(context.Background()).Where(repo.LlmWorkspaceDocument.WorkspaceSlug.Eq(workspaceSlug)).Delete()

	return err
}
