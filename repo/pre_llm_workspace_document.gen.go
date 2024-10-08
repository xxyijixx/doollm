// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package repo

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"doollm/repo/model"
)

func newLlmWorkspaceDocument(db *gorm.DB, opts ...gen.DOOption) llmWorkspaceDocument {
	_llmWorkspaceDocument := llmWorkspaceDocument{}

	_llmWorkspaceDocument.llmWorkspaceDocumentDo.UseDB(db, opts...)
	_llmWorkspaceDocument.llmWorkspaceDocumentDo.UseModel(&model.LlmWorkspaceDocument{})

	tableName := _llmWorkspaceDocument.llmWorkspaceDocumentDo.TableName()
	_llmWorkspaceDocument.ALL = field.NewAsterisk(tableName)
	_llmWorkspaceDocument.ID = field.NewInt64(tableName, "id")
	_llmWorkspaceDocument.WorkspaceID = field.NewInt64(tableName, "workspace_id")
	_llmWorkspaceDocument.WorkspaceSlug = field.NewString(tableName, "slug")
	_llmWorkspaceDocument.DocumentID = field.NewInt64(tableName, "document_id")
	_llmWorkspaceDocument.CreatedAt = field.NewTime(tableName, "created_at")

	_llmWorkspaceDocument.fillFieldMap()

	return _llmWorkspaceDocument
}

type llmWorkspaceDocument struct {
	llmWorkspaceDocumentDo llmWorkspaceDocumentDo

	ALL           field.Asterisk
	ID            field.Int64
	WorkspaceID   field.Int64
	WorkspaceSlug field.String
	DocumentID    field.Int64
	CreatedAt     field.Time

	fieldMap map[string]field.Expr
}

func (l llmWorkspaceDocument) Table(newTableName string) *llmWorkspaceDocument {
	l.llmWorkspaceDocumentDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l llmWorkspaceDocument) As(alias string) *llmWorkspaceDocument {
	l.llmWorkspaceDocumentDo.DO = *(l.llmWorkspaceDocumentDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *llmWorkspaceDocument) updateTableName(table string) *llmWorkspaceDocument {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt64(table, "id")
	l.WorkspaceID = field.NewInt64(table, "workspace_id")
	l.WorkspaceSlug = field.NewString(table, "slug")
	l.DocumentID = field.NewInt64(table, "document_id")
	l.CreatedAt = field.NewTime(table, "created_at")

	l.fillFieldMap()

	return l
}

func (l *llmWorkspaceDocument) WithContext(ctx context.Context) ILlmWorkspaceDocumentDo {
	return l.llmWorkspaceDocumentDo.WithContext(ctx)
}

func (l llmWorkspaceDocument) TableName() string { return l.llmWorkspaceDocumentDo.TableName() }

func (l llmWorkspaceDocument) Alias() string { return l.llmWorkspaceDocumentDo.Alias() }

func (l llmWorkspaceDocument) Columns(cols ...field.Expr) gen.Columns {
	return l.llmWorkspaceDocumentDo.Columns(cols...)
}

func (l *llmWorkspaceDocument) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *llmWorkspaceDocument) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 5)
	l.fieldMap["id"] = l.ID
	l.fieldMap["workspace_id"] = l.WorkspaceID
	l.fieldMap["slug"] = l.WorkspaceSlug
	l.fieldMap["document_id"] = l.DocumentID
	l.fieldMap["created_at"] = l.CreatedAt
}

func (l llmWorkspaceDocument) clone(db *gorm.DB) llmWorkspaceDocument {
	l.llmWorkspaceDocumentDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l llmWorkspaceDocument) replaceDB(db *gorm.DB) llmWorkspaceDocument {
	l.llmWorkspaceDocumentDo.ReplaceDB(db)
	return l
}

type llmWorkspaceDocumentDo struct{ gen.DO }

type ILlmWorkspaceDocumentDo interface {
	gen.SubQuery
	Debug() ILlmWorkspaceDocumentDo
	WithContext(ctx context.Context) ILlmWorkspaceDocumentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILlmWorkspaceDocumentDo
	WriteDB() ILlmWorkspaceDocumentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILlmWorkspaceDocumentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILlmWorkspaceDocumentDo
	Not(conds ...gen.Condition) ILlmWorkspaceDocumentDo
	Or(conds ...gen.Condition) ILlmWorkspaceDocumentDo
	Select(conds ...field.Expr) ILlmWorkspaceDocumentDo
	Where(conds ...gen.Condition) ILlmWorkspaceDocumentDo
	Order(conds ...field.Expr) ILlmWorkspaceDocumentDo
	Distinct(cols ...field.Expr) ILlmWorkspaceDocumentDo
	Omit(cols ...field.Expr) ILlmWorkspaceDocumentDo
	Join(table schema.Tabler, on ...field.Expr) ILlmWorkspaceDocumentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILlmWorkspaceDocumentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILlmWorkspaceDocumentDo
	Group(cols ...field.Expr) ILlmWorkspaceDocumentDo
	Having(conds ...gen.Condition) ILlmWorkspaceDocumentDo
	Limit(limit int) ILlmWorkspaceDocumentDo
	Offset(offset int) ILlmWorkspaceDocumentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILlmWorkspaceDocumentDo
	Unscoped() ILlmWorkspaceDocumentDo
	Create(values ...*model.LlmWorkspaceDocument) error
	CreateInBatches(values []*model.LlmWorkspaceDocument, batchSize int) error
	Save(values ...*model.LlmWorkspaceDocument) error
	First() (*model.LlmWorkspaceDocument, error)
	Take() (*model.LlmWorkspaceDocument, error)
	Last() (*model.LlmWorkspaceDocument, error)
	Find() ([]*model.LlmWorkspaceDocument, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LlmWorkspaceDocument, err error)
	FindInBatches(result *[]*model.LlmWorkspaceDocument, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.LlmWorkspaceDocument) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILlmWorkspaceDocumentDo
	Assign(attrs ...field.AssignExpr) ILlmWorkspaceDocumentDo
	Joins(fields ...field.RelationField) ILlmWorkspaceDocumentDo
	Preload(fields ...field.RelationField) ILlmWorkspaceDocumentDo
	FirstOrInit() (*model.LlmWorkspaceDocument, error)
	FirstOrCreate() (*model.LlmWorkspaceDocument, error)
	FindByPage(offset int, limit int) (result []*model.LlmWorkspaceDocument, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILlmWorkspaceDocumentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l llmWorkspaceDocumentDo) Debug() ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Debug())
}

