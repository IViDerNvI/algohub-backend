package substore

var oss SubStore

type SubStore interface {
	AwsStore() AwsStore
}

func SetSubStore(Substore SubStore) {
	oss = Substore
}

func GetSubStore() SubStore {
	return oss
}
