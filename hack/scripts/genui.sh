#!/usr/bin/env bash

[ -d "ui" ] && (
pushd ui
  [ ! -d "node_modules" ] && yarn install
  yarn build
popd
)