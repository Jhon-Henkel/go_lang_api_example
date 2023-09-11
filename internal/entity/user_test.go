package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Joe Miller", "joe@tes.com", "12345678")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Joe Miller", user.Name)
	assert.Equal(t, "joe@tes.com", user.Email)
	assert.NotEmpty(t, user.Password) 
	assert.NotEmpty(t, user.ID) 
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Joe Miller", "joe@tes.com", "12345678")
	assert.Nil(t, err)
	assert.True(t, user.ComparePassword("12345678"))
	assert.False(t, user.ComparePassword("123456789"))
	assert.NotEqual(t, "12345678", user.Password)
}