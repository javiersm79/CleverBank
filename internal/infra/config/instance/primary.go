package instance

import (
	"cleverbank/internal/infra/primary/controller"
	"github.com/rendis/lightms"
)

func AccountMovementControllerInstance() *controller.AccountMovementController {
	return controller.GetAccountMovementController(GetAccountMovementUseCase())
}

func AccountInfoControllerInstance() *controller.AccountInfoController {

	return controller.GetAccountInfoController(GetAccountInfoUseCase())
}

func GetControllerPrimaryProcessInstance() lightms.PrimaryProcess {
	return controller.GetControllerInstance([]controller.ControllerRunnable{
		AccountMovementControllerInstance(),
		AccountInfoControllerInstance(),
	})
}
