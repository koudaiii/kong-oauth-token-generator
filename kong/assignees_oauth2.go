package kong

import (
	_ "fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type AssigneesOAuth2List struct {
	AssigneesOAuth2 []AssigneesOAuth2 `json:"data,omitempty"`
	Total           int               `json:"total,omitempty"`
}

type AssigneesOAuth2 struct {
	AccessToken  string `json:"access_token,omitempty"`
	ID           string `json:"id,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	CredentialID string `json:"credential_id,omitempty"`
	CreatedAt    int    `json:"created_at,omitempty"`
	ExpiresIn    int    `json:"expires_in"`
}

// Services

// ConfigService provides methods for creating and reading issues.
type AssigneesOAuth2Service struct {
	sling *sling.Sling
}

// NewCOAuth2Service returns a new OAuth2Service.
func NewAssigneesOAuth2Service(httpClient *http.Client, config *config.KongConfiguration) *AssigneesOAuth2Service {
	return &AssigneesOAuth2Service{
		sling: sling.New().Client(httpClient).Base(config.KongAdminURL + "oauth2_tokens/"),
	}
}

func (s *AssigneesOAuth2Service) Create(params *AssigneesOAuth2) (*AssigneesOAuth2, *http.Response, error) {
	assigneesOAuth2 := new(AssigneesOAuth2)
	resp, err := s.sling.New().Post("http://localhost:8001/oauth2_tokens").BodyJSON(params).ReceiveSuccess(assigneesOAuth2)
	return assigneesOAuth2, resp, err
}

func (s *AssigneesOAuth2Service) Get(oauth2ID string) (*AssigneesOAuth2, *http.Response, error) {
	assigneesOAuth2 := new(AssigneesOAuth2)
	resp, err := s.sling.New().Path(oauth2ID).ReceiveSuccess(assigneesOAuth2)
	return assigneesOAuth2, resp, err
}

func (s *AssigneesOAuth2Service) List() (*AssigneesOAuth2List, *http.Response, error) {
	assigneesOAuth2List := new(AssigneesOAuth2List)
	resp, err := s.sling.New().ReceiveSuccess(assigneesOAuth2List)
	return assigneesOAuth2List, resp, err
}

func (s *AssigneesOAuth2Service) Update(params *AssigneesOAuth2) (*AssigneesOAuth2, *http.Response, error) {
	assigneesOAuth2 := new(AssigneesOAuth2)
	resp, err := s.sling.New().Patch("http://localhost:8001/oauth2_tokens/" + params.ID).BodyJSON(params).ReceiveSuccess(assigneesOAuth2)
	return assigneesOAuth2, resp, err
}

func (s *AssigneesOAuth2Service) Delete(oauth2ID string) (string, *http.Response, error) {
	var message string
	resp, err := s.sling.New().Delete("http://localhost:8001/oauth2_tokens/" + oauth2ID).ReceiveSuccess(message)
	return message, resp, err
}
