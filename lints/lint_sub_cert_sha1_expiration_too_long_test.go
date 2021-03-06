// lint_sub_cert_sha1_expiration_too_long_test.go
package lints

import (
	"testing"
)

func TestRsaSha1TooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/sha1ExpireAfter2017.cer"
	desEnum := Warn
	out, _ := Lints["w_sub_cert_sha1_expiration_too_long"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRsaSha1NotTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/sha1ExpirePrior2017.cer"
	desEnum := Pass
	out, _ := Lints["w_sub_cert_sha1_expiration_too_long"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
