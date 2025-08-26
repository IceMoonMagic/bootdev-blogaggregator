#!/usr/bin/env zsh

mkdir -p /run/user/1000/postgresql/
postgres -D .postgresql.conf 2> /dev/null &
