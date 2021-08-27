package index

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"mimo/internal/app/riro/cfg"
	"mimo/internal/utils/errors"
	"path"
)

type RiRo struct {
	RiRoRead
	option cfg.Option
}

type KnowledgeRepo struct {
	Meta Meta
	Knowledge map[string]Knowledge `json:"knowledge"`
}
type FileRepo struct {
	Meta Meta
	File map[string]File `json:"file"`
}

type Meta struct {
	Total int
	Size  string
}

// 本地知识库详情清单
type Knowledge struct {
	Kid			 int64   `json:"kid"`
	Name         string  `json:"name"`
	Cover        string  `json:"cover"`
	CreationTime string  `toml:"creation_time";json:"creation_time"`
	Brief        string  `json:"brief"`
	Store        string  `json:"store"`
}

// 本地某知识库的文件详情清单
type File struct {
	Fid			 int64   `json:"fid"`
	Name         string  `json:"name"`
	Cover        string  `json:"cover"`
	CreationTime string  `toml:"creation_time";json:"creation_time"`
	Brief        string  `json:"brief"`
	Store        string  `json:"store"`
}

type RiRoRead interface {
	ListAllKnowledge()
	ListOneKnowledgeDoc()
	readIndex()
}

func NewRiRo(opt func() cfg.Option) *RiRo {
	return &RiRo{option: opt()}
}

// 读取知识库清单的索引
func (riro *RiRo) readKnowledgeListIndex(repoCfgIndex int) (KnowledgeRepo, error) {
	knowledgeRepo := KnowledgeRepo{}
	if p, err:=riro.option.GetKnowledgeIndexPath(repoCfgIndex);err != nil {
		return KnowledgeRepo{}, errors.Wrap(err, "读取知识库清单的索引")
	} else {
		if _, err := toml.DecodeFile(p, &knowledgeRepo); err != nil {
			return KnowledgeRepo{}, errors.Wrap(err, "读取知识库清单的索引")
		}
	}
	return knowledgeRepo, nil
}

// 读取某知识库文件清单的索引
func (riro *RiRo) readOneKnowledgeFileListIndex(repoCfgIndex int, kid int64) (FileRepo, error) {
	fileRepo := FileRepo{}
	if p, err:=riro.GetOneKnowledgeIndexPath(repoCfgIndex, kid);err != nil {
		return FileRepo{}, errors.Wrap(err, "读取某知识库文件清单的索引")
	} else {
		if _, err := toml.DecodeFile(p, &fileRepo); err != nil {
			return FileRepo{}, errors.Wrap(err, "读取某知识库文件清单的索引")
		}
	}
	return fileRepo, nil
}

// 获取某个知识库的索引中的信息
func (riro *RiRo) GetOneKnowledgeInfoFromIndex(repoCfgIndex int, kid int64) Knowledge {
	allK, _ := riro.readKnowledgeListIndex(repoCfgIndex)
	for _kid, k := range allK.Knowledge {
		if _kid == fmt.Sprint(kid) {
			return k
		}
	}
	return Knowledge{}
}

// 列出所有知识库清单
func (riro *RiRo) ListAllLocalKnowledge() []map[string]interface{} {
	allK, _ := riro.readKnowledgeListIndex(0)
	var allKs []map[string]interface{}

	for i:=1; i <= allK.Meta.Total; i++ {
		k, _ := json.Marshal(allK.Knowledge[fmt.Sprint(i)])
		mRet := map[string]interface{}{}
		fmt.Printf("%v", string(k))
		_ = json.Unmarshal(k, &mRet)
		allKs = append(allKs, mRet)
	}

	return allKs
}

// 列出所有某知识库 文件清单
func (riro *RiRo) ListAllLocalOneKnowledgeFile(kid int64) []map[string]interface{} {
	allK, _ := riro.readOneKnowledgeFileListIndex(0, kid)
	var allKs []map[string]interface{}

	for i:=1; i <= allK.Meta.Total; i++ {
		k, _ := json.Marshal(allK.File[fmt.Sprint(i)])
		mRet := map[string]interface{}{}
		fmt.Printf("%v", string(k))
		_ = json.Unmarshal(k, &mRet)
		allKs = append(allKs, mRet)
	}
	return allKs
}

// 获取某个知识库的目录
func (riro *RiRo) GetOneKnowledgeDir(repoCfgIndex int, kid int64) (string, error) {
	k := riro.GetOneKnowledgeInfoFromIndex(repoCfgIndex, kid)
	if p, err:= riro.option.GetKnowledgePath(repoCfgIndex);err != nil {
		return "", errors.Wrap(err, "获取某个知识库的目录")
	} else {
		return path.Join(p, k.Name + "___" + k.CreationTime), nil
	}
}
// 获取某个知识库的目录
func (riro *RiRo) GetOneKnowledgeIndexPath(repoCfgIndex int, kid int64) (string, error) {
	if p, err:= riro.GetOneKnowledgeDir(repoCfgIndex, kid);err != nil {
		return "", errors.Wrap(err, "获取某个知识库的目录")
	} else {
		return path.Join(p, cfg.FileIndex), nil
	}
}