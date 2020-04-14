#!/bin/bash

name=$1
if [ -z "$name" ];then
  echo "Usage: $0 <name>"
  exit 1
fi

find . -type f -name "*.go" | xargs sed -i 's#gin-template#'$name'#g'
mv gin-template.conf $name.conf