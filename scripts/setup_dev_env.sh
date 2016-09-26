#!/bin/bash

if [ "$(which docker)" == "" ]; then
	echo "You need docker to hack on tessera, please install before running this script"
	exit 1
fi

if [ "$(which go)" == "" ]; then
	echo "You need go to hack on tessera, please install before running this script."
	exit 1
fi

if [ "$(which npm)" == "" ]; then
	echo "You need npm if you want to work on the frontend."
fi


echo "Starting database..."
docker run --name tessera-dev -e POSTGRES_PASSWORD=tessera -e POSTGRES_DB=tessera_dev -p 5432:5432 -d postgres

sleep 5
echo "Creating database..."
go run seeds.go
