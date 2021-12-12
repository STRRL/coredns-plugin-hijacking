package hijacking

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
)

var log = clog.NewWithPlugin("hihacking")

func init() { plugin.Register(name, setup) }

func setup(c *caddy.Controller) error {
	config := dnsserver.GetConfig(c)

	remainingArgs := c.RemainingArgs()
	log.Info(remainingArgs)

	config.AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Hijacking{
			Next: next,
			Zone: config.Zone,
		}
	})

	return nil
}
