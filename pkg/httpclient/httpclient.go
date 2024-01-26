package httpclient

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/koor-tech/data-control-center/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("http_client",
	fx.Provide(New),
)

func New(logger *zap.Logger, cfg *config.Config) (*http.Client, error) {
	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	for _, certFile := range cfg.Certs.CACerts {
		// Read cert file
		certs, err := os.ReadFile(certFile)
		if err != nil {
			return nil, fmt.Errorf("failed to append %q to RootCAs: %w", certFile, err)
		}

		// Append cert to the system pool
		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			log.Println("No certs appended, using system certs only")
		}
	}

	// Trust the augmented cert pool in our client
	config := &tls.Config{
		InsecureSkipVerify: cfg.Certs.InsecureSkipVerify,
		RootCAs:            rootCAs,
	}

	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	return client, nil
}
