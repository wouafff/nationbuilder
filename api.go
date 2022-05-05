package nationbuilder

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const apiVersion = "v1"
const debug = false
const defaultLimit = 50

type nationbuilderURL struct {
	u url.URL
}

func (n *nationbuilderURL) setQuery(key string, val string) {
	q := n.u.Query()
	q.Set(key, val)
	n.u.RawQuery = q.Encode()
}

func (n *nationbuilderURL) setLimit(limit int) {
	n.setQuery("limit", strconv.Itoa(limit))
}

func (n *nationbuilderURL) setToken(token string) {
	n.setQuery("access_token", token)
}

func (n *nationbuilderURL) extendPath(path string) {
	if len(path) > 0 {
		if string(path[0]) != "/" {
			n.u.Path += "/"
		}

		n.u.Path += path
	}
}

func (n *nationbuilderURL) String() string {
	return n.u.String()
}

type Client struct {
	Slug    string
	ApiKey  string
	baseURL *nationbuilderURL
	c       *http.Client
}

func (n *Client) getRequest(method string, path string, options *Options) *apiRequest {
	b := *n.baseURL
	b.extendPath(path)

	if options != nil {
		options.setQuery(&b.u)
	}

	return &apiRequest{
		url:    b.String(),
		method: method,
	}

}

// By default http.DefaultClient is used to make requests but if you need to set additional options
// such as a proxy or you are running on Google App Engine, then you may want to supply a different
// http client
func (n *Client) SetHTTPClient(c *http.Client) {
	n.c = c
}

func NewClientForOauth(slug string) (*Client, error) {
	u, err := url.Parse(fmt.Sprintf("https://%s.nationbuilder.com", slug))
	if err != nil {
		return nil, err
	}

	nbURL := &nationbuilderURL{
		u: *u,
	}

	return &Client{
		Slug:    slug,
		ApiKey:  "",
		baseURL: nbURL,
		c:       http.DefaultClient,
	}, nil
}

// Creates a new Nationbuilder Client
func NewClient(slug string, key string) (*Client, error) {
	u, err := url.Parse(fmt.Sprintf("https://%s.nationbuilder.com/api/%s", slug, apiVersion))
	if err != nil {
		return nil, err
	}

	nbURL := &nationbuilderURL{
		u: *u,
	}

	nbURL.setToken(key)

	return &Client{
		Slug:    slug,
		ApiKey:  key + "&fire_webhooks=false",
		baseURL: nbURL,
		c:       http.DefaultClient,
	}, nil
}
