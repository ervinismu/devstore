package schema

type LoginReq struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,alphanum" json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LogoutReq struct {
	UserID int `json:"user_id"`
}

type RefreshTokenReq struct {
	UserID       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResp struct {
	AccessToken string `json:"access_token"`
}
