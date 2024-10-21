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

func newProjectUser(db *gorm.DB, opts ...gen.DOOption) projectUser {
	_projectUser := projectUser{}

	_projectUser.projectUserDo.UseDB(db, opts...)
	_projectUser.projectUserDo.UseModel(&model.ProjectUser{})

	tableName := _projectUser.projectUserDo.TableName()
	_projectUser.ALL = field.NewAsterisk(tableName)
	_projectUser.ID = field.NewInt64(tableName, "id")
	_projectUser.ProjectID = field.NewInt64(tableName, "project_id")
	_projectUser.Userid = field.NewInt64(tableName, "userid")
	_projectUser.Owner = field.NewInt32(tableName, "owner")
	_projectUser.TopAt = field.NewTime(tableName, "top_at")
	_projectUser.CreatedAt = field.NewTime(tableName, "created_at")
	_projectUser.UpdatedAt = field.NewTime(tableName, "updated_at")

	_projectUser.fillFieldMap()

	return _projectUser
}

type projectUser struct {
	projectUserDo

	ALL       field.Asterisk
	ID        field.Int64
	ProjectID field.Int64
	Userid    field.Int64
	Owner     field.Int32
	TopAt     field.Time
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (p projectUser) Table(newTableName string) *projectUser {
	p.projectUserDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p projectUser) As(alias string) *projectUser {
	p.projectUserDo.DO = *(p.projectUserDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *projectUser) updateTableName(table string) *projectUser {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.ProjectID = field.NewInt64(table, "project_id")
	p.Userid = field.NewInt64(table, "userid")
	p.Owner = field.NewInt32(table, "owner")
	p.TopAt = field.NewTime(table, "top_at")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *projectUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *projectUser) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["id"] = p.ID
	p.fieldMap["project_id"] = p.ProjectID
	p.fieldMap["userid"] = p.Userid
	p.fieldMap["owner"] = p.Owner
	p.fieldMap["top_at"] = p.TopAt
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p projectUser) clone(db *gorm.DB) projectUser {
	p.projectUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p projectUser) replaceDB(db *gorm.DB) projectUser {
	p.projectUserDo.ReplaceDB(db)
	return p
}

type projectUserDo struct{ gen.DO }

type IProjectUserDo interface {
	gen.SubQuery
	Debug() IProjectUserDo
	WithContext(ctx context.Context) IProjectUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProjectUserDo
	WriteDB() IProjectUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProjectUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProjectUserDo
	Not(conds ...gen.Condition) IProjectUserDo
	Or(conds ...gen.Condition) IProjectUserDo
	Select(conds ...field.Expr) IProjectUserDo
	Where(conds ...gen.Condition) IProjectUserDo
	Order(conds ...field.Expr) IProjectUserDo
	Distinct(cols ...field.Expr) IProjectUserDo
	Omit(cols ...field.Expr) IProjectUserDo
	Join(table schema.Tabler, on ...field.Expr) IProjectUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProjectUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProjectUserDo
	Group(cols ...field.Expr) IProjectUserDo
	Having(conds ...gen.Condition) IProjectUserDo
	Limit(limit int) IProjectUserDo
	Offset(offset int) IProjectUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectUserDo
	Unscoped() IProjectUserDo
	Create(values ...*model.ProjectUser) error
	CreateInBatches(values []*model.ProjectUser, batchSize int) error
	Save(values ...*model.ProjectUser) error
	First() (*model.ProjectUser, error)
	Take() (*model.ProjectUser, error)
	Last() (*model.ProjectUser, error)
	Find() ([]*model.ProjectUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectUser, err error)
	FindInBatches(result *[]*model.ProjectUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProjectUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProjectUserDo
	Assign(attrs ...field.AssignExpr) IProjectUserDo
	Joins(fields ...field.RelationField) IProjectUserDo
	Preload(fields ...field.RelationField) IProjectUserDo
	FirstOrInit() (*model.ProjectUser, error)
	FirstOrCreate() (*model.ProjectUser, error)
	FindByPage(offset int, limit int) (result []*model.ProjectUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProjectUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p projectUserDo) Debug() IProjectUserDo {
	return p.withDO(p.DO.Debug())
}

func (p projectUserDo) WithContext(ctx context.Context) IProjectUserDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p projectUserDo) ReadDB() IProjectUserDo {
	return p.Clauses(dbresolver.Read)
}

func (p projectUserDo) WriteDB() IProjectUserDo {
	return p.Clauses(dbresolver.Write)
}

func (p projectUserDo) Session(config *gorm.Session) IProjectUserDo {
	return p.withDO(p.DO.Session(config))
}

func (p projectUserDo) Clauses(conds ...clause.Expression) IProjectUserDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p projectUserDo) Returning(value interface{}, columns ...string) IProjectUserDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p projectUserDo) Not(conds ...gen.Condition) IProjectUserDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p projectUserDo) Or(conds ...gen.Condition) IProjectUserDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p projectUserDo) Select(conds ...field.Expr) IProjectUserDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p projectUserDo) Where(conds ...gen.Condition) IProjectUserDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p projectUserDo) Order(conds ...field.Expr) IProjectUserDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p projectUserDo) Distinct(cols ...field.Expr) IProjectUserDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p projectUserDo) Omit(cols ...field.Expr) IProjectUserDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p projectUserDo) Join(table schema.Tabler, on ...field.Expr) IProjectUserDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p projectUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProjectUserDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p projectUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IProjectUserDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p projectUserDo) Group(cols ...field.Expr) IProjectUserDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p projectUserDo) Having(conds ...gen.Condition) IProjectUserDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p projectUserDo) Limit(limit int) IProjectUserDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p projectUserDo) Offset(offset int) IProjectUserDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p projectUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectUserDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p projectUserDo) Unscoped() IProjectUserDo {
	return p.withDO(p.DO.Unscoped())
}

func (p projectUserDo) Create(values ...*model.ProjectUser) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p projectUserDo) CreateInBatches(values []*model.ProjectUser, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p projectUserDo) Save(values ...*model.ProjectUser) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p projectUserDo) First() (*model.ProjectUser, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectUser), nil
	}
}

func (p projectUserDo) Take() (*model.ProjectUser, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectUser), nil
	}
}

func (p projectUserDo) Last() (*model.ProjectUser, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectUser), nil
	}
}

func (p projectUserDo) Find() ([]*model.ProjectUser, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProjectUser), err
}

func (p projectUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectUser, err error) {
	buf := make([]*model.ProjectUser, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p projectUserDo) FindInBatches(result *[]*model.ProjectUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p projectUserDo) Attrs(attrs ...field.AssignExpr) IProjectUserDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p projectUserDo) Assign(attrs ...field.AssignExpr) IProjectUserDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p projectUserDo) Joins(fields ...field.RelationField) IProjectUserDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p projectUserDo) Preload(fields ...field.RelationField) IProjectUserDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p projectUserDo) FirstOrInit() (*model.ProjectUser, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectUser), nil
	}
}

func (p projectUserDo) FirstOrCreate() (*model.ProjectUser, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectUser), nil
	}
}

func (p projectUserDo) FindByPage(offset int, limit int) (result []*model.ProjectUser, count int64, err error) {
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

func (p projectUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p projectUserDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p projectUserDo) Delete(models ...*model.ProjectUser) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *projectUserDo) withDO(do gen.Dao) *projectUserDo {
	p.DO = *do.(*gen.DO)
	return p
}
