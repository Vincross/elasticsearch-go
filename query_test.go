package query

import (
	"encoding/json"
	"fmt"
	"testing"

	. "elasticsearch-go/query"
)

func TestTermQuery(t *testing.T) {
	q1 := TermQuery("title", "test").Boost(1)
	data, _ := json.MarshalIndent(q1, "", "    ")
	fmt.Println(string(data))
}

func TestBoolQuery(t *testing.T) {
	q1 := TermQuery("title", "test").Boost(1)
	q2 := BoolQuery().MustNot(q1).Boost(2).QueryName("querytest")

	data, _ := json.MarshalIndent(q2, "", "    ")
	fmt.Println(string(data))
}

func TestSpanNearQuery(t *testing.T) {
	q1 := SpanTermQuery("title", "test1")

	q3 := SpanNearQuery().
		Clause(q1).
		Clause(SpanTermQuery("title", "test2").Boost(3)).
		Boost(2).
		QueryName("querytest")

	data, _ := json.MarshalIndent(q3, "", "    ")
	fmt.Println(string(data))
}

func TestDisMaxQuery(t *testing.T) {
	q1 := TermQuery("title", "test1")

	q3 := DisMaxQuery().
		Add(q1).
		Add(TermQuery("title", "test2").Boost(3)).
		Boost(2).
		QueryName("querytest").
		TieBreaker(1)

	data, _ := json.MarshalIndent(q3, "", "    ")
	fmt.Println(string(data))
}

func TestMatchAllQuery(t *testing.T) {
	q1 := MatchAllQuery().Boost(2)
	data, _ := json.MarshalIndent(q1, "", "    ")
	fmt.Println(string(data))
}

func TestConstantScoreQuery(t *testing.T) {
	q1 := ConstantScoreQuery(TermQuery("title", "test1")).Boost(2)
	data, _ := json.MarshalIndent(q1, "", "    ")
	fmt.Println(string(data))
}
