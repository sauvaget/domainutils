package domainutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		answer := IsValid(tc.domain)
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
		answer := ExtractTld(tc.domain)
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
		answer := IsValidTld(tc.tld)
		assert.Equal(t, tc.expect, answer)
	}
}
