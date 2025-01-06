package config

type MissingEnvVariableError struct{}

func (m *MissingEnvVariableError) Error() string {
	return "missing environment variable for config"
}
