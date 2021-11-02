package store

import (
	"github.com/jackc/pgx/v4"
	"go-articles/model"
	"go-articles/pkg/id"
	"math/rand"
	"time"
)

type ArticleRepository struct {
	store *Store
}

func (r *ArticleRepository) Create(a *model.Article) (string, error) {
	rand.Seed(time.Now().UnixNano())

	uid := id.GenerateID("")
	if _, err := r.store.dbConn.Exec(
		r.store.dbCtx, "INSERT INTO articles (id, state, title, text) VALUES ($1, $2, $3, $4)",
		uid,
		a.State,
		a.Title,
		a.Text,
	); err != nil {
		return "", err
	}
	return uid, nil
}

func (r *ArticleRepository) ListArticles() ([]model.Article, error){

	articles := []model.Article{}
	rows, err := r.store.dbConn.Query(
		r.store.dbCtx, "SELECT id, state, title, text FROM articles")		;
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrArticlesNotFound
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var a model.Article
		err := rows.Scan(&a.ID, &a.State, &a.Title, &a.Text)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles,nil
}

func (r *ArticleRepository) FindByState(state string) (*model.Article, error){
	a := &model.Article{}
	if err := r.store.dbConn.QueryRow(
		r.store.dbCtx, "SELECT id, state, title, text FROM articles WHERE state = $1",
		state).Scan(&a.ID, &a.State, &a.Title, &a.Text); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrArticlesNotFound
		}
		return nil, err
	}
	return a,nil
}

func (r *ArticleRepository) FindById(id string) (*model.Article, error){
	a := &model.Article{}
	if err := r.store.dbConn.QueryRow(
		r.store.dbCtx, "SELECT id, state, title, text FROM articles WHERE state = $1",
		id).Scan(&a.ID, &a.State, &a.Title, &a.Text); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrArticlesNotFound
		}
		return nil, err
	}
	return a,nil
}