#
# Makefile
#
# Copyright (c) 2016 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
VERSION = snapshot

default: build

.PHONY: asset
asset:
	go-bindata -pkg command -o command/assets.go assets

.PHONY: build
build: asset
	goxc -os="darwin linux windows" -d=pkg -pv=$(VERSION)

.PHONY: test
test: asset
	go test -v ./...

.PHONY: get-deps
get-deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get github.com/ttacon/chalk
	go get github.com/urfave/cli
	go get github.com/tcnksm/go-gitconfig
	go get gopkg.in/yaml.v2
