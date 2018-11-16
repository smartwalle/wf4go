package wf4go

import (
	"encoding/json"
	"github.com/smartwalle/xid"
)

// --------------------------------------------------------------------------------
type Process struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	StartTaskId string           `json:"start_task_id"`
	EndTaskId   string           `json:"end_task_id"`
	TaskList    map[string]*Task `json:"task_list"`
	FlowList    map[string]*Flow `json:"flow_list"`
}

func LoadProcess(s string) (p *Process, err error) {
	err = json.Unmarshal([]byte(s), &p)
	return p, err
}

func NewProcess(name string) *Process {
	var p = &Process{}
	p.Id = xid.NewXID().Hex()
	p.Name = name
	p.TaskList = make(map[string]*Task)
	p.FlowList = make(map[string]*Flow)
	return p
}

func (this *Process) String() string {
	ds, _ := json.MarshalIndent(this, "", " ")
	return string(ds)
}

func (this *Process) AddTask(t *Task) {
	if t == nil {
		return
	}
	t.ProcessId = this.Id
	this.TaskList[t.TaskId] = t
}

func (this *Process) AddStartTask(t *Task) {
	if t == nil {
		return
	}
	this.AddTask(t)
	this.StartTaskId = t.TaskId
}

func (this *Process) AddEndTask(t *Task) {
	if t == nil {
		return
	}
	this.AddTask(t)
	this.EndTaskId = t.TaskId
}

func (this *Process) Link(name string, source, target *Task, c ...*Condition) *Flow {
	if source == nil || target == nil {
		return nil
	}
	if source.TaskId == target.TaskId {
		return nil
	}
	this.AddTask(source)
	this.AddTask(target)

	var f = &Flow{}
	f.ProcessId = this.Id
	f.FlowId = xid.NewXID().Hex()
	f.FlowName = name
	f.SourceTaskId = source.TaskId
	f.TargetTaskId = target.TaskId
	f.ConditionList = c
	this.FlowList[f.FlowId] = f
	return f
}

// --------------------------------------------------------------------------------
type TaskType int

const (
	WF_TASK_TYPE_START          TaskType = iota // 开始任务
	WF_TASK_TYPE_END                            // 结束任务
	WF_TASK_TYPE_ASSIGNEE                       // 默认，需要由人员处理
	WF_TASK_TYPE_EXCLUSIVE                      // 自动流转到下一级任务节点中匹配的任务（执行单一任务）
	WF_TASK_TYPE_PARALLEL_FORK                  // 自动流转到下一级任务节点中所有的任务（执行所有任务）
	WF_TASK_TYPE_PARALLEL_JOIN                  // 汇集上一级节点中所有任务，等待上一级所有相关的任务完成后，自动流转到下一级任务
	WF_TASK_TYPE_INCLUSIVE_FORK                 // 自动流转到下一级任务节点中匹配的任务（执行多个任务）
	WF_TASK_TYPE_INCLUSIVE_JOIN                 // 汇集上一级节点中所有任务，等待上一级所有相关的任务完成后，自动流转到下一级任务
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
	t.TaskType = WF_TASK_TYPE_ASSIGNEE
	return t
}

func NewStartTask(name string) *Task {
	var t = NewTask(name)
	t.TaskType = WF_TASK_TYPE_START
	return t
}

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

func NewParallelForkTask(name string) *Task {
	var t = NewTask(name)
	t.TaskType = WF_TASK_TYPE_PARALLEL_FORK
	return t
}

func NewParallelJoinTask(name string) *Task {
	var t = NewTask(name)
	t.TaskType = WF_TASK_TYPE_PARALLEL_JOIN
	return t
}

// --------------------------------------------------------------------------------
type Flow struct {
	ProcessId     string       `json:"process_id"`
	FlowId        string       `json:"flow_id"`
	FlowName      string       `json:"flow_name"`
	SourceTaskId  string       `json:"source_task_id"`
	TargetTaskId  string       `json:"target_task_id"`
	ConditionList []*Condition `json:"condition_list,omitempty"`
}

type Condition struct {
	FlowId        string `json:"flow_id"`
	ConditionType int    `json:"condition_type"`
	Expression    string `json:"expression"`
}

func NewCondition(cType int, expression string) *Condition {
	var fc = &Condition{}
	fc.ConditionType = cType
	fc.Expression = expression
	return fc
}
