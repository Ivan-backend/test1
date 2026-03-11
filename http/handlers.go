package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"main/subscribes"
	"main/users"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	subscribes *subscribes.List
	users      *users.List
}

func NewHTTPHandlers(users *users.List, subscribes *subscribes.List) *HTTPHandlers {
	return &HTTPHandlers{
		subscribes: subscribes,
		users:      users,
	}
}

func (h *HTTPHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	user := users.NewUser(uuid.New())

	if err := h.users.CreateUser(user); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, nil) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *HTTPHandlers) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.users.GetUsers()
	b, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *HTTPHandlers) HandleCreateNoteSubscribe(w http.ResponseWriter, r *http.Request) {
	var subscribeDTO SubscribeDTO
	if err := json.NewDecoder(r.Body).Decode(&subscribeDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	subscribe := subscribes.NewSubscribe(subscribeDTO.ServiceName, subscribeDTO.Price, subscribeDTO.UserId, subscribeDTO.DateStart)
	if err := h.subscribes.CreateNoteSubscribe(subscribe); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, nil) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(subscribe, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *HTTPHandlers) HandleGetSubscribes(w http.ResponseWriter, r *http.Request) {
	subscribes := h.subscribes.GetSubscribes()
	b, err := json.MarshalIndent(subscribes, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *HTTPHandlers) HandleSortServiceNameSubscribe(w http.ResponseWriter, r *http.Request) {
	// userId := mux.Vars(r)["user-id"]
	serviceName := mux.Vars(r)["service-name"]
	// priceAllSubscribes := h.subscribes.PriceAllSubscribes(serviceName, userId)

	sortServiceNameSubscribe, priceAllSubscribes := h.subscribes.SortServiceNameSubscribe(serviceName)
	// sortServiceNameSubscribe := h.subscribes.SortServiceNameSubscribe(serviceName)

	b, err := json.MarshalIndent(sortServiceNameSubscribe, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

	p, err := json.Marshal(priceAllSubscribes)
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(p); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *HTTPHandlers) HandleSortServiceUserIdSubscribe(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user-id"]
	// serviceName := mux.Vars(r)["service-name"]
	// priceAllSubscribes := h.subscribes.PriceAllSubscribes(serviceName, userId)
	SortServiceUserIdSubscribe, priceAllSubscribes := h.subscribes.SortServiceUserIdSubscribe(userId)

	b, err := json.MarshalIndent(SortServiceUserIdSubscribe, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

	p, err := json.Marshal(priceAllSubscribes)
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(p); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

// func (h *HTTPHandlers) HandlePriceAllSubscribes(w http.ResponseWriter, r *http.Request) {
// 	userId := mux.Vars(r)["user-id"]
// 	serviceName := mux.Vars(r)["service-name"]
// 	priceAllSubscribes := h.subscribes.PriceAllSubscribes(serviceName, userId)

// 	params := r.URL.Query()
// 	if params.Get("service-name") != "" {
// 		sortServiceNameSubscribe := h.subscribes.SortServiceNameSubscribe(serviceName)

// 		b, err := json.MarshalIndent(sortServiceNameSubscribe, "", "	")
// 		if err != nil {
// 			panic(err)
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		if _, err := w.Write(b); err != nil {
// 			fmt.Println("failed to write http response", err)
// 			return
// 		}
// 	}

// 	if params.Get("user-id") != "" {
// 		SortServiceUserIdSubscribe := h.subscribes.SortServiceUserIdSubscribe(userId)

// 		b, err := json.MarshalIndent(SortServiceUserIdSubscribe, "", "	  ")
// 		if err != nil {
// 			panic(err)
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		if _, err := w.Write(b); err != nil {
// 			fmt.Println("failed to write http response", err)
// 			return
// 		}
// 	}

// 	b, err := json.MarshalIndent(priceAllSubscribes, "", "	")
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	if _, err := w.Write(b); err != nil {
// 		fmt.Println("failed to write http response", err)
// 		return
// 	}
// }
