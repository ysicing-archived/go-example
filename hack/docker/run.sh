#!/usr/bin/env bash

if [ "$1" = "bash" ]; then
  exec /bin/bash
fi

exec /opt/go/go-example server "$@"