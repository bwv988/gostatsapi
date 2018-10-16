package min

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/bwv988/gostatsapi/src/lib/srv"
	"github.com/bwv988/gostatsapi/src/lib/types"
	"github.com/bwv988/gostatsapi/src/lib/utils"
)

/*
FIXME
*/
func Routes() *chi.Mux {
	router := chi.NewRouter()

	// Use a RegExp to filter requests.
	router.Post("/{topn:[0-9]+}", PostMin)

	return router
}

/*
PostMin - Handle the POST request.
*/
func PostMin(w http.ResponseWriter, r *http.Request) {

	// Very basic parameter checking.
	topn, err := strconv.Atoi(chi.URLParam(r, "topn"))
	if err != nil {
		render.Render(w, r, srverrors.ErrInvalidRequest(errors.New("invalid request parameter")))
		return
	}

	if r.Body == nil {
		render.Render(w, r, srverrors.ErrInvalidRequest(errors.New("missing parameter")))
		return
	}

	var dataBody mytypes.DataBody
	jsonErr := json.NewDecoder(r.Body).Decode(&dataBody)
	if jsonErr != nil {
		render.Render(w, r, srverrors.ErrInvalidRequest(errors.New("no data given")))
		return
	}

	// FIXME: Error handling.
	dataRaw := dataBody.Data[:]
	data := convert.StringtoFloat64(dataRaw)

	l := len(data)
	if l > 0 {
		sort.Float64s(data)
		w.Write([]byte(fmt.Sprintf("%f", data[0:topn])))
	} else {
		render.Render(w, r, srverrors.ErrInvalidRequest(errors.New("invalid data")))
		return
	}
}
