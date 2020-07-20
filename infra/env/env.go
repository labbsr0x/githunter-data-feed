package env

import (
	"flag"

	"github.com/sirupsen/logrus"
)

// Environment Environment
type Environment struct {
	GithubGraphQLURL string
	GitlabGraphQLURL string
	ServerPort       int
}

var environment Environment

// Configuration for environment variables
func Config() {

	logLevel := flag.String("log-level", "info", "debug, info, warning, error")
	graphqlGithubURL := flag.String("graphql-github-url", "https://api.github.com/graphql", "The GraphQL Github API URL.")
	graphqlGitlabURL := flag.String("graphql-gitlab-url", "https://gitlab.com/api/graphql", "The GraphQL GitLab API URL.")
	serverPort := flag.Int("server-port", 3001, "The server port")
	flag.Parse()

	switch *logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
		break
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
		break
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
		break
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.Debugf("log level set: %s", logLevel)

	environment.GithubGraphQLURL = *graphqlGithubURL
	logrus.Debugf("GraphQL Github API URL at %s", graphqlGithubURL)

	environment.GitlabGraphQLURL = *graphqlGitlabURL
	logrus.Debugf("GraphQL GitLab API URL at %s", graphqlGitlabURL)

	environment.ServerPort = *serverPort
	logrus.Debugf("The server port:%d", serverPort)

}

// Get env from external
func Get() Environment {
	return environment
}
