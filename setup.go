package ace

import (
	"net/http"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

func init() {
	caddy.RegisterPlugin("ace", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

// setup configures a new Markdown middleware instance.
func setup(c *caddy.Controller) error {
	aceconfigs, err := aceParse(c)
	if err != nil {
		return err
	}

	cfg := httpserver.GetConfig(c)

	ace := Ace{
		Root:       cfg.Root,
		FileSys:    http.Dir(cfg.Root),
		Configs:    aceconfigs,
		IndexFiles: []string{"index.ace"},
	}

	cfg.AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		ace.Next = next
		return ace
	})

	return nil
}

func aceParse(c *caddy.Controller) ([]*Config, error) {
	var aceconfigs []*Config

	ace := &Config{
		Path:       "",
		Extensions: make(map[string]struct{}),
	}

	for c.Next() {

		args := c.RemainingArgs()

		switch len(args) {
		case 0:
			continue
		case 1:
			ace.Path = args[0]
		default:
			continue
		}

	}

	aceconfigs = append(aceconfigs, ace)

	return aceconfigs, nil
}
