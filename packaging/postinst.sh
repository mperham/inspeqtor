#!/bin/sh

set -e

if [ -x /etc/service/inspeqtor/run ]; then
  sv restart inspeqtor
else
  cat <<"TXT"
 _                            _
(_)_ __  ___ _ __   ___  __ _| |_ ___  _ __
| | '_ \/ __| '_ \ / _ \/ _  | __/ _ \| '__|
| | | | \__ \ |_) |  __/ (_| | || (_) | |
|_|_| |_|___/ .__/ \___|\__, |\__\___/|_|
            |_|            |_|

Please configure your notification settings in /etc/inspeqtor/inspeqtor.conf and
then start Inspeqtor with 'sudo sv start inspeqtor'.
TXT
fi
