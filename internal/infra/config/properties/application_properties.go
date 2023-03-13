package properties

import "sync"

var (
	applicationPropertiesOnce     sync.Once
	applicationPropertiesInstance *ApplicationProperties
)

func GetApplicationProperties() *ApplicationProperties {
	applicationPropertiesOnce.Do(func() {
		applicationPropertiesInstance = &ApplicationProperties{}
	})

	return applicationPropertiesInstance
}

type ApplicationProperties struct {
	Application `yaml:"application"`
}

type Application struct {
	SecretKeyJWT string `yaml:"secret-key-jwt"`
}
