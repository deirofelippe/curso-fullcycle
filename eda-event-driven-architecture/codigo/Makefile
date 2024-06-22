build:
	@go build -o ./cmd/walletcore/main ./cmd/walletcore

debug: build
	@dlv debug --headless --api-version=2 --log --log-output=debugger,dap,rpc --listen=127.0.0.1:4000 exec ./cmd/walletcore/main