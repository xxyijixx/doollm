package file

import (
	"context"
	"doollm/clients/anythingllm"
	"doollm/clients/anythingllm/system"
	"doollm/config"
	"doollm/repo"
	"doollm/repo/model"
	linktype "doollm/service/document/type"
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
	Update()
}

type FileServiceImpl struct{}

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
	if document.LastModifiedAt.Equal(file.UpdatedAt) {
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
	return anythingllmClient.RemoveDocument(system.RemoveDocumentParams{
		Names: []string{document.Location},
	})

}
