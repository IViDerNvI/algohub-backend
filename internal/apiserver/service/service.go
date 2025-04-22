package service

import (
	"github.com/ividernvi/algohub/internal/apiserver/cache"
	"github.com/ividernvi/algohub/internal/apiserver/objstore"
	"github.com/ividernvi/algohub/internal/apiserver/store"
)

type Service interface {
	Users() UserService
	Posts() PostService
	Problems() ProblemService
	Comments() CommentService
	Submits() SubmitService
	Subscribes() SubscribeService
	Likes() LikeService
	Solutions() SolutionService
}

type service struct {
	store store.Store
	cache cache.Cache
	minio objstore.ObjStore
}

func NewService(store store.Store) Service {
	return &service{
		store: store,
		cache: cache.CacheFactory(),
		minio: *objstore.GetObjStore(),
	}
}

func (s *service) Users() UserService {
	return newUserService(s)
}

func (s *service) Posts() PostService {
	return newPostService(s)
}

func (s *service) Problems() ProblemService {
	return newProblemService(s)
}

func (s *service) Submits() SubmitService {
	return newSubmitService(s)
}

func (s *service) Subscribes() SubscribeService {
	return newSubscribeService(s)
}

func (s *service) Likes() LikeService {
	return newLikeService(s)
}

func (s *service) Comments() CommentService {
	return newCommentService(s)
}

func (s *service) Solutions() SolutionService {
	return newSolutionService(s)
}
