package server

import (
	"flag"
	"time"
)

var (
	Options struct {
		Id                     string
		Zone                   string
		ApiHttpAddr            string
		ApiHttpsAddr           string
		Store                  string
		ManagerStore           string
		PidFile                string
		CertFile               string
		KeyFile                string
		LogFile                string
		LogLevel               string
		CrashLogFile           string
		InfluxServer           string
		InfluxDbName           string
		ShowVersion            bool
		Ratelimit              bool
		DisableMetrics         bool
		EnableGzip             bool
		DryRun                 bool
		EnableAccessLog        bool
		EnableHttpPanicRecover bool
		Debug                  bool
		HttpHeaderMaxBytes     int
		LogRotateSize          int
		MaxMsgTagLen           int
		MaxClients             int
		MaxRequestPerConn      int // to make load balancer distribute request even for persistent conn
		ReporterInterval       time.Duration
		ManagerRefresh         time.Duration
		HttpReadTimeout        time.Duration
		HttpWriteTimeout       time.Duration
	}
)

func ParseFlags() {
	flag.StringVar(&Options.Id, "id", "", "kateway id, the id must be unique within a host")
	flag.StringVar(&Options.Zone, "zone", "", "kafka zone name")
	flag.StringVar(&Options.ApiHttpAddr, "http", ":9195", "http bind addr")
	flag.StringVar(&Options.ApiHttpsAddr, "https", "", "https bind addr")
	flag.StringVar(&Options.LogLevel, "level", "trace", "log level")
	flag.StringVar(&Options.LogFile, "log", "stdout", "log file, default stdout")
	flag.StringVar(&Options.CrashLogFile, "crashlog", "", "crash log")
	flag.StringVar(&Options.CertFile, "certfile", "", "cert file path")
	flag.StringVar(&Options.PidFile, "pid", "", "pid file")
	flag.StringVar(&Options.KeyFile, "keyfile", "", "key file path")
	flag.StringVar(&Options.Store, "store", "kafka", "backend store")
	flag.StringVar(&Options.ManagerStore, "mstore", "mysql", "store integration with manager")
	flag.StringVar(&Options.InfluxServer, "influxdbaddr", "", "influxdb server address for the metrics reporter")
	flag.StringVar(&Options.InfluxDbName, "influxdbname", "pubsub", "influxdb db name")
	flag.BoolVar(&Options.ShowVersion, "version", false, "show version and exit")
	flag.BoolVar(&Options.Debug, "debug", false, "enable debug mode")
	flag.BoolVar(&Options.EnableAccessLog, "accesslog", false, "en(dis)able access log")
	flag.BoolVar(&Options.DryRun, "dryrun", false, "dry run mode")
	flag.BoolVar(&Options.EnableGzip, "gzip", true, "enable http response gzip")
	flag.BoolVar(&Options.Ratelimit, "raltelimit", false, "enable rate limit")
	flag.BoolVar(&Options.EnableHttpPanicRecover, "httppanic", true, "enable http handler panic recover")
	flag.BoolVar(&Options.DisableMetrics, "metricsoff", false, "disable metrics reporter")
	flag.IntVar(&Options.HttpHeaderMaxBytes, "maxheader", 4<<10, "http header max size in bytes")
	flag.IntVar(&Options.MaxRequestPerConn, "maxreq", -1, "max request per connection")
	flag.IntVar(&Options.LogRotateSize, "logsize", 10<<30, "max unrotated log file size")
	flag.IntVar(&Options.MaxClients, "maxclient", 100000, "max concurrent connections")
	flag.DurationVar(&Options.HttpReadTimeout, "httprtimeout", time.Minute*5, "http server read timeout")
	flag.DurationVar(&Options.HttpWriteTimeout, "httpwtimeout", time.Minute, "http server write timeout")
	flag.DurationVar(&Options.ReporterInterval, "report", time.Second*10, "reporter flush interval")
	flag.DurationVar(&Options.ManagerRefresh, "manrefresh", time.Minute*5, "manager integration refresh interval")

	flag.Parse()
}

func ValidateFlags() {

}
