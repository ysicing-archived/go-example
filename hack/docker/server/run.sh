#!/usr/bin/env bash

if [ "$1" = "bash" -o "$1" = "sh" ]; then
  exec /bin/bash
fi

exec /opt/go/go-example server "$@"