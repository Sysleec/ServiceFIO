package people

import (
	"net/http"
	"strconv"

	"github.com/Sysleec/ServiceFIO/internal/converter"
	resp "github.com/Sysleec/ServiceFIO/internal/utils"
	"github.com/rs/zerolog/log"
)

func (s *Implementation) Get(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}
	filter := r.URL.Query().Get("filter")

	log.Debug().Msgf("filter: %s, page: %d, limit: %d", filter, page, limit)

	people, err := s.peopleService.Get(r.Context(), filter, page, limit)
	if err != nil {
		resp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msgf("got %d people", len(people))
	resp.RespondWithJSON(w, http.StatusOK, converter.ToSqlcPersonSliceFromModelPeopleSlice(people))
}
