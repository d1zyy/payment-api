#!/usr/bin/env bash

SERVICE_NAME="payment-api"
PROJECT_DIR="$HOME/Project/go/payment-api"
BINARY_PATH="/opt/payment-api/payment-api"
BACKUP_PATH="$BINARY_PATH.previous"
URL="${1:-http://localhost/health}"

if [ -z "$BINARY_PATH" ] || [ -z "$BACKUP_PATH" ]; then 

	[ -z "$BINARY_PATH" ] && echo 'BINARY_PATH is empty' >&2
	[ -z "$BACKUP_PATH" ] && echo 'BACKUP_PATH is empty' >&2
	
	exit 1
fi

if [ ! -d "$PROJECT_DIR" ]; then
	echo 'Empty variable PROJECT_DIR' >&2
	exit 1
fi


rollback_binary() {
	if sudo systemctl stop "$SERVICE_NAME"; then
		echo 'Successfully stop programm'
	else
		echo 'Failed to stop programm' >&2
		exit 1
	fi

	if sudo cp -p "$BACKUP_PATH" "$BINARY_PATH"; then
		echo 'Old binary restored!'
	else
		echo 'Failed to restored old binary' >&2
		exit 1
	fi

	if sudo systemctl start "$SERVICE_NAME"; then
		echo 'Successfully start program!'
	else
		echo 'Failed to start program' >&2
		exit 1
	fi

	if rollback_micro; then
	    echo 'Old version restored and healthy'
	else
	    echo 'Old version health check failed' >&2
	    exit 1
	fi
}

rollback_micro() {
	for attempt in {1..5}; do 
		echo "Health check: $attempt"

		if curl --fail --silent --show-error --max-time 5 -o /dev/null "$URL"; then
			return 0
		fi

		if [ "$attempt" -lt 5 ]; then
			sleep 1
		fi
	done

	return 1 
} 

if cd "$PROJECT_DIR"; then
	echo 'Found current directory'

else 
	echo 'Not found current director' >&2
	exit 1
fi

if [ ! -f "$BINARY_PATH" ]; then 
	echo "Installed binary not found: $BINARY_PATH" >&2
	exit 1
fi

if [ "$#" -gt 1 ]; then
	echo "Usage: $0 [health-url]" >&2
	exit 1
fi

if go build -o "$SERVICE_NAME" ./cmd; then
	echo 'Complete build'

else 
	echo 'Fail complete' >&2
	exit 1
fi



if sudo cp -p "$BINARY_PATH" "$BACKUP_PATH"; then
	echo 'Copy complete!'

else 
	echo 'Copy failed' >&2
	exit 1
fi



if sudo systemctl stop "$SERVICE_NAME"; then
	echo 'Successfully stop old!'

else 
	echo 'Failed to stop and install!' >&2
	exit 1
fi


if sudo install -o root -g root -m 755 ./"$SERVICE_NAME" /opt/"$SERVICE_NAME"/"$SERVICE_NAME"; then
	echo 'Successfully install'
	
	if sudo systemctl start "$SERVICE_NAME"; then
		echo 'Command start program - successfully!'

	else 
		echo 'Failed to start new program' >&2
		rollback_binary
		exit 1
	fi

else 
	echo 'Failed to install new' >&2
	rollback_binary
	exit 1
fi

if rollback_micro; then
	echo 'New version healthly'
	exit 0

else 
	echo 'New version not health' >&2
	rollback_binary
	exit 1
fi


