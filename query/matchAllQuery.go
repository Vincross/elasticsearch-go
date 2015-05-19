package query

type matchAllQuery struct {
	boost *float64 `json:"boost"`
}

func MatchAllQuery() *matchAllQuery {
	return &matchAllQuery{}
}

func (self *matchAllQuery) Boost(boost float64) *matchAllQuery {
	self.boost = &boost
	return self
}

func (self *matchAllQuery) MarshalJSON() ([]byte, error) {
	return toJson(
		wrapper(
			"match_all",
			convertStruct(*self),
		),
	)
}
