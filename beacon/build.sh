#! /bin/bash
echo Building...
# Windows
env GOOS=windows GOARCH=amd64 go build -o beacon_win_64.exe . && echo "Built Windows 64 bit!"
env GOOS=windows GOARCH=386 go build -o beacon_win_32.exe . && echo "Built Windows 32 bit!"
# Linux
env GOOS=linux GOARCH=amd64 go build -o beacon_linux_64 . && echo "Built Linux 64 bit!"
env GOOS=linux GOARCH=386 go build -o beacon_linux_32 . && echo "Built Linux 32 bit!"