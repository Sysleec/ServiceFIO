package people

import (
	"fmt"
	"net/http"
	"strconv"

	resp "github.com/Sysleec/ServiceFIO/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (s *Implementation) Delete(w http.ResponseWriter, r *http.Request) {
	deleteIDstr := chi.URLParam(r, "id")
	//convert to int
	deleteID, err := strconv.Atoi(deleteIDstr)
	if err != nil {
		fmt.Println("Error converting string to int")
	}

	peopleID, err := s.peopleService.Delete(r.Context(), int64(deleteID))
	if err != nil {
		resp.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	log.Info().Msgf("deleted people with id %d", peopleID)

	res := make(map[string]int64)
	res["deleted"] = peopleID

	resp.RespondWithJSON(w, http.StatusOK, res)
}
