package instance

import (
	"cleverbank/internal/infra/primary/controller"
	"github.com/rendis/lightms"
)

func AccountMovementControllerInstance() *controller.AccountMovementController {
	return controller.GetAccountMovementController(GetAccountMovementUseCase())
}

func GetControllerPrimaryProcessInstance() lightms.PrimaryProcess {
	return controller.GetControllerInstance([]controller.ControllerRunnable{
		AccountMovementControllerInstance(),
	})
}
