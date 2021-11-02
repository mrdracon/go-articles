package store_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-articles/model"
	"go-articles/store"
	"testing"
)

func TestArticleRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, TestCfg.DSN)

	defer teardown("articles")

	uid, err := s.Articles().Create(&model.Article{
		ID:    "abc",
		State: "NEW",
		Title: "11",
		Text:  "22",
	})
	assert.NoError(t, err)
	fmt.Println(uid)
}

func TestArticleRepository_FindById(t *testing.T) {
	s, teardown := store.TestStore(t, TestCfg.DSN)

	defer teardown("articles")
	state := "NEW"
	_, err := s.Articles().FindByState(state)
	assert.Error(t, err)

	_, _ = s.Articles().Create(&model.Article{
		ID:    "abc",
		State: "NEW",
		Title: "11",
		Text:  "22",	})

	u, err := s.Articles().FindByState(state)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestArticleRepository_ListArticles(t *testing.T) {
	s, teardown := store.TestStore(t, TestCfg.DSN)

	defer teardown("articles")
	state := "NEW"

	_, _ = s.Articles().Create(&model.Article{
		ID:    "abc",
		State: state,
		Title: "11",
		Text:  "11",	})

	_, _ = s.Articles().Create(&model.Article{
		ID:    "cba",
		State: state,
		Title: "22",
		Text:  "22",	})

	u, err := s.Articles().ListArticles()
	assert.NoError(t, err)
	assert.NotNil(t, u)
}