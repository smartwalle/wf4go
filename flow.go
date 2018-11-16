package wf4go

import (
	"fmt"
	"github.com/smartwalle/conv4go"
	"strings"
)

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

func (this *Condition) Exec(data map[string]interface{}) bool {
	return false
}

type Exp struct {
	f expFunc
}

func exec(v string, data map[string]interface{}) bool {
	if data == nil {
		return false
	}

	v = strings.TrimSpace(v)
	if strings.Index(v, "&&") != -1 {
		var nvs = strings.Split(v, "&&")
		for _, nv := range nvs {
			if exec(nv, data) == false {
				return false
			}
		}
		return true
	}

	if strings.Index(v, "||") != -1 {
		var nvs = strings.Split(v, "||")
		for _, nv := range nvs {
			if exec(nv, data) == true {
				return true
			}
		}
		return false
	}

	for mk, mv := range m {
		if strings.Index(v, mk) != -1 {
			var nvs = strings.Split(v, mk)
			if len(nvs) < 2 {
				return false
			}
			var p1 = strings.TrimSpace(nvs[0])
			var p2 = strings.TrimSpace(nvs[1])

			var v1 interface{}
			var v2 interface{}

			if v, ok := data[p1]; ok {
				v1 = v
			} else {
				v1 = p1
			}

			if v, ok := data[p2]; ok {
				v2 = v
			} else {
				v2 = p2
			}
			return mv(v1, v2)
		}
	}

	return false
}

type expFunc func(vs ...interface{}) bool

var m = map[string]expFunc{
	">":  gt,
	">=": gte,
	"<":  lt,
	"<=": lte,
	"=":  eq,
	"!=": neq,
}

// gt >
func gt(vs ...interface{}) bool {
	var v1 = vs[0]
	var v2 = vs[1]
	switch vv1 := v1.(type) {
	case int:
		return vv1 > conv4go.Int(v2)
	case int8:
		return vv1 > conv4go.Int8(v2)
	case int16:
		return vv1 > conv4go.Int16(v2)
	case int32:
		return vv1 > conv4go.Int32(v2)
	case int64:
		return vv1 > conv4go.Int64(v2)
	case uint:
		return vv1 > conv4go.Uint(v2)
	case uint8:
		return vv1 > conv4go.Uint8(v2)
	case uint16:
		return vv1 > conv4go.Uint16(v2)
	case uint32:
		return vv1 > conv4go.Uint32(v2)
	case uint64:
		return vv1 > conv4go.Uint64(v2)
	case float32:
		return vv1 > conv4go.Float32(v2)
	case float64:
		return vv1 > conv4go.Float64(v2)
	case string:
		return vv1 > conv4go.String(v2)
	}
	return false
}

// gte >=
func gte(vs ...interface{}) bool {
	var v1 = vs[0]
	var v2 = vs[1]
	switch vv1 := v1.(type) {
	case int:
		return vv1 >= conv4go.Int(v2)
	case int8:
		return vv1 >= conv4go.Int8(v2)
	case int16:
		return vv1 >= conv4go.Int16(v2)
	case int32:
		return vv1 >= conv4go.Int32(v2)
	case int64:
		return vv1 >= conv4go.Int64(v2)
	case uint:
		return vv1 >= conv4go.Uint(v2)
	case uint8:
		return vv1 >= conv4go.Uint8(v2)
	case uint16:
		return vv1 >= conv4go.Uint16(v2)
	case uint32:
		return vv1 >= conv4go.Uint32(v2)
	case uint64:
		return vv1 >= conv4go.Uint64(v2)
	case float32:
		return vv1 >= conv4go.Float32(v2)
	case float64:
		return vv1 >= conv4go.Float64(v2)
	case string:
		return vv1 >= conv4go.String(v2)
	}
	return false
}

