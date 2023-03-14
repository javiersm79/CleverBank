package controller

import (
	Account "cleverbank/internal/core/domain/account"
	"cleverbank/internal/core/usecase"
	"cleverbank/internal/infra/primary/middleware"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type AccountInfoUseCaseMock struct {
	AccountInfoUseCase usecase.AccountInfoUseCase
	TestError          bool
}

func CreateAccountInfoUseCaseMock(testError bool) usecase.AccountInfoUseCase {
	return AccountInfoUseCaseMock{
		TestError: testError,
	}
}

func (u AccountInfoUseCaseMock) Handle(accountNumber string) (Account.AccountDetails, error) {
	if u.TestError {
		return Account.AccountDetails{}, errors.New("Error Testing")
	}

	return createAccountDetails(), nil

}

func TestCoExportRunController(t *testing.T) {

	type args struct {
		uc         usecase.AccountInfoUseCase
		queryParam string
		ginEngine  *gin.Engine
		authorized bool
	}

	type wantResponse struct {
		code     int
		response Account.AccountDetails
		err      error
	}

	tests := []struct {
		name string
		args args
		want wantResponse
	}{
		{
			name: "#1 Test AccountInfoController Endpoint /v1/balance OK",
			args: args{
				uc:         CreateAccountInfoUseCaseMock(false),
				queryParam: "?account=001-1234-457",
				ginEngine:  gin.Default(),
				authorized: true,
			},
			want: wantResponse{
				code:     http.StatusOK,
				response: createAccountDetails(),
				err:      nil,
			},
		},
		{
			name: "#2 Test AccountInfoController Endpoint /v1/balance BadRequest",
			args: args{
				uc:         CreateAccountInfoUseCaseMock(false),
				queryParam: "?",
				ginEngine:  gin.Default(),
				authorized: true,
			},
			want: wantResponse{
				code:     http.StatusBadRequest,
				response: createAccountDetails(),
				err:      nil,
			},
		},
		{
			name: "#3 Test AccountInfoController Endpoint /v1/balance Unauthorized",
			args: args{
				uc:         CreateAccountInfoUseCaseMock(false),
				queryParam: "?account=001-1234-457",
				ginEngine:  gin.Default(),
				authorized: false,
			},
			want: wantResponse{
				code:     http.StatusUnauthorized,
				response: createAccountDetails(),
				err:      nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			middleware.IsAuthorized = func(userToken string) (bool, error) {
				var err error

				if tt.args.authorized {
					err = nil
				} else {
					err = fmt.Errorf("invalid Token")
				}

				return tt.args.authorized, err
			}

			controller := AccountInfoController{tt.args.uc}
			controller.RunController(tt.args.ginEngine)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/v1/balance"+tt.args.queryParam, nil)
			req.Header.Add("Authorization", "anytoken")

			tt.args.ginEngine.ServeHTTP(w, req)

			if tt.want.code != w.Code {
				t.Errorf("AccountInfoController (/v1/balance) = %v, want %v", w.Code, tt.want.code)
			}
		})
	}

}

func createAccountDetails() Account.AccountDetails {
	return Account.AccountDetails{
		Id:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		ClientId:  2,
		Type:      "corriente",
		Number:    "123-456-789",
		Balance:   1000,
	}
}
