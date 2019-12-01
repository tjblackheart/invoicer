package app

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tjblackheart/Invoicer/backend/pkg/models"
)

func (app Application) validateToken(ts string) error {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.cfg.Secret), nil
	})

	if err == nil && token.Valid {
		return nil
	}

	return err
}

func (app Application) getUUID(r *http.Request) (string, error) {
	uuid := ""
	ts, err := app.getTokenAsString(r)

	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.cfg.Secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uuid = claims["uuid"].(string)
	}

	return uuid, nil
}

func (app Application) generateToken(u *models.User) (ts string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"uuid":  u.UUID,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})

	if ts, err = token.SignedString([]byte(app.cfg.Secret)); err != nil {
		return
	}

	return
}

func (app Application) getTokenAsString(r *http.Request) (string, error) {
	ts := r.Header.Get("Authorization")
	if len(ts) == 0 {
		return "", errors.New("Missing auth header")
	}

	ts = strings.Replace(ts, "Bearer ", "", 1)

	return ts, nil
}
