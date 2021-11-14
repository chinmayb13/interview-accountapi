package handlers

import (
	"encoding/json"
	"fmt"
	"interview-accountapi-demo/models"
	"io/ioutil"
	"net/http"
)

//constants for the urls of downstream API
var (
	posturl   = "https://api.staging-form3.tech/v1/organisation/accounts"
	geturl    = "https://api.staging-form3.tech/v1/organisation/accounts/%s"
	deleteurl = "https://api.staging-form3.tech/v1/organisation/accounts/%s?version=%s"
)

//handler function for POST req
func CreateHandler(client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req, err := http.NewRequest(http.MethodPost, posturl, r.Body) //create a new http POST request for downstream API
		// here we are directly passing the req body to downstream API

		if err != nil { //error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := client.Do(req) //call the downstream API

		if err != nil { //error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if resp.StatusCode < 200 || resp.StatusCode > 204 { // if response code is not success
			http.Error(w, "response returned non 2xx status", resp.StatusCode)
			return
		}
		defer resp.Body.Close() //close the body before returning

		body, err := ioutil.ReadAll(resp.Body) // read response body

		if err != nil { //error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		respData := &models.PostReq{} // define response structure

		err = json.Unmarshal(body, respData) //unmmarshal the response body to the response structure

		if err != nil { //error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated) //update statuscode of the response to be returned
		fmt.Fprint(w, respData)           //write the data to writer

	}
}

//handler func for GET account request
func GetHandler(client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accID := r.URL.Query().Get("account_id") // get query param value for account_id

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(geturl, accID), nil) //create a new http GET request for downstream API

		if err != nil { //error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := client.Do(req) //call downstream API

		if err != nil { //error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if resp.StatusCode < 200 || resp.StatusCode > 204 { //if response code is not success
			http.Error(w, "response returned non 2xx status", resp.StatusCode)
			return
		}

		defer resp.Body.Close() //close resp body before returning

		body, err := ioutil.ReadAll(resp.Body) //read response body

		if err != nil { //error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		respData := &models.PostReq{} //define response structure

		err = json.Unmarshal(body, respData) //unmarshal the response body to response struct

		if err != nil { //error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK) //update the response status code
		fmt.Fprint(w, respData)      //write to the writer

	}
}

//handler func for DELETE account request
func DeleteHandler(client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accID := r.URL.Query().Get("account_id") //get query param value of account_id
		version := r.URL.Query().Get("version")  //get query param value of version

		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(deleteurl, accID, version), nil) //create new http DELETE request for downstream API

		if err != nil { //error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := client.Do(req) //call downstream API

		if err != nil { //error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if resp.StatusCode < 200 || resp.StatusCode > 204 { //if response code is not success
			http.Error(w, "response returned non 2xx status", resp.StatusCode)
			return
		}

		w.WriteHeader(http.StatusNoContent) //update the status code of the response to be returned
		fmt.Fprint(w, "Success")            //write to the writer
	}
}
