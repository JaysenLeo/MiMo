package cfg

import "path"

type Option struct {
	repository []string
}
var (
	KnowledgeRoot = "knowledge"
	KnowledgeIndex = ".repository.index.knowledge.toml"
)
func NewRiRoOption(repository []string) *Option {
	return &Option{repository: repository}
}

// 获取知识库的仓库地址
func (option *Option) GetKnowledgePath(index int) string {
	return path.Join(option.repository[index], KnowledgeRoot, KnowledgeIndex)
}
