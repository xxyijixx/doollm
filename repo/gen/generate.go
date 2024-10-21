package main

import (
	"doollm/repo/model"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../",
		Mode:    gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface, // generate mode
	})

	// dsn := config.EnvConfig.GetDSN()

	// g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.File{}, model.FileContent{}, model.FileUser{}, model.Report{}, model.ReportReceive{}, model.User{}, model.LlmDocument{}, model.LlmWorkspace{}, model.LlmWorkspaceDocument{})
	g.ApplyBasic(model.Project{}, model.ProjectColumn{}, model.ProjectFlow{}, model.ProjectFlowItem{}, model.ProjectTask{}, model.ProjectTaskContent{}, model.ProjectTaskFile{}, model.ProjectTaskUser{}, model.ProjectTaskVisibilityUser{}, model.ProjectUser{})
	g.ApplyBasic(model.HistoryChat{}, model.WorkspacePermission{})
	// Generate Type Safe API with Dynamic SQL defined on Querier interface
	// g.ApplyInterface(func(Querier) {}, model.UserToken{}, model.Video{}, model.User{}, model.Comment{})

	// Generate the code
	g.Execute()
}
