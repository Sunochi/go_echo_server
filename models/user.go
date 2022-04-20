package models

import (
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/clause"

	d "go_example_server/database"
)

type User struct {
	ID      uint   `csv:"id"`
	Name    string `csv:"name" gorm:"type:varchar(24);not null;default:''"`
	Phone   string `csv:"phone" gorm:"type:varchar(24);not null;default:''"`
	Address string `csv:"address" gorm:"type:varchar(255);not null;default:''"`
}

// TableName overrides the table name
func (User) TableName() string {
	return "user"
}

// Log出力の定義。example. log.AppLog.Info("test log:", zap.Object("UserObject", userStruct))
func (u User) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddUint("id", u.ID)
	enc.AddString("name", u.Name)
	enc.AddString("phone", u.Phone)
	enc.AddString("address", u.Address)
	return nil
}

func (u *User) Create() error {
	db := d.GetDB()
	err := db.Create(&u).Error
	return err
}

func (User) FetchByName(name string) ([]User, error) {
	var users []User
	db := d.GetDB()
	err := db.Where("name = ?", name).Find(&users).Error
	return users, err
}

func (User) Save(u *[]User) error {
	var repo = d.GetDB()
	// upsert
	err := repo.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "phone", "address"}),
	}).Create(&u).Error
	return err
}

func (User) FetchAll() ([]User, error) {
	var users []User
	var repo = d.GetDB()
	err := repo.Find(&users).Error
	return users, err
}
