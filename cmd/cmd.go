package cmd

import (
	"github.com/ckeyer/diego/api"
	"github.com/ckeyer/diego/global"
	"github.com/ckeyer/diego/storage"
	"github.com/ckeyer/logrus"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

var (
	addr          string
	debug         bool
	dataDir       string
	redisEndpoint string
	redisDB       int

	rootCmd = cobra.Command{
		Use:   "diego",
		Short: "diego 版本发布系统",
		Run:   runServe,
	}
)

func init() {
	rootCmd.Flags().BoolVarP(&global.Debug, "debug", "D", false, "debug level")

	rootCmd.Flags().StringVarP(&addr, "addr", "s", ":8080", "web server listenning address.")
	rootCmd.Flags().StringVarP(&dataDir, "data-dir", "d", "/data", "data storage directory.")
	rootCmd.Flags().StringVar(&redisEndpoint, "redis-endpoint", "127.0.0.1:6379", "redis address.")
	rootCmd.Flags().IntVar(&redisDB, "redis-db", 0, "redis db.")
}

// runServe start http server.
func runServe(cmd *cobra.Command, args []string) {
	redisOpts := []redis.DialOption{
		redis.DialDatabase(redisDB),
	}

	conn, err := redis.Dial("tcp", redisEndpoint, redisOpts...)
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Infof("listenning at %s", addr)
	if err := api.Serve(addr, storage.NewRedisStorager(conn)); err != nil {
		logrus.Error(err)
	}
}

// Execute cmd main
func Execute() {
	rootCmd.Execute()
}
