run:
	go run .


build_all: build_mac_arm64 build_mac_x86_64 build_linux_arm64 build_linux_x86_64 build_win_arm64 build_win_x86_64
	echo "done"

build_mac_arm64:
	GOOS=darwin GOARCH=arm64 go build -o out.mac.arm .
	
build_mac_x86_64:
	GOOS=darwin GOARCH=amd64 go build -o out.mac.intel .

build_linux_arm64:
	GOOS=linux GOARCH=arm64 go build -o out.linux.arm64 .

build_linux_x86_64:
	GOOS=linux GOARCH=amd64 go build -o out.linux.amd64 .

build_win_arm64:
	GOOS=windows GOARCH=arm64 go build -o out.win.arm64.exe .

build_win_x86_64:
	GOOS=windows GOARCH=amd64 go build -o out.win.x86.exe .

