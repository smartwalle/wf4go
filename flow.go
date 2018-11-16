package wf4go

import "fmt"

// --------------------------------------------------------------------------------
type Flow struct {
	ProcessId     string       `json:"process_id"`
	FlowId        string       `json:"flow_id"`
	FlowName      string       `json:"flow_name"`
	SourceTaskId  string       `json:"source_task_id"`
	TargetTaskId  string       `json:"target_task_id"`
	TargetTask    *Task        `json:"-"`
	ConditionList []*Condition `json:"condition_list,omitempty"`
}

func (this *Flow) String() string {
	return fmt.Sprintf("%s-%s", this.FlowId, this.FlowName)
}

// --------------------------------------------------------------------------------
type Condition struct {
	Expression string `json:"expression"`
}

func NewCondition(expression string) *Condition {
	var fc = &Condition{}
	fc.Expression = expression
	return fc
}

func (this *Condition) String() string {
	return fmt.Sprintf("%s", this.Expression)
}
