package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type MapMessage map[string]string

// message from error
const (
	errEnServerError            = "Internal server error"
	errIdServerError            = "Terjadi kesalahan pada server"
	errEnGeneral                = "Failed"
	errIdGeneral                = "Gagal"
	successEnGeneral            = "Success"
	successIdGeneral            = "Sukses"
	errEnUndefined              = "Undefined error"
	errIdUndefined              = "Error tidak terdefinisi"
	errEnBadRequest             = "Invalid request"
	errIdBadRequest             = "Request tidak sah"
	errEnRegisEmailNotAvailable = "Email not available"
	errIdRegisEmailNotAvailable = "Email tidak dapat digunakan"
	errEnUserNotFound           = "User with email not found"
	errIdUserNotFound           = "User dengan email tidak ditemukan"
	errEnPwdNotMatch            = "Password not match"
	errIdPwdNotMatch            = "Password tidak cocok"
	errEnTokenExpired           = "Token expired"
	errIdTokenExpired           = "Token expired"
	errEnTokenInvalid           = "Token invalid"
	errIdTokenInvalid           = "Token invalid"
	errEnNotFound               = "Data not found"
	errIdNotFound               = "Data tidak ditemukan"
	errEnUnauthorized           = "Unauthorized"
	errIdUnauthorized           = "Unauthorized"
	errEnInvalidParameter       = "Invalid parameter ID"
	errIdInvalidParameter       = "Parameter ID tidak sah"
	errEnOnlyCandidate          = "Only Candidate can process action"
	errIdOnlyCandidate          = "Hanya candidate yang dapat melakukan aksi"
	ErrEnReviewFound            = "Candidate can review company once"
	ErrIdReviewFound            = "Kandidat hanya boleh membuat 1 review tiap company"
	successEnCreated            = "Success Created"
)

// status code
const (
	statusServerError            = http.StatusInternalServerError
	statusOk                     = http.StatusOK
	statusBadRequest             = http.StatusBadRequest
	statusValidatorFail          = http.StatusBadRequest
	statusRegisEmailNotAvailable = http.StatusBadRequest
	statusUserNotFound           = http.StatusNotFound
	statusPwdNotMatch            = http.StatusBadRequest
	statusTokenExpired           = http.StatusBadRequest
	statusTokenInvalid           = http.StatusBadRequest
	statusNotFound               = http.StatusNotFound
	statusUnautorized            = http.StatusUnauthorized
	statusInvalidParameterID     = http.StatusBadRequest
	statusOnlyCandidate          = http.StatusBadRequest
	statusReviewFound            = http.StatusBadRequest
	statusCreated                = http.StatusCreated
)

// constant increasing sequences
const (
	ErrorServerError = 0
	ErrorBadRequest  = iota + 1
	SuccesOk
	ErrorValidation
	ErrorRegisEmail
	ErrorUserNotFound
	ErrorPwdNotMatch
	ErrorTokenExpired
	ErrorTokenInvalid
	ErrorNotFound
	ErrorUnauthorized
	ErrorInvalidParameter
	ErrorOnlyCandidate
	ErrorReviewFound
	SuccessCreated
)

type ReturningValue struct {
	Status  int         `json:"status"`
	Message MapMessage  `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Err     string      `json:"err,omitempty"`
}

func add(status int, en string, id string) ReturningValue {
	return ReturningValue{
		Status: status,
		Message: MapMessage{
			"en": en,
			"id": id,
		},
		Err:  "",
		Data: nil,
	}
}

var mapping = map[int]ReturningValue{
	ErrorServerError:      add(statusServerError, errEnServerError, errIdServerError),
	ErrorBadRequest:       add(statusBadRequest, errEnBadRequest, errIdBadRequest),
	SuccesOk:              add(statusOk, successEnGeneral, successIdGeneral),
	ErrorValidation:       add(statusValidatorFail, "", ""),
	ErrorRegisEmail:       add(statusRegisEmailNotAvailable, errEnRegisEmailNotAvailable, errIdRegisEmailNotAvailable),
	ErrorUserNotFound:     add(statusUserNotFound, errEnUserNotFound, errIdUserNotFound),
	ErrorPwdNotMatch:      add(statusPwdNotMatch, errEnPwdNotMatch, errIdPwdNotMatch),
	ErrorTokenExpired:     add(statusTokenExpired, errEnTokenExpired, errIdTokenExpired),
	ErrorTokenInvalid:     add(statusTokenInvalid, errEnTokenInvalid, errIdTokenInvalid),
	ErrorNotFound:         add(statusNotFound, errEnNotFound, errIdNotFound),
	ErrorUnauthorized:     add(statusUnautorized, errEnUnauthorized, errIdUnauthorized),
	ErrorInvalidParameter: add(statusInvalidParameterID, errEnInvalidParameter, errIdInvalidParameter),
	ErrorOnlyCandidate:    add(statusOnlyCandidate, errEnOnlyCandidate, errIdOnlyCandidate),
	ErrorReviewFound:      add(statusReviewFound, ErrEnReviewFound, ErrIdReviewFound),
	SuccessCreated:        add(statusCreated, successEnCreated, successEnCreated),
}

func NewErrork(code int) error {
	return errors.New(strconv.Itoa(code))
}

func Result(w http.ResponseWriter, code int) {
	ResultWithData(w, code, nil)
}

func ResultWithData(w http.ResponseWriter, code int, data interface{}) {
	result := mapping[code]
	if data != nil {
		result.Data = data
	}

	write(w, result)
}

func ResultError(w http.ResponseWriter, code int, err error) {
	result := mapping[code]

	if result.Message["en"] == "" || result.Message["id"] == "" {
		result.Message = MapMessage{
			"en": err.Error(),
			"id": err.Error(),
		}
	}

	if err != nil {
		result.Err = err.Error()
	}

	write(w, result)
}

func write(w http.ResponseWriter, result ReturningValue) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(result.Status)

	err := json.NewEncoder(w).Encode(&result)
	if err != nil {
		panic(err)
	}
}
