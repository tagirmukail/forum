package forum

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	jsoniter "github.com/json-iterator/go"
	"github.com/tagirmukail/forum/internal/dto"
)

// @Summary Create Topic
// @Description create topic
// @Tags Topic
// @Accept  json
// @Produce json
// @Param Topic body dto.TopicRequest true "Topic"
// @Success 200 {object} dto.Topic "Success operation"
// @Router /api/v1/topics [post]
//
func (s *Service) createTopic(w http.ResponseWriter, r *http.Request) {
	wr, _ := w.(*customResponseWriter)

	var topicDTO dto.TopicRequest

	err := jsoniter.NewDecoder(r.Body).Decode(&topicDTO)
	if err != nil {
		wr.ERR = err

		http.Error(w, requestBodyDeserializationFailed, http.StatusBadRequest)

		return
	}

	result, err := s.l.CreateTopic(r.Context(), topicDTO)
	if err != nil {
		errMessage := "topic creation failed"
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

// @Summary List Topics
// @Description list topics
// @Tags Topic
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} dto.Topics "Success operation"
// @Router /api/v1/topics [get]
//
func (s *Service) listTopics(w http.ResponseWriter, r *http.Request) {
	wr, _ := w.(*customResponseWriter)

	var (
		err    error
		limit  int
		offset int
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

	var topics dto.Topics

	topics, err = s.l.ListTopics(r.Context(), limit, offset)
	if err != nil {
		errMessage := "topics list failed"
		wr.ERR = err

		http.Error(w, errMessage, http.StatusInternalServerError)

		return
	}

	s.setJSONContentType(w)

	err = jsoniter.NewEncoder(w).Encode(topics)
	if err != nil {
		wr.ERR = err

		http.Error(w, respSerializationFailed, http.StatusBadRequest)

		return
	}
}

// @Summary Get Topic
// @Description get topic
// @Tags Topic
// @Produce json
// @Param topic_id path string true "Topic identification"
// @Success 200 {object} dto.TopicDetailed "Success operation"
// @Router /api/v1/topics/{topic_id} [get]
//
func (s *Service) getTopic(w http.ResponseWriter, r *http.Request) {
	wr, _ := w.(*customResponseWriter)

	topicID := mux.Vars(r)["topic_id"]

	result, err := s.l.GetTopic(r.Context(), topicID)
	if err != nil {
		errMessage := "get topic failed"
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
