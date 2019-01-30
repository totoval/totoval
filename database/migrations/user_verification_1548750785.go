package migration

import (
	"Wallet/app/models"
	"github.com/jinzhu/gorm"
)

func init(){
	migratorList = append(migratorList, &UserVerification1548750785{})
}

type UserVerification1548750785 struct{
	MigratorIdentify
	MigrationUtils
}

func (*UserVerification1548750785) Up(db *gorm.DB) {
	db.CreateTable(&models.UserVerification{})
}

func (*UserVerification1548750785) Down(db *gorm.DB) {
	db.DropTableIfExists(&models.UserVerification{})
}
