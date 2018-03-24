package cmd

import (
	"github.com/ckeyer/diego/api"
	"github.com/ckeyer/logrus"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

var (
	addr      string
	debug     bool
	dataDir   string
	redisAddr string
	redisDB   int

	rootCmd = cobra.Command{
		Use:   "diego",
		Short: "diego 版本发布系统",
		Run:   runServe,
	}
)

func init() {
	rootCmd.Flags().StringVarP(&addr, "addr", "s", ":8080", "web server listenning address.")
	rootCmd.Flags().BoolVarP(&debug, "debug", "D", false, "debug level")
	rootCmd.Flags().StringVarP(&dataDir, "data-dir", "d", "/data", "data storage directory.")
	rootCmd.Flags().StringVar(&redisAddr, "redis-addr", "127.0.0.1", "redis address.")
	rootCmd.Flags().IntVar(&redisDB, "redis-db", 0, "redis db.")
}

// runServe start http server.
func runServe(cmd *cobra.Command, args []string) {
	redisOpts := []redis.DialOption{
		redis.DialDatabase(redisDB),
	}
	conn, err := redis.Dial("tcp", redisAddr, redisOpts...)
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Infof("listenning at %s", addr)
	if err := api.Serve(addr); err != nil {
		logrus.Error(err)
	}
}

// Execute cmd main
func Execute() {
	rootCmd.Execute()
}
