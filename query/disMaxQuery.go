package query

type disMaxQuery struct {
	Queries    []Query  `json:"queries"`
	boost      *float64 `json:"boost"`
	tieBreaker *float64 `json:"tie_breaker"`
	queryName  *string  `json:"_name"`
}

func DisMaxQuery() *disMaxQuery {
	return &disMaxQuery{}
}

func (self *disMaxQuery) Add(query Query) *disMaxQuery {
	self.Queries = append(self.Queries, query)
	return self
}

func (self *disMaxQuery) Boost(boost float64) *disMaxQuery {
	self.boost = &boost
	return self
}

func (self *disMaxQuery) TieBreaker(tieBreaker float64) *disMaxQuery {
	self.tieBreaker = &tieBreaker
	return self
}

func (self *disMaxQuery) QueryName(queryName string) *disMaxQuery {
	self.queryName = &queryName
	return self
}

func (self *disMaxQuery) MarshalJSON() ([]byte, error) {
	return toJson(
		wrapper(
			"dis_max",
			convertStruct(*self),
		),
	)
}
