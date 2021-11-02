package store

import (
	"github.com/jackc/pgx/v4"
	"go-articles/model"
	"go-articles/pkg/id"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (string, error) {
	uid := id.GenerateID("")
	if _, err := r.store.dbConn.Exec(
		r.store.dbCtx, "INSERT INTO users (id, email, encrypted_password) VALUES ($1, $2, $3)",
		uid,
		u.Email,
		u.EncryptedPassword,
	); err != nil {
		return "", err
	}
	return uid, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error){
	u := &model.User{}
	if err := r.store.dbConn.QueryRow(
		r.store.dbCtx, "SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return u,nil
}

func (r *UserRepository) FindById(id string) (*model.User, error){
	u := &model.User{}
	if err := r.store.dbConn.QueryRow(
		r.store.dbCtx, "SELECT id, email, encrypted_password FROM users WHERE id = $1",
		id).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return u,nil
}