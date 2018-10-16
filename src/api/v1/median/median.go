package median

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gonum.org/v1/gonum/stat"

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
	router.Post("/", PostMedian)

	return router
}

/*
PostMedian - Handle the POST request.
*/
func PostMedian(w http.ResponseWriter, r *http.Request) {

	// Very basic parameter checking.
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
		median := stat.Quantile(0.5, stat.Empirical, data, nil)
		w.Write([]byte(fmt.Sprintf("%f", median)))
	} else {
		render.Render(w, r, srverrors.ErrInvalidRequest(errors.New("invalid data")))
		return
	}
}
