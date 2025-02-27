.PHONY: build debug run clean release man install uninstall

build:
	@env CGO_ENABLED=0 go build -o bin/cipher -trimpath .

debug: build
	@dlv exec ./bin/cipher

run: build
	@./bin/cipher

clean:
	@rm -r bin

release:
	@env CGO_ENABLED=0 GOOS="windows" GOARCH="amd64" go build -o bin/cipher_windows_amd64.exe -ldflags="-s -w -extldflags=-static" -trimpath .
	@env CGO_ENABLED=0 GOOS="windows" GOARCH="arm64" go build -o bin/cipher_windows_arm64.exe -ldflags="-s -w -extldflags=-static" -trimpath .
	@env CGO_ENABLED=0 GOOS="darwin" GOARCH="amd64" go build -o bin/cipher_macos_amd64 -ldflags="-s -w -extldflags=-static" -trimpath .
	@env CGO_ENABLED=0 GOOS="darwin" GOARCH="arm64" go build -o bin/cipher_macos_arm64 -ldflags="-s -w -extldflags=-static" -trimpath .
	@env CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -o bin/cipher_linux_amd64 -ldflags="-s -w -extldflags=-static" -trimpath .
	@env CGO_ENABLED=0 GOOS="linux" GOARCH="arm64" go build -o bin/cipher_linux_arm64 -ldflags="-s -w -extldflags=-static" -trimpath .
	@env CGO_ENABLED=0 GOOS="freebsd" GOARCH="amd64" go build -o bin/cipher_freebsd_amd64 -ldflags="-s -w -extldflags=-static" -trimpath .
	@env CGO_ENABLED=0 GOOS="openbsd" GOARCH="amd64" go build -o bin/cipher_openbsd_amd64 -ldflags="-s -w -extldflags=-static" -trimpath .

man:
	@pandoc -s -t man -o cipher.1 MAN.md

install: cipher.1
	@install -d /usr/local/share/man/man1/
	@install -m 644 cipher.1 /usr/local/share/man/man1/
	@gzip -f /usr/local/share/man/man1/cipher.1

uninstall:
	@rm -f /usr/local/share/man/man1/cipher.1.gz