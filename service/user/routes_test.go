package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nikhilsharma270027/API-Cart-GO/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockerUserStore{}
	Handler := NewHandler(userStore)

	// 1st test
	t.Run("should fail if the user payload is invalid",
		func(t *testing.T) {
			payload := types.RegisterUserPayload{
				FirstName: "user",
				LastName:  "123",
				Email:     "",
				Password:  "asd",
			}
			marshalled, _ := json.Marshal(payload)
			req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			router.HandleFunc("/register", Handler.handlerRegister)
			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusBadRequest {
				t.Errorf("expected status code %d , got %d", http.StatusBadRequest, rr.Code)
			}
		})

}

type mockerUserStore struct{}

func (m *mockerUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}
func (m *mockerUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (m *mockerUserStore) CreateUser(types.User) error {
	return nil
}
