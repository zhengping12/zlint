// lint_ext_san_uniform_resource_identifier_present.go
/************************************************************************************************************
7.1.4.2.1. Subject Alternative Name Extension
Certificate Field: extensions:subjectAltName
Required/Optional:  Required
Contents:  This extension MUST contain at least one entry.  Each entry MUST be either a dNSName containing
the Fully‐Qualified Domain Name or an iPAddress containing the IP address of a server.  The CA MUST
confirm that the Applicant controls the Fully‐Qualified Domain Name or IP address or has been granted the
right to use it by the Domain Name Registrant or IP address assignee, as appropriate.
Wildcard FQDNs are permitted.
*************************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANURI struct {
	// Internal data here
}

func (l *SANURI) Initialize() error {
	return nil
}

func (l *SANURI) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SANOID)
}

func (l *SANURI) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.URIs != nil {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_uniform_resource_identifier_present",
		Description:   "The Subject Alternate Name extension must contain only dnsName and ipaddress name types.",
		Providence:    "CAB: 7.1.4.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &SANURI{}})
}
