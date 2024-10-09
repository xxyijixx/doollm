package document

import (
	"context"
	"doollm/clients/anythingllm"
	"doollm/clients/anythingllm/system"
	"doollm/clients/anythingllm/workspace"
	"doollm/repo"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var anythingllmClient = anythingllm.NewClient()

type DocumentService interface{}

type DocumentServiceImpl struct {
}

func (d *DocumentServiceImpl) UploadAndSave() {

}

func (d *DocumentServiceImpl) UploadAndRemove() {

}

func (d *DocumentServiceImpl) Remove(documentId int64) {

	ctx := context.Background()

	document, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.ID.Eq(documentId)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Debugf("未找到相关文档信息")
		}
		return
	}

	anythingllmClient.RemoveDocument(system.RemoveDocumentParams{
		Names: []string{document.Location},
	})

	resultInfo, err := repo.LlmWorkspaceDocument.WithContext(ctx).Where(repo.LlmWorkspaceDocument.DocumentID.Eq(document.ID)).Delete()
	if err != nil {
		return
	}
	_ = resultInfo
}

// RemoveAndUpdateWorkspace
func (d *DocumentServiceImpl) RemoveAndUpdateWorkspace(documentId int64, newLocation, oldLocation string) error {
	anythingllmClient.RemoveDocument(system.RemoveDocumentParams{
		Names: []string{oldLocation},
	})

	workspaceDocuments, err := repo.LlmWorkspaceDocument.WithContext(context.Background()).
		Where(repo.LlmWorkspaceDocument.DocumentID.Eq(documentId)).
		Find()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	for _, workDocument := range workspaceDocuments {
		anythingllmClient.UpdateEmbeddings(workDocument.WorkspaceSlug, workspace.UpdateEmbeddingsParams{
			Adds: []string{newLocation},
		})
	}
	return nil
}

// RemoveAndUpdateWorkspace
func (d *DocumentServiceImpl) RemoveAll(removeLinkType string) error {
	log.Warnf("正在清除类型为[%s]的文档", removeLinkType)
	ctx := context.Background()
	documents, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.LinkType.Eq(removeLinkType)).Find()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	locations := make([]string, len(documents))
	ids := make([]int64, len(documents))
	for i, document := range documents {
		locations[i] = document.Location
		ids[i] = document.ID
	}
	anythingllmClient.RemoveDocument(system.RemoveDocumentParams{
		Names: locations,
	})

	repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.ID.In(ids...)).Delete()
	repo.LlmWorkspaceDocument.WithContext(ctx).Where(repo.LlmWorkspaceDocument.DocumentID.In(ids...)).Delete()
	return nil
}
