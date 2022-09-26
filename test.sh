# This script performs the test using ginkgo
mkdir -p .coverage
go test -v $1 -cover -coverprofile=.coverage/coverage.out
go tool cover -html=.coverage/coverage.out -o .coverage/coverage.html
open .coverage/coverage.html