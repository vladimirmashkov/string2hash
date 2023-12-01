compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux		GOARCH=386		${GOROOT}/bin/go build -o bin/string2hash_nix_x86 cmd/string2hash/main.go
	GOOS=linux		GOARCH=amd64	${GOROOT}/bin/go build -o bin/string2hash_nix_x64 cmd/string2hash/main.go
	GOOS=linux		GOARCH=arm64	${GOROOT}/bin/go build -o bin/string2hash_nix_arm cmd/string2hash/main.go
	GOOS=darwin		GOARCH=amd64	${GOROOT}/bin/go build -o bin/string2hash_osx_x64 cmd/string2hash/main.go
	GOOS=darwin		GOARCH=arm64	${GOROOT}/bin/go build -o bin/string2hash_osx_arm cmd/string2hash/main.go
	GOOS=windows	GOARCH=386		${GOROOT}/bin/go build -o bin/string2hash_x86.exe cmd/string2hash/main.go
	GOOS=windows	GOARCH=amd64	${GOROOT}/bin/go build -o bin/string2hash_x64.exe cmd/string2hash/main.go
	GOOS=windows	GOARCH=arm64	${GOROOT}/bin/go build -o bin/string2hash_arm.exe cmd/string2hash/main.go