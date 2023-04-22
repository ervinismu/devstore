package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/ervinismu/devstore/internal/app/schema"
	"github.com/ervinismu/devstore/internal/pkg/reason"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type TokenGenerator interface {
	GenerateAccessToken(userID int) (string, time.Time, error)
	GenerateRefreshToken(userID int) (string, time.Time, error)
}

type AuthRepository interface {
	Find(userID int, refreshToken string) (model.Auth, error)
	Create(auth model.Auth) error
	DeleteAllByUserID(userID int) error
}

type SessionService struct {
	userRepo   UserRepository
	authRepo   AuthRepository
	tokenMaker TokenGenerator
}

func NewSessionService(userRepo UserRepository, authRepo AuthRepository, tokenMaker TokenGenerator) *SessionService {
	return &SessionService{
		userRepo:   userRepo,
		authRepo:   authRepo,
		tokenMaker: tokenMaker,
	}
}

func (svc *SessionService) Login(req *schema.LoginReq) (schema.LoginResp, error) {
	var resp schema.LoginResp

	existingUser, _ := svc.userRepo.GetByEmail(req.Email)
	if existingUser.ID <= 0 {
		return resp, errors.New(reason.UserNotFound)
	}

	isVerified := svc.verifyPassword(existingUser.HashedPassword, req.Password)
	if !isVerified {
		return resp, errors.New(reason.FailedLogin)
	}

	accessToken, _, err := svc.tokenMaker.GenerateAccessToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("accesstoken creation : %w", err))
		return resp, errors.New(reason.FailedLogin)
	}
	refreshToken, expiredAt, err := svc.tokenMaker.GenerateRefreshToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("refreshToken creation : %w", err))
		return resp, errors.New(reason.FailedLogin)
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	err = svc.saveRefreshToken(model.Auth{
		UserID:    existingUser.ID,
		Token:     refreshToken,
		AuthType:  "refresh_token",
		ExpiredAt: expiredAt,
	})
	if err != nil {
		log.Error(fmt.Errorf("error SessionService - Login : %w", err))
		return resp, errors.New(reason.FailedLogin)
	}

	return resp, nil
}

func (svc *SessionService) Refresh(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error) {
	var resp schema.RefreshTokenResp

	existingUser, _ := svc.userRepo.GetByID(req.UserID)
	if existingUser.ID <= 0 {
		return resp, errors.New(reason.UserNotFound)
	}

	auth, err := svc.authRepo.Find(existingUser.ID, req.RefreshToken)
	if err != nil || auth.ID < 0 {
		log.Error(fmt.Errorf("error SessionService - refresh : %w", err))
		return resp, errors.New(reason.InvalidRefreshToken)
	}

	accessToken, _, _ := svc.tokenMaker.GenerateAccessToken(existingUser.ID)

	resp.AccessToken = accessToken
	return resp, nil
}

func (svc *SessionService) Logout(req *schema.LogoutReq) error {
	err := svc.authRepo.DeleteAllByUserID(req.UserID)
	if err != nil {
		log.Error(fmt.Errorf("delete all user session : %w", err))
		return err
	}
	return nil
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
