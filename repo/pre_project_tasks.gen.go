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

func newProjectTask(db *gorm.DB, opts ...gen.DOOption) projectTask {
	_projectTask := projectTask{}

	_projectTask.projectTaskDo.UseDB(db, opts...)
	_projectTask.projectTaskDo.UseModel(&model.ProjectTask{})

	tableName := _projectTask.projectTaskDo.TableName()
	_projectTask.ALL = field.NewAsterisk(tableName)
	_projectTask.ID = field.NewInt64(tableName, "id")
	_projectTask.ParentID = field.NewInt64(tableName, "parent_id")
	_projectTask.ProjectID = field.NewInt64(tableName, "project_id")
	_projectTask.ColumnID = field.NewInt64(tableName, "column_id")
	_projectTask.DialogID = field.NewInt64(tableName, "dialog_id")
	_projectTask.FlowItemID = field.NewInt64(tableName, "flow_item_id")
	_projectTask.FlowItemName = field.NewString(tableName, "flow_item_name")
	_projectTask.Name = field.NewString(tableName, "name")
	_projectTask.Color = field.NewString(tableName, "color")
	_projectTask.Desc = field.NewString(tableName, "desc")
	_projectTask.StartAt = field.NewTime(tableName, "start_at")
	_projectTask.EndAt = field.NewTime(tableName, "end_at")
	_projectTask.ArchivedAt = field.NewTime(tableName, "archived_at")
	_projectTask.ArchivedUserid = field.NewInt64(tableName, "archived_userid")
	_projectTask.ArchivedFollow = field.NewInt32(tableName, "archived_follow")
	_projectTask.CompleteAt = field.NewTime(tableName, "complete_at")
	_projectTask.Userid = field.NewInt64(tableName, "userid")
	_projectTask.Visibility = field.NewInt32(tableName, "visibility")
	_projectTask.PLevel = field.NewInt32(tableName, "p_level")
	_projectTask.PName = field.NewString(tableName, "p_name")
	_projectTask.PColor = field.NewString(tableName, "p_color")
	_projectTask.Sort = field.NewInt32(tableName, "sort")
	_projectTask.Loop = field.NewString(tableName, "loop")
	_projectTask.LoopAt = field.NewTime(tableName, "loop_at")
	_projectTask.CreatedAt = field.NewTime(tableName, "created_at")
	_projectTask.UpdatedAt = field.NewTime(tableName, "updated_at")
	_projectTask.DeletedAt = field.NewField(tableName, "deleted_at")
	_projectTask.DeletedUserid = field.NewInt64(tableName, "deleted_userid")

	_projectTask.fillFieldMap()

	return _projectTask
}

type projectTask struct {
	projectTaskDo projectTaskDo

	ALL            field.Asterisk
	ID             field.Int64
	ParentID       field.Int64
	ProjectID      field.Int64
	ColumnID       field.Int64
	DialogID       field.Int64
	FlowItemID     field.Int64
	FlowItemName   field.String
	Name           field.String
	Color          field.String
	Desc           field.String
	StartAt        field.Time
	EndAt          field.Time
	ArchivedAt     field.Time
	ArchivedUserid field.Int64
	ArchivedFollow field.Int32
	CompleteAt     field.Time
	Userid         field.Int64
	Visibility     field.Int32
	PLevel         field.Int32
	PName          field.String
	PColor         field.String
	Sort           field.Int32
	Loop           field.String
	LoopAt         field.Time
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	DeletedUserid  field.Int64

	fieldMap map[string]field.Expr
}

