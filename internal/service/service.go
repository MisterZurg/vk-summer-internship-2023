package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/machinebox/graphql"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"vk-summer-internship-2023/internal/model"
)

type Service struct {
	logger *zerolog.Logger
	client HTTPClient
}

type Params struct {
	Cmd string
	Arg string
}

func (s *Service) Do(params Params, t time.Time) (string, error) {
	switch params.Cmd {
	case "cwUser":
		userInfo, err := s.CodeWarsUser(params.Arg)
		if err != nil {
			s.logger.Fatal().Err(err).Msg("Configuration error")
		}
		return fmt.Sprintf("%s \n %s", userInfo.Username, userInfo.Honor), nil
	case "cwProblem":
		problemInfo, err := s.CodeWarsCodeChallenge(params.Arg)
		if err != nil {
			fmt.Println(err)
		}
		return problemInfo.Name, nil

	case "lcUser":
		userInfo, err := s.LeetCodeUser(params.Arg)
		if err != nil {
			fmt.Println(err)
		}
		// statText := "User " + username + " not found."
		if userInfo.MatchedUser.Username != "" {
			return "User " + params.Arg + " not found.", nil
		}

		msg := "Stats for user " + userInfo.MatchedUser.Username
		for _, val := range userInfo.MatchedUser.SubmitStats.AcSubmissionNum {
			msg += "\n" + val.Difficulty + ": " + strconv.Itoa(val.Count) + " tasks"
		}
		return msg, nil
	case "cfUser":
		userInfo, err := s.CodeForcesUser(params.Arg)
		if err != nil {
			fmt.Println(err)
		}
		msg := "Stats for user " + userInfo.Result[0].Handle
		msg += "\n Contributed " + strconv.Itoa(userInfo.Result[0].Contribution) + " times"
		msg += "\n Has " + strconv.Itoa(userInfo.Result[0].FriendOfCount) + " friends"
		return msg, nil
	}
	return "", nil
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(logger *zerolog.Logger) *Service {
	return &Service{
		logger: logger,
		client: &http.Client{
			Timeout: time.Duration(time.Minute),
		},
	}
}

const (
	CodeWarsUserFormat          = "https://www.codewars.com/api/v1/users/%s"
	CodeWarsCodeChallengeFormat = "https://www.codewars.com/api/v1/code-challenges/%s"
	LeetCodeEndpoint            = "https://leetcode.com/graphql"
	CodeForcesUserFormat        = "https://codeforces.com/api/user.info?locale=en&handles=%s"
)

func (s *Service) CodeWarsUser(username string) (*model.CodeWarsUser, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(CodeWarsUserFormat, username), nil)
	if err != nil {
		s.logger.Error().Err(err).Msg("Bad Request")
	}

	res, err := s.client.Do(req)
	if err != nil {
		s.logger.Error().Err(err).Msg("Bad Response")
	}
	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		s.logger.Error().Err(err).Msg("Cannot Read Response Body")
	}

	if res.StatusCode == http.StatusOK {
		u := model.CodeWarsUser{}
		err = json.Unmarshal(raw, &u)
		if err != nil {
			s.logger.Error().Err(err).Msg("Cannot Unmarshal data")
		}
		return &u, nil
	}
	return &model.CodeWarsUser{}, nil
}

func (s *Service) CodeWarsCodeChallenge(title string) (*model.CodeWarsProblem, error) {
	// preprocess from title to slug
	slug := strings.ReplaceAll(title, " ", "-")

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(CodeWarsCodeChallengeFormat, slug), nil)
	if err != nil {
		s.logger.Error().Err(err).Msg("Bad Request")
	}

	res, err := s.client.Do(req)
	if err != nil {
		s.logger.Error().Err(err).Msg("Cannot Read Response Body")
	}
	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if res.StatusCode == http.StatusOK {
		pb := model.CodeWarsProblem{}
		err = json.Unmarshal(raw, &pb)
		if err != nil {
			s.logger.Error().Err(err).Msg("Cannot Unmarshal data")
		}
		return &pb, nil
	}

	return &model.CodeWarsProblem{}, nil
}

var requestUser = `query getUserProfile($username: String!) {
	matchedUser(username: $username) {
		username
		submitStats: submitStatsGlobal {
			acSubmissionNum {
				difficulty
				count
				submissions
			}
		}
	}
}`

func (s *Service) LeetCodeUser(username string) (*model.LeetCodeUserData, error) {
	ctx := context.Background()
	client := graphql.NewClient(LeetCodeEndpoint)
	req := graphql.NewRequest(requestUser)

	req.Var("username", username)

	u := model.LeetCodeUserData{}

	err := client.Run(ctx, req, &u)
	if err != nil {
		s.logger.Fatal().Err(err).Msg("Cannot Process Grapql resp")
	}

	return &u, nil
}

func (s *Service) CodeForcesUser(handle string) (*model.CodeForcesUser, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(CodeForcesUserFormat, handle), nil)
	if err != nil {
		s.logger.Error().Err(err).Msg("Bad Request")
	}

	res, err := s.client.Do(req)
	if err != nil {
		s.logger.Error().Err(err).Msg("Bad Response")
	}
	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		s.logger.Error().Err(err).Msg("Cannot Read Response Body")
	}

	if res.StatusCode == http.StatusOK {
		u := model.CodeForcesUser{}
		err = json.Unmarshal(raw, &u)
		if err != nil {
			s.logger.Error().Err(err).Msg("Cannot Unmarshal Data")
		}
		return &u, nil
	} else {
		return &model.CodeForcesUser{}, err
	}
}
