package document

import (
	"context"
	"doollm/clients/anythingllm"
	"doollm/clients/anythingllm/system"
	"doollm/clients/anythingllm/workspace"
	"doollm/repo"

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

// RemoveAndUpdateWorkspace
func (d *DocumentServiceImpl) RemoveAndUpdateWorkspace(documentId int64, newLocation, oldLocation string) error {
	anythingllmClient.RemoveDocument(system.RemoveDocumentParams{
		Names: []string{oldLocation},
	})

	workspaceDocuments, err := repo.LlmWorkspaceDocument.WithContext(context.Background()).Where(repo.LlmWorkspaceDocument.DocumentID.Eq(documentId)).Find()
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
func (d *DocumentServiceImpl) RemoveAll() error {
	documents, err := repo.LlmDocument.WithContext(context.Background()).Find()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	locations := make([]string, len(documents))
	for i, document := range documents {
		locations[i] = document.Location
	}
	anythingllmClient.RemoveDocument(system.RemoveDocumentParams{
		Names: locations,
	})

	repo.LlmDocument.WithContext(context.Background()).Where(repo.LlmDocument.LinkParantId.Eq(0)).Delete()
	return nil
}
