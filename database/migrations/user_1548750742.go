package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/totoval/framework/database/migration"
	"time"
)

func init() {
	migration.AddMigrator(&User1548750742{})
}

type User struct {
	ID    uint   `gorm:"column:user_id;primary_key;auto_increment"`
	Name  string `gorm:"column:user_name;type:varchar(100)"`
	Email string `gorm:"column:user_email;type:varchar(100);unique_index"`
	//Telephone  string     `gorm:"column:user_telephone;type:varchar(100);unique_index"`
	Password string `gorm:"column:user_password;type:varchar(100)"`
	//VerifiedAt time.Time  `gorm:"column:user_verified_at"`
	CreatedAt *time.Time  `gorm:"column:user_created_at"`
	UpdatedAt time.Time  `gorm:"column:user_updated_at"`
	DeletedAt *time.Time `gorm:"column:user_deleted_at"`
}

type User1548750742 struct {
	migration.MigratorIdentify
	migration.MigrationUtils
}

func (*User1548750742) Up(db *gorm.DB) *gorm.DB {
	db = db.CreateTable(&User{})
	return db
}

func (*User1548750742) Down(db *gorm.DB) *gorm.DB {
	db = db.DropTableIfExists(&User{})
	return db
}
