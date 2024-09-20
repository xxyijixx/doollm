package task

import (
	"context"
	"doollm/clients/anythingllm"
	"doollm/clients/anythingllm/system"
	"doollm/repo"
	"doollm/repo/model"
	linktype "doollm/service/document/type"
	"doollm/service/workspace"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TaskService interface {
	Traversal()
	UploadWorkspace()
	Update()
}

type TaskServiceImpl struct {
}

var anythingllmClient = anythingllm.NewClient()
var workspaceService = &workspace.WorkspaceServiceImpl{}

type ProjectTaskHandle struct {
	ctx              context.Context
	project          *model.Project
	task             *model.ProjectTask
	rowTask          *TaskRowText
	extras           *TaskDocumentExtras
	userMap          *map[int64]*model.User
	projectColumnMap *map[int64]*model.ProjectColumn
	attachment       []*model.ProjectTaskFile
}

func (t *TaskServiceImpl) Traversal() {
	ctx := context.Background()
	proojects, err := repo.Project.WithContext(ctx).Find()
	if err != nil {
		return
	}
	users, err := repo.User.WithContext(ctx).Find()
	if err != nil {
		return
	}
	// 构建 UserId 到 User 的映射
	userMap := make(map[int64]*model.User)
	for _, user := range users {
		userMap[user.Userid] = user
	}
	for _, project := range proojects {
		handleProject(ctx, project, &userMap)
	}

}

func (t *TaskServiceImpl) UploadWorkspace() {
	ctx := context.Background()
	documents, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.TASK)).Find()
	if err != nil {
		return
	}
	for _, document := range documents {
		extras := TaskDocumentExtras{}
		if err := json.Unmarshal([]byte(document.LinkExtras), &extras); err != nil {
			continue
		}
		for _, owner := range extras.Owner {
			workspaceService.Upload(owner, document.ID)
		}

	}
}

func (t *TaskServiceImpl) Update() {

}

