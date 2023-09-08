/**
 * @Author: tbb
 * @Email: 645381379@qq.com
 * @Date: 2023/9/5 18:47
 */
package main

import (
	"data-dog/log"
	"git.code.oa.com/tencentcloud-serverless/scf_common/recov"
)

func main() {

	recov.RegisterGlobalHandler(func(panicError interface{}) {
		stack := recov.DumpStacktrace(3, panicError)
		log.Log.Fatal().Msgf("panic, stack: %s", stack)
	})

	//配置初始化

}
