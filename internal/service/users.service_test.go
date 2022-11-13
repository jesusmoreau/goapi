package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"goapi/internal/entity"
	"goapi/internal/repository"
	"testing"
)

func TestRegisterUser(t *testing.T) {

	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			Email:         "test@test.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_UserAlreadyExists",
			Email:         "test@exists.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	repo := &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(&entity.User{Email: "test@exists.com"}, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			//	t.Parallel()
			repo.Mock.Test(t)

			s := New(repo)

			err := s.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %, got %v", tc.ExpectedError, err)
			}
		})
	}
}
