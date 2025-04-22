package service

import (
	"context"

	"github.com/ividernvi/algohub/internal/apiserver/cache"
	"github.com/ividernvi/algohub/internal/apiserver/objstore"
	"github.com/ividernvi/algohub/internal/apiserver/store"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

type UserService interface {
	Create(ctx context.Context, user *v1.User, opts *v1.CreateOptions) error
	Get(ctx context.Context, username string, opts *v1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.UserList, error)
	Update(ctx context.Context, user *v1.User, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts *v1.DeleteOptions) error
	Verify(ctx context.Context, token string, opts *v1.VerifyOptions) (*v1.User, error)
	Authorize(ctx context.Context, token, id string, opts *v1.VerifyOptions) error
	Logout(ctx context.Context, token string, opts *v1.VerifyOptions) error
	GetAvatar(ctx context.Context, userID string, opts *v1.GetOptions) ([]byte, error)
	PutAvatar(ctx context.Context, userID string, data []byte, opts *v1.UpdateOptions) error
}

type userService struct {
	store store.Store
	cache cache.Cache
	minio objstore.ObjStore
}

func newUserService(s *service) *userService {
	return &userService{
		store: s.store,
		cache: s.cache,
		minio: s.minio,
	}
}

func (s *userService) Create(ctx context.Context, user *v1.User, opts *v1.CreateOptions) error {
	return s.store.Users().Create(ctx, user, opts)
}

func (s *userService) Get(ctx context.Context, username string, opts *v1.GetOptions) (*v1.User, error) {
	return s.store.Users().Get(ctx, username, opts)
}

func (s *userService) List(ctx context.Context, opts *v1.ListOptions) (*v1.UserList, error) {
	return s.store.Users().List(ctx, opts)
}

func (s *userService) Update(ctx context.Context, user *v1.User, opts *v1.UpdateOptions) error {
	return s.store.Users().Update(ctx, user, opts)
}

func (s *userService) Delete(ctx context.Context, username string, opts *v1.DeleteOptions) error {
	err := s.store.Users().Delete(ctx, username, opts)
	return err
}

func (s *userService) Verify(ctx context.Context, token string, opts *v1.VerifyOptions) (*v1.User, error) {
	if token[:7] == "Bearer " {
		opts.IsBearer = true
	}

	if token[:6] == "Basic " {
		opts.IsBasic = true
	}

	return s.store.Users().Verify(ctx, token, opts)
}

func (s *userService) Authorize(ctx context.Context, token, id string, opts *v1.VerifyOptions) error {
	u, err := s.Verify(ctx, token, opts)
	if err != nil || u.UserName != id {
		return core.ErrNoAuthorization
	}

	return nil
}

func (s *userService) Logout(ctx context.Context, token string, opts *v1.VerifyOptions) error {
	if s.cache.Tokens().Get(ctx, token, nil) != nil {
		return core.ErrNoAuthorization
	}
	return nil
}

func (s *userService) GetAvatar(ctx context.Context, userID string, opts *v1.GetOptions) ([]byte, error) {
	return s.minio.Avators().Get(ctx, userID, opts)
}

func (s *userService) PutAvatar(ctx context.Context, userID string, data []byte, opts *v1.UpdateOptions) error {
	return s.minio.Avators().Put(ctx, userID, data, opts)
}
