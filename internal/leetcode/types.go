package leetcode

const NotAvailable = "N/A"

type ProfileStats struct {
	Summary        ProfileSummary
	TotalSolved    int
	LanguageCounts []LanguageCount
}

type ProfileSummary struct {
	Username    string
	DisplayName string
	Ranking     string
	Reputation  string
	ProfileURL  string
}

type LanguageCount struct {
	Language string
	Solved   int
}

type graphQLResponse struct {
	Data   *profileData   `json:"data"`
	Errors []graphQLError `json:"errors,omitempty"`
}

type graphQLError struct {
	Message string `json:"message"`
}

type profileData struct {
	MatchedUser *matchedUser `json:"matchedUser"`
}

type matchedUser struct {
	Username             string                  `json:"username"`
	Profile              *profile                `json:"profile"`
	SubmitStatsGlobal    *submitStatsGlobal      `json:"submitStatsGlobal"`
	LanguageProblemCount *[]languageProblemCount `json:"languageProblemCount"`
}

type profile struct {
	RealName   *string `json:"realName"`
	Ranking    *int    `json:"ranking"`
	Reputation *int    `json:"reputation"`
}

type submitStatsGlobal struct {
	ACSubmissionNum []submissionCount `json:"acSubmissionNum"`
}

type submissionCount struct {
	Difficulty  string `json:"difficulty"`
	Count       int    `json:"count"`
	Submissions int    `json:"submissions"`
}

type languageProblemCount struct {
	LanguageName   string `json:"languageName"`
	ProblemsSolved int    `json:"problemsSolved"`
}
