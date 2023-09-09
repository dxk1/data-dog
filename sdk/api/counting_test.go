/**
 * Package api
 * @Author: tbb
 * @Date: 2023/9/9 16:51
 */
package api

import (
	"data-dog/log"
	"sync/atomic"
	"testing"
	"time"
)

func TestCounting(t *testing.T) {
	log.Init()
	conf := &SwitchConfig{
		ForceIgnore:         false,
		OverSecondQpsIgnore: 500,
	}

	var t1, f1 int64

	for i := 0; i < 1000; i++ {
		go func(i int, t, f int64) {
			pass := CountingPass("test1", conf)
			if pass {
				atomic.AddInt64(&t1, 1)
			} else {
				atomic.AddInt64(&f1, 1)
			}
			//log.Log.Error().Msgf("%d type %v", i, pass)
		}(i, t1, f1)
	}

	time.Sleep(1 * time.Second)
	log.Log.Error().Msgf("t:%d , f:%d", t1, f1)
	atomic.StoreInt64(&t1, 0)
	atomic.StoreInt64(&f1, 0)

	for i := 0; i < 1000; i++ {
		go func(i int, t, f int64) {
			pass := CountingPass("test1", conf)
			if pass {
				atomic.AddInt64(&t1, 1)
			} else {
				atomic.AddInt64(&f1, 1)
			}
			//log.Log.Error().Msgf("%d type %v", i, pass)
		}(i, t1, f1)
	}

	time.Sleep(2 * time.Second)
	log.Log.Error().Msgf("3s after t:%d , f:%d", t1, f1)
}
