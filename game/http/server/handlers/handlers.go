package handlers

import (
	"context"
	"encoding/json"
	"example.com/m/internal/service"
	"example.com/m/pkg/life"
	"fmt"
	"net/http"
	"os"
)

// создадим новый тип для добавления middleware к обработчикам
type Decorator func(http.Handler) http.Handler

// объект для хранения состояния игры
type LifeStates struct {
	service.LifeService
}

func New(ctx context.Context,
	lifeService service.LifeService,
) (http.Handler, error) {
	serveMux := http.NewServeMux()

	lifeState := LifeStates{
		LifeService: lifeService,
	}

	serveMux.HandleFunc("/nextstate", lifeState.nextState)
	serveMux.HandleFunc("/setstate", lifeState.setState)
	serveMux.HandleFunc("/reset", lifeState.setState)

	return serveMux, nil
}

// функция добавления middleware
func Decorate(next http.Handler, ds ...Decorator) http.Handler {
	decorated := next
	for d := len(ds) - 1; d >= 0; d-- {
		decorated = ds[d](decorated)
	}

	return decorated
}

// получение очередного состояния игры
func (ls *LifeStates) nextState(w http.ResponseWriter, r *http.Request) {
	worldState := ls.LifeService.NewState()

	writeJSON(w, worldState)
}

type randBody struct {
	fill int
}

func (ls *LifeStates) setState(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var v randBody
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	file, err := os.OpenFile("state.cfg", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d%", v.fill))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	worldState := ls.LifeService.SetRandomState(v.fill)
	writeJSON(w, worldState)
}

func writeJSON(w http.ResponseWriter, worldState *life.World) {
	err := json.NewEncoder(w).Encode(worldState.Cells)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