func (p projectTask) Table(newTableName string) *projectTask {
	p.projectTaskDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p projectTask) As(alias string) *projectTask {
	p.projectTaskDo.DO = *(p.projectTaskDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *projectTask) updateTableName(table string) *projectTask {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.ParentID = field.NewInt64(table, "parent_id")
	p.ProjectID = field.NewInt64(table, "project_id")
	p.ColumnID = field.NewInt64(table, "column_id")
	p.DialogID = field.NewInt64(table, "dialog_id")
	p.FlowItemID = field.NewInt64(table, "flow_item_id")
	p.FlowItemName = field.NewString(table, "flow_item_name")
	p.Name = field.NewString(table, "name")
	p.Color = field.NewString(table, "color")
	p.Desc = field.NewString(table, "desc")
	p.StartAt = field.NewTime(table, "start_at")
	p.EndAt = field.NewTime(table, "end_at")
	p.ArchivedAt = field.NewTime(table, "archived_at")
	p.ArchivedUserid = field.NewInt64(table, "archived_userid")
	p.ArchivedFollow = field.NewInt32(table, "archived_follow")
	p.CompleteAt = field.NewTime(table, "complete_at")
	p.Userid = field.NewInt64(table, "userid")
	p.Visibility = field.NewInt32(table, "visibility")
	p.PLevel = field.NewInt32(table, "p_level")
	p.PName = field.NewString(table, "p_name")
	p.PColor = field.NewString(table, "p_color")
	p.Sort = field.NewInt32(table, "sort")
	p.Loop = field.NewString(table, "loop")
	p.LoopAt = field.NewTime(table, "loop_at")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.DeletedUserid = field.NewInt64(table, "deleted_userid")

	p.fillFieldMap()

	return p
}

func (p *projectTask) WithContext(ctx context.Context) IProjectTaskDo {
	return p.projectTaskDo.WithContext(ctx)
}

func (p projectTask) TableName() string { return p.projectTaskDo.TableName() }

func (p projectTask) Alias() string { return p.projectTaskDo.Alias() }

func (p projectTask) Columns(cols ...field.Expr) gen.Columns { return p.projectTaskDo.Columns(cols...) }

func (p *projectTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *projectTask) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 28)
	p.fieldMap["id"] = p.ID
	p.fieldMap["parent_id"] = p.ParentID
	p.fieldMap["project_id"] = p.ProjectID
	p.fieldMap["column_id"] = p.ColumnID
	p.fieldMap["dialog_id"] = p.DialogID
	p.fieldMap["flow_item_id"] = p.FlowItemID
	p.fieldMap["flow_item_name"] = p.FlowItemName
	p.fieldMap["name"] = p.Name
	p.fieldMap["color"] = p.Color
	p.fieldMap["desc"] = p.Desc
	p.fieldMap["start_at"] = p.StartAt
	p.fieldMap["end_at"] = p.EndAt
	p.fieldMap["archived_at"] = p.ArchivedAt
	p.fieldMap["archived_userid"] = p.ArchivedUserid
	p.fieldMap["archived_follow"] = p.ArchivedFollow
	p.fieldMap["complete_at"] = p.CompleteAt
	p.fieldMap["userid"] = p.Userid
	p.fieldMap["visibility"] = p.Visibility
	p.fieldMap["p_level"] = p.PLevel
	p.fieldMap["p_name"] = p.PName
	p.fieldMap["p_color"] = p.PColor
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["loop"] = p.Loop
	p.fieldMap["loop_at"] = p.LoopAt
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["deleted_userid"] = p.DeletedUserid
}

func (p projectTask) clone(db *gorm.DB) projectTask {
	p.projectTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p projectTask) replaceDB(db *gorm.DB) projectTask {
	p.projectTaskDo.ReplaceDB(db)
	return p
}

type projectTaskDo struct{ gen.DO }

type IProjectTaskDo interface {
	gen.SubQuery
	Debug() IProjectTaskDo
	WithContext(ctx context.Context) IProjectTaskDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProjectTaskDo
	WriteDB() IProjectTaskDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProjectTaskDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProjectTaskDo
	Not(conds ...gen.Condition) IProjectTaskDo
	Or(conds ...gen.Condition) IProjectTaskDo
	Select(conds ...field.Expr) IProjectTaskDo
	Where(conds ...gen.Condition) IProjectTaskDo
	Order(conds ...field.Expr) IProjectTaskDo
	Distinct(cols ...field.Expr) IProjectTaskDo
	Omit(cols ...field.Expr) IProjectTaskDo
	Join(table schema.Tabler, on ...field.Expr) IProjectTaskDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProjectTaskDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProjectTaskDo
	Group(cols ...field.Expr) IProjectTaskDo
	Having(conds ...gen.Condition) IProjectTaskDo
	Limit(limit int) IProjectTaskDo
	Offset(offset int) IProjectTaskDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectTaskDo
	Unscoped() IProjectTaskDo
	Create(values ...*model.ProjectTask) error
	CreateInBatches(values []*model.ProjectTask, batchSize int) error
	Save(values ...*model.ProjectTask) error
	First() (*model.ProjectTask, error)
	Take() (*model.ProjectTask, error)
	Last() (*model.ProjectTask, error)
	Find() ([]*model.ProjectTask, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectTask, err error)
	FindInBatches(result *[]*model.ProjectTask, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProjectTask) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProjectTaskDo
	Assign(attrs ...field.AssignExpr) IProjectTaskDo
	Joins(fields ...field.RelationField) IProjectTaskDo
	Preload(fields ...field.RelationField) IProjectTaskDo
	FirstOrInit() (*model.ProjectTask, error)
	FirstOrCreate() (*model.ProjectTask, error)
	FindByPage(offset int, limit int) (result []*model.ProjectTask, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProjectTaskDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p projectTaskDo) Debug() IProjectTaskDo {
	return p.withDO(p.DO.Debug())
}

func (p projectTaskDo) WithContext(ctx context.Context) IProjectTaskDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p projectTaskDo) ReadDB() IProjectTaskDo {
	return p.Clauses(dbresolver.Read)
}

