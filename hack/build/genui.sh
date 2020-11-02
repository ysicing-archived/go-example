#!/usr/bin/env bash

[ -d "ui" ] && (
pushd ui
  yarn build
popd
)