package query

type termQuery struct {
	Name      string      `json:"-"`
	Value     interface{} `json:"value"`
	boost     *float64    `json:"boost"`
	queryName *string     `json:"_name"`

	simplified bool `json:"-"`
}

func TermQuery(name string, value interface{}) *termQuery {
	return &termQuery{Name: name, Value: value, simplified: true}
}

func (self *termQuery) QueryName(queryName string) *termQuery {
	self.queryName = &queryName
	self.simplified = false
	return self
}

func (self *termQuery) Boost(boost float64) *termQuery {
	self.boost = &boost
	self.simplified = false
	return self
}

func (self *termQuery) MarshalJSON() ([]byte, error) {
	if self.simplified {
		return toJson(wrapper(
			"term",
			map[string]interface{}{
				self.Name: self.Value,
			},
		),
		)
	}
	return toJson(
		wrapper(
			"term",
			wrapper(
				self.Name,
				convertStruct(*self),
			),
		),
	)
}
