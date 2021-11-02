package store_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-articles/model"
	"go-articles/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, TestCfg.DSN)

	defer teardown("users")

	uid, err := s.User().Create(&model.User{
		Email:             "user@example.org",
	})
	assert.NoError(t, err)
	fmt.Println(uid)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, TestCfg.DSN)

	defer teardown("users")
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	_, _ = s.User().Create(&model.User{
		Email:             "user@example.org",
	})

	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}