func (p projectTaskDo) WriteDB() IProjectTaskDo {
	return p.Clauses(dbresolver.Write)
}

func (p projectTaskDo) Session(config *gorm.Session) IProjectTaskDo {
	return p.withDO(p.DO.Session(config))
}

func (p projectTaskDo) Clauses(conds ...clause.Expression) IProjectTaskDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p projectTaskDo) Returning(value interface{}, columns ...string) IProjectTaskDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p projectTaskDo) Not(conds ...gen.Condition) IProjectTaskDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p projectTaskDo) Or(conds ...gen.Condition) IProjectTaskDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p projectTaskDo) Select(conds ...field.Expr) IProjectTaskDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p projectTaskDo) Where(conds ...gen.Condition) IProjectTaskDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p projectTaskDo) Order(conds ...field.Expr) IProjectTaskDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p projectTaskDo) Distinct(cols ...field.Expr) IProjectTaskDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p projectTaskDo) Omit(cols ...field.Expr) IProjectTaskDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p projectTaskDo) Join(table schema.Tabler, on ...field.Expr) IProjectTaskDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p projectTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProjectTaskDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p projectTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) IProjectTaskDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p projectTaskDo) Group(cols ...field.Expr) IProjectTaskDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p projectTaskDo) Having(conds ...gen.Condition) IProjectTaskDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p projectTaskDo) Limit(limit int) IProjectTaskDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p projectTaskDo) Offset(offset int) IProjectTaskDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p projectTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProjectTaskDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p projectTaskDo) Unscoped() IProjectTaskDo {
	return p.withDO(p.DO.Unscoped())
}

func (p projectTaskDo) Create(values ...*model.ProjectTask) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p projectTaskDo) CreateInBatches(values []*model.ProjectTask, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p projectTaskDo) Save(values ...*model.ProjectTask) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p projectTaskDo) First() (*model.ProjectTask, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectTask), nil
	}
}

func (p projectTaskDo) Take() (*model.ProjectTask, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectTask), nil
	}
}

func (p projectTaskDo) Last() (*model.ProjectTask, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectTask), nil
	}
}

func (p projectTaskDo) Find() ([]*model.ProjectTask, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProjectTask), err
}

func (p projectTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProjectTask, err error) {
	buf := make([]*model.ProjectTask, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p projectTaskDo) FindInBatches(result *[]*model.ProjectTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p projectTaskDo) Attrs(attrs ...field.AssignExpr) IProjectTaskDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p projectTaskDo) Assign(attrs ...field.AssignExpr) IProjectTaskDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p projectTaskDo) Joins(fields ...field.RelationField) IProjectTaskDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p projectTaskDo) Preload(fields ...field.RelationField) IProjectTaskDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p projectTaskDo) FirstOrInit() (*model.ProjectTask, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectTask), nil
	}
}

func (p projectTaskDo) FirstOrCreate() (*model.ProjectTask, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProjectTask), nil
	}
}

func (p projectTaskDo) FindByPage(offset int, limit int) (result []*model.ProjectTask, count int64, err error) {
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

func (p projectTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p projectTaskDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p projectTaskDo) Delete(models ...*model.ProjectTask) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *projectTaskDo) withDO(do gen.Dao) *projectTaskDo {
	p.DO = *do.(*gen.DO)
	return p
}