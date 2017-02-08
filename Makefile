#
# Makefile
#
# Copyright (c) 2017 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
VERSION = snapshot

default: build

.PHONY: build
build:
	goxc -d=pkg -pv=$(VERSION)

.PHONY: release
release:
	ghr  -u jkawamoto  v$(VERSION) pkg/$(VERSION)

.PHONY: test
test: 
	go test -v ./...

.PHONY: get-deps
get-deps:
	go get -d -t -v .
