package utils

import (
	"encoding/json"
	"net/http"
)

type netException struct {
    ErrorName   string `json:"error"`
    ErrorCode   int    `json:"code"`
}

func NetThrowError(w http.ResponseWriter, msg string, code int) {
    w.WriteHeader(code)
    newExcp := netException{
        ErrorName: msg,
        ErrorCode: code,
    }

    json.NewEncoder(w).Encode(newExcp)
    return
}




