// SPDX-FileCopyrightText: Copyright 2025 Carabiner Systems, Inc
// SPDX-License-Identifier: Apache-2.0

// Package main implements the Grumpy Connector 2000 or GC2K for short.
// The whole purpose of GC2K is to absolutely never connect to a proxy
// or handle a JWT but will gladly insult you if you try.
package main

import (
	"fmt"
	"os"

	// A note on module versions:
	// For this demo to work you need to start with the following versions
	// of the dependencies:

	"github.com/carabiner-dev/lab-vexable-repo/pkg/insult"

	// go-jose 4.0.2
	// Just linkning the module ensures CVE-2025-27144 becomes visible to the
	// scanner, but since we don't use it, the vulnerability is never called.
	_ "github.com/go-jose/go-jose/v4"

	// x/net 0.35.0
	"golang.org/x/net/proxy"
)

// main tries to connect
func main() {
	address := "[::1%25.example.com]:80"
	if len(os.Args) > 2 {
		address = os.Args[1]
	}

	// This calls the Dial function of the proxy.PerHost struct which is affected
	// by CVE-2025-22870.
	// //
	// But note that we are not really affected as we using a fake dialer that
	// whose purpose is to insult you and always fails to connect:
	ph := proxy.NewPerHost(insult.NewDialer(), insult.NewDialer())

	// Open a connection with the InsultDialer:
	_, err := ph.Dial("tcp", address)

	// Say hi.
	fmt.Printf("\n\nError: %v\n\n", err)
}
