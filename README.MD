# coredns-plugin-hijacking

This project is a plugin for CoreDNS that allows hijacking DNS requests.

This plugin would response for the `A` DNS query with an wildcard domain, and each record should correspond to the IP for the Kubernertes worker node.

## Usage

When I configuring the ingress for the Kubernetes cluster, I found that I should make up a virtual IP for all the node IP.

But I do not want to configure that vip, so there is a trick solution for that: response watch Kubernetes worker node IP for the ingress domain.

At first, I have used `dnsmasq` for a while, but it would response all the DNS request for `kubernetes.cluster.example.com` with the static record, which makes my cert-manager issuers not working. So I need tool that could only hijack `A` record, but forward other type (like `TXT`) to the upstream DNS server.

## Config Example

```text
kubernetes.cluster.example.com:8053 {
    hijacking {
        record * A 10.233.233.1
        record * A 10.233.233.2
        record * A 10.233.233.3
        record * A 10.233.233.4
    }
    forward . 1.1.1.1
    log
    errors
}
```
