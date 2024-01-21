package people

import (
	"net/http"

	"github.com/Sysleec/ServiceFIO/internal/converter"
	resp "github.com/Sysleec/ServiceFIO/internal/utils"
	"github.com/rs/zerolog/log"
)

// Create creates a new user
func (s *Implementation) Create(w http.ResponseWriter, r *http.Request) {
	mPeopl, err := converter.ToPeopleFromReq(r)
	if err != nil {
		resp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	people, err := s.peopleService.Create(r.Context(), mPeopl)
	if err != nil {
		resp.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	log.Info().Msgf("created person with id %d", people.ID)

	resp.RespondWithJSON(w, http.StatusCreated, converter.ToSqlcPersonFromModelPeople(people))
}
