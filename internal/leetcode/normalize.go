package leetcode

import (
	"fmt"
	"net/url"
	"strconv"
)

func normalizeProfileStats(user *matchedUser) (ProfileStats, error) {
	if user == nil || user.Profile == nil || user.Username == "" {
		return ProfileStats{}, classify(ErrorKindMissingStats, fmt.Errorf("missing profile summary"))
	}
	if user.SubmitStatsGlobal == nil {
		return ProfileStats{}, classify(ErrorKindMissingStats, fmt.Errorf("missing submission stats"))
	}
	if user.LanguageProblemCount == nil {
		return ProfileStats{}, classify(ErrorKindMissingStats, fmt.Errorf("missing language stats"))
	}

	totalSolved, ok := totalSolvedCount(user.SubmitStatsGlobal.ACSubmissionNum)
	if !ok {
		return ProfileStats{}, classify(ErrorKindMissingStats, fmt.Errorf("missing total solved count"))
	}

	return ProfileStats{
		Summary: ProfileSummary{
			Username:    user.Username,
			DisplayName: optionalString(user.Profile.RealName),
			Ranking:     optionalInt(user.Profile.Ranking),
			Reputation:  optionalInt(user.Profile.Reputation),
			ProfileURL:  profileURL(user.Username),
		},
		TotalSolved:    totalSolved,
		LanguageCounts: normalizeLanguageCounts(*user.LanguageProblemCount),
	}, nil
}

func totalSolvedCount(counts []submissionCount) (int, bool) {
	for _, count := range counts {
		if count.Difficulty == "All" {
			return count.Count, true
		}
	}
	return 0, false
}

func normalizeLanguageCounts(counts []languageProblemCount) []LanguageCount {
	languages := make([]LanguageCount, 0, len(counts))
	for _, count := range counts {
		languages = append(languages, LanguageCount{
			Language: count.LanguageName,
			Solved:   count.ProblemsSolved,
		})
	}
	return languages
}

func optionalString(value *string) string {
	if value == nil || *value == "" {
		return NotAvailable
	}
	return *value
}

func optionalInt(value *int) string {
	if value == nil {
		return NotAvailable
	}
	return strconv.Itoa(*value)
}

func profileURL(username string) string {
	return "https://leetcode.com/u/" + url.PathEscape(username) + "/"
}
