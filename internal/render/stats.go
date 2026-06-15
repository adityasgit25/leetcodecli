package render

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mattn/go-runewidth"

	"github.com/adityasgit25/leetcodecli/internal/leetcode"
)

const minValueWidth = 12

func RenderStatsWithWidthDetector(stats leetcode.ProfileStats, detector WidthDetector) (string, error) {
	return RenderStats(stats, ResolveWidth(detector))
}

func RenderStats(stats leetcode.ProfileStats, width int) (string, error) {
	if width <= 0 {
		width = FallbackWidth
	}

	valueWidth := summaryValueWidth(width)
	languageWidth := languageNameWidth(width)

	var output strings.Builder
	appendTable(&output, "Profile Summary", table.Row{"Field", "Value"}, []table.Row{
		{"Username", fit(stats.Summary.Username, valueWidth)},
		{"Display Name", fit(stats.Summary.DisplayName, valueWidth)},
		{"Ranking", fit(stats.Summary.Ranking, valueWidth)},
		{"Reputation", fit(stats.Summary.Reputation, valueWidth)},
		{"Profile URL", fit(stats.Summary.ProfileURL, valueWidth)},
	})
	output.WriteString("\n\n")

	appendTable(&output, "Total Solved Count", table.Row{"Metric", "Count"}, []table.Row{
		{"Total Solved", fmt.Sprintf("%d", stats.TotalSolved)},
	})
	output.WriteString("\n\n")

	languageRows := make([]table.Row, 0, len(stats.LanguageCounts))
	if len(stats.LanguageCounts) == 0 {
		languageRows = append(languageRows, table.Row{"No languages found", "0"})
	} else {
		for _, language := range stats.LanguageCounts {
			languageRows = append(languageRows, table.Row{
				fit(language.Language, languageWidth),
				fmt.Sprintf("%d", language.Solved),
			})
		}
	}
	appendTable(&output, "Language Breakdown", table.Row{"Language", "Solved"}, languageRows)
	output.WriteString("\n")

	return output.String(), nil
}

func appendTable(output *strings.Builder, title string, header table.Row, rows []table.Row) {
	output.WriteString(title)
	output.WriteString("\n")

	writer := table.NewWriter()
	writer.SetStyle(table.StyleDefault)
	writer.AppendHeader(header)
	for _, row := range rows {
		writer.AppendRow(row)
	}
	output.WriteString(writer.Render())
}

func summaryValueWidth(width int) int {
	valueWidth := width - 28
	if valueWidth < minValueWidth {
		return minValueWidth
	}
	return valueWidth
}

func languageNameWidth(width int) int {
	nameWidth := width - 18
	if nameWidth < minValueWidth {
		return minValueWidth
	}
	return nameWidth
}

func fit(value string, maxWidth int) string {
	if maxWidth <= 0 || runewidth.StringWidth(value) <= maxWidth {
		return value
	}
	return runewidth.Truncate(value, maxWidth, "...")
}
