#!/bin/bash

if [ -z "$1" ]; then
	echo "No user supplied"
	exit 1
fi

BIN="uramd"
TOKEN="uuram"
USER="$1"

ADDRESS=$($BIN keys show "$USER" -a --keyring-backend test)

$BIN query bank balances --denom $TOKEN -o json -- "$ADDRESS"
