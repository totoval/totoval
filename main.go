package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	c "github.com/totoval/framework/config"
	"github.com/totoval/framework/database"
	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/http/middleware"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"sync"
	"totoval/config"
	"totoval/resources/lang/validation_translator"
	"totoval/routes"
)

func init() {
	config.Initialize()
	database.Initialize()
	m.Initialize()
	validation_translator.Initialize() // an translation must contains resources/lang/xx.json file (then a resources/lang/validation_translator/xx.go)
}

func main() {

	// upgrade gin validator v8 to v9
	binding.Validator = new(defaultValidator)

	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//r.Use(middleware.RequestLogger())

	r.Use(middleware.Locale())

	routes.Register(r)

	r.Run(":" + c.GetString("app.port"))
}

// gin validator v8 to v9
type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
