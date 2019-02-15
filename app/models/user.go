package models

import (
	"time"
	"github.com/totoval/framework/model"
)

type User struct {
	model.BaseModel
	ID         uint       `gorm:"column:user_id;primary_key;auto_increment"`
	Name       string     `gorm:"column:user_name;type:varchar(100)"`
	Email      string     `gorm:"column:user_email;type:varchar(100);unique_index"`
	Telephone  string     `gorm:"column:user_telephone;type:varchar(100);unique_index"`
	Password   string     `gorm:"column:user_password;type:varchar(100)"`
	VerifiedAt time.Time  `gorm:"column:user_verified_at"`
	CreatedAt  time.Time  `gorm:"column:user_created_at"`
	UpdatedAt  time.Time  `gorm:"column:user_updated_at"`
	DeletedAt  *time.Time `gorm:"column:user_deleted_at"`
}

func (user *User) User() *User {
	model.DB().Where("user_id = ?", 1).Find(user)
	return user
}

func (*User) GetObjArr()         {} //@todo     public function getObjArr(?array $filter_arr = [], ?array $sort_arr = null, ?int $limit = null, bool $with_trashed = false): Collection;
func (*User) GetObjArrPaginate() {} //@todo     public function getObjArrPaginate(int $per_page, ?array $filter_arr = [], ?array $sort_arr = null, bool $with_trashed = false): LengthAwarePaginator;
