#!/bin/bash

if ! [ -e "$2" ]; then
    echo "ファイルが存在しません。"
fi

mysql -h 192.168.10.30 -u root -p $1 < $2
