package implementation

import (
	Account "cleverbank/internal/core/domain/account"
	"cleverbank/internal/core/port"
	"errors"
	"reflect"
	"testing"
	"time"
)

var timenow = time.Now()

type AccountCreationPersistencePort struct {
	PersistencePort port.AccountCreationPort
	TestError       bool
}

func (c AccountCreationPersistencePort) CreateAccount(newAccountRequest Account.NewAccountReq) (Account.AccountDetails, error) {
	if c.TestError {
		return Account.AccountDetails{}, errors.New("Error Testing")
	}
	return createAccountDetails(), nil
}

func TestNewAccountCreationService(t *testing.T) {
	persistencePort := AccountCreationPersistencePort{}

	getService := GetAccountCreationService(persistencePort)
	if getService == nil {
		t.Error("Expect GetAccountCreationService not to be nil but got:", getService)
	}
}

func TestHandleAccountCreationService(t *testing.T) {

	type args struct {
		newAccountRequest Account.NewAccountReq
		port              port.AccountCreationPort
	}

	type wantResponse struct {
		response Account.AccountDetails
		err      error
	}

	test := []struct {
		name string
		args args
		want wantResponse
	}{
		{
			name: "#1 Test Handle AccountCreationService",
			args: args{
				newAccountRequest: createNewAccountRequest(),
				port:              createAccountCreationPersistencePort(false),
			},
			want: wantResponse{
				response: createAccountDetails(),
				err:      nil,
			},
		},
		{
			name: "#2 Test Handle AccountCreationService - Error",
			args: args{
				newAccountRequest: createNewAccountRequest(),
				port:              createAccountCreationPersistencePort(true),
			},
			want: wantResponse{
				response: Account.AccountDetails{},
				err:      errors.New("Error Testing"),
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			serv := AccountCreationService{tt.args.port}
			result, err := serv.Handle(tt.args.newAccountRequest)

			if err == nil {
				if !reflect.DeepEqual(result, tt.want.response) {
					t.Errorf("AccountCreationService Hanlde() Result= [%v], Expected= [%v]", result, tt.want.response)
				}
			} else {
				if err.Error() != tt.want.err.Error() {
					t.Errorf("Handle \nExpected: [%v], \nResult: [%v]", tt.want.err.Error(), err.Error())
				}
			}
		})

	}

}

func createAccountDetails() Account.AccountDetails {
	return Account.AccountDetails{
		Id:        1,
		CreatedAt: timenow,
		UpdatedAt: timenow,
		DeletedAt: nil,
		ClientId:  2,
		Type:      "corriente",
		Number:    "123-456-789",
		Balance:   1000,
	}
}

func createNewAccountRequest() Account.NewAccountReq {
	return Account.NewAccountReq{
		ClientId: 2,
		Type:     "corriente",
	}
}

func createAccountCreationPersistencePort(testError bool) port.AccountCreationPort {
	return AccountCreationPersistencePort{
		TestError: testError,
	}
}
