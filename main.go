package main

import (
	"flag"
	"github.com/things-go/go-socks5"
	"log"
	"onsite/internal/domainlist"
	"onsite/internal/list"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "proxy: ", log.LstdFlags)

	flag.Parse()

	filename := "config.yaml"
	if args := flag.Args(); len(args) > 1 {
		filename = args[1]
	}

	config, err := ParseConfig(filename)
	if err != nil {
		logger.Fatalf("Unable to read configuration file %#v: %v", filename, err)
	}

	ips := list.New(config.AllowedIP, config.ForbiddenIP)
	ports := list.New(config.AllowedPorts, config.ForbiddenPorts)
	domains := domainlist.New(config.AllowedDomains, config.ForbiddenDomains)

	for _, rule := range config.Rules {
		if rule.Deny != "" && len(rule.Ports) == 0 {
			domains.Deny(rule.Deny)
		}

		if rule.Allow != "" {
			domains.Allows(rule.Allow)
		}
	}

	resolver := &Resolver{
		logger:  logger,
		parent:  &socks5.DNSResolver{},
		domains: domains,
	}

	if config.AllowResolvedIPs {
		resolver.ips = ips
	}

	rule := &Rule{
		logger: logger,
		ips:    ips,
		ports:  ports,
	}

	// Create a SOCKS5 server
	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(logger)),
		socks5.WithRule(rule),
		socks5.WithResolver(resolver),
	)

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", ":8000"); err != nil {
		panic(err)
	}
}
