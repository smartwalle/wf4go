package wf4go

import (
	"fmt"
	"github.com/smartwalle/xid"
)

// --------------------------------------------------------------------------------
type TaskType int

const (
	WF_TASK_TYPE_END       TaskType = 1000 // 结束任务（流程结束）
	WF_TASK_TYPE_DEFAULT   TaskType = 2000 // 默认，需要由人员处理
	WF_TASK_TYPE_EXCLUSIVE TaskType = 3000 // 自动流转到下一级任务节点中匹配的任务（执行单一任务）

	//WF_TASK_TYPE_PARALLEL_FORK TaskType = 4000 // 自动流转到下一级任务节点中所有的任务（执行所有任务）
	//WF_TASK_TYPE_PARALLEL_JOIN TaskType = 5000 // 汇集上一级节点中所有任务，等待上一级所有相关的任务完成后，自动流转到下一级任务
	//
	//WF_TASK_TYPE_INCLUSIVE_FORK TaskType = 6000 // 自动流转到下一级任务节点中匹配的任务（执行多个任务）
	//WF_TASK_TYPE_INCLUSIVE_JOIN TaskType = 7000 // 汇集上一级节点中所有任务，等待上一级所有相关的任务完成后，自动流转到下一级任务
)

type Task struct {
	ProcessId string   `json:"process_id"`
	TaskId    string   `json:"task_id"`
	TaskName  string   `json:"task_name"`
	TaskType  TaskType `json:"task_type"`
	Assignee  string   `json:"assignee"`
}

func NewTask(name string) *Task {
	var t = &Task{}
	t.TaskId = xid.NewXID().Hex()
	t.TaskName = name
	t.TaskType = WF_TASK_TYPE_DEFAULT
	return t
}

//func NewStartTask(name string) *Task {
//	var t = NewTask(name)
//	t.TaskType = WF_TASK_TYPE_START
//	return t
//}

func NewEndTask(name string) *Task {
	var t = NewTask(name)
	t.TaskType = WF_TASK_TYPE_END
	return t
}

func NewExclusiveTask(name string) *Task {
	var t = NewTask(name)
	t.TaskType = WF_TASK_TYPE_EXCLUSIVE
	return t
}

//func NewParallelForkTask(name string) *Task {
//	var t = NewTask(name)
//	t.TaskType = WF_TASK_TYPE_PARALLEL_FORK
//	return t
//}
//
//func NewParallelJoinTask(name string) *Task {
//	var t = NewTask(name)
//	t.TaskType = WF_TASK_TYPE_PARALLEL_JOIN
//	return t
//}

func (this *Task) String() string {
	return fmt.Sprintf("%s-%s", this.TaskId, this.TaskName)
}
