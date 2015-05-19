package query

import ()

type boolQuery struct {
	MustClauses        []Query  `json:"must"`
	MustNotClauses     []Query  `json:"must_not"`
	ShouldClauses      []Query  `json:"should"`
	boost              *float64 `json:"boost"`
	disableCoord       *bool    `json:"disable_coord"`
	minimumShouldMatch *string  `json:"minimum_should_match"`
	adjustPureNegative *bool    `json:"adjust_pure_negative"`
	queryName          *string  `json:"_name"`
}

func BoolQuery() *boolQuery {
	return &boolQuery{}
}

func (self *boolQuery) Must(queries ...Query) *boolQuery {
	self.MustClauses = append(self.MustClauses, queries...)
	return self
}

func (self *boolQuery) MustNot(queries ...Query) *boolQuery {
	self.MustNotClauses = append(self.MustNotClauses, queries...)
	return self
}

func (self *boolQuery) Should(queries ...Query) *boolQuery {
	self.ShouldClauses = append(self.ShouldClauses, queries...)
	return self
}

func (self *boolQuery) Boost(boost float64) *boolQuery {
	self.boost = &boost
	return self
}

func (self *boolQuery) DisableCoord(disableCoord bool) *boolQuery {
	self.disableCoord = &disableCoord
	return self
}

func (self *boolQuery) MinimumShouldMatch(minimumShouldMatch string) *boolQuery {
	self.minimumShouldMatch = &minimumShouldMatch
	return self
}

func (self *boolQuery) AdjustPureNegative(adjustPureNegative bool) *boolQuery {
	self.adjustPureNegative = &adjustPureNegative
	return self
}

func (self *boolQuery) QueryName(queryName string) *boolQuery {
	self.queryName = &queryName
	return self
}

func (self *boolQuery) MarshalJSON() ([]byte, error) {
	return toJson(wrapper("bool", convertStruct(*self)))
}
