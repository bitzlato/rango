package main

import (
	"flag"
	"fmt"

	"github.com/openware/rango/pkg/auth"
)

var (
	p2p_uid        = flag.Int("p2p_uid", 123, "P2P user ID")
	uid            = flag.String("uid", "IDABC0000001", "UID")
	email          = flag.String("email", "admin@barong.io", "Email")
	role           = flag.String("role", "admin", "Role")
	level          = flag.Int("level", 3, "Level")
	email_verified = flag.Bool("email_verified", true, "Is Email Verified")
	// port           = flag.String("port", "7070", "Port to bind")
	// sub            = flag.String("sub", "auth0|user@example.com", "Auth subject (typicaly from auth0)")
)

func main() {
	flag.Parse()

	ks, _ := auth.LoadOrGenerateKeys("config/rsa-key", "config/rsa-key.pub")
	t, _ := auth.ForgeToken(*p2p_uid, *uid, *email, *role, *level, *email_verified, ks.PrivateKey, nil)
	fmt.Print(t)
}
