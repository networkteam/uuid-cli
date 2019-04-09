package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
)

func main() {
	v := flag.Int("v", 4, "UUID version (1, 2, 3 or 4), defaults to 4")
	domainStr := flag.String("d", "", "Use organisation (org), UID (person) or GID (group) for V2 UUID")
	ns := flag.String("ns", "", "Namespace UUID for V3 UUID")

	flag.Parse()

	args := flag.Args()

	var (
		u   uuid.UUID
		err error
	)
	switch *v {
	case 1:
		u, err = uuid.NewV1()
	case 2:
		var domain byte
		switch *domainStr {
		case "org":
			domain = uuid.DomainOrg
		case "group":
			domain = uuid.DomainGroup
		case "person":
			domain = uuid.DomainPerson
		default:
			_, _ = fmt.Fprintf(os.Stderr, "Domain (-d) must be one of: org, group or person\n")
			os.Exit(4)
		}
		u, err = uuid.NewV2(domain)
	case 3:
		if *ns == "" {
			_, _ = fmt.Fprintf(os.Stderr, "V3 UUID requires namespace (-ns) to be set\n")
			os.Exit(3)
		}
		nsUuid, err := uuid.FromString(*ns)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Could not parse namespace UUID: %v\n", err)
			os.Exit(2)
		}
		if len(args) != 1 {
			_, _ = fmt.Fprintf(os.Stderr, "V3 UUID expects a single argument: [name]\n")
			os.Exit(3)
		}
		name := args[0]
		u = uuid.NewV3(nsUuid, name)
	case 4:
		u, err = uuid.NewV4()
	case 5:
		if *ns == "" {
			_, _ = fmt.Fprintf(os.Stderr, "V3 UUID requires namespace (-ns) to be set\n")
			os.Exit(3)
		}
		nsUuid, err := uuid.FromString(*ns)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Could not parse namespace UUID: %v\n", err)
			os.Exit(2)
		}
		if len(args) != 1 {
			_, _ = fmt.Fprintf(os.Stderr, "V5 UUID expects a single argument: [name]\n")
			os.Exit(3)
		}
		name := args[0]
		u = uuid.NewV5(nsUuid, name)
	}

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Could not generate UUID: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(u)
}
