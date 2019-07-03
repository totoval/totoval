package policies

import (
	"github.com/totoval/framework/model"
)

type userPolicy struct {
}

func NewUserPolicy() *userPolicy {
	return &userPolicy{}
}

func (up *userPolicy) Before() *bool {
	return nil
}
func (up *userPolicy) Create(userIF model.IUser) bool {
	return true
}
func (up *userPolicy) Update(userIF model.IUser) bool      { return true }
func (up *userPolicy) Delete(userIF model.IUser) bool      { return true }
func (up *userPolicy) ForceDelete(userIF model.IUser) bool { return true }
func (up *userPolicy) View(userIF model.IUser) bool {
	//currentUser := userIF.Value().(*models.User)
	return true
}
func (up *userPolicy) Restore(userIF model.IUser) bool { return true }
