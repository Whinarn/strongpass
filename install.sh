#!/bin/bash

version=$(git describe --abbrev=0)
git_commit=$(git rev-parse --short=10 HEAD)
build_date=$(date "+%Y-%m-%d %T +%Z")
go install -a -ldflags="-X 'github.com/whinarn/strongpass/internal/version.Version=$version' -X 'github.com/whinarn/strongpass/internal/version.GitCommit=$git_commit' -X 'github.com/whinarn/strongpass/internal/version.BuildDate=$build_date'" ./cmd/strongpass/
