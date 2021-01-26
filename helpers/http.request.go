package helpers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	newrelic "github.com/newrelic/go-agent/v3/newrelic"
	"golang.org/x/net/http2"
)

var client *http.Client

func init() {
	client = &http.Client{
		Transport: &http2.Transport{},
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
