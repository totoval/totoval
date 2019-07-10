package controllers

import (
	"net/http"

	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/helpers/ptr"
	"github.com/totoval/framework/helpers/toto"
	"github.com/totoval/framework/http/controller"
	"github.com/totoval/framework/http/middleware"
	"github.com/totoval/framework/model"
	"github.com/totoval/framework/policy"
	"github.com/totoval/framework/request"

	"totoval/app/models"
	"totoval/app/policies"
)

type User struct {
	controller.BaseController
}

func (*User) LogOut(c *request.Context) {
	if err := middleware.Revoke(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, toto.V{})
	return
}

func (u *User) Info(c *request.Context) {
	if u.Scan(c) {
		return
	}
	user := u.User().Value().(*models.User)

	if permit, _ := u.Authorize(c, policies.NewUserPolicy(), policy.ActionView); !permit {
		c.JSON(http.StatusForbidden, toto.V{"error": policy.UserNotPermitError{}.Error()})
		return
	}

	user.Password = ptr.String("") // remove password value for response rendering
	c.JSON(http.StatusOK, toto.V{"data": user})
	return
}

func (*User) AllUser(c *request.Context) {
	user := &models.User{}
	outArr, err := user.ObjArr([]model.Filter{}, []model.Sort{}, 0, false)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, toto.V{"data": outArr.([]models.User)})
	return
}

func (*User) PaginateUser(c *request.Context) {
	user := &models.User{}
	pagination, err := user.ObjArrPaginate(c, 25, []model.Filter{}, []model.Sort{}, 0, false)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, toto.V{"data": toto.V{"item": pagination.ItemArr(), "totalPage": pagination.LastPage(), "currentPage": pagination.CurrentPage(), "count": pagination.Count(), "total": pagination.Total()}})
	return
}

func (*User) Update(c *request.Context) {
	var id uint
	id = 14
	user := models.User{
		ID: &id,
	}
	if err := m.H().First(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err.Error()})
		return
	}

	name := "t2222es123t"
	modifyUser := models.User{
		Name: &name,
	}
	if err := m.H().Save(&user, modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, toto.V{"data": user})
	return

	// m.Transaction(func() {
	// 	fmt.Println(id)
	// 	panic(123)
	// }, 3)
}
func (*User) Delete(c *request.Context) {
	var id uint
	id = 14
	user := models.User{
		ID: &id,
	}
	if err := m.H().Delete(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, toto.V{"data": true})
	return
}
func (*User) DeleteTransaction(c *request.Context) {
	defer func() { // handle transaction error
		if err := recover(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err.(error).Error()})
			return
		}
	}()

	var id uint
	id = 14
	user := models.User{
		ID: &id,
	}
	m.Transaction(func(h *m.Helper) {
		user.SetTX(h.DB()) // important
		if err := h.Delete(&user, false); err != nil {
			panic(err)
		}
	}, 1)

	c.JSON(http.StatusOK, toto.V{"data": true})
	return
}
func (*User) Restore(c *request.Context) {
	var id uint
	id = 14
	modifyUser := models.User{
		ID: &id,
	}

	if err := m.H().Restore(&modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, toto.V{"data": true})
	return
}
