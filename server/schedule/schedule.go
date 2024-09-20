package schedule

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type TaskFunc func()

type ScheduledTask struct {
	cron    *cron.Cron
	mu      sync.Mutex
	running bool
	tasks   map[cron.EntryID]TaskFunc // 存储任务的 entryID 和 TaskFunc
}

func NewScheduledTask() *ScheduledTask {
	return &ScheduledTask{
		cron:  cron.New(cron.WithSeconds()),
		tasks: make(map[cron.EntryID]TaskFunc),
	}
}

// AddTask 添加一个新的定时任务，返回 entryID 和错误
func (st *ScheduledTask) AddTask(cronExpr string, taskFunc TaskFunc) (cron.EntryID, error) {
	st.mu.Lock()
	defer st.mu.Unlock()

	entryID, err := st.cron.AddFunc(cronExpr, taskFunc)
	if err != nil {
		return 0, err
	}
	st.tasks[entryID] = taskFunc
	return entryID, nil
}

// RemoveTask 根据 entryID 删除任务
func (st *ScheduledTask) RemoveTask(entryID cron.EntryID) {
	st.mu.Lock()
	defer st.mu.Unlock()

	if _, exists := st.tasks[entryID]; exists {
		delete(st.tasks, entryID) // 从任务列表删除
		st.cron.Remove(entryID)   // 从 cron 删除
	}
}

func (st *ScheduledTask) Start() {
	st.mu.Lock()
	defer st.mu.Unlock()

	if !st.running {
		st.cron.Start()
		st.running = true
	}
}

func (st *ScheduledTask) Stop() {
	st.mu.Lock()
	defer st.mu.Unlock()

	if st.running {
		st.cron.Stop()
		st.running = false
	}
}

func (st *ScheduledTask) IsRunning() bool {
	st.mu.Lock()
	defer st.mu.Unlock()

	return st.running
}
