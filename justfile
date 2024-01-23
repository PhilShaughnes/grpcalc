# https://cheatography.com/linux-china/cheat-sheets/justfile/

set dotenv-load

cov := `mktemp`

default:
	@just --list

# using this: https://grpc.io/docs/languages/go/quickstart/

# make go proto gen stuff available - only needed if GOBIN isn't set
_pexport:
	export PATH="$PATH:$(go env GOBIN)"
	# export PATH="$PATH:$(go env GOPATH)/bin"

# rerun go server when any go file changes
watch:
	ls **/*.go | entr -rc go run ./...

# run grpcui - assumes app is running
ui:
	grpcui --plaintext localhost:8080

# generate proto grpc and serialization code
pgen PROTOFILE: _pexport
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		{{ PROTOFILE }}

# generate proto grpc and serialization code
gen: _pexport
	protoc --go_out=. \
		--go-grpc_out=. \
		--proto_path=proto \
		proto/*.proto

# # run tests
# test FILES:
# 	go test -race {{ FILES }}
#
# _testcov FILES:
# 	go test -coverprofile={{ cov }} {{ FILES }}
#
# # show test coverage as cli %
# cover FILES: (_testcov FILES) && clean
# 	go tool cover -func={{ cov }}
#
# # show test coverage in browser view
# cover-web FILES: (_testcov FILES) && clean
# 	go tool cover -html={{ cov }}
#
# # clear coverage files
# clean:
# 	unlink {{ cov }}
#
# # stress test
# #
