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

type UserGorm struct {
	gorm.Model
	FirstName string
	LastName  string
}

// 由泛型结构体统一处理
type SQLExecutor[T any] struct {
	db *gorm.DB
}

// 通用的处理器
// ！！！最关键的是通过泛型的使用！！！
// ！！！代码整体上少写了很多 Create 和 Get。尤其是当还需要添加更多 strcut 时，节省的代码量会非常多！！！
func NewSQLExecutor[T any](db *gorm.DB) *SQLExecutor[T] {
	return &SQLExecutor[T]{
		db: db,
	}
}

func (e *SQLExecutor[T]) Create(t *T) error {
	return e.db.Create(t).Error
}

func (e *SQLExecutor[T]) Get(id uint) (*T, error) {
	var t T
	err := e.db.Where("id = ?", id).First(&t).Error
	return &t, err
}

// 这个函数只是为了与使用接口对比而设计的，真实用法不是这样的。
// func handlingData[T any](e *SQLExecutor[T], p *T) {
// 	e.Create(p)
// 	fmt.Println(e.Get(1))
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

	// 泛型的最佳时机，就是当两个“对象”的某些“方法”的行为完全一样时，我们可以通过泛型来声明这些方法，以防止重复编写完全相同的代码。
	// 就像下面这段代码中，ProductGorm 与 UserGorm 的两个方法代码其实是完全相同的，如果使用 interface{}，那么我们需要重复写 Create() 与 Get() 方法。
	// 其实，本质上 ProductGorm 与 UserGorm 所需要执行的 SQL 是完全一样，不一样的只是其中的列名而已。
	ep := NewSQLExecutor[ProductGorm](db)
	ep.Create(p)
	fmt.Println(ep.Get(1))

	eu := NewSQLExecutor[UserGorm](db)
	eu.Create(u)
	fmt.Println(eu.Get(1))

	// 使用 handlingData() 只是为了与不使用泛型的示例中，添加了接口的使用示例保持一致，以方便对比。
	// handlingData(NewSQLExecutor[ProductGorm](db), p)
	// handlingData(NewSQLExecutor[UserGorm](db), u)
}
