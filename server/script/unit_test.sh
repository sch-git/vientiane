go test -v -json -cover -coverprofile cover.out ../... -coverpkg ../... ../...

go tool cover -html=./cover.out -o coverage.html

go tool cover -func ./cover.out