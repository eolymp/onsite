package main

import "testing"

func TestMatchDomain(t *testing.T) {
	tt := []struct {
		Pattern string
		Domain  string
		Match   bool
	}{
		{Pattern: "eolymp.com", Domain: "eolymp.com", Match: true},
		{Pattern: "eolymp.com", Domain: "www.eolymp.com", Match: false},
		{Pattern: "*.eolymp.com", Domain: "www.eolymp.com", Match: true},
		{Pattern: "*.eolymp.com", Domain: "sl.www.eolymp.com", Match: false},
		{Pattern: "*.*.eolymp.com", Domain: "sl.www.eolymp.com", Match: true},
		{Pattern: "www.*.com", Domain: "www.eolymp.com", Match: true},
		{Pattern: "www.*.com", Domain: "www.basecamp.eolymp.com", Match: false},
	}

	for _, tc := range tt {
		t.Run(tc.Pattern, func(t *testing.T) {
			got := MatchDomain(tc.Domain, tc.Pattern)

			if got != tc.Match {
				if tc.Match {
					t.Fatalf("Domain %#v must match pattern %#v, but it does not", tc.Domain, tc.Pattern)
				} else {
					t.Fatalf("Domain %#v must NOT match pattern %#v, but it does", tc.Domain, tc.Pattern)
				}
			}
		})
	}
}
