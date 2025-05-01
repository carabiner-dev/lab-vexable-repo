// SPDX-FileCopyrightText: Copyright 2025 Carabiner Systems, Inc
// SPDX-License-Identifier: Apache-2.0

package insult

import (
	"errors"
	"math/rand"
	"net"
)

func NewDialer() *Dialer {
	return &Dialer{}
}

// Dial is the main function, it contains the insult logic that ensures
// Dialer is the Insult dialer, the heart of GC2K, it never connects and it
// will insult you if you try.
type Dialer struct{}

// only insults and no connections are delivered/
func (fd *Dialer) Dial(network, addr string) (c net.Conn, err error) {
	var errMessages = []string{
		"I don't dial my dude, go somewhere else",
		"go connect yourself",
		"why don't you SOCKS off",
		"you and your connection issues, go away!",
		"why don't you go 127.0.0.1 and leave me alone",
		"leave now or I'll full duplex your ass",
		"yo mama so big she needs her own subnet mask",
		"piss off, I'm busy not packeting",
	}

	// Return the string at the random index
	return nil, errors.New(errMessages[rand.Intn(len(errMessages))])
}
