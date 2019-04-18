package migrations

import (
	"github.com/jinzhu/gorm"

	"github.com/totoval/framework/database/migration"
	"github.com/totoval/framework/model"

	"time"
)

func init() {
	migration.AddMigrator(&CreateUserTable1548750742{})
}

type User struct {
	ID    *uint   `gorm:"column:user_id;primary_key;auto_increment"`
	Name  *string `gorm:"column:user_name;type:varchar(100)"`
	Email *string `gorm:"column:user_email;type:varchar(100);unique_index"`
	//Telephone  *string     `gorm:"column:user_telephone;type:varchar(100);unique_index"`
	Password *string `gorm:"column:user_password;type:varchar(100)"`
	//VerifiedAt time.Time  `gorm:"column:user_verified_at"`
	CreatedAt *time.Time `gorm:"column:user_created_at"`
	UpdatedAt time.Time  `gorm:"column:user_updated_at"`
	DeletedAt *time.Time `gorm:"column:user_deleted_at"`
	model.BaseModel
}

func (u *User) TableName() string {
	return u.SetTableName("user")
}

type CreateUserTable1548750742 struct {
	migration.MigratorIdentify
	migration.MigrationUtils
}

func (*CreateUserTable1548750742) Up(db *gorm.DB) *gorm.DB {
	db = db.CreateTable(&User{})
	return db
}

func (*CreateUserTable1548750742) Down(db *gorm.DB) *gorm.DB {
	db = db.DropTableIfExists(&User{})
	return db
}
