package leetcode

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestClientBuildsGraphQLRequestWithUsernameVariable(t *testing.T) {
	var captured *http.Request
	var capturedBody []byte
	client := NewClient(WithHTTPClient(&http.Client{
		Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
			captured = request
			var err error
			capturedBody, err = io.ReadAll(request.Body)
			if err != nil {
				t.Fatalf("ReadAll request body: %v", err)
			}
			return jsonResponse(http.StatusOK, `{"data":{"matchedUser":{"username":"alice","profile":{"realName":"Alice","ranking":123,"reputation":7},"submitStatsGlobal":{"acSubmissionNum":[]},"languageProblemCount":[]}}}`)
		}),
	}))

	_, err := client.fetchProfileData(context.Background(), "alice")
	if err != nil {
		t.Fatalf("fetchProfileData returned error: %v", err)
	}

	if captured == nil {
		t.Fatal("no request captured")
	}
	if captured.Method != http.MethodPost {
		t.Fatalf("method = %q, want POST", captured.Method)
	}
	if captured.URL.String() != defaultGraphQLEndpoint {
		t.Fatalf("url = %q, want %q", captured.URL.String(), defaultGraphQLEndpoint)
	}
	if got := captured.Header.Get("Content-Type"); got != "application/json" {
		t.Fatalf("Content-Type = %q, want application/json", got)
	}
	if got := captured.Header.Get("User-Agent"); got != defaultUserAgent {
		t.Fatalf("User-Agent = %q, want %q", got, defaultUserAgent)
	}

	var payload graphQLRequest
	if err := json.Unmarshal(capturedBody, &payload); err != nil {
		t.Fatalf("request body is not JSON: %v", err)
	}
	if payload.Variables["username"] != "alice" {
		t.Fatalf("variables username = %q, want alice", payload.Variables["username"])
	}
	if strings.Contains(payload.Query, "alice") {
		t.Fatalf("query interpolated username: %q", payload.Query)
	}
	if !strings.Contains(payload.Query, "$username") {
		t.Fatalf("query = %q, want username variable", payload.Query)
	}
}

func TestClientClassifiesGraphQLFailures(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantKind   ErrorKind
	}{
		{
			name:       "not found",
			statusCode: http.StatusOK,
			body:       `{"data":{"matchedUser":null}}`,
			wantKind:   ErrorKindNotFound,
		},
		{
			name:       "rate limited",
			statusCode: http.StatusTooManyRequests,
			body:       `{}`,
			wantKind:   ErrorKindRateLimited,
		},
		{
			name:       "access blocked",
			statusCode: http.StatusForbidden,
			body:       `{}`,
			wantKind:   ErrorKindRateLimited,
		},
		{
			name:       "endpoint failure",
			statusCode: http.StatusBadGateway,
			body:       `{}`,
			wantKind:   ErrorKindEndpointFailure,
		},
		{
			name:       "graphql errors",
			statusCode: http.StatusOK,
			body:       `{"errors":[{"message":"resolver failed"}],"data":{"matchedUser":null}}`,
			wantKind:   ErrorKindUnavailable,
		},
		{
			name:       "malformed json",
			statusCode: http.StatusOK,
			body:       `{`,
			wantKind:   ErrorKindMalformedResponse,
		},
		{
			name:       "missing data object",
			statusCode: http.StatusOK,
			body:       `{}`,
			wantKind:   ErrorKindMalformedResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(WithHTTPClient(&http.Client{
				Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
					return jsonResponse(tt.statusCode, tt.body)
				}),
			}))

			_, err := client.fetchProfileData(context.Background(), "alice")

			if !IsErrorKind(err, tt.wantKind) {
				t.Fatalf("error = %v, want kind %s", err, tt.wantKind)
			}
		})
	}
}

func TestClientClassifiesTransportFailure(t *testing.T) {
	client := NewClient(WithHTTPClient(&http.Client{
		Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			return nil, errors.New("network unavailable")
		}),
	}))

	_, err := client.fetchProfileData(context.Background(), "alice")

	if !IsErrorKind(err, ErrorKindEndpointFailure) {
		t.Fatalf("error = %v, want endpoint failure", err)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func jsonResponse(statusCode int, body string) (*http.Response, error) {
	return &http.Response{
		StatusCode: statusCode,
		Header:     make(http.Header),
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}, nil
}
