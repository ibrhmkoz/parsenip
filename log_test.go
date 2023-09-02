package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseOnLog(t *testing.T) {
	format := `
    CA Name{:e}: {:s:CAName}
    Template Name{:e}: {:s:TemplateName}
{:i}
    msPKI-Certificate-Name-Flag{:e}: {:a:MSPKICertificateNameFlag}
    mspki-enrollment-flag{:e}: {:a:MSPKIEnrollmentFlag}
{:i}
    pkiextendedkeyusage{:e}: {:a:PKIExtendedKeyUsage}
    mspki-certificate-application-policy{:e}: {:a:MSPKICertificateApplicationPolicy}`

	dat, err := os.ReadFile("Certify.log")
	if err != nil {
		panic(err)
	}
	target := string(dat)

	expected := map[string]interface{}{}

	result, err := Parse(format, target)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
