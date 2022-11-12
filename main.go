package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
	"goapi/database"
	"goapi/settings"
	"log"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New),
		fx.Invoke(
			func(s *settings.Settings) {
				log.Println(s)
			},
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * from USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)
	app.Run()
}
