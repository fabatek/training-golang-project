package testhelpers

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func JSONRequest(ctx context.Context, t *testing.T, httpMethod, path string, requestBody io.Reader, urlParamsCallback func(rctx *chi.Context), handlerCallback func(req *http.Request, rec *httptest.ResponseRecorder)) (*http.Response, interface{}) {
	if handlerCallback == nil {
		t.Fatalf("handler callback cannot be nil")
	}

	r := chi.NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest(httpMethod, ts.URL+path, requestBody)
	if err != nil {
		t.Fatalf("Unable to create request. Error: %v", err)
	}

	rctx := chi.NewRouteContext()
	if urlParamsCallback != nil {
		urlParamsCallback(rctx)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	}

	rec := httptest.NewRecorder()
	handlerCallback(req, rec)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response: %+v. Error: %v", resp.Body, err)
	}
	// API only return HTTP status code without response body
	if string(body) == "" {
		return resp, nil
	}

	var jsonBody interface{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		t.Fatalf("unable to unmarshal response body: %v. Error: %v", string(body), err)
	}

	return resp, jsonBody
}

// LoadJSONFixture is a convenient method to load JSON fixture from a file
func LoadJSONFixture(t *testing.T, path string) interface{} {
	fixture, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("Unable to read %s. Err: %v", path, err)
	}

	var JSONFixture interface{}
	if err = json.Unmarshal(fixture, &JSONFixture); err != nil {
		t.Fatalf("unable to unmarshal JSON fixture: %v. Err: %v", string(fixture), err)
	}

	return JSONFixture
}
