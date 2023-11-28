package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&UserGorm{}, &ProductGorm{})
}

type ProductGorm struct {
	gorm.Model
	Name  string
	Price uint
}

func (p *ProductGorm) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

// 由于该 struct 必须要实现 interface，所以，该 struct 的方法不像泛型可以直接返回一个泛型类型 T。只能返回 interface{}
func (p *ProductGorm) Get(db *gorm.DB, id uint) (interface{}, error) {
	var pg ProductGorm
	err := db.Where("id = ?", id).First(&pg).Error
	return pg, err
}

type UserGorm struct {
	gorm.Model
	FirstName string
	LastName  string
}

// 不使用泛型的情况下，每个新的 struct 都要写一遍 Create 和 Get，哪怕 Create 和 Get 中的逻辑一样。
func (u *UserGorm) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *UserGorm) Get(db *gorm.DB, id uint) (interface{}, error) {
	var ug UserGorm
	err := db.Where("id = ?", id).First(&ug).Error
	return ug, err
}

// 在这个示例中，接口只是为了与 handlingData 搭配展示类似泛型一样的参数传递。并不一定要使用接口
// type SQLExecutor interface {
// 	Create(db *gorm.DB) error
// 	Get(db *gorm.DB, id uint) (interface{}, error)
// }

// func handlingData(e SQLExecutor, db *gorm.DB, id uint) {
// 	e.Create(db)
// 	fmt.Println(e.Get(db, id))
// }

func main() {
	initDB()

	p := &ProductGorm{
		Name:  "product",
		Price: 100,
	}

	u := &UserGorm{
		FirstName: "first",
		LastName:  "last",
	}

	// 不使用泛型，需要给每个结构体都编写一遍相同的方法。
	p.Create(db)
	fmt.Println(p.Get(db, 1))

	u.Create(db)
	fmt.Println(u.Get(db, 1))

	// 并不需要 handlingData() 函数来演示不使用泛型的效果，
	// 只是用了接口后，添加该函数可以展现出接口中类似泛型一样的可以传递任意类型的效果
	// handlingData(p, db, 1)
	// handlingData(u, db, 1)
}
