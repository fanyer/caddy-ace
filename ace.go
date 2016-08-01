package ace

import (
	"fmt"
	"net/http"
	"path"

	"github.com/mholt/caddy/caddyhttp/httpserver"
	"github.com/yosssi/ace"
)

type Ace struct {
	// Server root
	Root string

	// Jail the requests to site root with a mock file system
	FileSys http.FileSystem

	// Next HTTP handler in the chain
	Next httpserver.Handler

	// The list of ace configurations
	Configs []*Config

	// The list of index files to try
	IndexFiles []string
}

type Config struct {

	// Base path to match
	Path string

	// List of extensions to consider as markdown files
	Extensions map[string]struct{}
}

// ServeHTTP implements the http.Handler interface.
func (a Ace) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	var cfg *Config
	for _, c := range a.Configs {
		if httpserver.Path(r.URL.Path).Matches(c.Path) { // not negated
			cfg = c
			break // or goto
		}
	}

	fmt.Println("ace!!")

	if cfg == nil {
		return a.Next.ServeHTTP(w, r) // exit early
	}

	// We only deal with HEAD/GET
	switch r.Method {
	case http.MethodGet, http.MethodHead:
	default:
		return http.StatusMethodNotAllowed, nil
	}

	fullpath := r.URL.Path
	basepath, filename := path.Split(fullpath)

	// fmt.Println(basepath)
	// fmt.Println(filename)

	if filename == "" {
		filename = "index"
	}

	tpl, err := ace.Load("."+basepath+"base", "."+basepath+filename, nil)

	if err != nil {
		fmt.Println(err)
		// http.NotFound(w, r)
		return a.Next.ServeHTTP(w, r)
	}

	data := map[string]interface{}{}
	if err := tpl.Execute(w, data); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return a.Next.ServeHTTP(w, r)
	}

	return http.StatusOK, nil
}

// ace is middleware to render templated files as the HTTP response.
