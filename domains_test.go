package domainutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		domain string
		expect bool
	}{
		{"google.com", true},
		{"google.de", true},
		{"google.ca", true},
		{"google.co.uk", true},
		{"google.de/", false},
		{"google.dedsl", false},
		{"google", false},
		{"google.", false},
		{"http://google.de", false},
		{"8.8.8.8", false},
	}

	for _, tc := range testCases {
		d := New(tc.domain)
		answer, err := d.IsValid()
		if err != nil {

		}
		assert.Equal(t, tc.expect, answer)
	}
}

func TestExtractTld(t *testing.T) {
	testCases := []struct {
		domain string
		expect string
	}{
		{"google.com", "com"},
		{"google.de", "de"},
		{"google.ca", "ca"},
		{"google.co.uk", "uk"},
		{"google.de/", "de/"},
		{"google.dedsl", "dedsl"},
		{"google", "google"},
		{"google.", ""},
		{"http://google.de", "de"},
		{"8.8.8.8", "8"},
	}

	for _, tc := range testCases {
		d := New(tc.domain)
		answer, err := d.ExtractTld()
		if err != nil {

		}
		assert.Equal(t, tc.expect, answer)
	}
}

func TestHasValidTld(t *testing.T) {
	testCases := []struct {
		domain string
		expect bool
	}{
		{"google.com", true},
		{"google.de", true},
		{"google.ca", true},
		{"google.co.uk", true},
		{"google.de/", false},
		{"google.dedsl", false},
		{"google", true},
		{"google.", false},
		{"http://google.de", true},
		{"8.8.8.8", false},
	}

	for _, tc := range testCases {
		d := New(tc.domain)
		answer, err := d.HasValidTld()
		if err != nil {

		}
		assert.Equal(t, tc.expect, answer)
	}
}

func TestIsValidTld(t *testing.T) {
	testCases := []struct {
		tld    string
		expect bool
	}{
		{"com", true},
		{"de", true},
		{"ca", true},
		{"uk", true},
		{"co.uk", false},
		{"de/", false},
		{"dedsl", false},
		{"google", true},
		{"google.", false},
		{"http://google.de", false},
		{"8.8.8.8", false},
	}

	for _, tc := range testCases {
		d := New("")
		answer, err := d.IsValidTld(tc.tld)
		if err != nil {

		}
		assert.Equal(t, tc.expect, answer)
	}
}
