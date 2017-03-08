#
# Makefile
#
# Copyright (c) 2017 Junpei Kawamoto
#
# This file is part of sss.
#
# sss is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# sss is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with sss.  If not, see <http://www.gnu.org/licenses/>.
#
VERSION = snapshot
SUBDIRS := client server
default: build
.PHONY: build release test get-deps proto docker $(SUBDIRS)

build: $(SUBDIRS)

$(SUBDIRS):
	$(MAKE) -C $@

release:
	ghr  -u itslab-kyushu  v$(VERSION) pkg/$(VERSION)

test: get-deps
	go test -v ./...

get-deps:
	for dir in $(SUBDIRS); do $(MAKE) -C "$$dir" get-deps; done

proto:
	protoc --go_out=plugins=grpc:. kvs/kvs.proto

docker:
	cd server && GOOS=linux GOARCH=amd64 go build -o sss-server
	docker build -t itslabq/sss -f dockerfile/Dockerfile .
	rm server/sss-server
