#!/bin/bash

name=$1
if [ -z "$name" ];then
  echo "Usage: $0 <name>"
  exit 1
fi

currPath=$(cd `dirname $0`;pwd)

find . -type f -name "*.go" | xargs sed -i 's#gin-template#'$name'#g'
sed -i 's#gin-template#'$name'#g' $currPath/go.mod
mv gin-template.conf $name.conf
