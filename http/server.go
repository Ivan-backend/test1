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
	router.Path("/subscribes/price").Queries("service-name", "{service-name}").Methods("GET").HandlerFunc(s.httpHandlers.HandleSortServiceNameSubscribe)
	router.Path("/subscribes/price").Queries("user-id", "{user-id}").Methods("GET").HandlerFunc(s.httpHandlers.HandleSortServiceUserIdSubscribe)
	router.Path("/subscribes").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetSubscribes)
	// router.Path("/subscribes/price").Methods("GET").Queries("service-name", "{service-name}").Queries("user-id", "{user-id}").HandlerFunc(s.httpHandlers.HandlePriceAllSubscribes)
	// router.Path("/subscribes/price").Methods("GET").Queries("user-id", "{user-id}").HandlerFunc(s.httpHandlers.HandlePriceAllSubscribes)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
