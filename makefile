build:
	mkdir -p build
	GOOS=linux GOARCH=amd64 go build -o build/jcalc ./cmd/jcalc
	tar cz -f build/jungle-calculator-linux64.tar.gz build/jcalc
	GOOS=linux GOARCH=arm64 go build -o build/jcalc ./cmd/jcalc
	tar cz -f build/jungle-calculator-linux_arm64.tar.gz build/jcalc
	GOOS=darwin GOARCH=amd64 go build -o build/jcalc ./cmd/jcalc
	tar cz -f build/jungle-calculator-macos64.tar.gz build/jcalc
	GOOS=darwin GOARCH=arm64 go build -o build/jcalc ./cmd/jcalc
	tar cz -f build/jungle-calculator-macos_arm64.tar.gz build/jcalc
	GOOS=windows GOARCH=amd64 go build -o build/jcalc.exe ./cmd/jcalc
	zip -r build/jungle-calculator-windows64.zip build/jcalc.exe
	GOOS=windows GOARCH=arm64 go build -o build/jcalc.exe ./cmd/jcalc
	zip -r build/jungle-calculator-windows_arm64.zip build/jcalc.exe
	cd build && rm jcalc jcalc.exe

testbuild:
	GOOS=linux GOARCH=amd64 go build ./cmd/jcalc
	GOOS=linux GOARCH=arm64 go build ./cmd/jcalc
	GOOS=darwin GOARCH=amd64 go build ./cmd/jcalc
	GOOS=darwin GOARCH=arm64 go build ./cmd/jcalc
	GOOS=windows GOARCH=amd64 go build ./cmd/jcalc
	GOOS=windows GOARCH=arm64 go build ./cmd/jcalc


.PHONY: build testbuild
