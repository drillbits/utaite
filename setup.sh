#!/bin/sh -eu

echo "Install tools..."

echo "go get -u github.com/golang/lint/golint"
go get -u github.com/golang/lint/golint

echo "go get -u golang.org/x/tools/cmd/goimports"
go get -u golang.org/x/tools/cmd/goimports

echo "go get -u github.com/constabulary/gb/..."
go get -u github.com/constabulary/gb/...

echo "go get -u code.palmstonegames.com/gb-gae"
go get -u code.palmstonegames.com/gb-gae

echo "Resolve dependencies..."

gb vendor restore
