package config

import (
	"fmt"
	"os"
	"path/filepath"
	"winstack_collect/collector/winstack"

	"gopkg.in/yaml.v3"
)

var globalConfig *Config

type Config struct {
	Logger       Logger            `yaml:"logger"`
	Socket       Socket            `yaml:"socket"`
	WinStackTask winstack.WinStack `yaml:"win_stack"`
}

// Logger 日志相关配置
type Logger struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

// Socket GSE 通讯相关的配置
type Socket struct {
	SocketPath string `yaml:"socket_path"`
}

// ParseConfig 待实现
func ParseConfig(f string) (*Config, error) {
	dir, _ := os.Getwd()
	confPath := filepath.Join(dir, f)
	// 读取不到配置文件的情况下直接退出程序
	content, err := os.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file, path: %s, error: %s\n", confPath, err)
	}
	// 解析错误的情况下也直接退出
	err = yaml.Unmarshal(content, &globalConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config file, error: %s\n", err)
	}
	return globalConfig, nil
}

func GetConfig() *Config {
	return globalConfig
}
