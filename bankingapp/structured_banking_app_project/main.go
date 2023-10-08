package main

import (
	"bankingapp/app"
	"bankingapp/components/log"
	"bankingapp/modules"
	"bankingapp/repository"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	log := log.GetLogger()
	db := app.NewDBConnection(*log)
	if db == nil {
		log.Print("DB connection failed")
	}
	defer func() {
		db.Close()
		log.Print("DB closed")
	}()

	var wg sync.WaitGroup
	var repository = repository.NewGormRepository()
	app := app.NewApp("Banking App", db, *log, &wg, repository)
	fmt.Println(app)
	app.Init()
	modules.RegisterModuleRoutes(app)

	go func() {
		err := app.StartServer()
		if err != nil {
			stopApp(app)
		}
	}()
	modules.Configure(app)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	stopApp(app)
	fmt.Print("All Working perfectly good")

}

func stopApp(app *app.App) {
	app.Stop()
	app.WG.Wait()
	log.GetLogger().Print("App stopped.")
	os.Exit(0)
}

// package main

// import (
// 	"bankingapp/components/guru_account"
// 	"bankingapp/components/guru_bank"
// 	guru_customer "bankingapp/components/guru_customer"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/handlers"
// 	"github.com/gorilla/mux"
// )

// func main() {
// 	fmt.Println("Main Called....")
// 	// var abbr string = guru_bank.GetAbbreviation("Bank of India")
// 	// fmt.Println(abbr)
// 	HandleMyRoutes()
// }

// func HandleMyRoutes() {
// 	headerOK := handlers.AllowCredentials()
// 	originsOK := handlers.AllowedOrigins([]string{"*"})
// 	methodsOK := handlers.AllowedMethods([]string{"Get", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"})

// 	mainRouter := mux.NewRouter().StrictSlash(true)
// 	subRouter := mainRouter.PathPrefix("/api/v1/banking-app").Subrouter()
// 	customerRouter := guru_customer.HandleRouter(subRouter)
// 	guru_bank.HandleRouter(subRouter)
// 	guru_account.HandleRouter(customerRouter)

// 	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(headerOK, originsOK, methodsOK)(mainRouter)))
// }
