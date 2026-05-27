package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	//no quieres las funciones solo quieres los efectos de initializacion de la libraria
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"snippetbox.alexarmenta.net/internal/models"
)

type application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
}

func main() {

	addr := flag.String("addr", "", "HTTP network address")
	dsn := flag.String("dsn", "", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if *addr == "" || *dsn == "" {
		logger.Error("Invalid input options")
		os.Exit(1)
	}

	openDB := func(dsn string) (*sql.DB, error) {
		db, err := sql.Open("mysql", dsn)

		if err != nil {
			return nil, err
		}

		err = db.Ping()

		if err != nil {
			db.Close()
			return nil, err
		}

		return db, nil
	}

	db, err := openDB(*dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	//this is a destructor
	defer db.Close()

	templateCache, err := newTemplateCache()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	app := &application{
		logger:        logger,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder:   formDecoder,
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)

}
