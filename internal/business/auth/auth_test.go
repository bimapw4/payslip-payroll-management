package auth_test

import (
	"context"
	"os"
	"payslips/internal/business/auth"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"payslips/internal/repositories/users"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthorization_Success_WithRealJWT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)

	repo := &repositories.Repository{
		Users: mockUsers,
	}

	b := auth.NewBusiness(repo)

	ctx := context.Background()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	mockUsers.EXPECT().
		GetUserByUsername(ctx, "testuser").
		Return(&presentations.Users{
			ID:       "user-1",
			Username: "testuser",
			Password: string(hashedPassword),
		}, nil)

	input := entity.Authorization{
		Username: "testuser",
		Password: "password123",
		// Type:     "authorization",
	}

	resp, err := b.Authorization(ctx, input)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.AccessToken)

	// Optional: parse token untuk verifikasi
	token, parseErr := jwt.Parse(resp.AccessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	assert.NoError(t, parseErr)
	assert.True(t, token.Valid)
}

func TestAuthorization_InvalidPassword_WithRealJWT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)

	repo := &repositories.Repository{
		Users: mockUsers,
	}

	b := auth.NewBusiness(repo)

	ctx := context.Background()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctPassword"), bcrypt.DefaultCost)

	mockUsers.EXPECT().
		GetUserByUsername(ctx, "testuser").
		Return(&presentations.Users{
			ID:       "user-1",
			Username: "testuser",
			Password: string(hashedPassword),
		}, nil)

	input := entity.Authorization{
		Username: "testuser",
		Password: "wrongPassword",
		// Type:     "authorization",
	}

	resp, err := b.Authorization(ctx, input)
	assert.ErrorIs(t, err, common.ErrUnauthorized)
	assert.Nil(t, resp)
}
