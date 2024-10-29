package cmd

import (
	"battle-of-kings/internal/repository"
	"battle-of-kings/internal/repository/redis"
	"battle-of-kings/internal/service"
	"battle-of-kings/internal/telegram"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the telegram bot",
	Run:   server,
}

func server(_ *cobra.Command, _ []string) {
	_ = godotenv.Load()

	// set up repositories
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_URL"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to initialize redis client")
	}
	accountRepository := repository.NewAccountRedisRepository(redisClient)

	// set up app
	app := service.NewApp(
		service.NewAccountService(accountRepository),
	)
	tg, err := telegram.NewTelegram(app, os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		logrus.WithError(err).Fatal("couldn't connect to telegram server")
	}
	tg.Start()
}

func init() {
	rootCmd.AddCommand(serveCmd)

}
