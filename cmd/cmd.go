package cmd

import (
	"github.com/ckeyer/diego/api"
	"github.com/ckeyer/logrus"
	"github.com/spf13/cobra"
)

var (
	addr  string
	debug bool

	rootCmd = cobra.Command{
		Use:   "diego",
		Short: "diego 版本发布系统",
		Run:   runServe,
	}
)

func init() {
	rootCmd.Flags().StringVarP(&addr, "addr", "s", ":8080", "web server listenning address.")
	rootCmd.Flags().BoolVarP(&debug, "debug", "D", false, "debug level")
}

// runServe start http server.
func runServe(cmd *cobra.Command, args []string) {
	logrus.Infof("listenning at %s", addr)
	err := api.Serve(addr)
	if err != nil {
		logrus.Error(err)
	}
}

// Execute cmd main
func Execute() {
	rootCmd.Execute()
}
