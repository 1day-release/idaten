build:
	go get github.com/aws/aws-lambda-go/lambda
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/slides-list api/slides-list.go

domain:
	sls create_domain --aws-profile 1day-release

deploy:
	sls deploy --aws-profile 1day-release
