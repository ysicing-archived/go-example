#!/usr/bin/env bash

which statik || go get -u github.com/rakyll/statik
statik -p assets -ns gexe -src=./ui/dist --dest .