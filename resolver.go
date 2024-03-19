package main

import (
	"context"
	"fmt"
	"github.com/things-go/go-socks5"
	"log"
	"net"
	"onsite/internal/domainlist"
	"onsite/internal/list"
)

type Resolver struct {
	parent  socks5.NameResolver
	domains *domainlist.DomainList
	ips     *list.List[string]
	logger  *log.Logger
}

func (r *Resolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	if !r.domains.Allows(name) {
		return ctx, nil, fmt.Errorf("domain \"%v\" is not allowed", name)
	}

	// call parent resolver to get an actual IP
	ctx, ip, err := r.parent.Resolve(ctx, name)
	if err != nil {
		return ctx, nil, err
	}

	// automatically whitelist all IP addresses resolved by DNS server
	if r.ips != nil && !r.ips.AllowHas(ip.String()) {
		r.logger.Printf("Domain %#v resolved to IP %#v, adding IP to allow list", name, ip.String())
		r.ips.Allow(ip.String())
	}

	return ctx, ip, nil
}
