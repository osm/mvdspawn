BIN ?= mvdspawn
VER ?= 1.0.0

${BIN}: $(wildcard *.go)
	go build -o $@

lint:
	golangci-lint run .

.PHONY: clean
clean:
	rm -rf ${BIN} ${BIN}-* *.zip

release:
	rm -rf mvdspawn
	mkdir mvdspawn
	go build -o mvdspawn/${BIN}
	cp README.md mvdspawn
	zip -X mvdspawn-${GOOS}-${GOARCH}-${VER}.zip mvdspawn/*
	rm -rf mvdspawn

releases: clean \
	release-darwin-amd64 \
	release-darwin-arm64 \
	release-linux-amd64 \
	release-linux-arm64 \
	release-windows-amd64 \
	release-windows-arm64

release-darwin-amd64:
	$(MAKE) GOOS=darwin GOARCH=amd64 release

release-darwin-arm64:
	$(MAKE) GOOS=darwin GOARCH=arm64 release

release-linux-amd64:
	$(MAKE) GOOS=linux GOARCH=amd64 release

release-linux-arm64:
	$(MAKE) GOOS=linux GOARCH=arm64 release

release-windows-amd64:
	$(MAKE) GOOS=windows GOARCH=amd64 BIN=${BIN}.exe release

release-windows-arm64:
	$(MAKE) GOOS=windows GOARCH=arm64 BIN=${BIN}.exe release
