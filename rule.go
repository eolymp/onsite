package main

import (
	"context"
	"github.com/things-go/go-socks5"
	"log"
	"onsite/internal/list"
)

type Rule struct {
	logger *log.Logger
	ips    *list.List[string]
	ports  *list.List[int]
}

func (r *Rule) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	ip, port := req.RawDestAddr.IP.String(), req.RawDestAddr.Port

	if !r.ips.Allows(ip) {
		r.logger.Printf("Destination IP %#v is not allowed", ip)
		return ctx, false
	}

	if !r.ports.Allows(port) {
		r.logger.Printf("Destination port %v is not allowed", port)
		return ctx, false
	}

	return ctx, true
}
