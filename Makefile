# Makefile to build the command lines and tests in Seele project.
# This Makefile doesn't consider Windows Environment. If you use it in Windows, please be careful.

all: release
api:
	go build -o ./build/dashboard-api ./cmd/api
	@echo "Done dashboard-api building debug"

release:
	go build -ldflags "-s -w" -o ./build/dashboard-api ./cmd/api
	@echo "Done dashboard-api building release"

.PHONY: api release
