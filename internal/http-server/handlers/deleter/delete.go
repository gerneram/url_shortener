package deleter

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
	resp "project/internal/lib/api/response"
)

type Deleter interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, URLDelete Deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.deleter"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("ignoring url deleter: missing alias")
			render.JSON(w, r, resp.Error("invalid url deleter: missing alias"))
			return
		}
		err := URLDelete.DeleteURL(alias)
		if err != nil {
			log.Info("ignoring url deleter: " + err.Error())
			render.JSON(w, r, resp.Error("url deleter: "+err.Error()))
			return
		}

		log.Info("deleter: success url: " + alias)
		render.JSON(w, r, resp.OK())
	}
}
