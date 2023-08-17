```sh
docker run --rm --platform linux/amd64 -v "$PWD":/usr/src/argument_checker -w /usr/src/argument_checker amd64/golang:1.20 /bin/sh -c "go mod tidy && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3 && GOOS="linux" GOARCH="amd64" go build -buildmode=plugin -o ./argument_checker.so ./plugin/main.go"

docker run --rm --platform linux/amd64 -v "$PWD":/usr/src/argument_checker -w /usr/src/argument_checker arm64v8/golang:1.20 /bin/sh -c "go mod tidy && go mod download && GOOS="linux" GOARCH="arm64" go build -buildmode=plugin -o ./argument_checker.so ./plugin/main.go"

go build -o '../github-actions-sample/.plugins/argument_checker.so' -buildmode=plugin plugin/main.go

go build -o '../github-actions-sample/.plugins/argument_checker.so' -buildmode=plugin plugin/main.go

# コンテナ起動
docker run -it -v "$PWD":/usr/src/argument_checker -w /usr/src/argument_checker arm64v8/golang:1.20
docker run -it --platform linux/arm64 -v "$PWD":/usr/src/argument_checker -w /usr/src/argument_checker arm64v8/golang:1.2
docker run -it --platform linux/amd64 -v "$PWD":/usr/src/argument_checker -w /usr/src/argument_checker amd64/golang:1.20


/home/runner/work/github-actions-sample/github-actions-sample/.plugins/argument_checker.so: ELF 64-bit LSB shared object, x86-64, version 1 (SYSV), dynamically linked, BuildID[sha1]=d756e94b5a3dc83b86b7d19f43adc333efe97adc, with debug_info, not stripped

.plugins/argument_checker.so: Mach-O 64-bit dynamically linked shared library arm64
```
