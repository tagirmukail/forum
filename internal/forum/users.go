package forum

import (
	"net/http"

	"github.com/tagirmukail/forum/internal/dto"

	jsoniter "github.com/json-iterator/go"
)

// @Summary Create User
// @Description create user
// @Tags User
// @Accept  json
// @Produce json
// @Param User body dto.UserRequest true "User"
// @Success 200 {object} dto.User "Success operation"
// @Router /api/v1/users [post]
//
func (s *Service) createUser(w http.ResponseWriter, r *http.Request) {
	wr, _ := w.(*customResponseWriter)

	var userDto dto.UserRequest

	err := jsoniter.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		wr.ERR = err

		http.Error(w, requestBodyDeserializationFailed, http.StatusBadRequest)

		return
	}

	var result dto.User

	result, err = s.l.CreateUser(r.Context(), userDto)
	if err != nil {
		errMessage := "user creation failed"
		wr.ERR = err

		http.Error(w, errMessage, http.StatusInternalServerError)

		return
	}

	s.setJSONContentType(w)

	err = jsoniter.NewEncoder(w).Encode(result)
	if err != nil {
		wr.ERR = err

		http.Error(w, respSerializationFailed, http.StatusBadRequest)

		return
	}
}
