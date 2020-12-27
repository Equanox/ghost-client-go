package ghostclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/TylerBrock/colorjson"
	"github.com/benchkram/errz"
	"github.com/equanox/ghost-client-go/content"
)

const (
	RoutePosts       string = "/content/posts/"
	RoutePostsById   string = "/content/posts/{id}/"
	RoutePostsBySlug string = "/content/posts/slug/{slug}/"

	RouteAuthors       string = "/content/authors/"
	RouteAuthorsById   string = "/content/authors/{id}/"
	RouteAuthorsBySlug string = "/content/authors/slug/{slug}/ "

	RouteTags       string = "/content/tags/"
	RouteTagsById   string = "/content/tags/{id}/"
	RouteTagsBySlug string = "/content/tags/slug/{slug}/"

	RoutePages       string = "/content/pages/"
	RoutePagesById   string = "/content/pages/{id}/"
	RoutePagesBySlug string = "/content/pages/slug/{slug}/"

	RouteSettings string = "/content/settings/"
)

// C client type for ghost api
type C struct {
	baseUrl *url.URL
	client  *http.Client

	basicAuthUsername string
	basicAuthPassword string

	// print ghost response as pretty-json
	// helpful to debug the ghost api responses.
	verbose bool
}

func New(opts ...Option) *C {
	client := &C{
		baseUrl: nil,
		client:  &http.Client{},

		verbose: false,
	}

	for _, opt := range opts {
		opt(client)
	}

	if client.baseUrl == nil {
		panic("baseurl not set")
	}

	return client
}

// ContentPosts call /posts/ api
// Example: http://localhost:3001/ghost/api/v3/content/posts/?key=62ebdfd5992678506b1e7ddbc3&include=authors,tags
func (c *C) ContentPosts(opts ...ContentOption) (posts *content.Posts, err error) {
	defer errz.Recover(&err)

	// copy the base so we can alter it securly
	u := *c.baseUrl
	u.Path = path.Join(u.Path, RoutePosts)
	for _, opt := range opts {
		opt(&u)
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	errz.Fatal(err)
	if c.basicAuthUsername != "" || c.basicAuthPassword != "" {
		req.SetBasicAuth(c.basicAuthUsername, c.basicAuthPassword)
	}

	resp, err := c.client.Do(req)
	errz.Fatal(err)

	body, err := ioutil.ReadAll(resp.Body)
	errz.Fatal(err)

	if c.verbose {
		var obj map[string]interface{}
		_ = json.Unmarshal(body, &obj)

		f := colorjson.NewFormatter()
		f.Indent = 2

		s, _ := f.Marshal(obj)
		fmt.Println(string(s))
	}

	posts = &content.Posts{}
	err = json.Unmarshal(body, posts)
	errz.Fatal(err)

	return posts, nil
}
