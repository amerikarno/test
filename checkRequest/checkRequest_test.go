package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



type mockCheck struct {
	mock.Mock
}

func (m *mockCheck) checkPermissions(reqPermission []string, allowedPermission []string) error {
	args := m.Called(reqPermission, allowedPermission)
	return args.Error(0)
}

// TestMain ...
func TestMain(t *testing.T) {
	t.Parallel()
	t.Run("Main", func(t *testing.T){

		reqPermission := []string{"PT", "CNX", "NEXT"}
		allowedPermission := []string{"PT"}

		mc := new(mockCheck)
		mc.On("checkPermissions",reqPermission, allowedPermission).Return(nil)
		err := CheckPermissions(reqPermission, allowedPermission)
		assert.Equal(t, nil, err)
		assert.NoError(t, err)
	})
	t.Run("Main", func(t *testing.T){

		// reqPermission := []string{"PT", "CNX", "NEXT"}
		allowedPermission := []string{"PT"}

		mc := new(mockCheck)
		mc.On("checkPermissions",mock.Anything, allowedPermission).Return(nil)
		err := CheckPermissions([]string{mock.Anything}, allowedPermission)
		assert.Equal(t, nil, err)
		assert.NoError(t, err)
	})
}