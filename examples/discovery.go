package main

import (
	"fmt"

	"github.com/vexxhost/oidc-utils/pkg/discovery"
)

func main() {
	issuer := "https://accounts.google.com"

	document, err := discovery.DocumentFromIssuer(issuer)
	if err != nil {
		panic(err)
	}

	userInfoEndpoint := document.UserinfoEndpoint

	fmt.Printf("UserInfo endpoint for %s is %s\n", issuer, userInfoEndpoint)
}
