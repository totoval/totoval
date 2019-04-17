package main

import (
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"

	"github.com/totoval/framework/cache"

	c "github.com/totoval/framework/config"
	"github.com/totoval/framework/database"
	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/http/middleware"
	"totoval/config"
	"totoval/resources/lang"
	"totoval/routes"
	"totoval/resources/views"
)

func init() {
	config.Initialize()
	cache.Initialize()
	database.Initialize()
	m.Initialize()
	lang.Initialize() // an translation must contains resources/lang/xx.json file (then a resources/lang/validation_translator/xx.go)
}

func main() {

	// upgrade gin validator v8 to v9
	binding.Validator = new(defaultValidator)

	r := gin.Default()

	if c.GetString("app.env") == "production" {
		r.Use(gin.Logger())

		r.Use(gin.Recovery())
	}

	if c.GetBool("app.debug") {
		r.Use(middleware.RequestLogger())
	}

	r.Use(middleware.Locale())

	routes.Register(r)

        views.Initialize(r)

	s := &http.Server{
		Addr:           ":" + c.GetString("app.port"),
		Handler:        r,
		ReadTimeout:    time.Duration(c.GetInt64("app.read_timeout_seconds")) * time.Second,
		WriteTimeout:   time.Duration(c.GetInt64("app.write_timeout_seconds")) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
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
