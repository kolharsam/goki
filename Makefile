.DEFAULT_GOAL=help

# run `help` to see all commands that are supported by this Makefile
help: ## display help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "     \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
install-server:  ## `go`-install's the server
	go install server/main.go
install-client:  ## `go`-install's the client
	go install client/main.go
make-server-exec:  ## makes a server executable (gox needs to be there)
	gox -os="darwin" github.com/kolharsam/goki/server
make-client-exec:  ## makes a client executable
	gox -os="darwin" github.com/kolharsam/goki/client
vet-server:  ## run go vet over server code
	go vet ./server
vet-client:  ## run go vet over client code
	go vet ./client
vet-common:  ## run go vet over common dir.
