package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"main/feature_postgres/simple_connection"
	"main/feature_postgres/simple_sql"
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

	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)

	if err != nil {
		panic(err)
	}

	if err := json.NewDecoder(r.Body).Decode(&subscribeDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	subscribe := subscribes.NewSubscribe(subscribeDTO.ServiceName, subscribeDTO.Price, subscribeDTO.UserId, subscribeDTO.DateStart)

	if err := simple_sql.InsertRow(ctx, conn, subscribeDTO.ServiceName, subscribeDTO.Price, subscribeDTO.UserId.String(), subscribeDTO.DateStart, time.Date(2026, time.March, 11, 10, 9, 8, 7, time.UTC)); err != nil {
		panic(err)
	}

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
	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}

	subscribes, err := simple_sql.GetSubscribes(ctx, conn)
	if err != nil {
		panic(err)
	}

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

func (h *HTTPHandlers) HandleGetByServiceNameSubscribe(w http.ResponseWriter, r *http.Request) {
	serviceName := mux.Vars(r)["service-name"]

	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}

	getServiceNameSubscribe, price, err := simple_sql.GetByServiceNameSubscribe(ctx, conn, serviceName)
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(getServiceNameSubscribe, "", "	")
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

	fmt.Println(price)
}

func (h *HTTPHandlers) HandleGetByServiceNameAndDateSubscribe(w http.ResponseWriter, r *http.Request) {
	serviceName := mux.Vars(r)["service-name"]
	dateSubscribeStart := mux.Vars(r)["date-subscribe-start"]
	dateSubscribeEnd := mux.Vars(r)["date-subscribe-end"]

	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}

	getServiceNameSubscribe, price, err := simple_sql.GetByServiceNameAndDateSubscribe(ctx, conn, serviceName, dateSubscribeStart, dateSubscribeEnd)
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(getServiceNameSubscribe, "", "	")
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

	fmt.Println(price)
}

func (h *HTTPHandlers) HandleGetByUserIdSubscribe(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user-id"]

	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}

	getByUserIdSubscribe, price, err := simple_sql.GetByUserIdSubscribe(ctx, conn, userId)
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(getByUserIdSubscribe, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

	fmt.Println(price)
}

func (h *HTTPHandlers) HandleGetByUserIdAndDateSubscribe(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user-id"]
	dateSubscribeStart := mux.Vars(r)["date-subscribe-start"]
	dateSubscribeEnd := mux.Vars(r)["date-subscribe-end"]

	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}

	getByUserIdSubscribe, price, err := simple_sql.GetByUserIdAndDateSubscribe(ctx, conn, userId, dateSubscribeStart, dateSubscribeEnd)
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(getByUserIdSubscribe, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

	fmt.Println(price)
}
