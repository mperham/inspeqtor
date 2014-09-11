#!/bin/sh

set -e

if [ "$1" = "purge" ]; then
    rm -rf /etc/inspeqtor
    rm -rf /var/log/inspeqtor
    rm -f /etc/service/inspeqtor
    rm -rf /etc/sv/inspeqtor
elif [ "$1" = "remove" ]; then
    rm -f /etc/service/inspeqtor
    rm -rf /etc/sv/inspeqtor
fi
