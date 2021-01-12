version := $(shell /bin/date "+%Y-%m-%d %H:%M")

build:
	go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" mctl.go
	$(if $(shell command -v upx), upx mctl)
mac:
	GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o mctl-darwin mctl.go
	$(if $(shell command -v upx), upx mctl-darwin)
win:
	GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o mctl.exe mctl.go
	$(if $(shell command -v upx), upx mctl.exe)
linux:
	GOOS=linux go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o mctl-linux mctl.go
	$(if $(shell command -v upx), upx mctl-linux)