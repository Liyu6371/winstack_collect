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
	WinStackTask winstack.WinStack `yaml:"win_stack"`
	Report       Report            `yaml:"reports"`
	Logger       Logger            `yaml:"logger"`
}

type Report struct {
	SocketPath    string `yaml:"socket_path"`
	ClusterDataId int32  `yaml:"cluster_metric_data_id"`
	HostDataId    int32  `yaml:"host_metric_data_id"`
	StorageDataId int32  `yaml:"storage_metric_data_id"`
	VMDataId      int32  `yaml:"vm_metric_data_id"`
}

type Logger struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
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
