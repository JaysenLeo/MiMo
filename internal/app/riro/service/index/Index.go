package index

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"mimo/internal/app/riro/cfg"
)

type RiRo struct {
	RiRoRead
	option cfg.Option
}

type KnowledgeRepo struct {
	Meta Meta
	Knowledge map[string]Knowledge
}
type Meta struct {
	Total int
	Size  string
}
type Knowledge struct {
	Name         string  `json:"name"`
	Cover        string  `json:"cover"`
	CreationTime string  `json:"creation_time"`
	Store        string  `json:"store"`
}

type RiRoRead interface {
	ListAllKnowledge()
	readIndex()
}

func NewRiRo(opt func() cfg.Option) *RiRo {
	return &RiRo{option: opt()}
}

func (riro *RiRo) readIndex() KnowledgeRepo {
	knowledgeRepo := KnowledgeRepo{}
	if _, err := toml.DecodeFile(riro.option.GetKnowledgePath(0), &knowledgeRepo); err != nil {
		log.Fatal(err)
	}
	return knowledgeRepo
}
func (riro *RiRo) ListAllKnowledge() []map[string]string {
	allK := riro.readIndex()
	var allKs []map[string]string

	for i:=1; i <= allK.Meta.Total; i++ {
		k, _ := json.Marshal(allK.Knowledge[fmt.Sprint(i)])
		mRet := map[string]string{}
		_ = json.Unmarshal(k, &mRet)
		allKs = append(allKs, mRet)
	}

	return allKs
}