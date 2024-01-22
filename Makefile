build:
	go clean -x
	CGO_CFLAGS="-I/usr/local/include" \
	LD_LIBRARY_PATH="/usr/local/lib" \
	CGO_LDFLAGS="-L/usr/local/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdactor -ltdapi -ltddb -ltdsqlite -ltdnet -ltdutils -lstdc++ -lssl -lcrypto -ldl -lz -lm" \
	GOOS=linux GOARCH=amd64 \
	go build -v \
	-ldflags "-w -s" \
	-tags netgo \
	-gcflags="-S -m" \
	-trimpath -mod=readonly -buildmode=pie \
	-a -installsuffix cgo -o y-z-a .


.PHONY: build