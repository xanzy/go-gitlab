package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPagesDomainsService_ListPagesDomains(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/pages/domains", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"domain": "ssl.domain.example",
				"url": "https://ssl.domain.example",
				"auto_ssl_enabled": false,
				"certificate": {
				  "subject": "/O=Example, Inc./OU=Example Origin CA/CN=Example Origin Certificate",
				  "expired": false,
				  "certificate": "-----BEGIN CERTIFICATE-----\n … \n-----END CERTIFICATE-----",
				  "certificate_text": "Certificate:\n … \n"
				}
			  }
			]
		`)
	})

	want := []*PagesDomain{{
		Domain:           "ssl.domain.example",
		AutoSslEnabled:   false,
		URL:              "https://ssl.domain.example",
		ProjectID:        0,
		Verified:         false,
		VerificationCode: "",
		EnabledUntil:     nil,
		Certificate: struct {
			Subject         string     `json:"subject"`
			Expired         bool       `json:"expired"`
			Expiration      *time.Time `json:"expiration"`
			Certificate     string     `json:"certificate"`
			CertificateText string     `json:"certificate_text"`
		}{
			Expired:         false,
			Expiration:      nil,
			Subject:         "/O=Example, Inc./OU=Example Origin CA/CN=Example Origin Certificate",
			Certificate:     "-----BEGIN CERTIFICATE-----\n … \n-----END CERTIFICATE-----",
			CertificateText: "Certificate:\n … \n",
		},
	}}

	pds, resp, err := client.PagesDomains.ListPagesDomains(5, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pds)

	pds, resp, err = client.PagesDomains.ListPagesDomains(5.01, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pds)

	pds, resp, err = client.PagesDomains.ListPagesDomains(5, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pds)

	pds, resp, err = client.PagesDomains.ListPagesDomains(7, nil)
	require.Error(t, err)
	require.Nil(t, pds)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestPagesDomainsService_ListAllPagesDomains(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/pages/domains", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"domain": "ssl.domain.example",
				"url": "https://ssl.domain.example",
				"project_id": 1337,
				"auto_ssl_enabled": false,
				"certificate": {
				  "expired": false
				}
			  }
			]
		`)
	})

	want := []*PagesDomain{{
		Domain:           "ssl.domain.example",
		AutoSslEnabled:   false,
		URL:              "https://ssl.domain.example",
		ProjectID:        1337,
		Verified:         false,
		VerificationCode: "",
		EnabledUntil:     nil,
		Certificate: struct {
			Subject         string     `json:"subject"`
			Expired         bool       `json:"expired"`
			Expiration      *time.Time `json:"expiration"`
			Certificate     string     `json:"certificate"`
			CertificateText string     `json:"certificate_text"`
		}{
			Expired:    false,
			Expiration: nil,
		},
	}}

	pds, resp, err := client.PagesDomains.ListAllPagesDomains(nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pds)

	pds, resp, err = client.PagesDomains.ListAllPagesDomains(nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pds)
}

