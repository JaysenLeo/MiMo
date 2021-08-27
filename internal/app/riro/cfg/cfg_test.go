package cfg

import (
	"fmt"
	"testing"
)

func TestGetKnowledgeIndexPath(t *testing.T) {
	riroOpt := NewRiRoOption([]string{"E:\\DevHub\\MiMo\\repository", })
	fmt.Println(riroOpt.GetKnowledgeIndexPath(0))
}

func TestGetKnowledgePath(t *testing.T) {
	riroOpt := NewRiRoOption([]string{"E:\\DevHub\\MiMo\\repository", })
	fmt.Println(riroOpt.GetKnowledgePath(0))
}