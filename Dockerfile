FROM golang:1.12.4-stretch as BUILDER
# create a working directory
WORKDIR /go/src/app
# add source code

COPY . .

ENV GO111MODULE "on" 
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/githunter

ENV GH_GRAPHQL_GITHUB_URL "https://api.github.com/graphql"
ENV GH_GRAPHQL_GITLAB_URL "https://gitlab.com/api/graphql"
ENV GH_API_GITLAB_URL "https://gitlab.com/api/v4"
ENV GH_LOG_LEVEL "warning"
ENV GH_SERVER_PORT 3001

USER root

FROM alpine

COPY startup.sh /
RUN chmod -R 777 /startup.sh

COPY --from=BUILDER /go/bin/githunter /bin/githunter

# run compiled go app
ENTRYPOINT [ "/bin/sh" ]
CMD [ "/startup.sh" ]