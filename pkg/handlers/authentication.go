package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tjblackheart/invoicer/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

// JWTSecret holds the signing secret
var JWTSecret string

// Credentials holds the login request
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login validates a login request.
func Login(w http.ResponseWriter, r *http.Request) {

	var c Credentials
	var u User
	var err error

	if err = json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if db.Conn.Preload("Settings").Where("email = ?", c.Email).First(&u).RecordNotFound() {
		http.Error(w, "Invalid credentials", http.StatusForbidden)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(c.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusForbidden)
		return
	}

	ts, err := generateToken(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"token": ts, "user": u})
}

// ValidateToken validates a JWT
func ValidateToken(ts string) error {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})

	if err == nil && token.Valid {
		return nil
	}

	return err
}

// GetTokenAsString returns the token from a request
func GetTokenAsString(r *http.Request) (ts string, err error) {
	ts = r.Header.Get("Authorization")

	if len(ts) == 0 {
		err = errors.New("Missing auth header")
		return
	}

	ts = strings.Replace(ts, "Bearer ", "", 1)

	return
}

// GetUUID extracts the UUID from JWT
func GetUUID(r *http.Request) (uuid string, err error) {
	ts, err := GetTokenAsString(r)
	if err != nil {
		return
	}

	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uuid = claims["uuid"].(string)
	}

	return
}

//

func generateToken(u *User) (ts string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"uuid":  u.UUID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		//"exp": time.Now().Add(time.Minute).Unix(),
	})

	if ts, err = token.SignedString([]byte(JWTSecret)); err != nil {
		return
	}

	return
}
