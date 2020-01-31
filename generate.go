package uuidcli

import (
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
)

func GenerateUUid(version int, domainStr string, ns string, name string) (u uuid.UUID, err error) {
	switch version {
	case 1:
		u, err = uuid.NewV1()
	case 2:
		var domain byte
		switch domainStr {
		case "org":
			domain = uuid.DomainOrg
		case "group":
			domain = uuid.DomainGroup
		case "person":
			domain = uuid.DomainPerson
		default:
			return u, NewExitCodeErr(errors.New("domain (-d) must be one of: org, group or person"), 4)
		}
		u, err = uuid.NewV2(domain)
	case 3:
		if ns == "" {
			return u, NewExitCodeErr(errors.New("version 3 UUID requires namespace (-ns) to be set"), 3)
		}
		nsUuid, err := uuid.FromString(ns)
		if err != nil {
			return u, NewExitCodeErr(fmt.Errorf("could not parse namespace UUID: %w", err), 2)
		}
		if name == "" {
			return u, NewExitCodeErr(errors.New("version 3 UUID requires single argument: [name]"), 3)
		}
		u = uuid.NewV3(nsUuid, name)
	case 4:
		u, err = uuid.NewV4()
	case 5:
		if ns == "" {
			return u, NewExitCodeErr(errors.New("version 3 UUID requires namespace (-ns) to be set"), 3)
		}
		nsUuid, err := uuid.FromString(ns)
		if err != nil {
			return u, NewExitCodeErr(fmt.Errorf("could not parse namespace UUID: %w", err), 2)
		}
		if name == "" {
			return u, NewExitCodeErr(errors.New("version 5 UUID requires single argument: [name]"), 3)
		}
		u = uuid.NewV5(nsUuid, name)
	default:
		return u, NewExitCodeErr(fmt.Errorf("unsupported UUID version: %d", version), 5)
	}

	if err != nil {
		return u, NewExitCodeErr(fmt.Errorf("failed to generate UUID: %w", err), 1)
	}

	return
}
