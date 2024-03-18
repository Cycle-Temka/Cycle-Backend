package user

import (
	"context"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"learning-gin/internal/entities"
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

func (usr User) Create(ctx context.Context, userCreate entities.UserCreate) (int, error) {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	var userID int

	// 10 - сложность хэш-функции
	// TODO. Хэширование убрать в сервисный слой
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

func (usr User) Get(ctx context.Context, userID int) (entities.User, error) {
	var user entities.User

	err := usr.db.QueryRowContext(ctx, `SELECT id, login, name FROM users WHERE users.id = $1;`,
		userID).Scan(&user.ID, &user.Login, &user.Name)
	if err != nil {
		return entities.User{}, err
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

func (usr User) UpdatePassword(ctx context.Context, userID int, newPassword string) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := transaction.ExecContext(ctx, `UPDATE users SET hashed_password = $1 WHERE users.id = $2`, newPassword, userID)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count != 1 {
		// TODO. По сути дублирование прошлого возврата. Добавить обработку ошибки count mismatch
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
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
		// TODO. Дублирование прошлого возврата. Добавить обработку ошибки count mismatch
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}
