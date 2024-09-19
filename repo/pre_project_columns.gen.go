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

func newProjectColumn(db *gorm.DB, opts ...gen.DOOption) projectColumn {
	_projectColumn := projectColumn{}

	_projectColumn.projectColumnDo.UseDB(db, opts...)
	_projectColumn.projectColumnDo.UseModel(&model.ProjectColumn{})

	tableName := _projectColumn.projectColumnDo.TableName()
	_projectColumn.ALL = field.NewAsterisk(tableName)
	_projectColumn.ID = field.NewInt64(tableName, "id")
	_projectColumn.ProjectID = field.NewInt64(tableName, "project_id")
	_projectColumn.Name = field.NewString(tableName, "name")
	_projectColumn.Color = field.NewString(tableName, "color")
	_projectColumn.Sort = field.NewInt32(tableName, "sort")
	_projectColumn.CreatedAt = field.NewTime(tableName, "created_at")
	_projectColumn.UpdatedAt = field.NewTime(tableName, "updated_at")
	_projectColumn.DeletedAt = field.NewField(tableName, "deleted_at")

	_projectColumn.fillFieldMap()

	return _projectColumn
}

type projectColumn struct {
	projectColumnDo projectColumnDo

	ALL       field.Asterisk
	ID        field.Int64
	ProjectID field.Int64
	Name      field.String
	Color     field.String
	Sort      field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (p projectColumn) Table(newTableName string) *projectColumn {
	p.projectColumnDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p projectColumn) As(alias string) *projectColumn {
	p.projectColumnDo.DO = *(p.projectColumnDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *projectColumn) updateTableName(table string) *projectColumn {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.ProjectID = field.NewInt64(table, "project_id")
	p.Name = field.NewString(table, "name")
	p.Color = field.NewString(table, "color")
	p.Sort = field.NewInt32(table, "sort")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *projectColumn) WithContext(ctx context.Context) IProjectColumnDo {
	return p.projectColumnDo.WithContext(ctx)
}

func (p projectColumn) TableName() string { return p.projectColumnDo.TableName() }

func (p projectColumn) Alias() string { return p.projectColumnDo.Alias() }

func (p projectColumn) Columns(cols ...field.Expr) gen.Columns {
	return p.projectColumnDo.Columns(cols...)
}

func (p *projectColumn) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *projectColumn) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 8)
	p.fieldMap["id"] = p.ID
	p.fieldMap["project_id"] = p.ProjectID
	p.fieldMap["name"] = p.Name
	p.fieldMap["color"] = p.Color
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p projectColumn) clone(db *gorm.DB) projectColumn {
	p.projectColumnDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p projectColumn) replaceDB(db *gorm.DB) projectColumn {
	p.projectColumnDo.ReplaceDB(db)
	return p
}

type projectColumnDo struct{ gen.DO }

type IProjectColumnDo interface {
	gen.SubQuery
	Debug() IProjectColumnDo
	WithContext(ctx context.Context) IProjectColumnDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProjectColumnDo
	WriteDB() IProjectColumnDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProjectColumnDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProjectColumnDo
	Not(conds ...gen.Condition) IProjectColumnDo
	Or(conds ...gen.Condition) IProjectColumnDo
	Select(conds ...field.Expr) IProjectColumnDo
	Where(conds ...gen.Condition) IProjectColumnDo
	Order(conds ...field.Expr) IProjectColumnDo
	Distinct(cols ...field.Expr) IProjectColumnDo
	Omit(cols ...field.Expr) IProjectColumnDo
	Join(table schema.Tabler, on ...field.Expr) IProjectColumnDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProjectColumnDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProjectColumnDo
	Group(cols ...field.Expr) IProjectColumnDo
	Having(conds ...gen.Condition) IProjectColumnDo
	Limit(limit int) IProjectColumnDo
	Offset(offset int) IProjectColumnDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectColumnDo
	Unscoped() IProjectColumnDo
	Create(values ...*model.ProjectColumn) error
	CreateInBatches(values []*model.ProjectColumn, batchSize int) error
	Save(values ...*model.ProjectColumn) error
	First() (*model.ProjectColumn, error)
	Take() (*model.ProjectColumn, error)
	Last() (*model.ProjectColumn, error)
	Find() ([]*model.ProjectColumn, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectColumn, err error)
	FindInBatches(result *[]*model.ProjectColumn, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProjectColumn) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProjectColumnDo
	Assign(attrs ...field.AssignExpr) IProjectColumnDo
	Joins(fields ...field.RelationField) IProjectColumnDo
	Preload(fields ...field.RelationField) IProjectColumnDo
	FirstOrInit() (*model.ProjectColumn, error)
	FirstOrCreate() (*model.ProjectColumn, error)
	FindByPage(offset int, limit int) (result []*model.ProjectColumn, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProjectColumnDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p projectColumnDo) Debug() IProjectColumnDo {
	return p.withDO(p.DO.Debug())
}

func (p projectColumnDo) WithContext(ctx context.Context) IProjectColumnDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p projectColumnDo) ReadDB() IProjectColumnDo {
	return p.Clauses(dbresolver.Read)
}

