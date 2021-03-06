package server

import (
	"database/sql"

	"github.com/elahe-dastan/record-appender/handler"
	"github.com/elahe-dastan/record-appender/store"

	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, db *sql.DB) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				str := store.SQLData{DB: db}
				api := handler.API{Store: str}
				api.Run()
			},
		},
	)
}
