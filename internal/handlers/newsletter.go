package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"

	"myBlog/internal/models"
	"myBlog/internal/utils"
)

var unverified_emails *utils.BiMap = utils.NewBiMap()

func validateEmail(email string) bool {
	utils.InfoLogger.Printf("Validating email: %s", email)
	// Define a regular expression for validating email addresses
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func emailRegistered(email string) (bool, error) {
	utils.InfoLogger.Printf("Checking if email is registered: %s", email)
	address := os.Getenv("API_ADDRESS") + "/newsletter-members?filters[email][$eq]=" + email

	req, _ := http.NewRequest("GET", address, nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.ErrorLogger.Printf("Error retrieving newsletter member: %v", err)
		return false, errors.New("trouble at retrieveing newsletter memeber")
	}

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if strings.Contains(string(body), `{"data":[]`) {
		return false, nil
	} else {
		return true, nil
	}
}

func regEmail(email string) error {
	utils.InfoLogger.Printf("Registering email: %s", email)
	address := os.Getenv("API_ADDRESS") + "/newsletter-members"
	data := "{\"data\":{\"email\":\"" + email + "\",\"uuid\":\"" + uuid.New().String() + "\"}}"

	req, err := http.NewRequest("POST", address, bytes.NewReader([]byte(data)))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		utils.ErrorLogger.Printf("Error creating registration request: %v", err)
		return errors.New("trouble in creating the registeration request")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.ErrorLogger.Printf("Error registering email: %v", err)
		return errors.New("trouble at registering the email")
	}

	if resp.StatusCode == 201 {
		return nil
	} else {
		utils.ErrorLogger.Printf("Failed to register email, status code: %d", resp.StatusCode)
		return errors.New("didn't success in adding in email" + " " + string(resp.StatusCode))
	}
}

func preRegEmail(email string) error {
	utils.InfoLogger.Printf("Pre-registering email: %s", email)
	registered, _ := emailRegistered(email)
	if registered {
		return nil
	}

	if _, ok := unverified_emails.Reverse[email]; ok {
		return errors.New("already sent a verification link to this email")
	}

	uuid_str := uuid.New().String()
	unverified_emails.Forward[uuid_str] = email
	unverified_emails.Reverse[email] = uuid_str
	time.AfterFunc(10*time.Minute, func() {
		delete(unverified_emails.Forward, uuid_str)
		delete(unverified_emails.Reverse, email)
	})

	return nil
}

func HandleNewsletterReg(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println("Handling newsletter registration")
	var v models.Email

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ErrorLogger.Println(`{"err": "Couldn't read all input"}`)
		return
	}

	err = json.Unmarshal(body, &v)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "{\"err\": \"Couldn't parse JSON\"}")
		return
	}
	if !validateEmail(v.Email) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "{\"err\": \"Invalid Email\"}")
		return
	}

	err = preRegEmail(v.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "{\"err\": "+err.Error()+"}")
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "{\"success\": \"true\"}")
}

func HandleVerification(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println("Handling email verification")
	uuid := r.URL.Path[len("/verifyEmail/"):]

	re := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	if !re.MatchString(uuid) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid UUID")
		utils.ErrorLogger.Printf("Invalid UUID: %s", uuid)
		return
	}

	if v, ok := unverified_emails.Forward[uuid]; ok {
		regEmail(v)
		fmt.Fprintln(w, `{"success": 1}`)
		utils.InfoLogger.Printf("Email verified and registered: %s", v)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid UUID")
		utils.ErrorLogger.Printf("Invalid UUID: %s", uuid)
		return
	}
}
