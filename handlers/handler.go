package handlers

import (
	"encoding/json"
	"fmt"
	"interview-accountapi-demo/models"
	"io/ioutil"
	"net/http"
)

var (
	posturl   = "https://api.staging-form3.tech/v1/organisation/accounts"
	geturl    = "https://api.staging-form3.tech/v1/organisation/accounts/%s"
	deleteurl = "https://api.staging-form3.tech/v1/organisation/accounts/%s?version=%s"
)

func CreateHandler(client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req, err := http.NewRequest(http.MethodPost, posturl, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if resp.StatusCode < 200 || resp.StatusCode > 204 {
			http.Error(w, "response returned non 2xx status", resp.StatusCode)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		respData := &models.AccountData{}
		err = json.Unmarshal(body, respData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, respData)

	}
}

func GetHandler(client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.RequestURI())
		accID := r.URL.Query().Get("account_id")
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(geturl, accID), nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(req.URL.RequestURI())
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if resp.StatusCode < 200 || resp.StatusCode > 204 {
			http.Error(w, "response returned non 2xx status", resp.StatusCode)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		respData := &models.AccountData{}
		err = json.Unmarshal(body, respData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, respData)

	}
}

func DeleteHandler(client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accID := r.URL.Query().Get("account_id")
		version := r.URL.Query().Get("version")

		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(deleteurl, accID, version), nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if resp.StatusCode < 200 || resp.StatusCode > 204 {
			http.Error(w, "response returned non 2xx status", resp.StatusCode)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, "Success")
	}
}
