package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/users").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateUser)
	router.Path("/users").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetUsers)
	router.Path("/subscribes").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateNoteSubscribe)
	router.Path("/subscribes/{service-name}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetByServiceNameSubscribe)
	router.Path("/subscribes/{service-name}/date").Queries("date-subscribe-start", "{date-subscribe-start}").Queries("date-subscribe-end", "{date-subscribe-end}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetByServiceNameAndDateSubscribe)
	router.Path("/subscribes/{user-id}").Queries("user-id", "{user-id}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetByUserIdSubscribe)
	router.Path("/subscribes/{user-id}/date").Queries("date-subscribe-start", "{date-subscribe-start}").Queries("date-subscribe-end", "{date-subscribe-end}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetByUserIdAndDateSubscribe)
	router.Path("/subscribes").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetSubscribes)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
