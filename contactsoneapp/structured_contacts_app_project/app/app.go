package app

import (
	"contactsoneapp/components/log"
	"contactsoneapp/repository"
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type ModuleConfig interface {
	TableMigration(wg *sync.WaitGroup)
}

type Controller interface {
	RegisterRoutes(router *mux.Router)
}

type App struct {
	sync.Mutex
	Name       string
	Router     *mux.Router
	DB         *gorm.DB
	Log        log.Log
	Server     *http.Server
	WG         *sync.WaitGroup
	Repository repository.Repository
}

func NewApp(name string, db *gorm.DB, log log.Log, wg *sync.WaitGroup, repo repository.Repository) *App {
	return &App{
		Name:       name,
		DB:         db,
		Log:        log,
		WG:         wg,
		Repository: repo,
	}

}

func NewDBConnection(log log.Log) *gorm.DB {
	url := fmt.Sprintf("%s:%s@tcp(localhost:3306)/contactapp?charset=utf8mb4&parseTime=true", ataa, btaa)

	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(3 * time.Minute)

	db.LogMode(true)

	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci")

	db.BlockGlobalUpdate(true)
	return db

}

func (app *App) Init() {
	app.initializeRouter()
	app.initializeServer()
}

func (app *App) initializeRouter() {
	app.Log.Print(app.Name + " App Route initializing")
	app.Router = mux.NewRouter().StrictSlash(true)
	app.Router = app.Router.PathPrefix("/api/v1/contactapp").Subrouter()
}
func (app *App) initializeServer() {
	headers := handlers.AllowedHeaders([]string{
		"Content-Type", "X-Total-Count", "token",
	})
	methods := handlers.AllowedMethods([]string{
		http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete, http.MethodOptions,
	})
	originOption := handlers.AllowedOriginValidator(app.CheckOrigin)

	app.Server = &http.Server{
		Addr:         "0.0.0.0:4000",
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(headers, methods, originOption)(app.Router),
	}
	app.Log.Print("Server Exposed on 4000")
}
func (app *App) CheckOrigin(origin string) bool {
	return true
}

func (app *App) StartServer() error {
	app.Log.Print("Server Time: ", time.Now())
	app.Log.Print("Server Running on port:4000")
	if err := app.Server.ListenAndServe(); err != nil {
		app.Log.Print("Listen and serve error: ", err)
		return err
	}
	return nil
}
func (app *App) RegisterControllerRoutes(controllers []Controller) {
	app.Lock()
	defer app.Unlock()
	for _, controller := range controllers {
		controller.RegisterRoutes(app.Router.NewRoute().Subrouter())
	}
}

func (app *App) MigrateTables(configs []ModuleConfig) {
	// app.WG.Add(len(configs))
	for _, config := range configs {
		config.TableMigration(app.WG)
		// app.WG.Done()
	}
	// app.WG.Wait()
	app.Log.Print("End of Migration")
}

func (app *App) Stop() {
	context, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	app.DB.Close()
	app.Log.Print("DB closed")

	err := app.Server.Shutdown(context)
	if err != nil {
		app.Log.Print("Fail to stop server")
		return

	}
	app.Log.Print("Server shutdown gracefully.")
}

var ataa string = "root"
var btaa string = "Test1234"
