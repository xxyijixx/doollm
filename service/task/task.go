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
	"reflect"
	"sort"
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

// Traversal 遍历全部项目，将所有任务进行上传
func (t *TaskServiceImpl) Traversal() {
	ctx := context.Background()
	// 查找所有项目
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

	// 上传用户工作区
	t.UploadWorkspace()
	// t.UpdateByTaskOwner()
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
		missingFromTasks, _, extras, err := compareTaskAndExtraOwners(ctx, document)
		if err != nil {
			continue
		}
		log.Infof("需要将文档[#%d]移除的用户: %+v", document.ID, missingFromTasks)
		// 从用户工作区移除
		for _, id := range missingFromTasks {
			workspaceService.RemoveDocument(id, document.ID)
		}
		// 添加到用户工作区
		for _, id := range extras.Owner {
			workspaceService.Upload(id, document.ID)
		}
	}
}

// Update 根据task的负责人信息进行更新
func (t *TaskServiceImpl) UpdateByTaskOwner() {
	log.Info("Start of task processing")
	ctx := context.Background()
	documents, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.TASK)).Find()
	if err != nil {
		log.Info("查询文档信息失败: ", err)
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
	for _, document := range documents {
		log.Infof("Start of task[%d] processing", document.LinkId)
		missingFromTasks, missingFromExtras, extras, err := compareTaskAndExtraOwners(ctx, document)
		if err != nil {
			continue
		}
		// 从用户工作区移除
		for _, id := range missingFromTasks {
			workspaceService.RemoveDocument(id, document.ID)
		}
		// 添加到用户工作区
		for _, id := range missingFromExtras {
			workspaceService.Upload(id, document.ID)
		}

		// 更新文档拓展信息
		jsonData, err := json.Marshal(extras)
		if err != nil {
			continue
		}
		_, err = repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.ID.Eq(document.ID)).Update(repo.LlmDocument.LinkExtras, string(jsonData))
		if err != nil {
			log.Debugf("更新任务文档拓展信息失败: %v", err)
		}
	}
	log.Info("End of task processing")
}

// Update 根据task的子项进行更新
func (t *TaskServiceImpl) UpdateBySubTask() {
	log.Info("Start of task processing")
	ctx := context.Background()
	documents, err := repo.LlmDocument.WithContext(ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.TASK)).Find()
	if err != nil {
		log.Info("查询文档信息失败: ", err)
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
	for _, document := range documents {

		log.Infof("Start of task[%d] processing", document.LinkId)
		// 需要额外写一个更新逻辑，引用原来部分

	}
	log.Info("End of task processing")
}

// compareTaskAndExtraOwners 获取需要
func compareTaskAndExtraOwners(ctx context.Context, document *model.LlmDocument) (missingFromTasks []int64, missingFromExtras []int64, extras TaskDocumentExtras, err error) {
	// 查询任务的负责人
	taskUsers, err := repo.ProjectTaskUser.WithContext(ctx).Where(
		repo.ProjectTaskUser.TaskID.Eq(document.LinkId),
		repo.ProjectTaskUser.Owner.Eq(1)).Find()
	if err != nil {
		return
	}

	// 提取任务用户的所有者 ID
	taskOwnerIds := make([]int64, len(taskUsers))
	for i, user := range taskUsers {
		taskOwnerIds[i] = user.Userid
	}

	// 解析文档中的额外信息
	if err = json.Unmarshal([]byte(document.LinkExtras), &extras); err != nil {
		return
	}

	// 创建两个映射表，用于快速查找
	taskOwnerIdSet := make(map[int64]bool, len(taskOwnerIds))
	extraOwnerIdSet := make(map[int64]bool, len(extras.Owner))

	for _, id := range taskOwnerIds {
		taskOwnerIdSet[id] = true
	}
	for _, id := range extras.Owner {
		extraOwnerIdSet[id] = true
	}

	// 找出 extras.Owner 中不在 taskOwnerIds 中的 ID
	for _, id := range extras.Owner {
		if _, exists := taskOwnerIdSet[id]; !exists {
			missingFromTasks = append(missingFromTasks, id)
		}
	}

	// 找出 taskOwnerIds 中不在 extras.Owner 中的 ID
	for _, id := range taskOwnerIds {
		if _, exists := extraOwnerIdSet[id]; !exists {
			missingFromExtras = append(missingFromExtras, id)
		}
	}

	extras.Owner = taskOwnerIds

	return
}

