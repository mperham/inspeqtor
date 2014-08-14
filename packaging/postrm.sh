#!/bin/sh

set -e

if [ "$1" = "purge" ]; then
    rm -rf /etc/inspeqtor
    rm -rf /var/log/inspeqtor
fi

rm -f /etc/service/inspeqtor
rm -rf /etc/sv/inspeqtor
