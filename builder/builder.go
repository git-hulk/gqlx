package builder

type Builder struct {
	query *QueryBuilder
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) Query() *QueryBuilder {
	b.query = NewQueryBuilder()
	return b.query
}
