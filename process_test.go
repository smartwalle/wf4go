package wf4go

//
//import (
//	"fmt"
//	"testing"
//)
//
//func TestNewProcess(t *testing.T) {
//	var p = NewProcess("请假流程")
//
//	var st = NewStartTask("开始")
//	p.AddStartTask(st)
//
//	var et = NewEndTask("结束")
//	p.AddEndTask(et)
//
//	var cd = NewExclusiveTask("检查天数")
//	p.AddTask(cd)
//
//	var mc = NewTask("经理审批")
//	p.AddTask(mc)
//
//	var bc = NewTask("老板审批")
//	p.AddTask(bc)
//
//	p.LinkTask("检查天数", st, cd)
//	p.LinkTask("小于等于3天", cd, mc, NewCondition("day <= 3"))
//	p.LinkTask("大于3天", cd, bc, NewCondition("day > 3"))
//	p.LinkTask("结束", mc, et)
//	p.LinkTask("结束", bc, et)
//
//	fmt.Println(p)
//
//	var nfs = p.NextFlows(cd.TaskId)
//
//	for _, f := range nfs {
//		fmt.Println(f.TargetTask)
//	}
//
//}

//func TestLoadProcess(t *testing.T) {
//	var s = `
//{
// "id": "5bed43d62fbbf7df5f3c0af7",
// "name": "请假流程",
// "start_task_id": "5bed43d62fbbf7df5f3c0af8",
// "end_task_id": "5bed43d62fbbf7df5f3c0af9",
// "task_list": {
//  "5bed43d62fbbf7df5f3c0af8": {
//   "process_id": "5bed43d62fbbf7df5f3c0af7",
//   "task_id": "5bed43d62fbbf7df5f3c0af8",
//   "task_name": "开始",
//   "task_type": 0,
//   "assignee": ""
//  },
//  "5bed43d62fbbf7df5f3c0af9": {
//   "process_id": "5bed43d62fbbf7df5f3c0af7",
//   "task_id": "5bed43d62fbbf7df5f3c0af9",
//   "task_name": "结束",
//   "task_type": 1,
//   "assignee": ""
//  },
//  "5bed43d62fbbf7df5f3c0afa": {
//   "process_id": "5bed43d62fbbf7df5f3c0af7",
//   "task_id": "5bed43d62fbbf7df5f3c0afa",
//   "task_name": "检查天数",
//   "task_type": 3,
//   "assignee": ""
//  },
//  "5bed43d62fbbf7df5f3c0afb": {
//   "process_id": "5bed43d62fbbf7df5f3c0af7",
//   "task_id": "5bed43d62fbbf7df5f3c0afb",
//   "task_name": "经理审批",
//   "task_type": 2,
//   "assignee": ""
//  },
//  "5bed43d62fbbf7df5f3c0afc": {
//   "process_id": "5bed43d62fbbf7df5f3c0af7",
//   "task_id": "5bed43d62fbbf7df5f3c0afc",
//   "task_name": "老板审批",
//   "task_type": 2,
//   "assignee": ""
//  }
// },
// "flow_list": {
//  "5bed43d62fbbf7df5f3c0afd": {
//   "process_id": "5bed43d62fbbf7df5f3c0af7",
//   "flow_id": "5bed43d62fbbf7df5f3c0afd",
//   "flow_name": "检查天数",
//   "source_task_id": "5bed43d62fbbf7df5f3c0af8",
//   "target_task_id": "5bed43d62fbbf7df5f3c0afa"
//  },
//  "5bed43d62fbbf7df5f3c0afe": {
//   "process_id": "5bed43d62fbbf7df5f3c0af7",
//   "flow_id": "5bed43d62fbbf7df5f3c0afe",
//   "flow_name": "小于等于3天",
//   "source_task_id": "5bed43d62fbbf7df5f3c0afa",
//   "target_task_id": "5bed43d62fbbf7df5f3c0afb"
//  },
//  "5bed43d62fbbf7df5f3c0aff": {
//   "process_id": "5bed43d62fbbf7df5f3c0af7",
//   "flow_id": "5bed43d62fbbf7df5f3c0aff",
//   "flow_name": "大于3天",
//   "source_task_id": "5bed43d62fbbf7df5f3c0afa",
//   "target_task_id": "5bed43d62fbbf7df5f3c0afc"
//  },
//  "5bed43d62fbbf7df5f3c0b00": {
//   "process_id": "5bed43d62fbbf7df5f3c0af7",
//   "flow_id": "5bed43d62fbbf7df5f3c0b00",
//   "flow_name": "结束",
//   "source_task_id": "5bed43d62fbbf7df5f3c0afb",
//   "target_task_id": "5bed43d62fbbf7df5f3c0af9"
//  },
//  "5bed43d62fbbf7df5f3c0b01": {
//   "process_id": "5bed43d62fbbf7df5f3c0af7",
//   "flow_id": "5bed43d62fbbf7df5f3c0b01",
//   "flow_name": "结束",
//   "source_task_id": "5bed43d62fbbf7df5f3c0afc",
//   "target_task_id": "5bed43d62fbbf7df5f3c0af9"
//  }
// }
//}
//`
//
//	var p, _ = LoadProcess(s)
//
//	p.Unlink(p.GetTask("5bed43d62fbbf7df5f3c0afa"), p.GetTask("5bed43d62fbbf7df5f3c0afc"))
//
//	fmt.Println(p)
//}
