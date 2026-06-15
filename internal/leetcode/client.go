package leetcode

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	defaultGraphQLEndpoint = "https://leetcode.com/graphql"
	defaultUserAgent       = "LeetcodeCLI/1.0"
	defaultTimeout         = 10 * time.Second
)

type Client struct {
	endpoint   string
	httpClient *http.Client
	userAgent  string
}

type Option func(*Client)

func NewClient(options ...Option) *Client {
	client := &Client{
		endpoint:   defaultGraphQLEndpoint,
		httpClient: &http.Client{Timeout: defaultTimeout},
		userAgent:  defaultUserAgent,
	}

	for _, option := range options {
		option(client)
	}
	if client.httpClient == nil {
		client.httpClient = &http.Client{Timeout: defaultTimeout}
	}
	if client.endpoint == "" {
		client.endpoint = defaultGraphQLEndpoint
	}
	if client.userAgent == "" {
		client.userAgent = defaultUserAgent
	}

	return client
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func WithEndpoint(endpoint string) Option {
	return func(client *Client) {
		client.endpoint = endpoint
	}
}

func WithUserAgent(userAgent string) Option {
	return func(client *Client) {
		client.userAgent = userAgent
	}
}

func (client *Client) FetchProfileStats(ctx context.Context, username string) (ProfileStats, error) {
	user, err := client.fetchProfileData(ctx, username)
	if err != nil {
		return ProfileStats{}, err
	}

	return normalizeProfileStats(user)
}

func (client *Client) fetchProfileData(ctx context.Context, username string) (*matchedUser, error) {
	request, err := client.newProfileRequest(ctx, username)
	if err != nil {
		return nil, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, classify(ErrorKindEndpointFailure, err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusForbidden || response.StatusCode == http.StatusTooManyRequests {
		return nil, classify(ErrorKindRateLimited, fmt.Errorf("leetcode returned HTTP %d", response.StatusCode))
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return nil, classify(ErrorKindEndpointFailure, fmt.Errorf("leetcode returned HTTP %d", response.StatusCode))
	}

	var payload graphQLResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, classify(ErrorKindMalformedResponse, err)
	}
	if len(payload.Errors) > 0 {
		return nil, classify(ErrorKindUnavailable, fmt.Errorf("leetcode returned GraphQL errors"))
	}
	if payload.Data == nil {
		return nil, classify(ErrorKindMalformedResponse, fmt.Errorf("missing data object"))
	}
	if payload.Data.MatchedUser == nil {
		return nil, classify(ErrorKindNotFound, fmt.Errorf("matched user not found"))
	}

	return payload.Data.MatchedUser, nil
}

func (client *Client) newProfileRequest(ctx context.Context, username string) (*http.Request, error) {
	payload := graphQLRequest{
		Query: userProfileQuery,
		Variables: map[string]string{
			"username": username,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, classify(ErrorKindMalformedResponse, err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, client.endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, classify(ErrorKindEndpointFailure, err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", client.userAgent)

	return request, nil
}
