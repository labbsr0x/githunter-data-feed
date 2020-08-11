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
	serverPort := flag.Int("server-port", 3002, "The server port")
	flag.Parse()

	configLogrus(logLevel)

	logrus.WithFields(logrus.Fields{
		"level-logging": *logLevel,
	}).Debug("Log level defined")

	environment.GithubGraphQLURL = *graphqlGithubURL
	logrus.WithFields(logrus.Fields{
		"url": *graphqlGithubURL,
	}).Debug("GraphQL Github API defined")

	environment.GitlabGraphQLURL = *graphqlGitlabURL
	logrus.WithFields(logrus.Fields{
		"url": *graphqlGitlabURL,
	}).Debug("GraphQL Gitlab API defined")

	environment.ServerPort = *serverPort
	logrus.WithFields(logrus.Fields{
		"port": *serverPort,
	}).Debug("The server port defined")

}

// Get env from external
func Get() Environment {
	return environment
}

func configLogrus(logLevel *string) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.999999",
	})

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
}
