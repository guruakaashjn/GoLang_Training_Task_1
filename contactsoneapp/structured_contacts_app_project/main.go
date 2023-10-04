package main

import (
	"contactsoneapp/app"
	"contactsoneapp/components/log"
	"contactsoneapp/modules"
	"contactsoneapp/repository"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// func main() {
// 	fmt.Println("Main Called....")
// 	handleMyRoutes()
// }

// func handleMyRoutes() {
// 	headerOK := handlers.AllowCredentials()
// 	originsOK := handlers.AllowedOrigins([]string{"*"})
// 	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"})

// 	mainRouter := mux.NewRouter().StrictSlash(true)
// 	subRouter := mainRouter.PathPrefix("/api/v1/contacts-one-app").Subrouter()
// 	userRouter := guru_user.HandleRouter(subRouter)
// 	contactRouter := guru_contacts.HandleRouter(userRouter)
// 	guru_contact_details.HandleRouter(contactRouter)
// 	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(headerOK, originsOK, methodsOK)(mainRouter)))
// }

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
	app := app.NewApp("Contact App", db, *log, &wg, repository)
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
	fmt.Print("All Working good")
}

func stopApp(app *app.App) {
	app.Stop()
	app.WG.Wait()
	log.GetLogger().Print("App stopped.")
	os.Exit(0)
}
