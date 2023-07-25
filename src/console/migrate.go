package console

import (
	"cake-store/src/config"
	"cake-store/src/database"

	"github.com/pressly/goose/v3"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run migrate database",
	Long:  "Start migrate database",
	Run:   migration,
}

func init() {
	migrateCmd.PersistentFlags().String("direction", "up", "migration direction up/down")
	RootCmd.AddCommand(migrateCmd)
}

func migration(cmd *cobra.Command, args []string) {
	direction := cmd.Flag("direction").Value.String()

	err := goose.SetDialect("mysql")
	if err != nil {
		log.Error(err)
	}
	goose.SetTableName("schema_migrations")
	sql := database.NewDB()
	if err != nil {
		log.WithField("DatabaseDSN", config.DBDSN()).Fatal("Failed to connect database: ", err)
	}
	defer sql.Close()

	var dir string = "./database/migration"
	if direction == "up" {
		err = goose.Up(sql, dir)
	} else {
		err = goose.Down(sql, dir)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"direction": direction}).
			Fatal("Failed to migrate database: ", err)
	}

	log.WithFields(log.Fields{
		"direction": direction,
	}).Info("Success applied migrations!")

}
