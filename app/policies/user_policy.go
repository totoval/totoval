package policies

import (
	"strconv"

	"github.com/totoval/framework/helpers/debug"
	"github.com/totoval/framework/model"
	"totoval/app/models"
)

type userPolicy struct {
}

func NewUserPolicy() *userPolicy {
	return &userPolicy{}
}

func (up *userPolicy) Before(IUser model.IUser, routeParamMap map[string]string) *bool {
	return nil
}
func (up *userPolicy) Create(IUser model.IUser, routeParamMap map[string]string) bool {
	return true
}
func (up *userPolicy) Update(IUser model.IUser, routeParamMap map[string]string) bool { return true }
func (up *userPolicy) Delete(IUser model.IUser, routeParamMap map[string]string) bool { return true }
func (up *userPolicy) ForceDelete(IUser model.IUser, routeParamMap map[string]string) bool {
	return true
}
func (up *userPolicy) View(IUser model.IUser, routeParamMap map[string]string) bool {
	// get current user
	currentUser := IUser.Value().(*models.User)
	debug.Dump(currentUser, routeParamMap)

	// get param user
	userIdStr, ok := routeParamMap["userId"]
	if !ok {
		return false
	}
	userIdUint, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return false
	}

	if *currentUser.ID != uint(userIdUint) {
		return false
	}

	return true
}
func (up *userPolicy) Restore(IUser model.IUser, routeParamMap map[string]string) bool { return true }
