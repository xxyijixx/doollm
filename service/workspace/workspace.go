package workspace

import (
	"context"
	"doollm/clients/anythingllm"
	wk "doollm/clients/anythingllm/workspace"
	"doollm/repo"
	"doollm/repo/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var anythingllmClient = anythingllm.NewClient()

type WorkspaceService interface {
	Verify(userid int64) bool
}

type WorkspaceServiceImpl struct {
}

func (w *WorkspaceServiceImpl) Verify(userid int64) bool {
	mockAuthUser := []int64{1, 6}
	for _, authUser := range mockAuthUser {
		if authUser == userid {
			return true
		}
	}
	return false
}

func (w *WorkspaceServiceImpl) Upload(userid int64, documentId int64) error {
	fmt.Println("正在上传", userid, documentId)
	var err error
	if !w.Verify(userid) {
		return fmt.Errorf("%v", "用户没有工作区权限")
	}
	ctx := context.Background()
	document, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.ID.Eq(documentId)).First()
	if err != nil {
		return err
	}
	workspace, err := repo.LlmWorkspace.WithContext(ctx).Where(repo.LlmWorkspace.Userid.Eq(userid)).First()
	if err != nil {
		return err
	}
	_, err = anythingllmClient.UpdateEmbeddings(workspace.Slug, wk.UpdateEmbeddingsParams{
		Adds: []string{document.Location},
	})
	if err != nil {
		return err
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

func (w *WorkspaceServiceImpl) Update(userid int64, documentId int64) bool {
	mockAuthUser := []int64{1, 6}
	for _, authUser := range mockAuthUser {
		if authUser == userid {
			return true
		}
	}
	return false
}

func (w *WorkspaceServiceImpl) Create(userid int64) {
	workspace, err := repo.LlmWorkspace.WithContext(context.Background()).Where(repo.LlmWorkspace.Userid.Eq(userid)).First()
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