func (l llmWorkspaceDocumentDo) WithContext(ctx context.Context) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l llmWorkspaceDocumentDo) ReadDB() ILlmWorkspaceDocumentDo {
	return l.Clauses(dbresolver.Read)
}

func (l llmWorkspaceDocumentDo) WriteDB() ILlmWorkspaceDocumentDo {
	return l.Clauses(dbresolver.Write)
}

func (l llmWorkspaceDocumentDo) Session(config *gorm.Session) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Session(config))
}

func (l llmWorkspaceDocumentDo) Clauses(conds ...clause.Expression) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l llmWorkspaceDocumentDo) Returning(value interface{}, columns ...string) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l llmWorkspaceDocumentDo) Not(conds ...gen.Condition) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l llmWorkspaceDocumentDo) Or(conds ...gen.Condition) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l llmWorkspaceDocumentDo) Select(conds ...field.Expr) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l llmWorkspaceDocumentDo) Where(conds ...gen.Condition) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l llmWorkspaceDocumentDo) Order(conds ...field.Expr) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l llmWorkspaceDocumentDo) Distinct(cols ...field.Expr) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l llmWorkspaceDocumentDo) Omit(cols ...field.Expr) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l llmWorkspaceDocumentDo) Join(table schema.Tabler, on ...field.Expr) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l llmWorkspaceDocumentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l llmWorkspaceDocumentDo) RightJoin(table schema.Tabler, on ...field.Expr) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l llmWorkspaceDocumentDo) Group(cols ...field.Expr) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l llmWorkspaceDocumentDo) Having(conds ...gen.Condition) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l llmWorkspaceDocumentDo) Limit(limit int) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l llmWorkspaceDocumentDo) Offset(offset int) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l llmWorkspaceDocumentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l llmWorkspaceDocumentDo) Unscoped() ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Unscoped())
}

func (l llmWorkspaceDocumentDo) Create(values ...*model.LlmWorkspaceDocument) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l llmWorkspaceDocumentDo) CreateInBatches(values []*model.LlmWorkspaceDocument, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l llmWorkspaceDocumentDo) Save(values ...*model.LlmWorkspaceDocument) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l llmWorkspaceDocumentDo) First() (*model.LlmWorkspaceDocument, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LlmWorkspaceDocument), nil
	}
}

func (l llmWorkspaceDocumentDo) Take() (*model.LlmWorkspaceDocument, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LlmWorkspaceDocument), nil
	}
}

func (l llmWorkspaceDocumentDo) Last() (*model.LlmWorkspaceDocument, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LlmWorkspaceDocument), nil
	}
}

func (l llmWorkspaceDocumentDo) Find() ([]*model.LlmWorkspaceDocument, error) {
	result, err := l.DO.Find()
	return result.([]*model.LlmWorkspaceDocument), err
}

func (l llmWorkspaceDocumentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LlmWorkspaceDocument, err error) {
	buf := make([]*model.LlmWorkspaceDocument, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l llmWorkspaceDocumentDo) FindInBatches(result *[]*model.LlmWorkspaceDocument, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l llmWorkspaceDocumentDo) Attrs(attrs ...field.AssignExpr) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l llmWorkspaceDocumentDo) Assign(attrs ...field.AssignExpr) ILlmWorkspaceDocumentDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l llmWorkspaceDocumentDo) Joins(fields ...field.RelationField) ILlmWorkspaceDocumentDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l llmWorkspaceDocumentDo) Preload(fields ...field.RelationField) ILlmWorkspaceDocumentDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l llmWorkspaceDocumentDo) FirstOrInit() (*model.LlmWorkspaceDocument, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LlmWorkspaceDocument), nil
	}
}

func (l llmWorkspaceDocumentDo) FirstOrCreate() (*model.LlmWorkspaceDocument, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LlmWorkspaceDocument), nil
	}
}

func (l llmWorkspaceDocumentDo) FindByPage(offset int, limit int) (result []*model.LlmWorkspaceDocument, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l llmWorkspaceDocumentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l llmWorkspaceDocumentDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l llmWorkspaceDocumentDo) Delete(models ...*model.LlmWorkspaceDocument) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *llmWorkspaceDocumentDo) withDO(do gen.Dao) *llmWorkspaceDocumentDo {
	l.DO = *do.(*gen.DO)
	return l
}
