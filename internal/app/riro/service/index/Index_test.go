package index

import (
	"fmt"
	"mimo/internal/app/riro/cfg"
	"testing"
)



func TestReadIndex(t *testing.T) {
	riro := NewRiRo(func() cfg.Option {
		opt := cfg.NewRiRoOption([]string{"E:\\DevHub\\MiMo\\repository" })
		return *opt
	})
	a := riro.ListAllLocalKnowledge()
	fmt.Printf("ListAllLocalKnowledge: %+v", a)
}

func TestListAllLocalOneKnowledgeFile(t *testing.T) {
	riro := NewRiRo(func() cfg.Option {
		opt := cfg.NewRiRoOption([]string{"E:\\DevHub\\MiMo\\repository" })
		return *opt
	})
	a := riro.ListAllLocalOneKnowledgeFile(1)
	fmt.Printf("ListAllLocalOneKnowledgeFile: %+v", a)
}

func TestGetOneKnowledgeInfoFromIndex(t *testing.T) {
	riro := NewRiRo(func() cfg.Option {
		opt := cfg.NewRiRoOption([]string{"E:\\DevHub\\MiMo\\repository" })
		return *opt
	})
	a := riro.GetOneKnowledgeInfoFromIndex(1, 1)
	fmt.Printf("GetOneKnowledgeInfoFromIndex: %+v", a)
}

func TestGetOneKnowledgeDir(t *testing.T) {
	riro := NewRiRo(func() cfg.Option {
		opt := cfg.NewRiRoOption([]string{"E:\\DevHub\\MiMo\\repository" })
		return *opt
	})
	if d, err := riro.GetOneKnowledgeDir(0,1);err != nil {
		fmt.Printf("err: %+v", err)
	} else {
		fmt.Printf("GetOneKnowledgeDir: %s", d)
	}


}