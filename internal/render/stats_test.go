package render

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/adityasgit25/leetcodecli/internal/leetcode"
	"github.com/mattn/go-runewidth"
)

func TestRenderStatsGoldenOutput(t *testing.T) {
	tests := []struct {
		name   string
		stats  leetcode.ProfileStats
		golden string
	}{
		{
			name:   "success",
			stats:  sampleStats(),
			golden: "stats_success_80.golden",
		},
		{
			name:   "long values",
			stats:  longValueStats(),
			golden: "stats_long_values_80.golden",
		},
		{
			name:   "empty languages",
			stats:  emptyLanguageStats(),
			golden: "stats_empty_languages_80.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := RenderStats(tt.stats, FallbackWidth)
			if err != nil {
				t.Fatalf("RenderStats returned error: %v", err)
			}

			assertGolden(t, tt.golden, output)
			for _, required := range []string{"Profile Summary", "Total Solved Count", "Language Breakdown"} {
				if !strings.Contains(output, required) {
					t.Fatalf("output missing %q:\n%s", required, output)
				}
			}
			for _, forbidden := range []string{"{\"", "[\"", ",", "CSV", "JSON"} {
				if strings.Contains(output, forbidden) {
					t.Fatalf("output contains machine-readable marker %q:\n%s", forbidden, output)
				}
			}
		})
	}
}

func TestRenderStatsWithWidthDetectorFallsBackTo80Columns(t *testing.T) {
	output, err := RenderStatsWithWidthDetector(sampleStats(), func() (int, error) {
		return 0, os.ErrInvalid
	})
	if err != nil {
		t.Fatalf("RenderStatsWithWidthDetector returned error: %v", err)
	}

	if !strings.Contains(output, "Profile Summary") {
		t.Fatalf("output missing profile summary:\n%s", output)
	}
}

func TestDetectWriterWidthFallsBackForNonFileWriter(t *testing.T) {
	if got := ResolveWidth(DetectWriterWidth(&bytes.Buffer{})); got != FallbackWidth {
		t.Fatalf("ResolveWidth(non-file writer) = %d, want %d", got, FallbackWidth)
	}
}

func TestResolveWidth(t *testing.T) {
	if got := ResolveWidth(func() (int, error) { return 120, nil }); got != 120 {
		t.Fatalf("ResolveWidth success = %d, want 120", got)
	}
	if got := ResolveWidth(func() (int, error) { return 0, os.ErrInvalid }); got != FallbackWidth {
		t.Fatalf("ResolveWidth failure = %d, want %d", got, FallbackWidth)
	}
	if got := ResolveWidth(func() (int, error) { return -1, nil }); got != FallbackWidth {
		t.Fatalf("ResolveWidth invalid = %d, want %d", got, FallbackWidth)
	}
}

func TestFitTruncatesUnicodeWithoutInvalidUTF8(t *testing.T) {
	got := fit("Go语言🙂with-extra-text", 8)

	if !utf8.ValidString(got) {
		t.Fatalf("fit returned invalid UTF-8: %q", got)
	}
	if runewidth.StringWidth(got) > 8 {
		t.Fatalf("fit width = %d, want <= 8 for %q", runewidth.StringWidth(got), got)
	}
	if !strings.HasSuffix(got, "...") {
		t.Fatalf("fit = %q, want ellipsis suffix", got)
	}
}

func sampleStats() leetcode.ProfileStats {
	return leetcode.ProfileStats{
		Summary: leetcode.ProfileSummary{
			Username:    "alice",
			DisplayName: "Alice Example",
			Ranking:     "123",
			Reputation:  "7",
			ProfileURL:  "https://leetcode.com/u/alice/",
		},
		TotalSolved: 42,
		LanguageCounts: []leetcode.LanguageCount{
			{Language: "Go", Solved: 10},
			{Language: "Python3", Solved: 32},
		},
	}
}

func longValueStats() leetcode.ProfileStats {
	return leetcode.ProfileStats{
		Summary: leetcode.ProfileSummary{
			Username:    "a-very-long-leetcode-username-that-still-needs-a-readable-table",
			DisplayName: leetcode.NotAvailable,
			Ranking:     leetcode.NotAvailable,
			Reputation:  leetcode.NotAvailable,
			ProfileURL:  "https://leetcode.com/u/a-very-long-leetcode-username-that-still-needs-a-readable-table/",
		},
		TotalSolved: 1234,
		LanguageCounts: []leetcode.LanguageCount{
			{Language: "TypeScriptWithExtraLongRuntimeName", Solved: 333},
			{Language: "C++", Solved: 222},
		},
	}
}

func emptyLanguageStats() leetcode.ProfileStats {
	stats := sampleStats()
	stats.LanguageCounts = []leetcode.LanguageCount{}
	return stats
}

func assertGolden(t *testing.T, name string, output string) {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("read golden %s: %v", name, err)
	}
	expected := normalizeLineEndings(string(data))
	actual := normalizeLineEndings(output)
	if actual != expected {
		t.Fatalf("output mismatch for %s\nwant:\n%s\ngot:\n%s", name, expected, actual)
	}
}

func normalizeLineEndings(value string) string {
	return strings.ReplaceAll(value, "\r\n", "\n")
}
