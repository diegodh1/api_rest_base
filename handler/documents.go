package handler

import (
	models "baseApi/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

//GetAllDocuments this function get all the documents at the DB
func GetAllDocuments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	documents := []models.DocumentType{}
	db.Where("document_type_status = ?", true).Find(&documents)
	respondJSON(w, http.StatusOK, documents)
}

//CreateDocuments this function creates a new document
func CreateDocuments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	document := models.DocumentType{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&document); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&document).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, document)
}

//UpdateDocuments this function set the state of a document
func UpdateDocuments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	documentID := vars["DocumentTypeID"]
	document := getDocumentOr404(db, documentID, w, r)
	if document == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&document); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&document).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, document)
}

func getDocumentOr404(db *gorm.DB, documentTypeID string, w http.ResponseWriter, r *http.Request) *models.DocumentType {
	document := models.DocumentType{}
	if err := db.First(&document, models.DocumentType{DocumentTypeID: documentTypeID}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &document
}
