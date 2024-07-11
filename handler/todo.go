package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// ServeHTTP implements http.Handler.
func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req *model.CreateTODORequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if req.Subject == `` {
			http.Error(w, "Subject is empty", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
		if err != nil {
			http.Error(w, "Failed to create TODO", http.StatusInternalServerError)
			return
		}

		resp := model.CreateTODOResponse{
			TODO: *todo,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

		return
	}
	if r.Method == http.MethodPut {
		var req *model.UpdateTODORequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if req.ID == 0 {
			http.Error(w, "ID is empty", http.StatusBadRequest)
			return
		}
		if req.Subject == `` {
			http.Error(w, "Subject is empty", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		todo, err := h.svc.UpdateTODO(ctx, req.ID, req.Subject, req.Description)
		if err != nil {
			http.Error(w, "Failed to updata TODO", http.StatusInternalServerError)
			return
		}
		if todo == nil {
			http.Error(w, "Id not found", http.StatusNotFound)
			return
		}

		resp := model.UpdateTODOResponse{
			TODO: *todo,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

		return
	}
	if r.Method == http.MethodGet {
		query := r.URL.Query()
		prevID, err := strconv.ParseInt((query.Get("prev_id")), 10, 64)
		if err != nil && query.Get("prev_id") != "" {
			http.Error(w, "prevID is empty", http.StatusBadRequest)
			return
		}
		size, err := strconv.ParseInt((query.Get("size")), 10, 64)
		if err != nil && query.Get("size") != "" {
			http.Error(w, "size is empty", http.StatusBadRequest)
			return
		}
		req := model.ReadTODORequest{
			PrevID: prevID,
			Size:   size,
		}

		ctx := r.Context()
		todos, err := h.svc.ReadTODO(ctx, req.PrevID, req.Size)
		if err != nil {
			http.Error(w, "Failed to read TODOs", http.StatusInternalServerError)
			return
		}

		resp := model.ReadTODOResponse{
			TODOs: todos,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

		return
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
