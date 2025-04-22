package store

type Store interface {
	Users() UserStore
	Posts() PostStore
	Comments() CommentStore
	Likes() LikeStore
	Submits() SubmitStore
	Problems() ProblemStore
	Subscribes() SubscribeStore
	Solutions() SolutionStore
	Close()
}

var cli *Store

func Factory() Store {
	return *cli
}

func SetFactory(store Store) {
	cli = &store
}
