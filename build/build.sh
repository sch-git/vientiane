CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./vientiane-api ../http/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./vientiane ../main.go

scp ./vientiane-api root@47.101.37.181:/home/nginx/vientiane-api/
scp ./vientiane root@47.101.37.181:/home/nginx/vientiane/


