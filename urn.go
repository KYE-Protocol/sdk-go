// KYE Protocol™ URN — public format.
//
//	kye:<entity-class>:<trust-domain>:<subclass>:<local>
//
// Full normative URN spec:
//
//	https://kye-protocol.github.io/protocol.html#urn-format
//	https://github.com/KYE-Protocol/id-format
package kye

import (
	"fmt"
	"regexp"
)

var urnRE = regexp.MustCompile(`^kye:([a-z][a-z0-9-]*):([a-z0-9.-]+):([a-z0-9-]+):([a-zA-Z0-9_-]+)$`)

// URN is a parsed KYE Protocol™ URN.
type URN struct {
	EntityClass string
	TrustDomain string
	Subclass    string
	Local       string
}

// ParseURN parses a URN string. Returns an error if malformed.
func ParseURN(s string) (URN, error) {
	m := urnRE.FindStringSubmatch(s)
	if m == nil {
		return URN{}, fmt.Errorf("invalid KYE URN: %q", s)
	}
	return URN{
		EntityClass: m[1],
		TrustDomain: m[2],
		Subclass:    m[3],
		Local:       m[4],
	}, nil
}

// IsValidURN reports whether s is a well-formed KYE URN.
func IsValidURN(s string) bool {
	return urnRE.MatchString(s)
}

// IsPublicClassURN reports whether s parses and its entity class is one
// of the public-vocabulary EntityClass* constants.
func IsPublicClassURN(s string) bool {
	u, err := ParseURN(s)
	if err != nil {
		return false
	}
	for _, c := range EntityClasses {
		if c == u.EntityClass {
			return true
		}
	}
	return false
}

// String returns the canonical URN form.
func (u URN) String() string {
	return fmt.Sprintf("kye:%s:%s:%s:%s", u.EntityClass, u.TrustDomain, u.Subclass, u.Local)
}

// Validate re-parses the URN; returns an error if any field is malformed.
func (u URN) Validate() error {
	_, err := ParseURN(u.String())
	return err
}
