package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/totoval/framework/helpers/ptr"

	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/model"
)

type User struct {
	ID    *uint   `gorm:"column:user_id;primary_key;auto_increment"`
	Name  *string `gorm:"column:user_name;type:varchar(100)"`
	Email *string `gorm:"column:user_email;type:varchar(100);unique_index"`
	//Telephone  string     `gorm:"column:user_telephone;type:varchar(100);unique_index"`
	Password *string `gorm:"column:user_password;type:varchar(100)"`
	//VerifiedAt mysql.NullTime  `gorm:"column:user_verified_at"`
	CreatedAt *time.Time `gorm:"column:user_created_at"`
	UpdatedAt time.Time  `gorm:"column:user_updated_at"`
	DeletedAt *time.Time `gorm:"column:user_deleted_at"`
	model.BaseModel
}

func (user *User) TableName() string {
	return user.SetTableName("user")
}

func (user *User) Default() interface{} {
	return User{
		Name: ptr.String(""),
	}
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
func (user *User) ObjArrPaginate(c *gin.Context, perPage uint, filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) (pagination model.Pagination, err error) {
	var outArr []User
	filter := model.Model(*m.H().Q(filterArr, sortArr, limit, withTrashed))
	return filter.Paginate(&outArr, c, perPage)
}
