// lint_cert_policy_conflicts_with_street.go
// If the Certificate asserts the policy identifier of 2.23.140.1.2.1, then it MUST NOT include
// organizationName, streetAddress, localityName, stateOrProvinceName, or postalCode in the Subject field.

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type certPolicyConflictsWithStreet struct {
	// Internal data here
}

func (l *certPolicyConflictsWithStreet) Initialize() error {
	return nil
}

func (l *certPolicyConflictsWithStreet) CheckApplies(cert *x509.Certificate) bool {
	return util.SliceContainsOID(cert.PolicyIdentifiers, util.BRDomainValidatedOID)
}

func (l *certPolicyConflictsWithStreet) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	var out ResultStruct
	if util.TypeInName(&cert.Subject, util.StreetAddressOID) {
		out.Result = Error
	} else {
		out.Result = Pass
	}
	return out, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_cab_dv_conflicts_with_street",
		Description:   "If certificate policy 2.23.140.1.2.1 (CA/B BR domain validated) is included, streetAddress must not be included in subject.",
		Providence:    "CAB: 7.1.6.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &certPolicyConflictsWithStreet{}})
}
