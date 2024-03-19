package main

import (
	"strings"
)

// MatchDomain against a pattern:
//   - "*"     matches all domains
//   - "xyz"   matches domain xyz
//   - "*.xyz" matches all subdomains of xyz
func MatchDomain(value string, pattern string) bool {
	if pattern == "*" {
		return true
	}

	if !strings.ContainsRune(pattern, '*') {
		return pattern == value
	}

	pp := strings.Split(pattern, ".")
	vp := strings.Split(value, ".")

	if len(pp) != len(vp) {
		return false
	}

	for i := 0; i < len(pp); i++ {
		if pp[i] != vp[i] && pp[i] != "*" {
			return false
		}
	}

	return true
}

type DomainList []string

func (l DomainList) Matches(name string) bool {
	for _, pattern := range l {
		if MatchDomain(name, pattern) {
			return true
		}
	}

	return false
}
