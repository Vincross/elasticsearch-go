package search

import (
	"fmt"
	"testing"

	"github.com/offercal/elasticsearch-go/query"
)

func TestClient(t *testing.T) {
	client := NewClient("https://offercal:offercal2015@es.offercal.com")
	res, _ := client.prepareSearch("articles").SetTypes("article").
		SetFrom(0).SetSize(10).SetQuery(query.TermQuery("title", "中国")).Do()
	for _, x := range res.GetHits() {
		fmt.Println(x.Id)
	}
}
