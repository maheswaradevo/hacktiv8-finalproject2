package dto

type UserSignInResponse struct {
	AccessToken string `json:"access_token"`
}

func NewUserSignInResponse(ac string) *UserSignInResponse {
	return &UserSignInResponse{
		AccessToken: ac,
	}
}
