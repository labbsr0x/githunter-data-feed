package env

import (
	"errors"
	"reflect"

	goEnv "github.com/Netflix/go-env"
	"github.com/rs/zerolog/log"
)

// Environment Environment
type Environment struct {
	LogLevel         string `env:"GH_LOG_LEVEL"`
	GithubGraphQLURL string `env:"GH_GIT_GRAPHQL_URL"`
	GitlabGraphQLURL string `env:"GH_GITLAB_GRAPHQL_URL"`

	Extras goEnv.EnvSet
	init   bool
}

var env Environment

// GetEnvironment variable
func Get() Environment {
	if !env.init {
		es, err := goEnv.UnmarshalFromEnviron(&env)
		if err != nil {
			log.Fatal().Err(err).Msg("Critical error during unmarshal env process")
		}
		log.Warn().Interface("env_vars", env).Msg("Environment variables loaded : OK")
		env.Extras = es
		env.init = true
		checkUndefinedEnvVar(env)
	}
	return env
}

// Config verifica se as variáveis obrigatórias foram definidas
func checkUndefinedEnvVar(environment Environment) {
	e := reflect.ValueOf(&environment).Elem()
	for i := 0; i < e.NumField(); i++ {
		if !e.Field(i).CanInterface() {
			continue
		}
		varName := e.Type().Field(i).Name
		varValue := e.Field(i).Interface()

		if varValue == "" || varValue == 0 {
			err := errors.New("required environment variable not set")
			log.Fatal().Err(err).Msgf("Error: environment variable %s was not set", varName)
		}

	}
}
