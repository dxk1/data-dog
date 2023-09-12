/**
 * @Author: tbb
 * @Email: 645381379@qq.com
 * @Date: 2023/9/5 18:47
 */
package main

import (
	"git.code.oa.com/tencentcloud-serverless/scf_common/recov"
	"github.com/dxk1/data-dog/conf"
	"github.com/dxk1/data-dog/log"
	"github.com/dxk1/data-dog/server"
	"os"
)

func main() {

	recov.RegisterGlobalHandler(func(panicError interface{}) {
		stack := recov.DumpStacktrace(3, panicError)
		log.Log.Fatal().Msgf("panic, stack: %s", stack)
	})

	//配置初始化
	_, err := conf.Init()
	if err != nil {
		log.Log.Fatal().Msgf("conf.Init err:%s ", err)
		os.Exit(-1)
	}
	//初始化server
	ser := server.Init()
	err = ser.Run()
	if err != nil {
		log.Log.Fatal().Msgf("server.Run err:%s ", err)
		os.Exit(-1)
	}

	log.Log.Info().Msgf("server start success")
}
