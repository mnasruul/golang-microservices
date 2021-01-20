package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserNotFoundInDatabase(t *testing.T) {
	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
}
