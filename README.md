# GitHunter-API
This project consumes the API of GitHub, GitLab and, in the future, another providers.

## Installation
Download and install [Go](https://golang.org/dl) to run the project.

## Run
After install, in the root directory write the command to run the api:
```bash
go run main.go --log-level=debug --graphql-github-url=https://api.github.com/graphql --graphql-gitlab-url=https://gitlab.com/api/graphql --server-port=3001
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
