package mysql

import (
	"sync"

	"github.com/ividernvi/algohub/internal/apiserver/store"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/db"
	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

var _ store.Store = &datastore{}
var (
	Once     sync.Once
	MySQLIns store.Store
)

func GetMySQLInstanceOr(opts db.DBOptions) (store.Store, error) {
	if MySQLIns != nil {
		return MySQLIns, nil
	}

	var err error
	Once.Do(func() {
		var database *gorm.DB
		database, err = db.New(opts)
		if err != nil {
			panic(err)
		}
		DropAndMigrate(database)

		MySQLIns = &datastore{db: database}
	})
	return MySQLIns, err
}

func (d *datastore) Users() store.UserStore {
	return newUserStore(d.db)
}

func (d *datastore) Posts() store.PostStore {
	return newPostStore(d.db)
}

func (d *datastore) Comments() store.CommentStore {
	return newCommentStore(d.db)
}

func (d *datastore) Likes() store.LikeStore {
	return newLikeStore(d.db)
}

func (d *datastore) Submits() store.SubmitStore {
	return newSubmitStore(d.db)
}

func (d *datastore) Problems() store.ProblemStore {
	return newProblemStore(d.db)
}

func (d *datastore) Subscribes() store.SubscribeStore {
	return newSubscribeStore(d.db)
}

func (d *datastore) Solutions() store.SolutionStore {
	return newSolutionStore(d.db)
}

func (d *datastore) Close() {
	sqlDB, err := d.db.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&v1.User{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&v1.Post{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&v1.Comment{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&v1.Like{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&v1.Subscribe{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&v1.Problem{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&v1.Solution{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&v1.Submit{})
	if err != nil {
		panic(err)
	}

}

func AutoDrop(db *gorm.DB) {
	err := db.Migrator().DropTable(&v1.User{})
	if err != nil {
		panic(err)
	}

	err = db.Migrator().DropTable(&v1.Post{})
	if err != nil {
		panic(err)
	}

	err = db.Migrator().DropTable(&v1.Comment{})
	if err != nil {
		panic(err)
	}

	err = db.Migrator().DropTable(&v1.Like{})
	if err != nil {
		panic(err)
	}

	err = db.Migrator().DropTable(&v1.Subscribe{})
	if err != nil {
		panic(err)
	}

	err = db.Migrator().DropTable(&v1.Problem{})
	if err != nil {
		panic(err)
	}

	err = db.Migrator().DropTable(&v1.Solution{})
	if err != nil {
		panic(err)
	}

	err = db.Migrator().DropTable(&v1.Submit{})
	if err != nil {
		panic(err)
	}
}

func DropAndMigrate(db *gorm.DB) {
	AutoDrop(db)
	AutoMigrate(db)

	DataInit(db)
}
