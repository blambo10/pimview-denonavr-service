#! /bin/bash
touch .env

if [ $# -ne 0 ]; then
  echo "===> Overriding env params with args ..."
  for var in "$@"
  do
    export "$var"
  done
fi

env | grep "REACT_APP" > .env

npm start