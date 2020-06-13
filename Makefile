.DEFAULT_GOAL=help

START_SERVER=go run server/main.go start 8080
START_CLIENT=go run client/main.go connect 8080

# run `help` to see all commands that are supported by this Makefile
help: ## display help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "     \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
run-server:  ## starts the goki-server on port 8080
	$(START_SERVER)
run-client:  ## starts a REPL for the client and connects with port 8080
	$(START_CLIENT)
install-server:  ## `go`-install's the server
	go install server/main.go
install-client:  ## `go`-install's the client
	go install client/main.go
make-server-exec:  ## makes a server executable (gox needs to be there)
	gox -os="darwin" github.com/kolharsam/goki/server
make-client-exec:  ## makes a client executable
	gox -os="darwin" github.com/kolharsam/goki/client
