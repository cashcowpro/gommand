# Provide env variables
-include .local.sh

dependencies:
	go get github.com/golang/dep/cmd/dep
	go get github.com/DATA-DOG/godog/cmd/godog
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega/...
	dep ensure -v

spec:
	ginkgo -v -trace ./...

feature:
	godog --format=pretty --strict
