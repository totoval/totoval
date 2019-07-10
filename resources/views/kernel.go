package views

import (
	"github.com/totoval/framework/request"
	"github.com/totoval/framework/view"
)

func Initialize(r *request.Engine) {
	view.Initialize(r)
}
