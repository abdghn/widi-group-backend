package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"product-order-be/api/auth"
	"product-order-be/api/models"
	"product-order-be/api/responses"
	"product-order-be/api/utils/formaterror"

	"github.com/gorilla/mux"
)

func (server *Server) CreatePuskesmas(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/form-data")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	puskesmas := models.Puskesmas{}
	err = json.Unmarshal(body, &puskesmas)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	puskesmas.Prepare()
	err = puskesmas.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// if uid != order.UserID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }
	puskesmasCreated, err := puskesmas.SavePuskesmas(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, puskesmasCreated.ID))
	responses.JSON(w, http.StatusCreated, puskesmasCreated)
}

// func (server *Server) UploadOrder(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	if err := r.ParseMultipartForm(1024); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	alias := r.FormValue("alias")

// 	uploadedFile, handler, err := r.FormFile("file")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer uploadedFile.Close()

// 	// dir, err := os.Getwd()
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	filename := handler.Filename
// 	if alias != "" {
// 		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
// 	}
// 	fileLocation := filepath.Join("files", filename)
// 	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer targetFile.Close()

// 	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	// uploadPath := "files/" + filename
// 	// res, err := json.Marshal(files)
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	responses.JSON(w, http.StatusCreated, fileLocation)
// }

func (server *Server) GetAllPuskesmas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/form-data")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	dataPuskesmas := models.Puskesmas{}

	dataAllPuskesmas, err := dataPuskesmas.FindAllPuskesmas(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataAllPuskesmas)
}

func (server *Server) GetPuskesmas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/form-data")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	dataPuskesmas := models.Puskesmas{}

	puskesmasReceived, err := dataPuskesmas.FindPuskesmasByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, puskesmasReceived)
}
func (server *Server) GetPuskesmasByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/form-data")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	dataPuskesmas := models.Puskesmas{}

	puskesmasReceived, err := dataPuskesmas.FindPuskesmasByUserId(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, puskesmasReceived)
}

func (server *Server) UpdatePuskesmas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/form-data")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	vars := mux.Vars(r)

	// Check if the post id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	_, err = auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the post exist
	dataPuskesmas := models.Puskesmas{}
	err = server.DB.Debug().Model(models.Puskesmas{}).Where("id = ?", pid).Take(&dataPuskesmas).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Puskesmas not found"))
		return
	}

	// If a user attempt to update a post not belonging to him
	// if uid != order.UserID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	puskesmasUpdate := models.Puskesmas{}
	err = json.Unmarshal(body, &puskesmasUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	// if uid != orderUpdate.UserID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	puskesmasUpdate.Prepare()
	err = puskesmasUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	puskesmasUpdate.ID = dataPuskesmas.ID //this is important to tell the model the post id to update, the other update field are set above

	puskesmasUpdated, err := puskesmasUpdate.UpdateAPuskesmas(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, puskesmasUpdated)
}

func (server *Server) DeletePuskesmas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	_, err = auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the post exist
	dataPuskesmas := models.Puskesmas{}
	err = server.DB.Debug().Model(models.Puskesmas{}).Where("id = ?", pid).Take(&dataPuskesmas).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this post?
	// if uid != order.UserID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	_, err = dataPuskesmas.DeleteAPuskesmas(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
