build:
	go get github.com/aws/aws-lambda-go/lambda
	env GOOS=linux go build -ldflags="-s -w" -o bin/slides-list api/slides-list/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/slides-get api/slides-get/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/slides-create api/slides-create/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/slides-delete api/slides-delete/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/slides-save api/slides-save/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-login api/users-login/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-logout api/users-logout/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-authorization api/users-authorization/main.go

domain:
	sls create_domain --aws-profile 1day-release

deploy:
	sls deploy --aws-profile 1day-release
