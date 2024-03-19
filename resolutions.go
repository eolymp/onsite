package main

import (
	"sync"
)

type Registry struct {
	lock sync.RWMutex
	ips  map[string]string
}

func (r *Registry) Add(domain string, ip string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.ips == nil {
		r.ips = map[string]string{}
	}

	r.ips[ip] = domain
}

func (r *Registry) Resolved(ip string) (domain string, ok bool) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	if r.ips == nil {
		return "", false
	}

	domain, ok = r.ips[ip]
	return
}
