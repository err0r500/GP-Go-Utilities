package helpers

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

var client *http.Client

func init() {
	client = &http.Client{
		// Transport: &http2.Transport{
		// 	TLSClientConfig: tlsClientConfig(),
		// },
	}
}

func tlsClientConfig() *tls.Config {
	var crt []byte
	var err error
	crtString := os.Getenv("TLS_LOCALHOST_PUBLIC_CRT")
	if crtString != "" {
		crt = []byte(crtString)
	} else {
		crt, err = ioutil.ReadFile("./cert/public.crt")
		if err != nil {
			log.Fatal(err)
		}
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		ServerName:         "localhost",
	}
}

// TLSServerConfig method
func TLSServerConfig() *tls.Config {
	var crt, key []byte
	var err error
	crtString := os.Getenv("TLS_LOCAL_PUBLIC_CRT")
	if crtString != "" {
		crt = []byte(crtString)
	} else {
		crt, err = ioutil.ReadFile("./cert/public.crt")
		if err != nil {
			log.Fatal(err)
		}
	}

	keyString := os.Getenv("TLS_LOCAL_PRIVATE_KEY")
	if keyString != "" {
		key = []byte(keyString)
	} else {
		key, err = ioutil.ReadFile("./cert/private.key")
		if err != nil {
			log.Fatal(err)
		}
	}

	cert, err := tls.X509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}
}

// GetJson method
func GetJson(c context.Context, url string, expectedCodeHTTPResponseCode int, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	txn := newrelic.FromContext(c)
	txn.AddAttribute("url_full", url)
	s := newrelic.StartExternalSegment(txn, req)
	resp, err := client.Do(req)
	s.Response = resp
	s.End()
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError, "Unexpected error occurred")
	}
	defer resp.Body.Close()
	if resp.StatusCode != expectedCodeHTTPResponseCode {
		var responseError map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&responseError); err != nil {
			txn.AddAttribute("err", err.Error())
			return NewHTTPError(err, resp.StatusCode, "Unexpected response from back API - StatusCode: "+strconv.Itoa(resp.StatusCode))
		}
		err = fmt.Errorf("Unexpected response from back API - %v: %v", resp.StatusCode, responseError)
		txn.AddAttribute("err", err.Error())
		return NewHTTPError(
			err,
			resp.StatusCode,
			"Unexpected response from back API - StatusCode: "+strconv.Itoa(resp.StatusCode),
		)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return NewHTTPError(err, http.StatusBadRequest, "Failed during parsing back api response")
	}

	return nil
}

// PostJson method
func PostJson(c context.Context, url string, expectedCodeHTTPResponseCode int, data, target interface{}) error {
	bodyBytes, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	txn := newrelic.FromContext(c)
	txn.AddAttribute("body", string(bodyBytes))
	s := newrelic.StartExternalSegment(txn, req)
	resp, err := client.Do(req)
	s.Response = resp
	s.End()
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError, "Unexpected error occurred")
	}
	defer resp.Body.Close()
	if resp.StatusCode != expectedCodeHTTPResponseCode {
		var responseError map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&responseError); err != nil {
			txn.AddAttribute("err", err.Error())
			return NewHTTPError(err, resp.StatusCode, "Unexpected response from back API - StatusCode: "+strconv.Itoa(resp.StatusCode))
		}
		err = fmt.Errorf("Unexpected response from back API - %v: %v", resp.StatusCode, responseError)
		txn.AddAttribute("err", err.Error())
		return NewHTTPError(
			err,
			resp.StatusCode,
			"Unexpected response from back API - StatusCode: "+strconv.Itoa(resp.StatusCode),
		)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return NewHTTPError(err, http.StatusBadRequest, "Failed during parsing back api response")
	}

	return nil
}

// DeleteJson method
func DeleteJson(c context.Context, url string, expectedCodeHTTPResponseCode int, data, target interface{}) error {
	bodyBytes, _ := json.Marshal(data)

	// Create request
	req, err := http.NewRequest("DELETE", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError, "Unexpected error occurred")
	}

	txn := newrelic.FromContext(c)
	txn.AddAttribute("body", string(bodyBytes))
	s := newrelic.StartExternalSegment(txn, req)
	req.Header.Set("Content-Type", "application/json")

	// Fetch Request
	resp, err := client.Do(req)
	s.Response = resp
	s.End()
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError, "Unexpected error occurred")
	}
	defer resp.Body.Close()
	if resp.StatusCode != expectedCodeHTTPResponseCode {
		var responseError map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&responseError); err != nil {
			txn.AddAttribute("err", err.Error())
			return NewHTTPError(err, resp.StatusCode, "Unexpected response from back API - StatusCode: "+strconv.Itoa(resp.StatusCode))
		}
		err = fmt.Errorf("Unexpected response from back API - %v: %v", resp.StatusCode, responseError)
		txn.AddAttribute("err", err.Error())
		return NewHTTPError(
			err,
			resp.StatusCode,
			"Unexpected response from back API - StatusCode: "+strconv.Itoa(resp.StatusCode),
		)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return NewHTTPError(err, http.StatusBadRequest, "Failed during parsing back api response")
	}

	return nil
}
