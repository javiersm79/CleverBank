package implementation

import (
	Account "cleverbank/internal/core/domain/account"
	"cleverbank/internal/core/port"
	"errors"
	"reflect"
	"testing"
	"time"
)

var timeinfonow = time.Now()

type AccountInfoPersistencePort struct {
	PersistencePort port.AccountInfoPort
	TestError       bool
}

func (c AccountInfoPersistencePort) GetBalance(accountNumber string) (Account.AccountDetails, error) {
	if c.TestError {
		return Account.AccountDetails{}, errors.New("Error Testing")
	}
	return createAccountDetails(), nil
}

func TestNewAccountInfoService(t *testing.T) {
	persistencePort := AccountInfoPersistencePort{}

	getService := GetAccountInfoService(persistencePort)
	if getService == nil {
		t.Error("Expect GetAccountInfoService not to be nil but got:", getService)
	}
}

func TestHandleAccountIfoService(t *testing.T) {

	type args struct {
		accountNumber string
		port          port.AccountInfoPort
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
			name: "#1 Test Handle AccountInfoService",
			args: args{
				accountNumber: "123-456-789",
				port:          createAccountInfoPersistencePort(false),
			},
			want: wantResponse{
				response: createAccountInfo(),
				err:      nil,
			},
		},
		{
			name: "#2 Test Handle AccountInfoService - Error",
			args: args{
				accountNumber: "123-456-789",
				port:          createAccountInfoPersistencePort(true),
			},
			want: wantResponse{
				response: Account.AccountDetails{},
				err:      errors.New("Error Testing"),
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			serv := AccountInfoService{tt.args.port}
			result, err := serv.Handle(tt.args.accountNumber)

			if err == nil {
				if !reflect.DeepEqual(result, tt.want.response) {
					t.Errorf("AccountInfoService Hanlde() Result= [%v], Expected= [%v]", result, tt.want.response)
				}
			} else {
				if err.Error() != tt.want.err.Error() {
					t.Errorf("Handle \nExpected: [%v], \nResult: [%v]", tt.want.err.Error(), err.Error())
				}
			}
		})

	}

}

func createAccountInfo() Account.AccountDetails {
	return Account.AccountDetails{
		Id:        1,
		CreatedAt: timeinfonow,
		UpdatedAt: timeinfonow,
		DeletedAt: nil,
		ClientId:  2,
		Type:      "corriente",
		Number:    "123-456-789",
		Balance:   1000,
	}
}

func createAccountInfoPersistencePort(testError bool) port.AccountInfoPort {
	return AccountInfoPersistencePort{
		TestError: testError,
	}
}
