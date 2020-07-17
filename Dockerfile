FROM golang:1.12.4-stretch
# create a working directory
WORKDIR /go/src/app
# add source code
ADD . .

ENV GO111MODULE "on" 
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/githunter

ENV GH_GRAPHQL_GITHUB_URL "https://api.github.com/graphql"
ENV GH_GRAPHQL_GITLAB_URL "https://gitlab.com/api/graphql"
ENV GH_LOG_LEVEL "warn"
ENV GH_SERVER_PORT "3001"

USER root
RUN ls -l

ADD startup.sh /
RUN chmod -R 777 /startup.sh

# run compiled go app
ENTRYPOINT [ "/bin/sh" ]
CMD [ "-C", "/startup.sh" ]