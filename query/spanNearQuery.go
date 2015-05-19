package query

type spanNearQuery struct {
	Clauses         []Query  `json:"clauses"`
	slop            *int     `json:"slop"`
	inOrder         *bool    `json:"in_order"`
	collectPayloads *bool    `json:"collect_payloads"`
	boost           *float64 `json:"boost"`
	queryName       *string  `json:"_name"`
}

func SpanNearQuery() *spanNearQuery {
	return &spanNearQuery{}
}

func (self *spanNearQuery) Clause(clause Query) *spanNearQuery {
	self.Clauses = append(self.Clauses, clause)
	return self
}

func (self *spanNearQuery) Slop(slop int) *spanNearQuery {
	self.slop = &slop
	return self
}

func (self *spanNearQuery) InOrder(inOrder bool) *spanNearQuery {
	self.inOrder = &inOrder
	return self
}

func (self *spanNearQuery) CollectPayloads(collectPayloads bool) *spanNearQuery {
	self.collectPayloads = &collectPayloads
	return self
}

func (self *spanNearQuery) QueryName(queryName string) *spanNearQuery {
	self.queryName = &queryName
	return self
}

func (self *spanNearQuery) Boost(boost float64) *spanNearQuery {
	self.boost = &boost
	return self
}

func (self *spanNearQuery) MarshalJSON() ([]byte, error) {
	return toJson(
		wrapper(
			"span_near",
			convertStruct(*self),
		),
	)
}
