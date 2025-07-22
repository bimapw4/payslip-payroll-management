package auth

import (
	"context"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type Contract interface {
	Authorization(ctx context.Context, payload entity.Authorization) (*presentations.AuthorizationResp, error)
}

type business struct {
	repo *repositories.Repository
	jwt  common.JwtCode
}

func NewBusiness(repo *repositories.Repository) Contract {
	return &business{
		repo: repo,
		jwt:  common.NewJwt(),
	}
}

func (b *business) Authorization(ctx context.Context, payload entity.Authorization) (*presentations.AuthorizationResp, error) {

	var (
		users *presentations.Users
		err   error
	)

	// if payload.Type == "authorization" {
	users, err = b.repo.Users.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(payload.Password))
	if err != nil {
		return nil, common.ErrUnauthorized
	}
	// }

	accesstoken, err := b.jwt.GenerateAuthorizartionCode(entity.Claim{
		UserID:   users.ID,
		Username: users.Username,
		IsAdmin:  users.IsAdmin,
	})
	if err != nil {
		return nil, err
	}

	// refreshToken := strings.ReplaceAll(uuid.NewString(), "-", "")

	return &presentations.AuthorizationResp{
		AccessToken: accesstoken,
		// RefreshToken: refreshToken,
	}, nil

}
