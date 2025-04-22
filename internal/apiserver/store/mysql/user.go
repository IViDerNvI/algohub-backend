package mysql

import (
	"context"
	"encoding/base64"
	"strings"

	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
	"github.com/ividernvi/algohub/pkg/util/bcryptutil"
	"github.com/ividernvi/algohub/pkg/util/jwtutil"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func newUserStore(db *gorm.DB) *userStore {
	return &userStore{db: db}
}

func (s *userStore) Create(ctx context.Context, user *v1.User, opts *v1.CreateOptions) error {
	return s.db.Create(user).Error
}

func (s *userStore) Get(ctx context.Context, username string, opts *v1.GetOptions) (*v1.User, error) {
	var user v1.User
	err := s.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (s *userStore) List(ctx context.Context, opts *v1.ListOptions) (*v1.UserList, error) {
	var users []v1.User
	var total int64

	query := opts.ApplyListOptions(s.db)
	err := query.Model(&v1.User{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = query.Offset(opts.Offset).Limit(opts.Limit).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &v1.UserList{
		ListMeta: v1.ListMeta{
			TotalItems: total,
		},
		Items: users,
	}, nil
}

func (s *userStore) Update(ctx context.Context, user *v1.User, opts *v1.UpdateOptions) error {
	return s.db.Model(&v1.User{}).Where("instance_id = ?", user.InstanceID).Updates(user).Error
}

func (s *userStore) Delete(ctx context.Context, username string, opts *v1.DeleteOptions) error {
	return s.db.Where("username = ?", username).Delete(&v1.User{}).Error
}

func (s *userStore) Verify(ctx context.Context, token string, opts *v1.VerifyOptions) (*v1.User, error) {
	if token == "" {
		return nil, core.ErrTokenInvalid
	}

	if opts.IsBasic {
		token = token[len("Basic "):]
		// Basic token is in the format "username:password" encoded in base64
		decoded, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			return nil, core.ErrTokenInvalid
		}

		// Split the decoded string into username and password
		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) != 2 {
			return nil, core.ErrTokenInvalid
		}

		username := parts[0]
		password := parts[1]

		var user v1.User
		err = s.db.Where("username = ?", username).First(&user).Error
		if err != nil {
			return nil, err
		}

		if bcryptutil.ComparePassword(user.Password, password) != nil {
			return nil, core.ErrUserVerify
		}

		return &user, nil
	}

	if opts.IsBearer {
		token = token[len("Bearer "):]

		claims, err := jwtutil.ValidateJWT(token)
		if err != nil {
			return nil, err
		}

		var user v1.User
		err = s.db.Where("username = ?", claims["username"]).First(&user).Error
		if err != nil {
			return nil, err
		}

		return &user, nil
	}

	return nil, core.ErrTokenUnsported
}
