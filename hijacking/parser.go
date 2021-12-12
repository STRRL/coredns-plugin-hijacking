package hijacking

import (
	"github.com/coredns/caddy"
	"net"
)

func parseConfig(c *caddy.Controller) ([]net.IP, error) {
	var result []net.IP
	for c.Next() {
		directive := c.Val()
		if directive != string(DirectiveRecord) {
			continue
		}

		remainingArgs := c.RemainingArgs()
		if remainingArgs[0] != domainWildcard {
			continue
		}
		if remainingArgs[1] != string(TypeA) {
			continue
		}

		if ip := net.ParseIP(remainingArgs[2]); ip != nil {
			result = append(result, ip)
			rule := []string{directive}
			rule = append(rule, remainingArgs...)
			log.Info("load rule: ", rule)
		}
	}
	return result, nil
}
