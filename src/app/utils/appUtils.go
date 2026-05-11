package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ValidadIdMongo(id string) (*bson.ObjectID, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID de mongo invalido")
	}
	return &objectID, nil
}

func ValidarMongoIdArray(id []string) (IDS []bson.ObjectID, err error) {
	var ids []bson.ObjectID
	for _, v := range id {
		objID, err := ValidadIdMongo(v)
		if err != nil {
			return nil, err
		}
		ids = append(ids, *objID)
	}
	return ids, nil
}
func FechaHoraBolivia() time.Time {
	fecha := time.Now()
	return fecha.Add(-4 * time.Hour)

}
func PaginadorHTTP(r *http.Request) (int, int, error) {
	query := r.URL.Query()
	paginaStr := query.Get("pagina")
	if paginaStr == "" {
		paginaStr = "1"
	}
	limiteStr := query.Get("limite")
	if limiteStr == "" {
		limiteStr = "20"
	}

	pagina, err := strconv.Atoi(paginaStr)
	if err != nil {
		return 0, 0, errors.New("Ingrese el numero pagina")
	}
	if pagina <= 0 {
		return 0, 0, errors.New("El pagina deve ser mayor a 0")
	}

	limite, err := strconv.Atoi(limiteStr)
	if err != nil {
		return 0, 0, errors.New("Ingrese el numero limite")
	}
	if limite <= 0 {
		return 0, 0, errors.New("El limite deve ser mayor a 0")
	}

	return pagina, limite, nil
}
func ResponseJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
func ValidarContrasena(password string) error {
	if len(password) < 8 {
		return errors.New("la contraseña debe tener al menos 8 caracteres")
	}
	mayuscula := regexp.MustCompile(`[A-Z]`)
	if !mayuscula.MatchString(password) {
		return errors.New("la contraseña debe contener al menos una letra mayúscula")
	}
	minuscula := regexp.MustCompile(`[a-z]`)
	if !minuscula.MatchString(password) {
		return errors.New("la contraseña debe contener al menos una letra minúscula")
	}
	numero := regexp.MustCompile(`[0-9]`)
	if !numero.MatchString(password) {
		return errors.New("la contraseña debe contener al menos un número")
	}
	simbolo := regexp.MustCompile(`[!@#~$%^&*()+|_]`)
	if !simbolo.MatchString(password) {
		return errors.New("la contraseña debe contener al menos un símbolo especial")
	}

	return nil
}
