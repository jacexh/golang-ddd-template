package option

type (
	// LoggerOption 日志配置模块
	LoggerOption struct {
		Level      string `default:"info"`
		Name       string
		Filename   string
		MaxSize    int  `default:"100" yaml:"max_size,omitempty" json:"max_size,omitempty"`
		MaxAge     int  `default:"7" yaml:"max_age,omitempty" json:"max_age,omitempty"`
		MaxBackups int  `default:"30" yaml:"max_backups,omitempty" json:"max_backups,omitempty"`
		LocalTime  bool `default:"true" yaml:"local_time,omitempty" json:"local_time,omitempty"`
		Compress   bool
	}

	// RouterOption 服务运行时配置
	RouterOption struct {
		Port            int  `default:"8080"`
		MergeLog        bool `default:"false" yaml:"merge_log" json:"merge_log"`
		DumpResponse    bool `default:"false" yaml:"dump_response" json:"dump_response"`
		LogStackIfPanic bool `default:"false" yaml:"log_stack_if_panic" json:"log_stack_if_panic"`
	}

	// DatabaseOption mysql数据库配置
	DatabaseOption struct {
		Host               string `default:"localhost"`
		Port               int    `default:"3306"`
		Username           string
		Password           string
		Database           string
		MaxOpenConnections int `default:"5" yaml:"max_open_connections" json:"max_open_connections"`
		MaxIdleConnections int `default:"3" yaml:"max_idle_connections" json:"max_idle_connections"`
	}

	// Option 配置入口
	Option struct {
		Title       string
		Description string
		Logger      LoggerOption
		Router      RouterOption
		Database    DatabaseOption
	}
)
