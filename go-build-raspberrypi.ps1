$env:GOOS="linux"
$env:GOARCH="arm"
$env:GOARM="6"

go build -ldflags "-X main.version=v1.1.4 -X main.revision=$(git rev-parse --short HEAD)" .
