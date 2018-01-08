package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"strings"

	"github.com/ckeyer/diego/api"
	"github.com/ckeyer/logrus"
	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/cobra"
)

var (
	addr                          string
	debug                         bool
	dataDir                       string
	etcdEndpoints                 string
	etcdCert, etcdCertKey, etcdCA string
	etcdUsername, etcdPassword    string

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
	rootCmd.Flags().StringVar(&etcdEndpoints, "etcd-endpoints", "http://127.0.0.1:2379/", "Endpoints for etcd.")
	rootCmd.Flags().StringVar(&etcdCert, "etcd-cert", "", "Cert file path for etcd.")
	rootCmd.Flags().StringVar(&etcdCA, "etcd-ca", "", "CA file path for etcd.")
	rootCmd.Flags().StringVar(&etcdCertKey, "etcd-certkey", "", "CertKey file path for etcd.")
	rootCmd.Flags().StringVar(&etcdUsername, "etcdUsername", "", "Username file path for etcd.")
	rootCmd.Flags().StringVar(&etcdPassword, "etcdPassword", "", "Password file path for etcd.")
}

// runServe start http server.
func runServe(cmd *cobra.Command, args []string) {
	etcdCfg := clientv3.Config{
		Endpoints: strings.Split(etcdEndpoints, ","),
	}
	if etcdCert != "" && etcdCertKey != "" {
		cer, err := tls.LoadX509KeyPair(etcdCert, etcdCertKey)
		if err != nil {
			logrus.Fatalf("start server failed, %s", err)
			return
		}
		tlsCfg := &tls.Config{
			Certificates: []tls.Certificate{cer},
		}

		if etcdCA != "" {
			caPem, err := ioutil.ReadFile(etcdCA)
			if err != nil {
				logrus.Fatalf("start server failed, %s", err)
				return
			}
			caPool := x509.NewCertPool()
			caPool.AppendCertsFromPEM(caPem)
			tlsCfg.RootCAs = caPool
		}
		logrus.Info("use tls transfer.")
		etcdCfg.TLS = tlsCfg
	}
	if etcdUsername != "" || etcdPassword != "" {
		etcdCfg.Username = etcdUsername
		etcdCfg.Password = etcdPassword
	}

	_ = etcdCfg

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
