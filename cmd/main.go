package main

import (
	"cleverbank/internal/infra/config/instance"
	"cleverbank/internal/infra/config/properties"
	"github.com/rendis/lightms"
)

func main() {
	lightms.SetPropFilePath("resources/properties.yml")
	lightms.AddProperty(properties.GetApplicationProperties())
	lightms.AddPrimary(instance.GetControllerPrimaryProcessInstance)
	lightms.Run()
}
