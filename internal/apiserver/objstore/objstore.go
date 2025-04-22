package objstore

var oss *ObjStore

type ObjStore interface {
	Avators() AvatorStore
}

func SetObjStore(objstore ObjStore) {
	oss = &objstore
}

func GetObjStore() *ObjStore {
	return oss
}
