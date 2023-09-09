/**
 * Package log
 * @Author: tbb
 * @Date: 2023/9/9 15:10
 */
package log

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	Init()
	Log.Error().Msgf("test 111:%s", "1")

	time.Sleep(time.Second)
}
