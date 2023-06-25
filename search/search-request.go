package search

type SearchQuery struct {
	Q                     string   `json:"q"`
	Offset                *int     `json:"offset"`
	Limit                 *int     `json:"limit"`
	Page                  *int     `json:"page"`
	HighlightPreTag       string   `json:"highlightPreTag"`
	HighlightPostTag      string   `json:"highlightPostTag"`
	ShowMatchesPosition   bool     `json:"showMatchesPosition"`
	Sort                  []string `json:"sort"`
	AttributesToHighlight []string `json:"attributesToHighlight"`
}

type SearchQueryBuilder struct {
	query SearchQuery
}

func NewSearchQueryBuilder() *SearchQueryBuilder {
	return &SearchQueryBuilder{query: SearchQuery{}}
}

func (b *SearchQueryBuilder) SetAttributesToHighlight(highlight []string) *SearchQueryBuilder {
	b.query.AttributesToHighlight = highlight
	return b
}

func (b *SearchQueryBuilder) SetShowMatchesPosition(show bool) *SearchQueryBuilder {
	b.query.ShowMatchesPosition = show
	return b
}

func (b *SearchQueryBuilder) SetQ(q string) *SearchQueryBuilder {
	b.query.Q = q
	return b
}

func (b *SearchQueryBuilder) SetOffset(offset int) *SearchQueryBuilder {
	b.query.Offset = &offset
	return b
}

func (b *SearchQueryBuilder) SetLimit(limit int) *SearchQueryBuilder {
	b.query.Limit = &limit
	return b
}

func (b *SearchQueryBuilder) SetPage(page int) *SearchQueryBuilder {
	b.query.Page = &page
	return b
}

func (b *SearchQueryBuilder) SetHighlightPreTag(highlightPreTag string) *SearchQueryBuilder {
	b.query.HighlightPreTag = highlightPreTag
	return b
}

func (b *SearchQueryBuilder) SetHighlightPostTag(highlightPostTag string) *SearchQueryBuilder {
	b.query.HighlightPostTag = highlightPostTag
	return b
}

func (b *SearchQueryBuilder) SetSort(sort []string) *SearchQueryBuilder {
	b.query.Sort = sort
	return b
}

func (b *SearchQueryBuilder) Build() SearchQuery {
	return b.query
}
