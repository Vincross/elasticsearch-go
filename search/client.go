package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/offercal/elasticsearch-go/query"
)

var client *http.Client = &http.Client{}

type Client struct {
	baseUrl string
}

func NewClient(baseUrl string) *Client {
	return &Client{baseUrl}
}

func (self *Client) PrepareSearch(indices ...string) *searcher {
	return newSearcher(self.baseUrl, indices)
}

type searcher struct {
	baseUrl string
	query   query.Query
	indices []string
	types   []string
	from    *int
	size    *int
}

func newSearcher(baseUrl string, indices []string) *searcher {
	return &searcher{baseUrl: baseUrl, indices: indices}
}

func (self *searcher) SetFrom(from int) *searcher {
	self.from = &from
	return self
}

func (self *searcher) SetSize(size int) *searcher {
	self.size = &size
	return self
}

func (self *searcher) SetTypes(types ...string) *searcher {
	self.types = append(self.types, types...)
	return self
}

func (self *searcher) SetQuery(query query.Query) *searcher {
	self.query = query
	return self
}

func (self *searcher) Do() (*SearchResponse, error) {
	req, err := http.NewRequest("GET", self.buildURL(), self.buildBody())
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	result := &SearchResponse{}

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(result)

	return result, nil
}

func (self *searcher) buildURL() string {
	url := self.baseUrl
	if len(self.indices) == 0 {
		url = url + "/*"
	} else {
		url = url + "/" + strings.Join(self.indices, ",")
	}
	if len(self.types) == 0 {
		url = url + "/*"
	} else {
		url = url + "/" + strings.Join(self.types, ",")
	}

	url += "/_search"
	if self.from == nil {
		url += "?from=0"
	} else {
		url += "?from=" + strconv.Itoa(*self.from)
	}
	if self.size != nil {
		url += "&size=" + strconv.Itoa(*self.size)
	}
	return url
}

func (self *searcher) buildBody() io.Reader {
	requestBody := map[string]interface{}{
		"query": self.query,
	}
	data, _ := json.Marshal(requestBody)
	reqBody := &bytes.Buffer{}
	reqBody.Write(data)

	return reqBody
}
