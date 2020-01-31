package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	uuidcli "github.com/networkteam/uuid-cli"
)

func main() {
	v := flag.Int("v", 4, "UUID version (1, 2, 3 or 4), in a nutshell:\n  1: machine\n  2: domain (org, person, group)\n  3: namespaced (MD5)\n  4: random\n  5: namespaced (SHA-1)\n ")
	domainStr := flag.String("d", "", "Use organisation (org), UID (person) or GID (group) for V2 UUID")
	ns := flag.String("ns", "", "Namespace UUID for V3/V5 UUID")
	hexEnc := flag.Bool("hex", false, "Encode UUID as hex number (without hyphens)")

	flag.Parse()

	args := flag.Args()

	var (
		name string
	)

	if len(args) == 1 {
		name = args[0]
	}

	u, err := uuidcli.GenerateUUid(*v, *domainStr, *ns, name)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		if exitCodeErr, ok := err.(uuidcli.ExitCodeErr); ok {
			os.Exit(exitCodeErr.ExitCode)
		}
		os.Exit(-1)
	}

	if *hexEnc {
		fmt.Println(hex.EncodeToString(u.Bytes()))
	} else {
		fmt.Println(u.String())
	}
}
