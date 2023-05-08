package model

type CodeForcesUser struct {
	Status string               `json:"status"`
	Result []CodeForcesUserData `json:"result"`
}

type CodeForcesUserData struct {
	Handle                  string `json:"handle"`                  // String. Codeforces user handle.
	Email                   string `json:"email,omitempty"`         //
	VKID                    string `json:"vkId,omitempty"`          // String. User id for VK social network. Shown only if user allowed to share his contact info.
	OpenId                  string `json:"openId,omitempty"`        // String. Shown only if user allowed to share his contact info.
	FirstName               string `json:"firstName,omitempty"`     // String. Localized. Can be absent.
	lastName                string `json:"lastName,omitempty"`      // String. Localized. Can be absent.
	Country                 string `json:"country,omitempty"`       // String. Localized. Can be absent.
	City                    string `json:"city,omitempty"`          // String. Localized. Can be absent.
	Organization            string `json:"organization,omitempty"`  // String. Localized. Can be absent.
	Contribution            int    `json:"contribution"`            // Integer. User contribution.
	Rank                    string `json:"rank"`                    // String. Localized.
	Rating                  int    `json:"rating"`                  // Integer.
	MaxRank                 string `json:"maxRank"`                 // String. Localized.
	MaxRating               int    `json:"maxRating"`               // Integer.
	LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds"`   // Integer. Time, when user was last seen online, in unix format.
	RegistrationTimeSeconds int    `json:"registrationTimeSeconds"` // Integer. Time, when user was registered, in unix format.
	FriendOfCount           int    `json:"friendOfCount"`           // Integer. Amount of users who have this user in friends.
	Avatar                  string `json:"avatar"`                  // String. User's avatar URL.
	TitlePhoto              string `json:"titlePhoto"`              // String. User's title photo URL.
}
