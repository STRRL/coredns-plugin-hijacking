package hijacking

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
)

type Directive string

const (
	DirectiveRecord Directive = "record"
)

var log = clog.NewWithPlugin(name)

func init() { plugin.Register(name, setup) }

func setup(c *caddy.Controller) error {
	ips, err := parseConfig(c)
	if err != nil {
		return err
	}
	config := dnsserver.GetConfig(c)
	config.AddPlugin(func(next plugin.Handler) plugin.Handler {
		hijacking := Hijacking{
			Next:    next,
			Zone:    config.Zone,
			Records: ips,
		}
		return hijacking
	})

	return nil
}
