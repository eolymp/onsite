package main

import (
	"flag"
	"github.com/things-go/go-socks5"
	"log"
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

	domainAllowList := DomainList{}
	domainDenyList := DomainList{}
	portAllowList := NewPortList()
	portDenyList := NewPortList()

	for _, rule := range config.Rules {
		// deny rule: if no ports are set, domain resolution will be disabled all together
		if rule.Deny != "" {
			if len(rule.Ports) == 0 {
				domainDenyList = append(domainDenyList, rule.Deny)
				portDenyList.Wildcard(rule.Deny)
			} else {
				portDenyList.Add(rule.Deny, rule.Ports...)
			}
		}

		// allow rule: if no ports are set, add wildcard allow port
		if rule.Allow != "" {
			domainAllowList = append(domainAllowList, rule.Allow)

			if len(rule.Ports) == 0 {
				portAllowList.Wildcard(rule.Allow)
			} else {
				portAllowList.Add(rule.Allow, rule.Ports...)
			}
		}
	}

	registry := &Registry{}
	resolver := &Resolver{
		logger:   logger,
		parent:   &socks5.DNSResolver{},
		registry: registry,
		allows:   domainAllowList,
		denies:   domainDenyList,
	}

	rule := &Rule{
		registry: registry,
		logger:   logger,
		allows:   portAllowList,
		denies:   portDenyList,
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