func (p projectColumnDo) WriteDB() IProjectColumnDo {
	return p.Clauses(dbresolver.Write)
}

func (p projectColumnDo) Session(config *gorm.Session) IProjectColumnDo {
	return p.withDO(p.DO.Session(config))
}

func (p projectColumnDo) Clauses(conds ...clause.Expression) IProjectColumnDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p projectColumnDo) Returning(value interface{}, columns ...string) IProjectColumnDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p projectColumnDo) Not(conds ...gen.Condition) IProjectColumnDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p projectColumnDo) Or(conds ...gen.Condition) IProjectColumnDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p projectColumnDo) Select(conds ...field.Expr) IProjectColumnDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p projectColumnDo) Where(conds ...gen.Condition) IProjectColumnDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p projectColumnDo) Order(conds ...field.Expr) IProjectColumnDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p projectColumnDo) Distinct(cols ...field.Expr) IProjectColumnDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p projectColumnDo) Omit(cols ...field.Expr) IProjectColumnDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p projectColumnDo) Join(table schema.Tabler, on ...field.Expr) IProjectColumnDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p projectColumnDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProjectColumnDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p projectColumnDo) RightJoin(table schema.Tabler, on ...field.Expr) IProjectColumnDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p projectColumnDo) Group(cols ...field.Expr) IProjectColumnDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p projectColumnDo) Having(conds ...gen.Condition) IProjectColumnDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p projectColumnDo) Limit(limit int) IProjectColumnDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p projectColumnDo) Offset(offset int) IProjectColumnDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p projectColumnDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectColumnDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p projectColumnDo) Unscoped() IProjectColumnDo {
	return p.withDO(p.DO.Unscoped())
}

func (p projectColumnDo) Create(values ...*model.ProjectColumn) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p projectColumnDo) CreateInBatches(values []*model.ProjectColumn, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p projectColumnDo) Save(values ...*model.ProjectColumn) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p projectColumnDo) First() (*model.ProjectColumn, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectColumn), nil
	}
}

func (p projectColumnDo) Take() (*model.ProjectColumn, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectColumn), nil
	}
}

func (p projectColumnDo) Last() (*model.ProjectColumn, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectColumn), nil
	}
}

func (p projectColumnDo) Find() ([]*model.ProjectColumn, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProjectColumn), err
}

func (p projectColumnDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectColumn, err error) {
	buf := make([]*model.ProjectColumn, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p projectColumnDo) FindInBatches(result *[]*model.ProjectColumn, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p projectColumnDo) Attrs(attrs ...field.AssignExpr) IProjectColumnDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p projectColumnDo) Assign(attrs ...field.AssignExpr) IProjectColumnDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p projectColumnDo) Joins(fields ...field.RelationField) IProjectColumnDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p projectColumnDo) Preload(fields ...field.RelationField) IProjectColumnDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p projectColumnDo) FirstOrInit() (*model.ProjectColumn, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectColumn), nil
	}
}

func (p projectColumnDo) FirstOrCreate() (*model.ProjectColumn, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectColumn), nil
	}
}

func (p projectColumnDo) FindByPage(offset int, limit int) (result []*model.ProjectColumn, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p projectColumnDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p projectColumnDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p projectColumnDo) Delete(models ...*model.ProjectColumn) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *projectColumnDo) withDO(do gen.Dao) *projectColumnDo {
	p.DO = *do.(*gen.DO)
	return p
}