#!/usr/bin/env bash
#Include shell font styles and some basic information
source ./style_info.cfg
source ./path_info.cfg
ulimit -n 200000

#Check if the service exists
#If it is exists,kill this process
check=$(ps aux | grep -w ./${bot_name} | grep -v grep | wc -l)
if [ $check -ge 1 ]; then
  oldPid=$(ps aux | grep -w ./${bot_name} | grep -v grep | awk '{print $2}')
    kill -9 ${oldPid}
fi
#Waiting port recycling
sleep 1
cd ${bot_binary_root}
nohup ./${bot_name}  >>../logs/bot.log 2>&1 &

#Check launched service process
sleep 3
check=$(ps aux | grep -w ./${bot_name} | grep -v grep | wc -l)
if [ $check -ge 1 ]; then
  allNewPid=$(ps aux | grep -w ./${bot_name} | grep -v grep | awk '{print $2}')
  echo -e ${SKY_BLUE_PREFIX}"SERVICE START SUCCESS"${COLOR_SUFFIX}
  echo -e ${SKY_BLUE_PREFIX}"SERVICE_NAME: "${COLOR_SUFFIX}${YELLOW_PREFIX}${bot_name}${COLOR_SUFFIX}
  echo -e ${SKY_BLUE_PREFIX}"PID: "${COLOR_SUFFIX}${YELLOW_PREFIX}${allNewPid}${COLOR_SUFFIX}
  echo -e ${SKY_BLUE_PREFIX}"LISTENING_PORT: "${COLOR_SUFFIX}${YELLOW_PREFIX}${COLOR_SUFFIX}
else
  echo -e ${YELLOW_PREFIX}${bot_name}${COLOR_SUFFIX}${RED_PREFIX}"SERVICE START ERROR, PLEASE CHECK openIM.log"${COLOR_SUFFIX}
fi