func TestPagesDomainsService_ListAllPagesDomains_StatusNotFound(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/pages/domains", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
	})

	pds, resp, err := client.PagesDomains.ListAllPagesDomains(nil)
	require.Error(t, err)
	require.Nil(t, pds)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestPagesDomainsService_GetPagesDomain(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/pages/domains/www.domain.example", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"domain": "www.domain.example",
			"url": "https://ssl.domain.example",
			"auto_ssl_enabled": false,
			"certificate": {
			  "subject": "/O=Example, Inc./OU=Example Origin CA/CN=Example Origin Certificate",
			  "expired": false,
			  "certificate": "-----BEGIN CERTIFICATE-----\n … \n-----END CERTIFICATE-----",
			  "certificate_text": "Certificate:\n … \n"
			}
		  }
		`)
	})

	want := &PagesDomain{
		Domain:           "www.domain.example",
		AutoSslEnabled:   false,
		URL:              "https://ssl.domain.example",
		ProjectID:        0,
		Verified:         false,
		VerificationCode: "",
		EnabledUntil:     nil,
		Certificate: struct {
			Subject         string     `json:"subject"`
			Expired         bool       `json:"expired"`
			Expiration      *time.Time `json:"expiration"`
			Certificate     string     `json:"certificate"`
			CertificateText string     `json:"certificate_text"`
		}{
			Expired:         false,
			Expiration:      nil,
			Subject:         "/O=Example, Inc./OU=Example Origin CA/CN=Example Origin Certificate",
			Certificate:     "-----BEGIN CERTIFICATE-----\n … \n-----END CERTIFICATE-----",
			CertificateText: "Certificate:\n … \n",
		},
	}

	pd, resp, err := client.PagesDomains.GetPagesDomain(5, "www.domain.example", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pd)

	pd, resp, err = client.PagesDomains.GetPagesDomain(5.01, "www.domain.example", nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pd)

	pd, resp, err = client.PagesDomains.GetPagesDomain(5, "www.domain.example", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pd)

	pd, resp, err = client.PagesDomains.GetPagesDomain(7, "www.domain.example", nil)
	require.Error(t, err)
	require.Nil(t, pd)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestPagesDomainsService_CreatePagesDomain(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/pages/domains", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"domain": "ssl.domain.example",
			"url": "https://ssl.domain.example",
			"auto_ssl_enabled": false,
			"certificate": {
			  "subject": "/O=Example, Inc./OU=Example Origin CA/CN=Example Origin Certificate",
			  "expired": false,
			  "certificate": "-----BEGIN CERTIFICATE-----\n … \n-----END CERTIFICATE-----",
			  "certificate_text": "Certificate:\n … \n"
			}
		  }
		`)
	})

	want := &PagesDomain{
		Domain:           "ssl.domain.example",
		AutoSslEnabled:   false,
		URL:              "https://ssl.domain.example",
		ProjectID:        0,
		Verified:         false,
		VerificationCode: "",
		EnabledUntil:     nil,
		Certificate: struct {
			Subject         string     `json:"subject"`
			Expired         bool       `json:"expired"`
			Expiration      *time.Time `json:"expiration"`
			Certificate     string     `json:"certificate"`
			CertificateText string     `json:"certificate_text"`
		}{
			Expired:         false,
			Expiration:      nil,
			Subject:         "/O=Example, Inc./OU=Example Origin CA/CN=Example Origin Certificate",
			Certificate:     "-----BEGIN CERTIFICATE-----\n … \n-----END CERTIFICATE-----",
			CertificateText: "Certificate:\n … \n",
		},
	}

	pd, resp, err := client.PagesDomains.CreatePagesDomain(5, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pd)

	pd, resp, err = client.PagesDomains.CreatePagesDomain(5.01, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pd)

	pd, resp, err = client.PagesDomains.CreatePagesDomain(5, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pd)

	pd, resp, err = client.PagesDomains.CreatePagesDomain(7, nil)
	require.Error(t, err)
	require.Nil(t, pd)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestPagesDomainsService_UpdatePagesDomain(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/pages/domains/ssl.domain.example", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"domain": "ssl.domain.example",
			"url": "https://ssl.domain.example",
			"auto_ssl_enabled": false,
			"certificate": {
			  "subject": "/O=Example, Inc./OU=Example Origin CA/CN=Example Origin Certificate",
			  "expired": false,
			  "certificate": "-----BEGIN CERTIFICATE-----\n … \n-----END CERTIFICATE-----",
			  "certificate_text": "Certificate:\n … \n"
			}
		  }
		`)
	})

	want := &PagesDomain{
		Domain:           "ssl.domain.example",
		AutoSslEnabled:   false,
		URL:              "https://ssl.domain.example",
		ProjectID:        0,
		Verified:         false,
		VerificationCode: "",
		EnabledUntil:     nil,
		Certificate: struct {
			Subject         string     `json:"subject"`
			Expired         bool       `json:"expired"`
			Expiration      *time.Time `json:"expiration"`
			Certificate     string     `json:"certificate"`
			CertificateText string     `json:"certificate_text"`
		}{
			Expired:         false,
			Expiration:      nil,
			Subject:         "/O=Example, Inc./OU=Example Origin CA/CN=Example Origin Certificate",
			Certificate:     "-----BEGIN CERTIFICATE-----\n … \n-----END CERTIFICATE-----",
			CertificateText: "Certificate:\n … \n",
		},
	}

	pd, resp, err := client.PagesDomains.UpdatePagesDomain(5, "ssl.domain.example", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pd)

	pd, resp, err = client.PagesDomains.UpdatePagesDomain(5.01, "ssl.domain.example", nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pd)

	pd, resp, err = client.PagesDomains.UpdatePagesDomain(5, "ssl.domain.example", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pd)

	pd, resp, err = client.PagesDomains.UpdatePagesDomain(7, "ssl.domain.example", nil)
	require.Error(t, err)
	require.Nil(t, pd)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestPagesDomainsService_DeletePagesDomain(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/pages/domains/ssl.domain.example", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.PagesDomains.DeletePagesDomain(5, "ssl.domain.example", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.PagesDomains.DeletePagesDomain(5.01, "ssl.domain.example", nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.PagesDomains.DeletePagesDomain(5, "ssl.domain.example", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.PagesDomains.DeletePagesDomain(7, "ssl.domain.example", nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
