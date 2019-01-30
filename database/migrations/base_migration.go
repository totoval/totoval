package migration

import (
	"Wallet/app/models"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
)

// contains all the migrators
var migratorList []Migrator

type Migrator interface {
	Up(db *gorm.DB)
	Down(db *gorm.DB)
	MigratorIdentifier
}
type MigratorIdentifier interface {
	Name(*Migrator) string
}
type MigratorIdentify struct {}
func (m *MigratorIdentify) Name(migrator *Migrator) string {
	return strings.Replace(reflect.TypeOf(*migrator).String(), "*migration.", "", 1)
}

type Migration struct {
	ID        uint   `gorm:"column:id;primary_key"`
	Migration string `gorm:"column:migration;type:varchar(255)"`
	Batch     int    `gorm:"column:batch;"`
}

// 建立migration表
func (m *Migration) up(db *gorm.DB) {
	tx := db.Begin()
	{
		tx.CreateTable(&m)
	}
	tx.Commit()
}

type MigrationUtils struct {
	db *gorm.DB
	Migration
}

func (m *MigrationUtils) SetDB(){
	m.db = models.DB()
}

// 项目初始化
func (m *MigrationUtils) Initialize() {
	m.Migration.up(m.db)
}

// 所有migration列表
func (m *MigrationUtils) migrationList() (migrationList []Migration) {
	m.db.Find(&migrationList)
	return
}

func (m *MigrationUtils) needMigrateList() (_migratorList []Migrator) {
	for _, migrator := range migratorList {
		found := false
		for _, migration := range m.migrationList() {
			if migrator.Name(&migrator) == migration.Migration {
				found = true
				break
			}
		}

		if !found {
			_migratorList = append(_migratorList, migrator)
		}
	}
	return
}

func (m *MigrationUtils) Migrate(){
	tx := m.db.Begin()
	{
		for _, migrator := range m.needMigrateList() {
			migrator.Up(m.db)
		}
	}
	tx.Commit()
}
func (m *MigrationUtils) Fresh(){

}
func (*MigrationUtils) Install(){
	//   --database[=DATABASE]  The database connection to use
	//@todo
	//  Create the migration repository
}
func (*MigrationUtils) Refresh(){

}
func (*MigrationUtils) Reset(){

}
func (*MigrationUtils) Rollback(){

}
func (*MigrationUtils) Status(){
	//+------+--------------------------------------------------------------+-------+
	//| Ran? | Migration                                                    | Batch |
	//+------+--------------------------------------------------------------+-------+
	//| Yes  | 2014_10_12_000000_create_users_table                         | 3     |
	//| Yes  | 2014_10_12_100000_create_password_resets_table               | 1     |
	//| Yes  | 2016_06_01_000001_create_oauth_auth_codes_table              | 1     |
	//| Yes  | 2016_06_01_000002_create_oauth_access_tokens_table           | 1     |
	//| Yes  | 2016_06_01_000003_create_oauth_refresh_tokens_table          | 1     |
	//| Yes  | 2016_06_01_000004_create_oauth_clients_table                 | 1     |
	//| Yes  | 2016_06_01_000005_create_oauth_personal_access_clients_table | 1     |
	//| Yes  | 2019_01_10_081308_create_user_verification_table             | 2     |
	//| Yes  | 2019_01_10_165704_create_data_area_table                     | 2     |
	//| No   | 2019_01_22_112905_create_customer_wechat_table               |       |
	//| No   | 2019_01_22_112909_create_customer_table                      |       |
	//+------+--------------------------------------------------------------+-------+
}
