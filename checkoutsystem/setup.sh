#!/bin/bash

if ! [ -x "$(command -v docker)" ]; then
  echo 'Error: Docker engine not installed' >&2
  exit 1
fi

#Build docker image
sudo docker build -t market checkoutsystem/Market_Docker/.
sudo docker build -t mongo checkoutsystem/Mongo_Docker/.

#Run container from the image
sudo docker run --name mongod -d mongo
