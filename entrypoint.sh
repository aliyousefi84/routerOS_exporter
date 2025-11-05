#!bin/bash

if [ -z "$ROUTEROS_USER" ] && [ -z "$ROUTEROS_PASSWORD" ] && [ -z "$ROUTEROS_ADDREESS" ];
then 
	echo "environment variable error"
	exit 1
fi

if [ ! -f /usr/local/bin/app ];
then
	echo "not exist"
	exit 1
else
	exec /usr/local/bin/app
fi



