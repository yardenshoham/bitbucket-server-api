package bitbucket

//go:generate wget https://dac-static.atlassian.com/server/bitbucket/9.3.swagger.v3.json?_v=1.637.0 -O bitbucket.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml bitbucket.json
//go:generate rm bitbucket.json