// lt <
func lt(vs ...interface{}) bool {
	var v1 = vs[0]
	var v2 = vs[1]
	switch vv1 := v1.(type) {
	case int:
		return vv1 < conv4go.Int(v2)
	case int8:
		return vv1 < conv4go.Int8(v2)
	case int16:
		return vv1 < conv4go.Int16(v2)
	case int32:
		return vv1 < conv4go.Int32(v2)
	case int64:
		return vv1 < conv4go.Int64(v2)
	case uint:
		return vv1 < conv4go.Uint(v2)
	case uint8:
		return vv1 < conv4go.Uint8(v2)
	case uint16:
		return vv1 < conv4go.Uint16(v2)
	case uint32:
		return vv1 < conv4go.Uint32(v2)
	case uint64:
		return vv1 < conv4go.Uint64(v2)
	case float32:
		return vv1 < conv4go.Float32(v2)
	case float64:
		return vv1 < conv4go.Float64(v2)
	case string:
		return vv1 < conv4go.String(v2)
	}
	return false
}

// lte <
func lte(vs ...interface{}) bool {
	var v1 = vs[0]
	var v2 = vs[1]
	switch vv1 := v1.(type) {
	case int:
		return vv1 <= conv4go.Int(v2)
	case int8:
		return vv1 <= conv4go.Int8(v2)
	case int16:
		return vv1 <= conv4go.Int16(v2)
	case int32:
		return vv1 <= conv4go.Int32(v2)
	case int64:
		return vv1 <= conv4go.Int64(v2)
	case uint:
		return vv1 <= conv4go.Uint(v2)
	case uint8:
		return vv1 <= conv4go.Uint8(v2)
	case uint16:
		return vv1 <= conv4go.Uint16(v2)
	case uint32:
		return vv1 <= conv4go.Uint32(v2)
	case uint64:
		return vv1 <= conv4go.Uint64(v2)
	case float32:
		return vv1 <= conv4go.Float32(v2)
	case float64:
		return vv1 <= conv4go.Float64(v2)
	case string:
		return vv1 <= conv4go.String(v2)
	}
	return false
}

// eq =
func eq(vs ...interface{}) bool {
	var v1 = vs[0]
	var v2 = vs[1]
	switch vv1 := v1.(type) {
	case int:
		return vv1 == conv4go.Int(v2)
	case int8:
		return vv1 == conv4go.Int8(v2)
	case int16:
		return vv1 == conv4go.Int16(v2)
	case int32:
		return vv1 == conv4go.Int32(v2)
	case int64:
		return vv1 == conv4go.Int64(v2)
	case uint:
		return vv1 == conv4go.Uint(v2)
	case uint8:
		return vv1 == conv4go.Uint8(v2)
	case uint16:
		return vv1 == conv4go.Uint16(v2)
	case uint32:
		return vv1 == conv4go.Uint32(v2)
	case uint64:
		return vv1 == conv4go.Uint64(v2)
	case float32:
		return vv1 == conv4go.Float32(v2)
	case float64:
		return vv1 == conv4go.Float64(v2)
	case string:
		return vv1 == conv4go.String(v2)
	}
	return false
}

// neq !=
func neq(vs ...interface{}) bool {
	var v1 = vs[0]
	var v2 = vs[1]
	switch vv1 := v1.(type) {
	case int:
		return vv1 != conv4go.Int(v2)
	case int8:
		return vv1 != conv4go.Int8(v2)
	case int16:
		return vv1 != conv4go.Int16(v2)
	case int32:
		return vv1 != conv4go.Int32(v2)
	case int64:
		return vv1 != conv4go.Int64(v2)
	case uint:
		return vv1 != conv4go.Uint(v2)
	case uint8:
		return vv1 != conv4go.Uint8(v2)
	case uint16:
		return vv1 != conv4go.Uint16(v2)
	case uint32:
		return vv1 != conv4go.Uint32(v2)
	case uint64:
		return vv1 != conv4go.Uint64(v2)
	case float32:
		return vv1 != conv4go.Float32(v2)
	case float64:
		return vv1 != conv4go.Float64(v2)
	case string:
		return vv1 != conv4go.String(v2)
	}
	return false
}
