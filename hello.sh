#!/bin/bash

NAME=World

if [[ $# -ge 1 ]]; then
    NAME=$@
fi

echo Hello $NAME