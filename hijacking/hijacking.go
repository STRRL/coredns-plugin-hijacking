package hijacking

import (
	"context"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
	"net"
)

const name = "hijacking"
const zeroTTL = 0

type DNSRecordType string

const (
	TypeA DNSRecordType = "A"
)

type Hijacking struct {
	Next       plugin.Handler
	Zone       string
	RecordType DNSRecordType
	Records    []net.IP
}

// ServeDNS implements the plugin.Handler interface.
func (it Hijacking) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	if matched := plugin.Zones([]string{it.Zone}).Matches(state.Name()); matched == "" {
		return plugin.NextOrFailure(it.Name(), it.Next, ctx, w, r)
	}

	var answer []dns.RR
	for _, item := range it.Records {
		answer = append(answer,
			&dns.A{
				Hdr: dns.RR_Header{
					Name:   state.Name(),
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    zeroTTL,
				},
				A: item,
			})
	}

	dnsMsg := dns.Msg{
		MsgHdr: dns.MsgHdr{
			Authoritative: true,
		},
		Answer: answer,
	}
	dnsMsg.SetReply(r)
	err := w.WriteMsg(&dnsMsg)
	return dns.RcodeSuccess, err
}

// Name implements the Handler interface.
func (it Hijacking) Name() string { return name }
