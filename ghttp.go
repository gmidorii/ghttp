package ghttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	HTTP  = "http"
	HTTPS = "https"
)

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

// GHttp is main struct
type GHttp struct {
	Protocol    string
	Domain      string
	Path        string
	Method      string
	Query       map[string]string
	Body        string
	ContentType string
}

// NewGHttp create GHttp struct
func NewGHttp(protocol, domain, path, method string) *GHttp {
	return &GHttp{
		Protocol: protocol,
		Domain:   domain,
		Path:     path,
		Method:   method,
		Query:    make(map[string]string),
	}
}

// PutQuery is putting query paramaeter.
// same key overwrite new value
func (g *GHttp) PutQuery(k, v string) {
	g.Query[k] = v
}

// PutContentType is putting Content-Type
func (g *GHttp) PutContentType(c string) {
	g.ContentType = c
}

// Exec is requesting server and reserving response.
func (g *GHttp) Exec() ([]byte, error) {
	switch g.Method {
	case GET:
		return execGet(*g)
	case POST:
		return execPost(*g)
	case DELETE:
	case PUT:
	default:
		return nil, fmt.Errorf("unexpected method: %v", g.Method)
	}
}

func (g GHttp) getURL() string {
	return fmt.Sprintf("%s://%s%s", g.Protocol, g.Domain, g.Path)
}

func branchStatusCode(resp http.Response) ([]byte, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return ioutil.ReadAll(resp.Body)
	case http.StatusBadRequest:
		return nil, fmt.Errorf("client bad request(status code: %v)", http.StatusBadRequest)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("internal server error (status code: %v", http.StatusInternalServerError)
	default:
		return nil, fmt.Errorf("status code: %v", resp.StatusCode)
	}
}
