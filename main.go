package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"

	"github.com/totoval/framework/cache"
	c "github.com/totoval/framework/config"
	"github.com/totoval/framework/database"
	"github.com/totoval/framework/graceful"
	"github.com/totoval/framework/helpers/log"
	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/helpers/zone"
	"github.com/totoval/framework/http/middleware"
	"github.com/totoval/framework/logs"
	"github.com/totoval/framework/queue"

	"totoval/app/events"
	"totoval/app/jobs"
	"totoval/app/listeners"
	"totoval/config"
	"totoval/resources/lang"
	"totoval/resources/views"
	"totoval/routes"
)

func init() {
	config.Initialize()
	logs.Initialize()
	zone.Initialize()
	cache.Initialize()
	database.Initialize()
	m.Initialize()
	lang.Initialize() // an translation must contains resources/lang/xx.json file (then a resources/lang/validation_translator/xx.go)
	queue.Initialize()
	jobs.Initialize()
	events.Initialize()
	listeners.Initialize()
}

// @caution cannot use config methods to get config in init function
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
		ReadTimeout:    zone.Duration(c.GetInt64("app.read_timeout_seconds")) * zone.Second,
		WriteTimeout:   zone.Duration(c.GetInt64("app.write_timeout_seconds")) * zone.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//j := &jobs.ExampleJob{}
	//j.SetParam(&pbs.ExampleJob{Query: "test", PageNumber: 111, ResultPerPage: 222})
	////j.SetDelay(5 * zone.Second)
	//err := job.Dispatch(j)
	//fmt.Println(err)

	//go hub.On("add-user-affiliation")  // go run artisan.go queue:listen add-user-affiliation

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer graceful.ShutDown(ctx)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", logs.Field{"error": err})
	}

	log.Info("Server exiting")
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
