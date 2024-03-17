package user

import (
	"context"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"learning-gin/internal/entitites"
	"learning-gin/internal/repository"
)

type User struct {
	db *sqlx.DB
}

func InitUserRepo(db *sqlx.DB) repository.User {
	return User{
		db: db,
	}
}

func (usr User) Create(ctx context.Context, userCreate entitites.UserCreate) (int, error) {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	var userID int

	// 10 - сложность хэш-функции
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userCreate.Password), 10)

	row := transaction.QueryRowContext(
		ctx, `INSERT INTO users (login, name, hashed_password) VALUES ($1, $2, $3) RETURNING id;`,
		userCreate.Login, userCreate.Name, hashedPassword)

	err = row.Scan(&userID)
	if err != nil {
		return 0, err
	}

	if err = transaction.Commit(); err != nil {
		return 0, err
	}

	return userID, nil
}

func (usr User) Get(ctx context.Context, userID int) (entitites.User, error) {
	var user entitites.User

	err := usr.db.QueryRowContext(ctx, `SELECT id, login, name FROM users WHERE users.id = $1;`,
		userID).Scan(&user.ID, &user.Login, &user.Name)
	if err != nil {
		return entitites.User{}, err
	}

	return user, nil
}

func (usr User) GetPassword(ctx context.Context, login string) (int, string, error) {
	var userID int
	var hashedPassword string

	err := usr.db.QueryRowContext(ctx, `SELECT id, hashed_password FROM users WHERE users.login = $1;`,
		login).Scan(&userID, &hashedPassword)
	if err != nil {
		return 0, "", err
	}

	return userID, hashedPassword, nil
}

func (usr User) Delete(ctx context.Context, userID int) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Что за метод ExecContext? Не нашел в документации нихуя
	result, err := transaction.ExecContext(ctx, `DELETE FROM users WHERE users.id = $1;`, userID)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count != 1 {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}
