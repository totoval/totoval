package models

import (
	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/helpers/ptr"
	"github.com/totoval/framework/helpers/zone"
	"github.com/totoval/framework/model"
)

type User struct {
	ID        *uint      `gorm:"column:user_id;primary_key;auto_increment"`
	Name      *string    `gorm:"column:user_name;type:varchar(100)"`
	Email     *string    `gorm:"column:user_email;type:varchar(100);unique_index;not null"`
	Password  *string    `gorm:"column:user_password;type:varchar(100);not null"`
	CreatedAt *zone.Time `gorm:"column:user_created_at"`
	UpdatedAt zone.Time  `gorm:"column:user_updated_at"`
	DeletedAt *zone.Time `gorm:"column:user_deleted_at"`
	model.BaseModel
}

func (user *User) TableName() string {
	return user.SetTableName("user")
}

func (user *User) Default() interface{} {
	return User{}
}

func (user *User) Scan(userId uint) error {
	newUser := User{
		ID: ptr.Uint(userId),
	}
	if err := m.H().First(&newUser, false); err != nil {
		return err
	}
	*user = newUser
	return nil
}
func (user *User) Value() interface{} {
	return user
}

func (user *User) User() *User {
	//model.DB().Where("user_id = ?", 1).Find(user)
	return user
}

func (user *User) ObjArr(filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (interface{}, error) {
	var outArr []User
	if err := m.H().Q(filterArr, sortArr, limit, withTrashed).Find(&outArr).Error; err != nil {
		return nil, err
	}
	return outArr, nil
}
func (user *User) ObjArrPaginate(c model.Context, perPage uint, filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (pagination model.Pagination, err error) {
	var outArr []User
	filter := model.Model(*m.H().Q(filterArr, sortArr, limit, withTrashed))
	return filter.Paginate(&outArr, c, perPage)
}
