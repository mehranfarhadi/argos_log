package bootstrap

import "crypto/tls"

func NewTlsConfig() (*tls.Config, error) {
	// For testing purposes, we're not loading any CA certificate
	// and allowing connections without server verification.

	return &tls.Config{
		InsecureSkipVerify: true, // Disable server certificate verification for testing
	}, nil
}
