/**
 * @Author: tbb
 * @Email: 645381379@qq.com
 * @Date: 2023/9/6 23:11
 */
package conf

import (
	"fmt"
	"git.code.oa.com/tencentcloud-serverless/scf_common/hotswap"
	"github.com/dxk1/data-dog/log"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

// Global global config
var Global *Config

// RainBowConfig ...
type RainBowConfig struct {
	Conf *Config `yaml:"config"`
}

// Config is the startup config
type Config struct {
	ListenPort int `yaml:"listen_port"`
}

// Init init config
func Init() (*Config, error) {
	log.Init()
	appID := os.Getenv("CONF_APPID")
	group := os.Getenv("CONF_GROUP")
	envName := os.Getenv("CONF_ENV_NAME")
	if appID == "" || group == "" || envName == "" {
		return nil, fmt.Errorf("config appID:%s, group:%s, envName:%s has empty value", appID, group, envName)
	}
	opts := &hotswap.Option{
		ConnectStr: "http://api.rainbow.woa.com:8080",
		AppID:      appID,
		Group:      group,
		EnvName:    envName,
	}
	watcher, err := hotswap.NewRainBowOperator(opts, &RainBowConfig{})
	if err != nil {
		return nil, err
	}
	watcher.AppendWatcher(&listener{})
	config, err := watcher.Init()
	if err != nil {
		return nil, err
	}
	Global = preHandle(config.(*RainBowConfig).Conf)
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
	if config.ListenPort == 0 {
		config.ListenPort = 10000
	}

	out, _ := yaml.Marshal(config)
	log.Log.Info().Msgf("load conf:%s", string(out))
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
