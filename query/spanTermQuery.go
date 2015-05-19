package query

type spanTermQuery struct {
	Name      string      `json:"-"`
	Value     interface{} `json:"value"`
	boost     *float64    `json:"boost"`
	queryName *string     `json:"_name"`
}

func SpanTermQuery(name string, value interface{}) *spanTermQuery {
	return &spanTermQuery{Name: name, Value: value}
}

func (self *spanTermQuery) QueryName(queryName string) *spanTermQuery {
	self.queryName = &queryName
	return self
}

func (self *spanTermQuery) Boost(boost float64) *spanTermQuery {
	self.boost = &boost
	return self
}

func (self *spanTermQuery) MarshalJSON() ([]byte, error) {
	return toJson(
		wrapper(
			"span_term",
			wrapper(
				self.Name,
				convertStruct(*self),
			),
		),
	)
}