// handleProject 处理项目
func handleProject(ctx context.Context, project *model.Project, userMap *map[int64]*model.User) {
	var err error
	// 查找项目中所有列
	projectColumns, err := repo.ProjectColumn.WithContext(ctx).Where(repo.ProjectColumn.ProjectID).Find()
	if err != nil {
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

	projectOwner, projectUser := getProjectUser(project.ID, userMap)
	// 处理项目下的任务
	for _, projectTask := range projectTasks {
		projectTaskHandle := &ProjectTaskHandle{
			ctx:     context.Background(),
			project: project,
			task:    projectTask,
			extras:  &TaskDocumentExtras{},
			rowTask: &TaskRowText{
				ProjectOwner: projectOwner,
				ProjectUser:  projectUser,
			},
			userMap:          userMap,
			projectColumnMap: &projectColumnMap,
		}
		projectTaskHandle.HandleTask()
	}

}

// getProjectUser 获取项目用户的名称
func getProjectUser(projectId int64, userMap *map[int64]*model.User) (projectOwner, projectUser string) {
	ctx := context.Background()
	projectUsers, err := repo.ProjectUser.WithContext(ctx).Where(repo.ProjectUser.ProjectID.Eq(projectId)).Find()
	if err != nil {
		log.Printf("Error query project user: %v", err)
		return
	}
	projectOwnerIds := make([]int64, 0)
	projectUserIds := make([]int64, len(projectUsers))
	// 获取项目负责人和项目成员ID
	for i, projectUser := range projectUsers {
		if projectUser.Owner == 1 {
			projectOwnerIds = append(projectOwnerIds, projectUser.Userid)
		}
		projectUserIds[i] = projectUser.Userid
	}

	// 获取成员名称
	projectOwner = GetUserNames(projectOwnerIds, userMap)
	projectUser = GetUserNames(projectUserIds, userMap)
	return
}

// HandleTask 处理项目任务
func (h *ProjectTaskHandle) HandleTask() {

	// 附件处理
	h.HandleTaskAttachment()
	projectColumnMap := *h.projectColumnMap
	h.rowTask.ProjectName = h.project.Name
	column, exist := projectColumnMap[h.task.ColumnID]
	if exist {
		h.rowTask.ColumnName = column.Name
	}
	h.findTaskOwner()
	h.rowTask.Priority = h.task.PName
	h.rowTask.StartAt = h.task.StartAt
	h.rowTask.EndAt = h.task.EndAt
	h.rowTask.TaskName = h.task.Name
	h.rowTask.TaskDescription = h.task.Desc
	h.rowTask.TaskAssistant = findTaskUser(h.task.ID, h.userMap, false)
	h.rowTask.CompleteAt = h.task.CompleteAt
	h.rowTask.Status = handleTaskStatus(h.task)

	h.extras.SubTask = make([]SubTaskExtras, 0)
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
		h.HandleSubTask(&subTaskRowText[i], subTask)
	}
	h.rowTask.SubComplete = completeNum
	h.rowTask.SubTask = subTaskRowText
	h.updateOrInsertDocument()
}

// handleSubTask 处理子任务
func (h *ProjectTaskHandle) HandleSubTask(subTaskRowText *SubTaskRowText, subTask *model.ProjectTask) {
	subTaskRowText.Name = subTask.Name
	subTaskRowText.Owner = findTaskUser(subTask.ID, h.userMap, true)
	subTaskRowText.StartAt = subTask.StartAt
	subTaskRowText.EndAt = subTask.EndAt
	subTaskRowText.CompleteAt = subTask.CompleteAt
	subTaskRowText.Status = handleTaskStatus(subTask)

	// 记录子任务的额外保存信息
	h.extras.SubTask = append(h.extras.SubTask, SubTaskExtras{
		TaskId:    subTask.ID,
		UpdatedAt: subTask.UpdatedAt,
	})
}

