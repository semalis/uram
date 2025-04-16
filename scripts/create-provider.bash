#!/bin/bash
set -ex
if [ -z "$1" ]; then
	echo "No user supplied"
	exit 1
fi

if [ -z "$2" ]; then
	echo "No bond supplied"
	exit 1
fi

if [ -z "$3" ]; then
	echo "No service supplied"
	exit 1
fi

PWD=$(dirname -- "${BASH_SOURCE[0]}")
BIN="uramd"
BIN_TX="arkeo"
USER="$1"
BOND="$2"
SERVICE="$3"

PUBKEY_RAW=$($BIN keys show "$USER" -p --keyring-backend test | jq -r .key)
PUBKEY=$($BIN debug pubkey-raw "$PUBKEY_RAW" | grep "Bech32 Acc" | awk '{ print $NF }')

./"$PWD"/bond-provider.bash "$USER" "$SERVICE" "$BOND"

$BIN tx $BIN_TX mod-provider -y -b block --from "$USER" --keyring-backend test -- "$PUBKEY" "$SERVICE" "http://localhost:3636/metadata.json" 1 1 10 5256000 10uarkeo 10uarkeo 10
