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
	if p != nil {
		for _, f := range p.FlowList {
			f.TargetTask = p.TaskList[f.TargetTaskId]
		}
	}
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

func (this *Process) GetFlow(id string) *Flow {
	return this.FlowList[id]
}

func (this *Process) NextFlows(taskId string) []*Flow {
	var nf []*Flow
	for _, f := range this.FlowList {
		if f.SourceTaskId == taskId {
			nf = append(nf, f)
		}
	}
	return nf
}

func (this *Process) NextTasks(taskId string) []*Task {
	var nt []*Task
	for _, f := range this.FlowList {
		if f.SourceTaskId == taskId {
			nt = append(nt, this.GetTask(f.TargetTaskId))
		}
	}
	return nt
}

func (this *Process) GetTask(taskId string) *Task {
	return this.TaskList[taskId]
}

func (this *Process) AddTask(t *Task) {
	if t == nil {
		return
	}
	t.ProcessId = this.Id
	this.TaskList[t.TaskId] = t
}

func (this *Process) GetStartTask() *Task {
	return this.GetTask(this.StartTaskId)
}

func (this *Process) AddStartTask(t *Task) {
	if t == nil {
		return
	}
	this.AddTask(t)
	this.StartTaskId = t.TaskId
}

func (this *Process) GetEndTask() *Task {
	return this.GetTask(this.EndTaskId)
}

func (this *Process) AddEndTask(t *Task) {
	if t == nil {
		return
	}
	this.AddTask(t)
	this.EndTaskId = t.TaskId
}

func (this *Process) LinkTask(name string, sourceTask, targetTask *Task, c ...*Condition) *Flow {
	if sourceTask == nil || targetTask == nil {
		return nil
	}
	if sourceTask.TaskId == targetTask.TaskId {
		return nil
	}
	this.AddTask(sourceTask)
	this.AddTask(targetTask)

	var f = &Flow{}
	f.ProcessId = this.Id
	f.FlowId = xid.NewXID().Hex()
	f.FlowName = name
	f.SourceTaskId = sourceTask.TaskId
	f.TargetTaskId = targetTask.TaskId
	f.TargetTask = targetTask
	f.ConditionList = c
	this.FlowList[f.FlowId] = f
	return f
}

func (this *Process) LinkWithTaskId(name string, sourceTaskId, targetTaskId string, c ...*Condition) *Flow {
	var source = this.GetTask(sourceTaskId)
	var target = this.GetTask(targetTaskId)
	return this.LinkTask(name, source, target, c...)
}

func (this *Process) Unlink(source, target *Task) *Flow {
	return this.UnlinkWithTaskId(source.TaskId, target.TaskId)
}

func (this *Process) UnlinkWithTaskId(sourceTaskId, targetTaskId string) *Flow {
	for _, f := range this.FlowList {
		if f.SourceTaskId == sourceTaskId && f.TargetTaskId == targetTaskId {
			f.TargetTask = nil
			delete(this.FlowList, f.FlowId)
			return f
		}
	}
	return nil
}

func (this *Process) RemoveFlow(f *Flow) {
	if f == nil {
		return
	}
	this.RemoveFlowWithId(f.FlowId)
}

func (this *Process) RemoveFlowWithId(flowId string) {
	delete(this.FlowList, flowId)
}

func (this *Process) RemoveTask(t *Task) {
	if t == nil {
		return
	}
	this.RemoveTaskWithId(t.TaskId)
}

func (this *Process) RemoveTaskWithId(taskId string) {
	delete(this.TaskList, taskId)
	for _, f := range this.FlowList {
		if f.SourceTaskId == taskId || f.TargetTaskId == taskId {
			delete(this.FlowList, f.FlowId)
		}
	}
}
