package service

import (
	"errors"
	"fmt"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/ervinismu/devstore/internal/app/repository"
	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	"github.com/ervinismu/devstore/internal/pkg/token"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type SessionService struct {
	userRepo repository.IUserRepository
	authRepo repository.IAuthRepository
}

func NewSessionService(userRepo repository.IUserRepository, authRepo repository.IAuthRepository) *SessionService {
	return &SessionService{userRepo: userRepo, authRepo: authRepo}
}

func (svc *SessionService) SignIn(req *schema.SignInReq) (schema.SignInResp, error) {
	var resp schema.SignInResp

	existingUser, _ := svc.userRepo.GetByEmail(req.Email)
	if existingUser.ID <= 0 {
		return resp, errors.New(reason.UserNotFound)
	}

	isVerified := svc.verifyPassword(existingUser.HashedPassword, req.Password)
	if !isVerified {
		return resp, errors.New(reason.FailedLogin)
	}

	accessToken, _, _ := token.GenerateAccessToken(existingUser.ID)
	refreshToken, expiredAt, _ := token.GenerateRefreshToken(existingUser.ID)

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	err := svc.saveRefreshToken(model.Auth{
		UserID:    existingUser.ID,
		Token:     refreshToken,
		AuthType:  "refresh_token",
		ExpiredAt: expiredAt,
	})
	if err != nil {
		log.Error(fmt.Errorf("error SessionService - SignIn : %w", err))
		return resp, errors.New(reason.FailedLogin)
	}

	return resp, nil
}

func (svc *SessionService) verifyPassword(hashPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd))
	return err == nil
}

func (svc *SessionService) saveRefreshToken(auth model.Auth) error {
	err := svc.authRepo.DeleteAllByUserID(auth.UserID)
	if err != nil {
		return fmt.Errorf("delete all user session : %w", err)
	}
	err = svc.authRepo.Create(auth)
	if err != nil {
		return fmt.Errorf("save auth : %w", err)
	}
	return nil
}
