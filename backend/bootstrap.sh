#! /bin/bash

if [ -z $RUN_PUBLISHER ]; then
   ./app run sub
else
   ./app run pub
fi
