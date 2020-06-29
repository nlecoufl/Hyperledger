#!/bin/bash
 
#Vérification du nombre d'arguments
#
if [[ $# -eq 0 || $# -gt 1 ]]; then
        echo "Usage : $0 [PROCESS_NAME]"
        exit 1
fi
 
TOTALMEM=`free -m|grep "Mem"|awk '{printf $2}'`
PROCESSPID=`pidof $1`
 
if [[ $PROCESSPID -ne "" ]]; then
	MEMORYPID=`ps aux |grep $PROCESSPID |grep -v grep|awk '{printf $4}'`
	REALMEMORY=`echo "${MEMORYPID}*${TOTALMEM}*0.01"|bc`
	echo $REALMEMORY
fi