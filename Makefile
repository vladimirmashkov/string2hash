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

archive:
	mv string2hash_nix_x86 string2hash && zip -9 string2hash_nix_x86.zip string2hash && rm string2hash
	mv string2hash_nix_x64 string2hash && zip -9 string2hash_nix_x64.zip string2hash && rm string2hash
	mv string2hash_nix_arm string2hash && zip -9 string2hash_nix_arm.zip string2hash && rm string2hash
	mv string2hash_osx_x64 string2hash && zip -9 string2hash_osx_x64.zip string2hash && rm string2hash
	mv string2hash_osx_arm string2hash && zip -9 string2hash_osx_arm.zip string2hash && rm string2hash
	mv string2hash_x86.exe string2hash.exe && zip -9 string2hash_x86.exe.zip string2hash.exe && rm string2hash.exe
	mv string2hash_x64.exe string2hash.exe && zip -9 string2hash_x64.exe.zip string2hash.exe && rm string2hash.exe
	mv string2hash_arm.exe string2hash.exe && zip -9 string2hash_arm.exe.zip string2hash.exe && rm string2hash.exe

