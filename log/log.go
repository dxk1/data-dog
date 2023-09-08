/**
 * @Author: tbb
 * @Email: 645381379@qq.com
 * @Date: 2023/9/6 22:57
 */
package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Log zerolog.Logger

// Init ...
func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	Log = log.Logger
}
