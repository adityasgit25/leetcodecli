package leetcode

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestNormalizeProfileStatsSuccess(t *testing.T) {
	stats := normalizeFixture(t, "profile_success.json")

	if stats.Summary.Username != "alice" {
		t.Fatalf("username = %q, want alice", stats.Summary.Username)
	}
	if stats.Summary.DisplayName != "Alice Example" {
		t.Fatalf("display name = %q, want Alice Example", stats.Summary.DisplayName)
	}
	if stats.Summary.Ranking != "123" {
		t.Fatalf("ranking = %q, want 123", stats.Summary.Ranking)
	}
	if stats.Summary.Reputation != "7" {
		t.Fatalf("reputation = %q, want 7", stats.Summary.Reputation)
	}
	if stats.Summary.ProfileURL != "https://leetcode.com/u/alice/" {
		t.Fatalf("profile url = %q, want https://leetcode.com/u/alice/", stats.Summary.ProfileURL)
	}
	if stats.TotalSolved != 42 {
		t.Fatalf("total solved = %d, want 42", stats.TotalSolved)
	}
	if len(stats.LanguageCounts) != 2 {
		t.Fatalf("language count length = %d, want 2", len(stats.LanguageCounts))
	}
	if stats.LanguageCounts[0] != (LanguageCount{Language: "Go", Solved: 10}) {
		t.Fatalf("first language = %#v, want Go/10", stats.LanguageCounts[0])
	}
}

func TestNormalizeProfileStatsNullableDisplayFields(t *testing.T) {
	user := decodeFixture(t, "profile_success.json").Data.MatchedUser
	user.Profile.RealName = nil
	user.Profile.Ranking = nil
	user.Profile.Reputation = nil

	stats, err := normalizeProfileStats(user)
	if err != nil {
		t.Fatalf("normalizeProfileStats returned error: %v", err)
	}

	if stats.Summary.DisplayName != NotAvailable {
		t.Fatalf("display name = %q, want %q", stats.Summary.DisplayName, NotAvailable)
	}
	if stats.Summary.Ranking != NotAvailable {
		t.Fatalf("ranking = %q, want %q", stats.Summary.Ranking, NotAvailable)
	}
	if stats.Summary.Reputation != NotAvailable {
		t.Fatalf("reputation = %q, want %q", stats.Summary.Reputation, NotAvailable)
	}
}

func TestNormalizeProfileStatsMandatoryValidation(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
		wantErr bool
	}{
		{name: "missing total", fixture: "profile_missing_total.json", wantErr: true},
		{name: "missing languages", fixture: "profile_missing_languages.json", wantErr: true},
		{name: "empty languages", fixture: "profile_empty_languages.json", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := decodeFixture(t, tt.fixture)
			stats, err := normalizeProfileStats(response.Data.MatchedUser)

			if tt.wantErr {
				if !IsErrorKind(err, ErrorKindMissingStats) {
					t.Fatalf("error = %v, want missing stats", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("normalizeProfileStats returned error: %v", err)
			}
			if stats.LanguageCounts == nil {
				t.Fatal("LanguageCounts = nil, want present empty slice")
			}
			if len(stats.LanguageCounts) != 0 {
				t.Fatalf("LanguageCounts length = %d, want 0", len(stats.LanguageCounts))
			}
		})
	}
}

func TestClientFetchProfileStatsReturnsNormalizedStats(t *testing.T) {
	client := NewClient(WithHTTPClient(&http.Client{
		Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			return jsonResponse(http.StatusOK, string(readFixture(t, "profile_success.json")))
		}),
	}))

	stats, err := client.FetchProfileStats(context.Background(), "alice")

	if err != nil {
		t.Fatalf("FetchProfileStats returned error: %v", err)
	}
	if stats.Summary.Username != "alice" {
		t.Fatalf("username = %q, want alice", stats.Summary.Username)
	}
	if stats.TotalSolved != 42 {
		t.Fatalf("total solved = %d, want 42", stats.TotalSolved)
	}
}

func TestClientFetchProfileStatsFixtureFailures(t *testing.T) {
	tests := []struct {
		name     string
		fixture  string
		wantKind ErrorKind
	}{
		{name: "profile not found", fixture: "profile_not_found.json", wantKind: ErrorKindNotFound},
		{name: "graphql error", fixture: "graphql_error.json", wantKind: ErrorKindUnavailable},
		{name: "missing total", fixture: "profile_missing_total.json", wantKind: ErrorKindMissingStats},
		{name: "missing languages", fixture: "profile_missing_languages.json", wantKind: ErrorKindMissingStats},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(WithHTTPClient(&http.Client{
				Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
					return jsonResponse(http.StatusOK, string(readFixture(t, tt.fixture)))
				}),
			}))

			_, err := client.FetchProfileStats(context.Background(), "alice")

			if !IsErrorKind(err, tt.wantKind) {
				t.Fatalf("error = %v, want kind %s", err, tt.wantKind)
			}
		})
	}
}

func normalizeFixture(t *testing.T, name string) ProfileStats {
	t.Helper()

	response := decodeFixture(t, name)
	stats, err := normalizeProfileStats(response.Data.MatchedUser)
	if err != nil {
		t.Fatalf("normalizeProfileStats(%s) returned error: %v", name, err)
	}

	return stats
}

func decodeFixture(t *testing.T, name string) graphQLResponse {
	t.Helper()

	var response graphQLResponse
	if err := json.Unmarshal(readFixture(t, name), &response); err != nil {
		t.Fatalf("decode fixture %s: %v", name, err)
	}

	return response
}

func readFixture(t *testing.T, name string) []byte {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("read fixture %s: %v", name, err)
	}

	return data
}
