build:
	cd cmd/jcalc && GOOS=linux GOARCH=amd64 go build -o build/jungle-calculator .
	cd cmd/jcalc && tar czf /build/jungle-calculator-linux64.tar.gz build/jungle-calculator
	cd cmd/jcalc && GOOS=linux GOARCH=arm64 go build -o build/jungle-calculator .
	cd cmd/jcalc && tar czf /build/jungle-calculator-linux_arm64.tar.gz build/jungle-calculator
	cd cmd/jcalc && GOOS=darwin GOARCH=amd64 go build -o build/jungle-calculator .
	cd cmd/jcalc && tar czf /build/jungle-calculator-macos64.tar.gz build/jungle-calculator
	cd cmd/jcalc && GOOS=darwin GOARCH=arm64 go build -o build/jungle-calculator .
	cd cmd/jcalc && tar czf /build/jungle-calculator-macos_arm64.tar.gz build/jungle-calculator
	cd cmd/jcalc && GOOS=windows GOARCH=amd64 go build -o build/jungle-calculator .
	cd cmd/jcalc && zip -r /build/jungle-calculator-windows64.zip build/jungle-calculator
	cd cmd/jcalc && GOOS=windows GOARCH=arm64 go build -o build/jungle-calculator .
	cd cmd/jcalc && zip -r /build/jungle-calculator-windows_arm64.zip build/jungle-calculator
	cd cmd/jcalc/build && rm jungle-calculator

.PHONY: build
