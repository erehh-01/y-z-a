build:
	go clean -x
	CGO_CFLAGS="-I/usr/local/include" LD_LIBRARY_PATH="/usr/local/lib" CGO_LDFLAGS="-L/usr/local/lib" \
	GOOS=linux GOARCH=amd64 \
	go build -v \
	-ldflags "-w -s" \
	-tags netgo \
	-gcflags="-S -m" \
	-trimpath -mod=readonly -buildmode=pie \
	-a -installsuffix cgo -o y-z-a .


.PHONY: build