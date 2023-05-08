package model

type LeetCodeUserData struct {
	MatchedUser UserStats `json:"matchedUser"`
}

type UserStats struct {
	Username    string      `json:"username"`
	SubmitStats SubmitStats `json:"submitStats"`
}

type SubmitStats struct {
	AcSubmissionNum []AcSubmissionNum `json:"acSubmissionNum"`
}

type AcSubmissionNum struct {
	Difficulty  string `json:"difficulty"`
	Count       int    `json:"count"`
	Submissions int    `json:"submissions"`
}
