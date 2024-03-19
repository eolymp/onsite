package main

import (
	"context"
	"github.com/things-go/go-socks5"
	"log"
)

type Rule struct {
	registry *Registry
	logger   *log.Logger
	allows   *PortList
	denies   *PortList
}

func (r *Rule) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	domain, ok := r.registry.Resolved(req.RawDestAddr.IP.String())
	if !ok {
		r.logger.Printf("Destination IP %v was never resolved by DNS server", req.DstAddr.IP)
		return ctx, false
	}

	if r.denies.Allows(domain, req.RawDestAddr.Port) {
		r.logger.Printf("Destination port %v for host %v (%v) is forbidden", req.RawDestAddr.Port, req.RawDestAddr.IP, domain)
		return ctx, false
	}

	if !r.allows.Empty() && !r.allows.Allows(domain, req.RawDestAddr.Port) {
		r.logger.Printf("Destination port %v for host %v (%v) is not allowed", req.RawDestAddr.Port, req.RawDestAddr.IP, domain)
		return ctx, false
	}

	return ctx, true
}
