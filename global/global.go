package global

import (
	tymysql "github.com/lw000/gocommon/db/mysql"
	"gjump/config"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

var (
	// ProjectConfig ...
	ProjectConfig *config.JSONConfig
	// DbGamedataSrv ...
	DbGamedataSrv *tymysql.Mysql
)

// config logrus log to local filesystem, with file rotation
func configLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d_%H%M",
		// rotatelogs.WithLinkName(baseLogPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge), // 文件最大保存时间
		// rotatelogs.WithRotationCount(365),  // 最多存365个文件
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)

	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	log.SetReportCaller(true)
	log.AddHook(lfHook)
}

// LoadGlobalConfig 加载工程全局配置
func LoadGlobalConfig() error {
	configLocalFilesystemLogger("log", "tweb", time.Hour*24*365, time.Hour*24)

	var err error
	ProjectConfig, err = config.LoadJSONConfig("conf/conf.json")
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// OpenMysql ...
func OpenMysql(cfg *tymysql.JsonConfig) (*tymysql.Mysql, error) {
	srv := &tymysql.Mysql{}
	if err := srv.OpenWithJsonConfig(cfg); err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("数据库连接成功")
	return srv, nil
}
