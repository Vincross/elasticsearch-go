package query

type constantScoreQuery struct {
	FilterBuilder Query    `json:"filter"`
	boost         *float64 `json:"boost"`
}

func ConstantScoreQuery(filterBuilder Query) *constantScoreQuery {
	return &constantScoreQuery{FilterBuilder: filterBuilder}
}

func (self *constantScoreQuery) Boost(boost float64) *constantScoreQuery {
	self.boost = &boost
	return self
}

func (self *constantScoreQuery) MarshalJSON() ([]byte, error) {
	return toJson(
		wrapper(
			"constant_score",
			convertStruct(*self),
		),
	)
}
