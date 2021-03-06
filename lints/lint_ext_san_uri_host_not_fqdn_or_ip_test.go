// lint_ext_san_uri_host_not_fqdn_or_ip_test.go
package lints

import (
	"testing"
)

func TestSANURIHostNotFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINotFQDN.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANURIHostWildcardFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIHostWildcardFQDN.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANURIHostWrongWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIHostWrongWildcard.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANURIHostAsterisk(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIHostAsterisk.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANURIHostFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIHostFQDN.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}


