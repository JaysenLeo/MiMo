package index

import (
	"fmt"
	"mimo/internal/app/riro/cfg"
	"testing"
)



func TestReadIndex(t *testing.T) {
	riro := NewRiRo(func() cfg.Option {
		opt := cfg.NewRiRoOption([]string{"E:\\DevHub\\MiMo\\repository", })
		return *opt
	})
	fmt.Printf("%+v", riro.ListAllKnowledge())
}
