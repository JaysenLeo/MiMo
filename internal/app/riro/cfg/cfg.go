package cfg

import (
	"fmt"
	"mimo/internal/utils/errors"
	"path"
)

type Option struct {
	repository []string
}
var (
	KnowledgeRoot =  "knowledge"
	KnowledgeIndex = ".repository.index.knowledge.toml"
	FileIndex =      ".repository.index.file.toml"
)

// repository: 仓库列表
func NewRiRoOption(repository []string) *Option {
	return &Option{repository: repository}
}

// 获取知识库的仓库的索引地址
// index: 0 代表默认仓库， 正常情况下默认仓库就是安装或者工程目录下的repository
func (option *Option) GetKnowledgeIndexPath(repoCfgIndex int) (string, error) {
	if p, err := option.GetKnowledgePath(repoCfgIndex);err != nil {
		return "", fmt.Errorf("GetKnowledgeIndexPath: %s", err)
	} else {
		return path.Join(p, KnowledgeIndex), nil
	}

}

// 获取知识库的仓库地址
func (option *Option) GetKnowledgePath(repoCfgIndex int) (string, error) {
	if repoCfgIndex <= len(option.repository) - 1 {
		return path.Join(option.repository[repoCfgIndex], KnowledgeRoot), nil
	}
	return "", errors.Wrap(errors.New("Not found"), "知识库源的索引超出配置范围")
}


