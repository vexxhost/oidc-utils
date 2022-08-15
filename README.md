# `oidc-utils`

This is a library that allows you to make some basic interactions with OpenID
connect APIs.

## `discovery`

This is a simple package that can allow you to retrieve values for OpenID
connect endpoints using discovery, it can also present raw data as well for
other storing purposes.

This module has been tested for the following OpenID connect providers for
succesful parsing:

- Google
- Microsoft
- Yahoo
- PayPal
- Salesforce
- Okta
- Keycloak

> **Note**
>
> This module enforces a strict parsing for JSON so it will error out if you
> actually try to parse a discovery URL that does not match the schema laid out,
> so we're happy to accept pull requests to handle them, however, we test for
> a large majority of different providers.

### Example

```go
package main

import (
	"github.com/vexxhost/oidc-utils/pkg/discovery"
)

func main() {
	issuer := "https://accounts.google.com"

	document, err := discovery.DocumentFromIssuer(issuer)
	if err != nil {
		panic(err)
	}

  // Do something with `document` or store `document.RawData`.
}
```
