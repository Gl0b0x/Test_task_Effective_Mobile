package app

import (
	"database/sql"

	_ "github.com/lib/pq" // need
	"github.com/sirupsen/logrus"

	server "EffectiveMobile"
	"EffectiveMobile/config"
	"EffectiveMobile/internal/handler"
	"EffectiveMobile/internal/repository"
	"EffectiveMobile/internal/service"
	"EffectiveMobile/pkg/postgres"
)

// Run application
func Run(cfg *config.Config) {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetLevel(logrus.DebugLevel)
	l.Debugf("Starting connection to database")
	l.Debug("DB_URL", cfg.DB.URL)
	db, err := postgres.New(cfg.DB.URL, cfg.DB.Driver)
	if err != nil {
		l.Fatalf("error connecting to database: %v", err)
	}
	l.Debugf("Successful connection to database")
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			l.Errorf("error close db: %v", err)
		}
	}(db)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, l)
	srv := new(server.Server)
	l.WithFields(logrus.Fields{
		"port": cfg.HTTP.Port,
	}).Debug("Starting server on")
	if err = srv.Run(cfg.HTTP.Port, handlers.InitRoutes()); err != nil {
		l.Fatalf("error server: %v", err)
	}
}
