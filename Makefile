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
.PHONY: build release test get-deps proto

build:
	goxc -d=pkg -pv=$(VERSION) -os="linux,darwin,windows,freebsd,openbsd"

release:
	ghr  -u itslab-kyushu  v$(VERSION) pkg/$(VERSION)

test:
	go test -v ./...

get-deps:
	go get -d -t -v .

proto:
	protoc --go_out=plugins=grpc:. kvs/kvs.proto
