package main

import (
	"context"
	"fmt"
	"github.com/things-go/go-socks5"
	"log"
	"net"
)

type Resolver struct {
	parent   socks5.NameResolver
	allows   DomainList
	denies   DomainList
	registry *Registry
	logger   *log.Logger
}

func (r *Resolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	if r.denies.Matches(name) {
		return ctx, nil, fmt.Errorf("domain \"%v\" is forbidden", name)
	}

	if len(r.allows) > 0 && !r.allows.Matches(name) {
		return ctx, nil, fmt.Errorf("domain \"%v\" is not allowed", name)
	}

	ctx, ip, err := r.parent.Resolve(ctx, name)
	if err != nil {
		return ctx, nil, err
	}

	r.registry.Add(name, ip.String())

	return ctx, ip, nil
}
