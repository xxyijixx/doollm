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

func newReport(db *gorm.DB, opts ...gen.DOOption) report {
	_report := report{}

	_report.reportDo.UseDB(db, opts...)
	_report.reportDo.UseModel(&model.Report{})

	tableName := _report.reportDo.TableName()
	_report.ALL = field.NewAsterisk(tableName)
	_report.ID = field.NewInt64(tableName, "id")
	_report.CreatedAt = field.NewTime(tableName, "created_at")
	_report.UpdatedAt = field.NewTime(tableName, "updated_at")
	_report.Title = field.NewString(tableName, "title")
	_report.Type = field.NewString(tableName, "type")
	_report.Userid = field.NewInt64(tableName, "userid")
	_report.Content = field.NewString(tableName, "content")
	_report.Sign = field.NewString(tableName, "sign")

	_report.fillFieldMap()

	return _report
}

type report struct {
	reportDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	Title     field.String
	Type      field.String
	Userid    field.Int64
	Content   field.String
	Sign      field.String

	fieldMap map[string]field.Expr
}

func (r report) Table(newTableName string) *report {
	r.reportDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r report) As(alias string) *report {
	r.reportDo.DO = *(r.reportDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *report) updateTableName(table string) *report {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.Title = field.NewString(table, "title")
	r.Type = field.NewString(table, "type")
	r.Userid = field.NewInt64(table, "userid")
	r.Content = field.NewString(table, "content")
	r.Sign = field.NewString(table, "sign")

	r.fillFieldMap()

	return r
}

func (r *report) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *report) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 8)
	r.fieldMap["id"] = r.ID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["title"] = r.Title
	r.fieldMap["type"] = r.Type
	r.fieldMap["userid"] = r.Userid
	r.fieldMap["content"] = r.Content
	r.fieldMap["sign"] = r.Sign
}

func (r report) clone(db *gorm.DB) report {
	r.reportDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r report) replaceDB(db *gorm.DB) report {
	r.reportDo.ReplaceDB(db)
	return r
}

type reportDo struct{ gen.DO }

type IReportDo interface {
	gen.SubQuery
	Debug() IReportDo
	WithContext(ctx context.Context) IReportDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReportDo
	WriteDB() IReportDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReportDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReportDo
	Not(conds ...gen.Condition) IReportDo
	Or(conds ...gen.Condition) IReportDo
	Select(conds ...field.Expr) IReportDo
	Where(conds ...gen.Condition) IReportDo
	Order(conds ...field.Expr) IReportDo
	Distinct(cols ...field.Expr) IReportDo
	Omit(cols ...field.Expr) IReportDo
	Join(table schema.Tabler, on ...field.Expr) IReportDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReportDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReportDo
	Group(cols ...field.Expr) IReportDo
	Having(conds ...gen.Condition) IReportDo
	Limit(limit int) IReportDo
	Offset(offset int) IReportDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReportDo
	Unscoped() IReportDo
	Create(values ...*model.Report) error
	CreateInBatches(values []*model.Report, batchSize int) error
	Save(values ...*model.Report) error
	First() (*model.Report, error)
	Take() (*model.Report, error)
	Last() (*model.Report, error)
	Find() ([]*model.Report, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Report, err error)
	FindInBatches(result *[]*model.Report, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Report) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReportDo
	Assign(attrs ...field.AssignExpr) IReportDo
	Joins(fields ...field.RelationField) IReportDo
	Preload(fields ...field.RelationField) IReportDo
	FirstOrInit() (*model.Report, error)
	FirstOrCreate() (*model.Report, error)
	FindByPage(offset int, limit int) (result []*model.Report, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReportDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r reportDo) Debug() IReportDo {
	return r.withDO(r.DO.Debug())
}

func (r reportDo) WithContext(ctx context.Context) IReportDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reportDo) ReadDB() IReportDo {
	return r.Clauses(dbresolver.Read)
}

func (r reportDo) WriteDB() IReportDo {
	return r.Clauses(dbresolver.Write)
}

func (r reportDo) Session(config *gorm.Session) IReportDo {
	return r.withDO(r.DO.Session(config))
}

func (r reportDo) Clauses(conds ...clause.Expression) IReportDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reportDo) Returning(value interface{}, columns ...string) IReportDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reportDo) Not(conds ...gen.Condition) IReportDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reportDo) Or(conds ...gen.Condition) IReportDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reportDo) Select(conds ...field.Expr) IReportDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reportDo) Where(conds ...gen.Condition) IReportDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reportDo) Order(conds ...field.Expr) IReportDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reportDo) Distinct(cols ...field.Expr) IReportDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reportDo) Omit(cols ...field.Expr) IReportDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reportDo) Join(table schema.Tabler, on ...field.Expr) IReportDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reportDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReportDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reportDo) RightJoin(table schema.Tabler, on ...field.Expr) IReportDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reportDo) Group(cols ...field.Expr) IReportDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reportDo) Having(conds ...gen.Condition) IReportDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reportDo) Limit(limit int) IReportDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reportDo) Offset(offset int) IReportDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReportDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reportDo) Unscoped() IReportDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reportDo) Create(values ...*model.Report) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reportDo) CreateInBatches(values []*model.Report, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reportDo) Save(values ...*model.Report) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reportDo) First() (*model.Report, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Report), nil
	}
}

func (r reportDo) Take() (*model.Report, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Report), nil
	}
}

func (r reportDo) Last() (*model.Report, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Report), nil
	}
}

func (r reportDo) Find() ([]*model.Report, error) {
	result, err := r.DO.Find()
	return result.([]*model.Report), err
}

func (r reportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Report, err error) {
	buf := make([]*model.Report, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reportDo) FindInBatches(result *[]*model.Report, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reportDo) Attrs(attrs ...field.AssignExpr) IReportDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reportDo) Assign(attrs ...field.AssignExpr) IReportDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reportDo) Joins(fields ...field.RelationField) IReportDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reportDo) Preload(fields ...field.RelationField) IReportDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reportDo) FirstOrInit() (*model.Report, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Report), nil
	}
}

func (r reportDo) FirstOrCreate() (*model.Report, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Report), nil
	}
}

func (r reportDo) FindByPage(offset int, limit int) (result []*model.Report, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r reportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reportDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reportDo) Delete(models ...*model.Report) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reportDo) withDO(do gen.Dao) *reportDo {
	r.DO = *do.(*gen.DO)
	return r
}
