package transport

import (
	"net/http"
	"user-api/transport/model/dto"
	"user-api/transport/model/transformation"
	"user-api/util"

	"github.com/go-chi/chi"
)

func getEmailFromURL(r *http.Request) string {
	email := chi.URLParam(r, "email")

	return email
}

func (handler *Handler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	var user dto.User
	err := util.UnmarshalRequest(request, &user)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	err = user.BirthDate.ValidateBirthDate()
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	err = handler.Service.CreateUser(request.Context(), transformation.ToServiceModel(user))
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusCreated, user)
}

func (handler *Handler) GetUser(writer http.ResponseWriter, request *http.Request) {
	user, err := handler.Service.GetUser(request.Context(), getEmailFromURL(request))
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusOK, transformation.ToDtoUser(user))
}

func (handler *Handler) ListUsers(writer http.ResponseWriter, request *http.Request) {
	users := handler.Service.ListUsers(request.Context())
	util.WriteResponse(writer, http.StatusOK, transformation.ToDtoUsers(users))
}

func (handler *Handler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var user dto.User
	err := util.UnmarshalRequest(request, &user)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	err = user.BirthDate.ValidateBirthDate()
	if err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	newUser, err := handler.Service.UpdateUser(request.Context(), getEmailFromURL(request), transformation.ToServiceModel(user))
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusOK, newUser)
}

func (handler *Handler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	err := handler.Service.DeleteUser(request.Context(), getEmailFromURL(request))
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusNoContent, nil)
}
