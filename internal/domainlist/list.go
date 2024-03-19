package domainlist

type DomainList struct {
	allows []string
	denies []string
}

func New(allow []string, deny []string) *DomainList {
	return &DomainList{allows: allow, denies: deny}
}

func (l *DomainList) Allow(name string) {
	l.allows = append(l.allows, name)
}

func (l *DomainList) Deny(name string) {
	l.denies = append(l.denies, name)
}

func (l *DomainList) Allows(name string) bool {
	// check deny patterns
	for _, pattern := range l.denies {
		if Match(name, pattern) {
			return false
		}
	}

	// check allow patterns
	if len(l.allows) == 0 {
		return true
	}

	for _, pattern := range l.allows {
		if Match(name, pattern) {
			return true
		}
	}

	return false
}
