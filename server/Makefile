build:
	go get github.com/aws/aws-lambda-go/lambda
	env GOOS=linux go build -ldflags="-s -w" -o bin/slides-list api/slides-list/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/slides-get api/slides-get/main.go

domain:
	sls create_domain --aws-profile 1day-release

deploy:
	sls deploy --aws-profile 1day-release
