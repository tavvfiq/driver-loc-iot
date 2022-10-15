build:
	@echo "building binary..."
	@go build -o ./build/driver-loc-iot main.go

clean:
	@echo "cleaning build.."
	@rm -rf ./build