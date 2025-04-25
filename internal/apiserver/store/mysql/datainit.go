package mysql

import (
	"time"

	v1 "github.com/ividernvi/algohub/model/v1"
	"gorm.io/gorm"
)

func DataInit(db *gorm.DB) {
	CreateAdminUser(db)
	CreateNormalUser(db)
	CreateNormalPost(db)
	CreateNormalComment(db)
	CreateReferenceComment(db)
	CreateNormalLike(db)
	CreateNormalSubscribe(db)
	CreateNormalProblem(db)
	CreateNormalSolution(db)
	CreateNormalSubmit(db)
}

func CreateAdminUser(db *gorm.DB) {
	db.Create(&v1.User{
		UserName:   "admin",
		Password:   "admin",
		Status:     "admin",
		NickName:   "admin",
		Email:      "admin@admin.com",
		Phone:      "1234567890",
		Avatar:     "abandoned",
		Bio:        "This is admin",
		Company:    "Admin Company",
		Location:   "Admin Location",
		ProfileURL: "https://example.com/profile",
	})

	db.Exec("UPDATE users SET instance_id = ? WHERE username = ?", 1, "admin")
}

func CreateNormalUser(db *gorm.DB) {
	db.Create(&v1.User{
		UserName:   "normal",
		Password:   "normal",
		Status:     "normal",
		NickName:   "normal",
		Email:      "normal@normal.com",
		Phone:      "0987654321",
		Avatar:     "abandoned",
		Bio:        "This is normal user",
		Company:    "Normal Company",
		Location:   "Normal Location",
		ProfileURL: "https://example.com/profile",
	})
	db.Exec("UPDATE users SET instance_id = ? WHERE username = ?", 2, "normal")
}

func CreateNormalPost(db *gorm.DB) {

	db.Create(&v1.Post{
		Title:   "Normal Post",
		Content: "This is a normal post",
		Author:  "normal",
		Tag:     "test,normal,example",
	})

	db.Create(&v1.Post{
		ObjMeta: v1.ObjMeta{
			InstanceID: 1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		Title:   "Another Normal Post",
		Content: "This is another normal post",
		Author:  "admin",
		Tag:     "test,admin,example",
	})

	db.Create(&v1.Post{
		Title:   "Solution Post",
		Content: "This is a solution post",
		Author:  "normal",
		Tag:     "solution:aPlusb,easy,math,gready",
	})

	db.Exec("UPDATE posts SET instance_id = ? WHERE title = ?", 1, "Normal Post")
	db.Exec("UPDATE posts SET instance_id = ? WHERE title = ?", 2, "Another Normal Post")
	db.Exec("UPDATE posts SET instance_id = ? WHERE title = ?", 3, "Solution Post")

}

func CreateNormalComment(db *gorm.DB) {
	db.Create(&v1.Comment{
		ObjMeta: v1.ObjMeta{
			InstanceID: 1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		Content:    "This is a normal comment",
		RefersItem: 1,
		RefersType: "post",
		SourceItem: 1,
		SourceType: "post",
		Auhtor:     "normal",
	})

	db.Create(&v1.Comment{
		Content:    "This is another normal comment",
		RefersItem: 1,
		RefersType: "post",
		SourceItem: 1,
		SourceType: "post",
		Auhtor:     "admin",
	})

	db.Exec("UPDATE comments SET instance_id = ? WHERE content = ?", 1, "This is a normal comment")
	db.Exec("UPDATE comments SET instance_id = ? WHERE content = ?", 2, "This is another normal comment")

}

func CreateReferenceComment(db *gorm.DB) {
	db.Create(&v1.Comment{
		ObjMeta: v1.ObjMeta{
			InstanceID: 2,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		Content:    "This is a reference comment",
		Auhtor:     "normal",
		RefersItem: 1,
		RefersType: "comment",
		SourceItem: 1,
		SourceType: "post",
	})

	db.Exec("UPDATE comments SET instance_id = ? WHERE content = ?", 3, "This is a reference comment")
}

func CreateNormalLike(db *gorm.DB) {
	db.Create(&v1.Like{
		ObjMeta: v1.ObjMeta{
			InstanceID: 1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		UserName: "normal",
		ItemType: "post",
		ItemID:   1,
	})

	db.Create(&v1.Like{
		ObjMeta: v1.ObjMeta{
			InstanceID: 2,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		UserName: "admin",
		ItemType: "post",
		ItemID:   1,
	})
	db.Exec("UPDATE likes SET instance_id = ? WHERE username = ?", 1, "normal")
	db.Exec("UPDATE likes SET instance_id = ? WHERE username = ?", 2, "admin")
}

func CreateNormalSubscribe(db *gorm.DB) {
	db.Create(&v1.Subscribe{
		ObjMeta: v1.ObjMeta{
			InstanceID: 1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		UserName: "normal",
		ItemType: "user",
		ItemName: "admin",
	})

	db.Exec("UPDATE subscribes SET instance_id = ? WHERE username = ?", 1, "normal")
}

func CreateNormalProblem(db *gorm.DB) {
	db.Create(&v1.Problem{
		ObjMeta: v1.ObjMeta{
			InstanceID: 1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		Unique_ID:   "aPlusb",
		Title:       "a + b",
		Descrition:  "given a and b, calculate a + b",
		Author:      "admin",
		TimeLimit:   1000,
		MemoryLimit: 125536,
		Tag:         "math,test,example,easy",
		Level:       1,
	})

	db.Exec("UPDATE problems SET instance_id = ? WHERE unique_id = ?", 1, "aPlusb")
}

func CreateNormalSolution(db *gorm.DB) {
	db.Create(&v1.Solution{
		ProblemID:  "aPlusb",
		TestData:   "1 2",
		TestResult: "3",
		Provider:   "admin",
	})

	db.Create(&v1.Solution{
		ProblemID:  "aPlusb",
		TestData:   "3 4",
		TestResult: "7",
		Provider:   "admin",
	})

	db.Create(&v1.Solution{
		ProblemID:  "aPlusb",
		TestData:   "5 6",
		TestResult: "11",
		Provider:   "admin",
	})

	db.Create(&v1.Solution{
		ProblemID:  "aPlusb",
		TestData:   "7 8",
		TestResult: "15",
		Provider:   "admin",
	})

	db.Exec("UPDATE solutions SET instance_id = ? WHERE data_test = ?", 1, "1,2")
	db.Exec("UPDATE solutions SET instance_id = ? WHERE data_test = ?", 2, "3,4")
	db.Exec("UPDATE solutions SET instance_id = ? WHERE data_test = ?", 3, "5,6")
	db.Exec("UPDATE solutions SET instance_id = ? WHERE data_test = ?", 4, "7,8")
}

func CreateNormalSubmit(db *gorm.DB) {
	db.Create(&v1.Submit{
		CodeText:  "int add(int a, int b) { return a + b; }",
		ProblemID: "aPlusb",
		Language:  "cpp",
		Author:    "normal",
		Status:    "pending",
	})

	db.Exec("UPDATE submits SET instance_id = ? WHERE author = ?", 1, "normal")
}
