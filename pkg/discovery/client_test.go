package discovery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testDocumentFromIssuer(t *testing.T, issuer string) {
	document, err := DocumentFromIssuer(issuer)

	if assert.NoError(t, err) {
		assert.Equal(t, issuer, document.Issuer)
	}
}

func TestDocumentFromIssuerGoogle(t *testing.T) {
	testDocumentFromIssuer(t, "https://accounts.google.com")
}
func TestDocumentFromIssuerMicrosoft(t *testing.T) {
	testDocumentFromIssuer(t, "https://login.microsoftonline.com/775527ff-9a37-4307-8b3d-cc311f58d925/v2.0")
}

func TestDocumentFromIssuerYahoo(t *testing.T) {
	testDocumentFromIssuer(t, "https://api.login.yahoo.com")
}

func TestDocumentFromIssuerPayPal(t *testing.T) {
	testDocumentFromIssuer(t, "https://www.paypal.com")
}

func TestDocumentFromIssuerSalesforce(t *testing.T) {
	testDocumentFromIssuer(t, "https://login.salesforce.com")
}

func TestDocumentFromIssuerOkta(t *testing.T) {
	testDocumentFromIssuer(t, "https://okta.okta.com")
}

func TestDocumentFromIssuerKeycloak(t *testing.T) {
	testDocumentFromIssuer(t, "https://id.vexxhost.com/auth/realms/external")
}
