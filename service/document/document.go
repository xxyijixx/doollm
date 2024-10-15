package document

import (
	"context"
	"doollm/clients/anythingllm"
	"doollm/clients/anythingllm/system"
	"doollm/clients/anythingllm/workspace"
	"doollm/repo"
	"doollm/repo/model"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var anythingllmClient = anythingllm.NewClient()

type DocumentService interface{}

type DocumentServiceImpl struct {
}

// Remove 移除文档信息和工作区的文档记录，llm工作区文档会在文档移除后自动清除
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

// RemoveAndUpdateWorkspace 移除旧文档内容，并重新添加到工作区
func (d *DocumentServiceImpl) RemoveAndUpdateWorkspace(documentId int64, newLocation, oldLocation string) error {

	// 移除旧文档
	err := anythingllmClient.RemoveDocument(system.RemoveDocumentParams{
		Names: []string{oldLocation},
	})
	if err != nil {
		log.Errorf("移除文档[#%d]失败: %v", documentId, err)
		return nil
	}

	// 查找需要更新的工作区文档
	workspaceDocuments, err := repo.LlmWorkspaceDocument.WithContext(context.Background()).
		Where(repo.LlmWorkspaceDocument.DocumentID.Eq(documentId)).
		Find()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	ctx := context.Background()
	// 更新工作区文档信息
	for _, workDocument := range workspaceDocuments {
		log.Infof("工作区[%s]更新后新增文档[%s]", workDocument.WorkspaceSlug, newLocation)
		resp, err := anythingllmClient.UpdateEmbeddings(workDocument.WorkspaceSlug, workspace.UpdateEmbeddingsParams{
			Adds: []string{newLocation},
		})
		if err != nil {
			log.Errorf("工作区[%s]更新后新增文档[%s]失败: %v", workDocument.WorkspaceSlug, newLocation, err)
			// 移除记录
			repo.LlmWorkspaceDocument.WithContext(ctx).Where(repo.LlmWorkspaceDocument.DocumentID.Eq(documentId),
				repo.LlmWorkspaceDocument.WorkspaceSlug.Eq(workDocument.WorkspaceSlug)).Delete()
			continue
		}
		flag := false
		for _, workspaceDocument := range resp.Workspace.Documents {
			if workspaceDocument.Docpath == newLocation {
				flag = true
				break
			}
		}
		if !flag {
			log.Debugf("文档[#%d]移动到工作区失败:", workDocument.DocumentID)
			// 移除记录
			repo.LlmWorkspaceDocument.WithContext(ctx).Where(repo.LlmWorkspaceDocument.DocumentID.Eq(documentId),
				repo.LlmWorkspaceDocument.WorkspaceSlug.Eq(workDocument.WorkspaceSlug)).Delete()
			continue
		}
	}
	return nil
}

// RemoveAndUpdateWorkspace
func (d *DocumentServiceImpl) RemoveAll(removeLinkType string) error {
	ctx := context.Background()
	llmDocumentDo := repo.LlmDocument.WithContext(ctx)
	if removeLinkType != "" {
		log.Warnf("正在清除类型为[%s]的文档", removeLinkType)
		llmDocumentDo.Where(repo.LlmDocument.LinkType.Eq(removeLinkType))
	} else {
		log.Warn("正在清除全部类型的文档")
	}

	documents, err := llmDocumentDo.Find()
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

// Clear 清除知识库中在数据库中没有保存的文档
func (d *DocumentServiceImpl) Clear() error {
	log.Debug("正在清除没有保存的文档信息")
	ctx := context.Background()
	documents, err := repo.LlmDocument.WithContext(ctx).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Warnf("查询文档信息错误: %v", err)
		return err
	}
	documentIdMap := make(map[string]*model.LlmDocument)
	for _, document := range documents {
		documentIdMap[document.DocID] = document
	}
	resp, err := anythingllmClient.QueryDocument()
	if err != nil {
		log.Warnf("查询知识库文档信息错误: %v", err)
		return err
	}
	notFoundNum := 0
	notFoundDocumentLocations := make([]string, 0)
	for _, folder := range resp.LocalFiles.Items {
		log.Debugf("当前文件夹: %s\n", folder.Name)
		for _, document := range folder.Items {

			if _, extst := documentIdMap[document.ID]; !extst {
				log.Debugf("当前文档ID: %s, 名称：%s, name:[%s] 不存在于数据库\n", document.ID, document.Title, document.Name)
				notFoundNum++
				notFoundDocumentLocations = append(notFoundDocumentLocations, fmt.Sprintf("%s/%s", folder.Name, document.Name))
			}
		}
	}
	log.Debugf("找到%d个文档不存在于数据库", notFoundNum)

	log.Warnf("正在从知识库移除%d个文档...", notFoundNum)
	// 移除文档
	_ = notFoundDocumentLocations
	// err = anythingllmClient.RemoveDocument(system.RemoveDocumentParams{
	// 	Names: notFoundDocumentLocations,
	// })

	return err
}
