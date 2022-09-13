package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"learn-golang/internal/config"
	"learn-golang/internal/driver"
	"learn-golang/internal/handlers"
	"learn-golang/internal/helpers"
	"learn-golang/internal/models"
	"learn-golang/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer func(SQL *sql.DB) {
		_ = SQL.Close()
	}(db.SQL)
	defer close(app.MailChan)

	fmt.Println("Starting mail listener...")
	listenForMail()

	// msg := models.MailData{
	//     To:      "john@do.ca",
	//     From:    "me@here.com",
	//     Subject: "Some subject",
	//     Content: "",
	// }
	//
	// app.MailChan <- msg

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=Admin123")
	if err != nil {
		log.Fatal("cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
