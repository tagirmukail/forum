package forum

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	jsoniter "github.com/json-iterator/go"
	"github.com/tagirmukail/forum/internal/dto"
)

// @Summary Create Comment
// @Description create comment
// @Tags Comment
// @Accept  json
// @Produce json
// @Param topic_id path string true "Topic ID"
// @Param Comment body dto.CommentRequest true "Comment"
// @Success 200 {object} dto.Comment "Success operation"
// @Router /api/v1/topics/{topic_id}/comments [post]
//
func (s *Service) createComment(w http.ResponseWriter, r *http.Request) {
	wr, _ := w.(*customResponseWriter)

	topicID := mux.Vars(r)["topic_id"]

	var commentDTO dto.CommentRequest

	err := jsoniter.NewDecoder(r.Body).Decode(&commentDTO)
	if err != nil {
		wr.ERR = err

		http.Error(w, requestBodyDeserializationFailed, http.StatusBadRequest)

		return
	}

	result, err := s.l.CreateComment(r.Context(), topicID, commentDTO)
	if err != nil {
		errMessage := "comment creation failed"
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

// @Summary List Comments
// @Description list comments
// @Tags Comment
// @Produce json
// @Param topic_id path string true "Topic ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} dto.Comments "Success operation"
// @Router /api/v1/topics/{topic_id}/comments [get]
//
func (s *Service) listComments(w http.ResponseWriter, r *http.Request) {
	wr, _ := w.(*customResponseWriter)

	var (
		err     error
		limit   int
		offset  int
		topicID = mux.Vars(r)["topic_id"]
	)

	limitParam := r.URL.Query().Get("limit")
	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			wr.ERR = err

			http.Error(w, invalidLimitParam, http.StatusBadRequest)

			return
		}
	}

	offsetParam := r.URL.Query().Get("offset")
	if offsetParam != "" {
		offset, err = strconv.Atoi(offsetParam)
		if err != nil {
			wr.ERR = err

			http.Error(w, invalidOffsetParam, http.StatusBadRequest)

			return
		}
	}

	result, err := s.l.ListComments(r.Context(), topicID, limit, offset)
	if err != nil {
		errMessage := "list comments failed"
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
