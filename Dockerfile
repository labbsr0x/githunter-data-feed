FROM golang:alpine
# create a working directory
WORKDIR /go/src/app
# add source code
ADD . .

ENV GH_GIT_GRAPHQL_URL=https://api.github.com/graphql
ENV GH_LOG_LEVEL=warn

# run main.go
CMD ["go", "run", "main.go"]