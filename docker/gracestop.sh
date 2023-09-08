echo "Function Stopping..."
BIN="/usr/local/services/data_dog/DataDog"

# 等待客户端北极星后端缓存失效
sleep 60

PROCESS_PID=`pidof ${BIN}`
echo `date`": kill pid:${PROCESS_PID}"
ret=0
for PID in ${PROCESS_PID}
do
	cmd="kill -s term $PID"
	echo ${cmd}
	eval ${cmd}
	let ret=ret+$?
done
echo `date`": kill ret:$ret"

# 检查进程优雅停止
timer=7200
while [[ true ]]; do
  pidof ${BIN}
	if [[ "$?" == 0 ]]; then
		echo "wait service stop."
		sleep 1
		let timer=timer-1
	else
		echo -e "\033[32m service stopped. \033[0m"
		exit 0
		# 预留5s缓冲用于调试使用过程中观察日志和进程等
		sleep 5
		break
	fi
	echo ${timer}
	if [[ ${timer} -le 0 ]]; then
		echo -e "\033[47;31m can't stop.\033[0m"
		exit 1
	fi
done
