package people

import (
	"log"
	"net/http"

	"github.com/Sysleec/ServiceFIO/internal/converter"
	resp "github.com/Sysleec/ServiceFIO/internal/utils"
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

	log.Printf("created people with id %d\n", people.ID)

	resp.RespondWithJSON(w, http.StatusOK, converter.ToSqlcPersonFromModelPeople(people))
}
