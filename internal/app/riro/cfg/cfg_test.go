package cfg

import (
	"fmt"
	"testing"
)

func TestGetKnowledgePath(t *testing.T) {
	riroOpt := NewRiRoOption([]string{"E:\\DevHub\\MiMo\\repository", })
	fmt.Println(riroOpt.GetKnowledgePath(0))
}