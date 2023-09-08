/**
 * @Author: tbb
 * @Email: 645381379@qq.com
 * @Date: 2023/9/6 23:11
 */
package conf

import (
	"fmt"
	"git.code.oa.com/tencentcloud-serverless/scf_common/hotswap"
	"gopkg.in/yaml.v3"
	"os"
	scflog "scf-log"
	"sync"
)

// Global global config
var Global *Config

// Config is the startup config
type Config struct {
}

// Init init config
func Init() (*Config, error) {
	appID := os.Getenv("EB_SCHEDULER_APPID")
	group := os.Getenv("EB_SCHEDULER_GROUP")
	envName := os.Getenv("EB_SCHEDULER_ENV_NAME")
	if appID == "" || group == "" || envName == "" {
		return nil, fmt.Errorf("config appID:%s, group:%s, envName:%s has empty value", appID, group, envName)
	}
	opts := &hotswap.Option{
		ConnectStr: "http://api.rainbow.woa.com:8080",
		AppID:      appID,
		Group:      group,
		EnvName:    envName,
	}
	watcher, err := hotswap.NewRainBowOperator(opts, &Config{})
	if err != nil {
		return nil, err
	}
	watcher.AppendWatcher(&listener{})
	config, err := watcher.Init()
	if err != nil {
		return nil, err
	}
	Global = preHandle(config.(*Config))
	return Global, nil
}

type listener struct {
}

// OnEvent is the configChange listener implementation
func (l *listener) OnEvent(data interface{}) {
	config := preHandle(data.(*Config))
	Global = config
	notify(config)
}

func preHandle(config *Config) *Config {
	out, _ := yaml.Marshal(config)
	scflog.Infof("load conf:%s", string(out))
	return config
}

// GetConfig return config
func GetConfig() *Config {
	return Global
}

func notify(config *Config) {
	for _, l := range listeners {
		_ = l.OnEvent(config)
	}
}

// Watcher define the interface of config update watcher
type Watcher interface {
	OnEvent(data *Config) error
}

var listeners []Watcher
var listenerMutex sync.RWMutex

// AppendConfigWatcher append config update listener
func AppendConfigWatcher(listener Watcher) {
	listenerMutex.Lock()
	defer listenerMutex.Unlock()
	listeners = append(listeners, listener)
}