// HandleTaskAttachment 处理任务附件信息，添加附件名称，附件内容未进行上传
func (h *ProjectTaskHandle) HandleTaskAttachment() {
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

// handleTaskStatus 处理任务的状态
func handleTaskStatus(task *model.ProjectTask) string {
	if task.FlowItemID == 0 {
		if task.CompleteAt.IsZero() {
			return "未完成"
		}
		return "已完成"
	}

	flowItemNames := strings.Split(task.FlowItemName, "|")
	if len(flowItemNames) > 0 {
		return flowItemNames[0]
	}

	return strings.ReplaceAll(task.FlowItemName, "|", "\\|")
}

// compareTaskExtras
func compareTaskExtras(taskExtras TaskDocumentExtras, documentTaskExtrasStr string) bool {
	documentTaskExtras := TaskDocumentExtras{}
	// 解析文档中的额外信息
	if err := json.Unmarshal([]byte(documentTaskExtrasStr), &documentTaskExtras); err != nil {
		return false
	}
	// 1. 比较 Owner 切片：排序后再比较
	if len(taskExtras.Owner) != len(documentTaskExtras.Owner) {
		return false
	}
	sort.Slice(taskExtras.Owner, func(i, j int) bool {
		return taskExtras.Owner[i] < taskExtras.Owner[j]
	})
	sort.Slice(documentTaskExtras.Owner, func(i, j int) bool {
		return documentTaskExtras.Owner[i] < documentTaskExtras.Owner[j]
	})
	if !reflect.DeepEqual(taskExtras.Owner, documentTaskExtras.Owner) {
		return false
	}

	// 2. 比较 SubTask 切片：按 TaskId 排序后再比较每个字段
	if len(taskExtras.SubTask) != len(documentTaskExtras.SubTask) {
		return false
	}
	sort.Slice(taskExtras.SubTask, func(i, j int) bool {
		return taskExtras.SubTask[i].TaskId < taskExtras.SubTask[j].TaskId
	})
	sort.Slice(documentTaskExtras.SubTask, func(i, j int) bool {
		return documentTaskExtras.SubTask[i].TaskId < documentTaskExtras.SubTask[j].TaskId
	})
	for i := range taskExtras.SubTask {
		if taskExtras.SubTask[i].TaskId != documentTaskExtras.SubTask[i].TaskId || !taskExtras.SubTask[i].UpdatedAt.Equal(documentTaskExtras.SubTask[i].UpdatedAt) {
			return false
		}
	}
	return true
}

func (h *ProjectTaskHandle) updateOrInsertDocument() error {
	var err error
	// 处理
	document, err := repo.LlmDocument.WithContext(h.ctx).Where(repo.LlmDocument.LinkType.Eq(linktype.TASK), repo.LlmDocument.LinkId.Eq(h.task.ID), repo.LlmDocument.LinkParantId.Eq(0)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("Error query document %v", err)
		return err
	}
	// if document != nil && document.LastModifiedAt.Equal(h.task.UpdatedAt) {
	// 	log.Debugf("Task[#%d]内容没有更新", h.task.ID)
	// 	if compareTaskExtras(*h.extras, document.LinkExtras) {
	// 		log.Debugf("Task[#%d]附加信息没有变更", h.task.ID)
	// 		return nil
	// 	}
	// }

	fileName := fmt.Sprintf("task-%d-%d-%d", h.project.ID, h.task.ID, time.Now().Unix())
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
			LinkExtras:         string(extras),
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

// findTaskOwner 查找任务负责人
func (h *ProjectTaskHandle) findTaskOwner() {
	taskUsers, err := repo.ProjectTaskUser.WithContext(h.ctx).Where(repo.ProjectTaskUser.TaskID.Eq(h.task.ID), repo.ProjectTaskUser.Owner.Eq(1)).Find()
	if err != nil {
		return
	}
	taskOwnerIds := make([]int64, len(taskUsers))
	for i, taskUser := range taskUsers {
		taskOwnerIds[i] = taskUser.Userid
	}
	// 任务负责人
	h.extras.Owner = taskOwnerIds
}

// findTaskUser 查找任务成员，包含负责人或协助人员
func findTaskUser(taskId int64, userMap *map[int64]*model.User, isOwner bool) string {
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

// GetUserNames 获取一组用户的名称，使用逗号进行分隔
func GetUserNames(userIds []int64, userMap *map[int64]*model.User) string {
	names := make([]string, len(userIds))
	uMap := *userMap
	for i, userId := range userIds {
		user := uMap[userId]

		names[i] = user.Nickname
		if names[i] == "" {
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
			sb.WriteString(fmt.Sprintf("- 附件%d %s \n", i+1, attachment))
		}
	}

	return sb.String()
}
