.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build: deps vet fmt

.PHONY: update-deps
update-deps:
	go get github.com/plantoncloud/planton-cloud-apis@latest
	go get github.com/plantoncloud/stack-job-runner-golang-sdk
	go get github.com/plantoncloud/pulumi-module-golang-commons
