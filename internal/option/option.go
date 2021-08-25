package option

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jacexh/multiconfig"
)

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
		Port    int `default:"8080"`
		Timeout int `default:"30"`
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
		Description string
		Logger      LoggerOption
		Router      RouterOption
		Database    DatabaseOption
	}
)

var (
	configFileType = "yml"
	configName     = "config"
	// searchInPaths 配置文件查找目录
	searchInPaths []string

	// environmentVariablesPrefix 项目环境变量前缀
	environmentVariablesPrefix = "{{.EnvironmentVariablesPrefix}}"
	// environmentVariableProfile 项目profile的环境变量名称
	environmentVariableProfile = environmentVariablesPrefix + "_PROJECT_PROFILE"
)

// SetConfigFileType 配置文件类型
func SetConfigFileType(t string) {
	if t != "" {
		configFileType = t
	}
}

// SetConfigName 配置名称
func SetConfigName(n string) {
	if n != "" {
		configName = n
	}
}

// AddConfigPath 添加配置文件目录
func AddConfigPath(path string) {
	if path == "" {
		return
	}
	if filepath.IsAbs(path) {
		searchInPaths = append(searchInPaths, filepath.Clean(path))
		return
	}
	fp, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	searchInPaths = append(searchInPaths, fp)
}

func HomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func getConfigName() string {
	profile := os.Getenv(environmentVariableProfile)
	if profile == "" {
		return configName
	}
	return fmt.Sprintf("%s_%s", configName, profile)
}

func findInDir(dir string, file string) string {
	fp := filepath.Join(dir, file)
	fi, err := os.Stat(fp)
	if err == nil && !fi.IsDir() {
		return fp
	}
	return ""
}

func findConfigFile() string {
	fp := fmt.Sprintf("%s.%s", getConfigName(), configFileType)
	for _, d := range searchInPaths {
		if p := findInDir(d, fp); p != "" {
			return p
		}
	}
	panic(errors.New("cannot find the config file"))
}

func LoadConfig() *Option {
	f := findConfigFile()
	opt := new(Option)
	loader := multiconfig.NewWithPathAndEnvPrefix(f, environmentVariablesPrefix)
	loader.MustLoad(opt)
	return opt
}

func init() {
	AddConfigPath("./conf")
}
