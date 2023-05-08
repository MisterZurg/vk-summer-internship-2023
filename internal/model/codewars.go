package model

type CodeWarsUser struct {
	Username            string      // string	Username of the user.
	Name                string      //	string	Name of the user.
	Honor               int         // number	Total honor points earned by the user.
	Clan                string      //	Name of the clan.
	LeaderboardPosition int         //number	The user's position on the overall leaderboard.
	skills              []string    // Array of skills entered by the user.
	ranks               interface{} // Ranks object with overall and language ranks.
	codeChallenges      interface{} //	Object with fields totalAuthored and totalCompleted for the number of authored and completed kata respectively.
}

type CodeWarsProblem struct {
	ID                 string      // ID of the kata.
	Name               string      // Name of the kata.
	Slug               string      // Slug of the kata.
	Url                string      // URL of the kata.
	Category           string      // Category of the kata.
	Description        string      // Description of the kata in Markdown.
	Tags               []string    // Array of tags associated with the kata.
	languages          []string    //	Array of language names the kata is available in.
	rank               interface{} // object?	Object describing the rank of the kata if approved.
	createdBy          interface{} // object	The author of the kata.
	publishedAt        interface{} // string	Date and time when the kata was first published.
	approvedBy         interface{} // object?	The approver of the kata.
	approvedAt         string      // Date and time when the kata was approved.
	totalCompleted     int         // number	Total number of completions.
	totalAttempts      int         // number	Total number of attempts.
	totalStars         int         // number	The number of bookmarks.
	voteScore          int         // number	The sum of all votes casted.
	contributorsWanted bool        // boolean	Whether to allow contributions.
	unresolved         interface{} // object	Object with fields issues and suggestions for the number of unresolved issues and suggestions respectively.
}
