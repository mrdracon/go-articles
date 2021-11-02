package store

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

var (
	ErrIncorrectEmailOrPassword = errors.New("incorrect email or password")
	ErrArticlesNotFound         = errors.New("articles not found")
	ErrUserNotFound         = errors.New("user not found")
)

type Store struct {
	DatabaseURL string
	dbConn *pgx.Conn
	dbCtx context.Context
	userRepository *UserRepository
	articlesRepository *ArticleRepository
}

func New(DSN string) *Store {
	return &Store{
		DatabaseURL: DSN,
	}
}

func (s *Store) Open() error {
	dbCtx := context.Background()

	dbConn, err := pgx.Connect(dbCtx, s.DatabaseURL)
	if err != nil {
		return err
	}

	if err := dbConn.Ping(dbCtx); err != nil {
		return err
	}

	s.dbCtx = dbCtx
	s.dbConn = dbConn
	return nil
}

func (s *Store) Close() {
	s.dbConn.Close(s.dbCtx)
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Articles() *ArticleRepository {
	if s.articlesRepository != nil {
		return s.articlesRepository
	}

	s.articlesRepository = &ArticleRepository{
		store: s,
	}

	return s.articlesRepository
}