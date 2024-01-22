@echo off
set ROOT_PATH=%cd%
set BUILD_PATH=%cd%\build
set LIB_PATH=%cd%\lib\runtime\windows\*
set APP_NAME=Y-Z-A.exe

if exist %BUILD_PATH% (
	rmdir /s /q %BUILD_PATH%
)


mkdir  %BUILD_PATH%
xcopy /s /i %LIB_PATH% %BUILD_PATH%\

set GOARCH=amd64
set APP_NAME=Y-Z-A-64.exe
set CGO_CFLAGS=-I%cd%\lib\build\windows\include
set LD_LIBRARY_PATH=-L%cd%\lib\build\windows\bin
set CGO_LDFLAGS=-L%cd%\lib\build\windows\bin -ltdjson
set CGO_ENABLED=1
set GOOS=windows

go build -v -ldflags "-w -s" -tags netgo -gcflags="-S -m" -trimpath -mod=readonly -buildmode=pie -a -installsuffix cgo -o  %BUILD_PATH%\%APP_NAME% .

cd %ROOT_PATH%