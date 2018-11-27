package wf4go

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var p = map[string]interface{}{
		"day":     30,
		"hour":    5,
		"my_hour": 10,
	}
	fmt.Println(exec("", p))
}
