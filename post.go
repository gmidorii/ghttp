package ghttp

import (
	"net/http"
	"strings"
)

func execPost(g GHttp) (body []byte, err error) {
	resp, err := http.Post(g.getURL(), g.ContentType, strings.NewReader(g.Body))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return branchStatusCode(*resp)
}
