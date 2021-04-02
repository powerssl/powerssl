package internal

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/square/go-jose.v2"
)

func jwksEndpoint(signKeys ...*rsa.PrivateKey) (func(w http.ResponseWriter, req *http.Request), error) {
	jsonWebKeys := make([]jose.JSONWebKey, len(signKeys))
	for i, signKey := range signKeys {
		jsonWebKey := jose.JSONWebKey{
			Algorithm: "RS256",
			Key:       signKey,
			Use:       "sig",
		}
		publicJSONWebKey := jsonWebKey.Public()
		thumbprint, err := publicJSONWebKey.Thumbprint(crypto.SHA1)
		if err != nil {
			return nil, err
		}
		publicJSONWebKey.KeyID = base64.URLEncoding.EncodeToString(thumbprint)
		jsonWebKeys[i] = publicJSONWebKey
	}
	jwks, err := json.Marshal(&jose.JSONWebKeySet{Keys: jsonWebKeys})
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/jwk+json")
		_, _ = fmt.Fprintln(w, string(jwks))
	}, nil
}
