package repository

import (
	"fmt"

	"github.com/ervinismu/devstore/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type AuthRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(DB *sqlx.DB) *AuthRepository {
	return &AuthRepository{DB: DB}
}

func (ar *AuthRepository) Create(auth model.Auth) error {
	var (
		sqlStatement = `
			INSERT INTO auths (token, auth_type, expired_at, user_id)
			VALUES ($1, $2, $3, $4)
		`
	)

	_, err := ar.DB.Exec(sqlStatement, auth.Token, auth.AuthType, auth.ExpiredAt, auth.UserID)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepository - Create : %w", err))
		return err
	}

	return nil
}

func (ar *AuthRepository) DeleteAllByUserID(userID int) error {
	var (
		sqlStatement = `
			DELETE FROM auths
			WHERE user_id = $1
		`
	)

	_, err := ar.DB.Exec(sqlStatement, userID)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepository - DeleteAllByUserID : %w", err))
		return err
	}

	return nil
}
