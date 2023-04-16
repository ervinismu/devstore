package schema

type SignInReq struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,alphanum" json:"password"`
}

type SignInResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
