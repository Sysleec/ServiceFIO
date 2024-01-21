package people

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sysleec/ServiceFIO/internal/converter"
	resp "github.com/Sysleec/ServiceFIO/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (s *Implementation) Update(w http.ResponseWriter, r *http.Request) {
	updateIDstr := chi.URLParam(r, "id")
	//convert to int
	updateID, err := strconv.Atoi(updateIDstr)
	if err != nil {
		fmt.Println("Error converting string to int")
	}

	mPeopl, err := converter.ToPeopleFromReq(r)
	if err != nil {
		resp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	people, err := s.peopleService.Update(r.Context(), updateID, mPeopl)
	if err != nil {
		resp.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	log.Info().Msgf("updated person with id %d", people.ID)

	resp.RespondWithJSON(w, http.StatusOK, converter.ToSqlcPersonFromModelPeople(people))
}
