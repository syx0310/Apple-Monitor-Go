package apple

// QueryBuilder 用于构建查询参数
type QueryBuilder struct {
	params map[string]string
}

// NewQueryBuilder 创建一个新的 QueryBuilder 实例
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		params: make(map[string]string),
	}
}

// Set 添加一个键值对到查询参数中
func (qb *QueryBuilder) Set(key, value string) *QueryBuilder {
	qb.params[key] = value
	return qb
}

// Build 返回构建好的查询参数映射
func (qb *QueryBuilder) Build() map[string]string {
	return qb.params
}
