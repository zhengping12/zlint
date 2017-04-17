// lint_subject_dn_leading_whitespace.go

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/util"
	"strings"
)

type SubjectDNLeadingSpace struct {
	// Internal data here
}

func (l *SubjectDNLeadingSpace) Initialize() error {
	return nil
}

func (l *SubjectDNLeadingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SubjectDNLeadingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	var subject pkix.RDNSequence
	if _, err := asn1.Unmarshal(c.RawSubject, &subject); err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	for _, rdn := range subject {
		if len(rdn) == 0 {
			continue
		}
		atv := rdn[0]
		value, ok := atv.Value.(string)
		if !ok {
			continue
		}

		t := atv.Type
		if len(t) == 4 && t[0] == 2 && t[1] == 5 && t[2] == 4 && util.IsAttributeInList(t[3]) {
			if strings.HasPrefix(value, " ") {
				return ResultStruct{Result: Warn}, nil
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_subject_dn_leading_whitespace",
		Description:   "AttributeValue in subject RelativeDistinguishedName sequence should not have leading whitespace",
		Providence:    "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SubjectDNLeadingSpace{}})
}
