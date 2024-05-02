#!/bin/bash

addlicense -f hack/licenses/default.tpl -ignore "web/**" -ignore "**/*.md" -ignore "vendor/**" -ignore "**/*.yml" -ignore "**/*.yaml" -ignore "**/*.sh" ./**
