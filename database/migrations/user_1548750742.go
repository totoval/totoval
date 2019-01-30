package migration

import (
	"Wallet/app/models"
	"github.com/jinzhu/gorm"
)

func init(){
	migratorList = append(migratorList, &User1548750742{})
}

type User1548750742 struct{
	MigratorIdentify
	MigrationUtils
}

func (*User1548750742) Up(db *gorm.DB) {
	db.CreateTable(&models.User{})
}

func (*User1548750742) Down(db *gorm.DB) {
	db.DropTableIfExists(&models.User{})
}
