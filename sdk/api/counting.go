/**
 * Package api
 * @Author: tbb
 * @Date: 2023/9/9 16:15
 */
package api

import (
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"
)

// GlobalCounting ...
var GlobalCounting sync.Map
var OverSecondQps sync.Map

// Counting ...
type Counting struct {
	Now       time.Time
	SecondQps int64
}

// SwitchConfig force pass limit switch config
type SwitchConfig struct {
	ForceIgnore         bool  `yaml:"force_ignore" json:"force_ignore"`
	OverSecondQpsIgnore int64 `yaml:"over_second_qps_ignore" json:"over_second_qps_ignore"`
}

// CountingPass ...
func CountingPass(funcId string, conf *SwitchConfig) bool {
	//强制不通过
	if conf.ForceIgnore {
		return false
	}

	con, loaded := GlobalCounting.LoadOrStore(funcId, &Counting{
		Now:       time.Now(),
		SecondQps: 1,
	})

	c := con.(*Counting)
	if !loaded {
		return false
	}

	now := time.Now()
	if now.Second() <= c.Now.Second() {
		atomic.AddInt64(&c.SecondQps, 1)
		if c.SecondQps <= conf.OverSecondQpsIgnore {
			return true
		} else {
			// 超过的函数信息
			RecordOver(funcId, c)
			return false
		}
	}

	//重入判断
	if now.Second() > c.Now.Second() {
		c.Now = now
		atomic.StoreInt64(&c.SecondQps, 1)
		return true
	}

	atomic.StoreInt64(&c.SecondQps, 1)
	if c.SecondQps <= conf.OverSecondQpsIgnore {
		return true
	} else {
		// 超过的函数信息
		RecordOver(funcId, c)
		return false
	}
}

// GetGlobalCounting ...
func GetGlobalCounting() map[string]*Counting {
	m := make(map[string]*Counting)
	GlobalCounting.Range(func(key, value interface{}) bool {
		m[key.(string)] = value.(*Counting)
		return true
	})

	return m
}

// RecordOver ...
func RecordOver(funcId string, counting *Counting) {
	con, loaded := OverSecondQps.LoadOrStore(funcId, &Counting{
		Now:       counting.Now,
		SecondQps: counting.SecondQps,
	})
	if loaded {
		con.(*Counting).Now = counting.Now
		con.(*Counting).SecondQps = counting.SecondQps
	}
}

// ReportOver ...
func ReportOver() string {
	m := make(map[string]*Counting)
	OverSecondQps.Range(func(key, value interface{}) bool {
		m[key.(string)] = value.(*Counting)
		return true
	})
	result, _ := json.Marshal(m)
	return string(result)
}
