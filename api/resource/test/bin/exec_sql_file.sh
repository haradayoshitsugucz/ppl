#!/bin/sh
echo `pwd`/../../resource/test/database/$1
mysql -h 127.0.0.1 --port 3307 -uroot test_purple --default-character-set=utf8mb4 --show-warnings < `pwd`/../../resource/test/database/$1
