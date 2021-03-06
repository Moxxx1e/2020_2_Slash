package repository

import (
	"context"
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_Slash/tools/logger"

	"github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	"github.com/go-park-mail-ru/2020_2_Slash/internal/user"
)

type UserPgRepository struct {
	dbConn *sql.DB
}

func NewUserPgRepository(conn *sql.DB) user.UserRepository {
	return &UserPgRepository{
		dbConn: conn,
	}
}

func (ur *UserPgRepository) Insert(user *models.User) error {
	tx, err := ur.dbConn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	row := tx.QueryRow(
		`INSERT INTO users(nickname, email, password, avatar, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		user.Nickname, user.Email, user.Password, user.Avatar, user.Role)

	err = row.Scan(&user.ID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			logger.Error(rollbackErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (ur *UserPgRepository) SelectByEmail(email string) (*models.User, error) {
	user := &models.User{}

	row := ur.dbConn.QueryRow(
		`SELECT id, nickname, email, password, avatar, role
		FROM users
		WHERE email=$1`, email)

	err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.Avatar, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserPgRepository) SelectByID(userID uint64) (*models.User, error) {
	user := &models.User{}
	row := ur.dbConn.QueryRow(
		`SELECT id, nickname, email, password, avatar, role
		FROM users
		WHERE id=$1`, userID)

	err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.Avatar, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserPgRepository) Update(user *models.User) error {
	tx, err := ur.dbConn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`UPDATE users
		SET nickname = $2, email = $3, password = $4, avatar = $5, role = $6
		WHERE id = $1;`,
		user.ID, user.Nickname, user.Email, user.Password, user.Avatar, user.Role)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			logger.Error(rollbackErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