// handleProject 处理项目
func handleProject(ctx context.Context, project *model.Project, userMap *map[int64]*model.User) {
	var err error
	projectColumns, err := repo.ProjectColumn.WithContext(ctx).Where(repo.ProjectColumn.ProjectID).Find()
	if err != nil {
		// 处理错误，比如记录日志或返回错误
		log.Printf("Error querying project columns: %v", err)
		return
	}
	projectColumnMap := make(map[int64]*model.ProjectColumn)
	for _, projectColumn := range projectColumns {
		projectColumnMap[projectColumn.ID] = projectColumn
	}
	// 处理工作流
	projectFlow, err := repo.ProjectFlow.WithContext(ctx).Where(repo.ProjectFlow.ProjectID.Eq(project.ID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	var flowItems []*model.ProjectFlowItem
	flowItemMap := make(map[int64]*model.ProjectFlowItem)
	if projectFlow != nil {
		flowItems, err = repo.ProjectFlowItem.WithContext(ctx).Where(repo.ProjectFlowItem.FlowID.Eq(projectFlow.ID)).Find()
		if err != nil {
			// 处理错误，比如记录日志或返回错误
			log.Printf("Error querying project flow items: %v", err)
		}

		for _, flowItem := range flowItems {
			flowItemMap[flowItem.ID] = flowItem
		}
	}

	// 查询该项目下的全部任务
	projectTasks, err := repo.ProjectTask.WithContext(ctx).Where(repo.ProjectTask.ProjectID.Eq(project.ID), repo.ProjectTask.ParentID.Eq(0)).Find()
	if err != nil {
		// 处理错误，比如记录日志或返回错误
		log.Printf("Error querying project task: %v", err)
		return
	}
	for _, projectTask := range projectTasks {
		projectTaskHandle := &ProjectTaskHandle{
			ctx:              context.Background(),
			project:          project,
			task:             projectTask,
			rowTask:          &TaskRowText{},
			userMap:          userMap,
			projectColumnMap: &projectColumnMap,
		}
		projectTaskHandle.handleTask()
	}

}

func (h *ProjectTaskHandle) handleTask() {

	// 附件处理
	h.handleTaskAttachment()

	projectColumnMap := *h.projectColumnMap
	h.rowTask.ProjectName = h.project.Name
	column, exist := projectColumnMap[h.task.ColumnID]
	if exist {
		h.rowTask.ColumnName = column.Name
	}
	h.rowTask.Priority = h.task.PName
	h.rowTask.StartAt = h.task.StartAt
	h.rowTask.EndAt = h.task.EndAt
	h.rowTask.TaskName = h.task.Name
	h.rowTask.TaskDescription = h.task.Desc
	h.rowTask.TaskOwner = FindTaskUser(h.task.ID, h.userMap, true)
	h.rowTask.TaskAssistant = FindTaskUser(h.task.ID, h.userMap, false)
	h.rowTask.CompleteAt = h.task.CompleteAt
	if h.task.FlowItemID == 0 {
		h.rowTask.Status = func(date time.Time) string {
			if date.IsZero() {
				return "未完成"
			}
			return "已完成"
		}(h.task.CompleteAt)
	} else {

		h.rowTask.Status = strings.ReplaceAll(h.task.FlowItemName, "|", "\\|")
	}
	h.FindProjectUser()
	subTasks, err := repo.ProjectTask.WithContext(h.ctx).Where(repo.ProjectTask.ParentID.Eq(h.task.ID)).Find()
	if err != nil {
		log.Printf("Error query sub task: %v", err)
		return
	}
	h.rowTask.SubNum = len(subTasks)
	completeNum := 0
	subTaskRowText := make([]SubTaskRowText, len(subTasks))
	for i, subTask := range subTasks {
		if !subTask.CompleteAt.IsZero() {
			completeNum += 1
		}
		subTaskRowText[i].Name = subTask.Name
		subTaskRowText[i].Owner = FindTaskUser(subTask.ID, h.userMap, true)
		subTaskRowText[i].StartAt = subTask.StartAt
		subTaskRowText[i].EndAt = subTask.EndAt
		subTaskRowText[i].CompleteAt = subTask.CompleteAt
		if subTask.FlowItemID == 0 {
			subTaskRowText[i].Status = func(date time.Time) string {
				if date.IsZero() {
					return "未完成"
				}
				return "已完成"
			}(h.task.CompleteAt)
		} else {

			subTaskRowText[i].Status = strings.ReplaceAll(subTask.FlowItemName, "|", "\\|")
		}
	}
	h.rowTask.SubComplete = completeNum
	h.rowTask.SubTask = subTaskRowText
	h.updateOrInsertDocument()
}

func (h *ProjectTaskHandle) handleTaskAttachment() {
	files, err := repo.ProjectTaskFile.WithContext(h.ctx).Where(repo.ProjectTaskFile.TaskID.Eq(h.task.ID)).Find()
	if err != nil {
		log.Printf("Error query task file %v", err)
		return
	}
	h.attachment = files
	h.rowTask.Attachment = make([]string, len(files))
	for i, file := range files {
		h.rowTask.Attachment[i] = file.Name
	}
}

func (h *ProjectTaskHandle) updateOrInsertDocument() error {
	var err error
	// 处理
	document, err := repo.LlmDocument.WithContext(h.ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.TASK), repo.LlmDocument.LinkId.Eq(h.task.ID), repo.LlmDocument.LinkParantId.Eq(0)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("Error query document %v", err)
		return err
	}
	if document != nil && document.LastModifiedAt.Equal(h.task.UpdatedAt) {
		log.Debugf("Task[#%d]内容没有更新", h.task.ID)
		return nil
	}

	fileName := "task-" + strconv.FormatInt(h.project.ID, 10) + "-" + strconv.FormatInt(h.task.ID, 10) + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	res, err := anythingllmClient.UploadFileFormString(generateMarkdown(*h.rowTask), fileName, "md")
	if err != nil || !res.Success {
		return err
	}

	if len(res.Documents) == 0 {
		return nil
	}
	doc := res.Documents[0]
	extras, _ := json.Marshal(h.extras)

	if document == nil {
		// 插入新文档
		log.Debugf("Task[#%d]内容没有上传", h.task.ID)
		newDocument := &model.LlmDocument{
			LinkType:           linktype.TASK,
			LinkId:             h.task.ID,
			LinkParantId:       0,
			LinkExtras:         string(extras),
			DocID:              doc.ID,
			Location:           doc.Location,
			Title:              doc.Title,
			Userid:             h.task.Userid,
			TokenCountEstimate: int64(doc.TokenCountEstimate),
			LastModifiedAt:     h.task.UpdatedAt,
			CreatedAt:          time.Now(),
		}
		return repo.LlmDocument.WithContext(h.ctx).Create(newDocument)
	}

	log.Debugf("Task[#%d]内容存在更新", h.task.ID)
	result, err := repo.LlmDocument.WithContext(h.ctx).
		Where(repo.LlmDocument.ID.Eq(document.ID)).
		Updates(&model.LlmDocument{
			LastModifiedAt:     h.task.UpdatedAt,
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

func (h *ProjectTaskHandle) FindProjectUser() {
	projectId := h.project.ID
	projectUsers, err := repo.ProjectUser.WithContext(context.Background()).Where(repo.ProjectUser.ProjectID.Eq(projectId)).Find()
	if err != nil {
		log.Printf("Error query project user: %v", err)
		return
	}
	projectOwnerIds := make([]int64, 0)
	projectUserIds := make([]int64, 0)
	for _, projectUser := range projectUsers {
		if projectUser.Owner == 1 {
			projectOwnerIds = append(projectOwnerIds, projectUser.Userid)
		}
		projectUserIds = append(projectUserIds, projectUser.Userid)

	}

	h.rowTask.ProjectOwner = GetUserNames(projectOwnerIds, h.userMap)
	h.rowTask.ProjectUser = GetUserNames(projectUserIds, h.userMap)

	taskUsers, err := repo.ProjectTaskUser.WithContext(context.Background()).Where(repo.ProjectTaskUser.TaskID.Eq(h.task.ID), repo.ProjectTaskUser.Owner.Eq(1)).Find()
	if err != nil {
		return
	}
	taskOwnerIds := make([]int64, len(taskUsers))
	for i, taskUser := range taskUsers {
		taskOwnerIds[i] = taskUser.Userid
	}
	h.extras = &TaskDocumentExtras{
		Owner: taskOwnerIds,
	}
}

func FindTaskUser(taskId int64, userMap *map[int64]*model.User, isOwner bool) string {
	taskUsers, err := repo.ProjectTaskUser.WithContext(context.Background()).Where(repo.ProjectTaskUser.TaskID.Eq(taskId)).Find()
	if err != nil {
		log.Printf("Error query task user: %v", err)
		return ""
	}
	taskOwnerIds := make([]int64, 0)
	for _, taskUser := range taskUsers {
		if (isOwner && taskUser.Owner == 1) || (!isOwner && taskUser.Owner == 0) {
			taskOwnerIds = append(taskOwnerIds, taskUser.Userid)
		}
	}
	return GetUserNames(taskOwnerIds, userMap)
}

// GetUserNames 获取一组用户的名称
func GetUserNames(userIds []int64, userMap *map[int64]*model.User) string {
	names := make([]string, len(userIds))
	uMap := *userMap
	for i, userId := range userIds {
		user := uMap[userId]
		if user.Nickname != "" {
			names[i] = user.Nickname
		} else {
			names[i] = user.Email
		}

	}
	return strings.Join(names, ",")

}

// generateMarkdown 将 TaskRowText 转换为 Markdown 格式的文本
func generateMarkdown(task TaskRowText) string {
	var sb strings.Builder

	// 标题
	sb.WriteString("# 任务详情\n\n")

	// 任务信息表格
	sb.WriteString("| 字段            | 内容                     |\n")
	sb.WriteString("|------------------|---------------------------|\n")
	sb.WriteString(fmt.Sprintf("| 项目名     | %s                |\n", task.ProjectName))
	sb.WriteString(fmt.Sprintf("| 项目负责人    | %s                |\n", task.ProjectOwner))
	sb.WriteString(fmt.Sprintf("| 项目成员     | %s                |\n", task.ProjectUser))
	sb.WriteString(fmt.Sprintf("| 所属列     | %s                |\n", task.ColumnName))
	sb.WriteString(fmt.Sprintf("| 任务名        | %s                |\n", task.TaskName))
	sb.WriteString(fmt.Sprintf("| 任务负责人       | %s                |\n", task.TaskOwner))
	sb.WriteString(fmt.Sprintf("| 任务协助人员       | %s                |\n", task.TaskAssistant))
	sb.WriteString(fmt.Sprintf("| 任务描述 | %s                |\n", task.TaskDescription))
	sb.WriteString(fmt.Sprintf("| 优先级         | %s                |\n", task.Priority))
	sb.WriteString(fmt.Sprintf("| 子任务数量          | %d                |\n", task.SubNum))
	sb.WriteString(fmt.Sprintf("| 子任务完成数量     | %d                |\n", task.SubComplete))
	sb.WriteString(fmt.Sprintf("|  状态      | %s                |\n", task.Status))
	if !task.StartAt.IsZero() {
		sb.WriteString(fmt.Sprintf("| 开始时间         | %s                         |\n", task.StartAt.Format(time.DateTime)))
	}
	if !task.EndAt.IsZero() {
		sb.WriteString(fmt.Sprintf("| 结束时间          | %s                         |\n", task.EndAt.Format(time.DateTime)))
	}
	if !task.CompleteAt.IsZero() {
		sb.WriteString(fmt.Sprintf("|  完成时间      | %s                         |\n", task.CompleteAt.Format(time.DateTime)))
	}

	// 子任务表格
	if len(task.SubTask) > 0 {
		sb.WriteString("\n## 子任务\n\n")
		sb.WriteString("| 名称      | 负责人     | 开始时间           | 结束时间              | 完成时间         |  状态         |\n")
		sb.WriteString("|-----------|-----------|---------------------|---------------------|---------------------|---------------------|\n")
	}

	for _, subTask := range task.SubTask {
		startAt := ""
		endAt := ""
		completeAt := ""
		if !subTask.StartAt.IsZero() {
			startAt = subTask.StartAt.Format(time.DateTime)
		}
		if !subTask.EndAt.IsZero() {
			endAt = subTask.EndAt.Format(time.DateTime)
		}
		if !subTask.CompleteAt.IsZero() {
			completeAt = subTask.CompleteAt.Format(time.DateTime)
		}
		sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
			subTask.Name,
			subTask.Owner,
			startAt,
			endAt,
			completeAt,
			subTask.Status,
		))
	}

	if len(task.Attachment) > 0 {
		sb.WriteString("\n## 附件\n\n")
		for i, attachment := range task.Attachment {
			sb.WriteString(fmt.Sprintf("%d. %s \n", i+1, attachment))
		}
	}

	return sb.String()
}
