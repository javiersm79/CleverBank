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

func AccountCreationControllerInstance() *controller.AccountCreationController {

	return controller.GetAccountCreationController(GetAccountCreationUseCase())
}

func AuthLoginControllerInstance() *controller.AuthLoginController {

	return controller.GetAuthLoginController(GetLoginUseCase())
}

func GetControllerPrimaryProcessInstance() lightms.PrimaryProcess {
	return controller.GetControllerInstance([]controller.ControllerRunnable{
		AccountMovementControllerInstance(),
		AccountInfoControllerInstance(),
		AccountCreationControllerInstance(),
		AuthLoginControllerInstance(),
	})
}
