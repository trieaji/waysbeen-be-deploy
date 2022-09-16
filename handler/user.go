package handler

import (
	"encoding/json"
	dto "golangfnl/dto/result"
	usersdto "golangfnl/dto/users"
	"golangfnl/models"
	"golangfnl/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users} //untuk terkait response error dan success.
	json.NewEncoder(w).Encode(response)
	//json.NewEncoder(w).Encode(users)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"]) //strconv.Atoi untuk mengubah nilai string ke integer. Vars untuk menangkap params yang berupa id.

	user, err := h.UserRepository.GetUser(id) //mengembalikan 2 nilai yaitu User dan error yang ada di repositories
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		//json.NewEncoder(w).Encode(users)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)}
	json.NewEncoder(w).Encode(response)
	//json.NewEncoder(w).Encode(users)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(usersdto.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil { //Untuk memvalidasi data apakah sudah dikirimkan atau belum atau untuk mengakali
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// data form pattern submit to pattern entity db user
	user := models.User{ //Untuk mengakali data yang akan kita kirimkan sesuai kebutuhan. Data yang diambil berasal dari (models.User)
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Status:   request.Status,
	}
	data, err := h.UserRepository.CreateUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Prosesnya hampir sama dengan CreateUser, yang membedakan hanya dto nya saja
	request := new(usersdto.UpdateUserRequest)                       //untuk update gk wajib semua datanya dikirimkan
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil { //Untuk memvalidasi data apakah sudah dikirimkan atau belum atau untuk mengakali
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//Membutuhkan id untuk dikirimkan
	id, _ := strconv.Atoi(mux.Vars(r)["id"]) //strconv.Atoi untuk mengubah nilai string ke integer. Vars untuk menangkap params yang berupa id.

	user := models.User{}
	user, _ = h.UserRepository.GetUser(id)

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	if request.Status != "" {
		user.Status = request.Status
	}

	data, err := h.UserRepository.UpdateUser(user, id) //Karena di repo membutuhkan dua parameter, maka mengembalikannya harus sesuai dengan yang ada di repo. Mengembalikan dua parameter yaitu (user,id).

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)

}

func (h *handler) DeleteteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	data, err := h.UserRepository.DeleteUser(user, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)

}

func convertResponse(u models.User) usersdto.UserResponse { //Untuk pemanggilan response yang ada di dto. Isi datanya sesuai dengan response yang ada di dto.
	return usersdto.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Status:   u.Status,
	}
}
