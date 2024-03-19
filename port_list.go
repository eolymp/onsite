package main

import "sync"

type portRule struct {
	domain  string // domain pattern
	ports   []int  // matching ports
	anyPort bool   // if true, all ports are matching
}

type portCacheKey struct {
	domain string
	port   int
}

type PortList struct {
	lock  sync.Mutex
	rules []portRule
	cache map[portCacheKey]bool
}

func NewPortList() *PortList {
	return &PortList{cache: map[portCacheKey]bool{}}
}

func (l *PortList) Add(domain string, ports ...int) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.rules = append(l.rules, portRule{domain: domain, ports: ports})
}

func (l *PortList) Wildcard(domain string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.rules = append(l.rules, portRule{domain: domain, anyPort: true})
}

func (l *PortList) Empty() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return len(l.rules) == 0
}

func (l *PortList) Allows(domain string, port int) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	key := portCacheKey{domain: domain, port: port}
	if v, ok := l.cache[key]; ok {
		return v
	}

	for _, rule := range l.rules {
		if !MatchDomain(domain, rule.domain) {
			continue
		}

		if rule.anyPort {
			l.cache[key] = true
			return true
		}

		for _, pp := range rule.ports {
			if pp != port {
				continue
			}

			l.cache[key] = true
			return true
		}
	}

	return false
}
