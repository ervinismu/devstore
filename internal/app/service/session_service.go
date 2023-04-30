package service

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetByEmail(email string) (model.User, error)
	GetByID(userID int) (model.User, error)
}

type AuthRepository interface {
	Create(auth model.Auth) error
	DeleteAllByUserID(userID int) error
	Find(userID int, refreshToken string) (model.Auth, error)
}

type TokenGenerator interface {
	GenerateAccessToken(userID int) (string, time.Time, error)
	GenerateRefreshToken(userID int) (string, time.Time, error)
}

type SessionService struct {
	userRepo   UserRepository
	authRepo   AuthRepository
	tokenMaker TokenGenerator
}

func NewSessionService(
	userRepo UserRepository,
	authRepo AuthRepository,
	tokenMaker TokenGenerator,
) *SessionService {
	return &SessionService{
		userRepo:   userRepo,
		authRepo:   authRepo,
		tokenMaker: tokenMaker,
	}
}

func (svc *SessionService) Login(req *schema.LoginReq) (schema.LoginResp, error) {
	var resp schema.LoginResp

	// find existing user by userID
	existingUser, _ := svc.userRepo.GetByEmail(req.Email)
	if existingUser.ID <= 0 {
		return resp, errors.New(reason.FailedLogin)
	}

	// verify password
	isVerified := svc.verifyPassword(existingUser.HashedPassword, req.Password)
	if !isVerified {
		return resp, errors.New(reason.FailedLogin)
	}

	// generate access token
	accessToken, _, err := svc.tokenMaker.GenerateAccessToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("access token creation : %w", err))
		return resp, errors.New(reason.FailedLogin)
	}

	// generate refresh token
	refreshToken, expiredAt, err := svc.tokenMaker.GenerateRefreshToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("refresh token creation : %w", err))
		return resp, errors.New(reason.FailedLogin)
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	// save refresh token
	authPayload := model.Auth{
		UserID:    existingUser.ID,
		Token:     refreshToken,
		AuthType:  "refresh_token",
		ExpiredAt: expiredAt,
	}
	err = svc.authRepo.Create(authPayload)
	if err != nil {
		log.Error(fmt.Errorf("refresh token saving : %w", err))
		return resp, errors.New(reason.FailedLogin)
	}

	return resp, nil
}

func (svc *SessionService) Logout(UserID int) error {
	err := svc.authRepo.DeleteAllByUserID(UserID)
	if err != nil {
		log.Error(fmt.Errorf("refresh token saving : %w", err))
		return errors.New(reason.FailedLogout)
	}
	return nil
}

func (svc *SessionService) Refresh(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error) {
	var resp schema.RefreshTokenResp

	existingUser, _ := svc.userRepo.GetByID(req.UserID)
	if existingUser.ID <= 0 {
		return resp, errors.New(reason.FailedRefreshToken)
	}

	auth, err := svc.authRepo.Find(existingUser.ID, req.RefreshToken)
	if err != nil || auth.ID < 0 {
		log.Error(fmt.Errorf("error SessionService - refresh : %w", err))
		return resp, errors.New(reason.FailedRefreshToken)
	}

	accessToken, _, _ := svc.tokenMaker.GenerateAccessToken(existingUser.ID)

	resp.AccessToken = accessToken
	return resp, nil
}

func (svc *SessionService) verifyPassword(hashPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd))
	return err == nil
}
