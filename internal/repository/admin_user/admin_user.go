package admin_user

import (
	"context"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"learning-gin/internal/entities"
	"learning-gin/internal/repository"
)

type AdminUser struct {
	db *sqlx.DB
}

func InitAdminUserRepo(db *sqlx.DB) repository.AdminUser {
	return AdminUser{
		db: db,
	}
}

func (usr AdminUser) Create(ctx context.Context, adminUserCreate entities.AdminUserCreate) (int, error) {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	var adminUserID int

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(adminUserCreate.Password), 10)

	row := transaction.QueryRowContext(
		ctx, `INSERT INTO admin_users (login, hashed_password) VALUES ($1, $2) RETURNING id`,
		adminUserCreate.Login, hashedPassword)

	err = row.Scan(&adminUserID)
	if err != nil {
		return 0, err
	}

	if err = transaction.Commit(); err != nil {
		return 0, err
	}

	return adminUserID, nil
}

func (usr AdminUser) Get(ctx context.Context, adminUserID int) (entities.AdminUser, error) {
	var adminUser entities.AdminUser

	err := usr.db.QueryRowContext(ctx, `SELECT id, login, hashed_password FROM admin_users WHERE id=$1`,
		adminUserID).Scan(&adminUser.ID, &adminUser.Login, &adminUser.Password)
	if err != nil {
		return entities.AdminUser{}, nil
	}

	return adminUser, nil
}

func (usr AdminUser) GetPassword(ctx context.Context, login string) (int, string, error) {
	var (
		adminUserID    int
		hashedPassword string
	)

	err := usr.db.QueryRowContext(ctx, `SELECT id, hashed_password FROM admin_users WHERE id=$1`,
		login).Scan(&adminUserID, &hashedPassword)
	if err != nil {
		return 0, "", err
	}

	return adminUserID, hashedPassword, nil
}

func (usr AdminUser) Delete(ctx context.Context, adminUserID int) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := transaction.ExecContext(ctx, `DELETE FROM admin_users WHERE id=$1`, adminUserID)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// TODO добавить ошибку count mismatch
	if count != 1 {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}
