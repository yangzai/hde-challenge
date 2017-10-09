package main

import (
	"os"
	"fmt"
	"time"
	"bytes"
	"net/http"
	"encoding/json"
	"encoding/base32"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const submissionUrl = "https://hdechallenge-solve.appspot.com/challenge/003/endpoint"
//const submissionUrl = "http://localhost:8080"
const secretSuffix  = "HDECHALLENGE003"

type Data struct {
	GithubUrl		string	`json:"github_url"`
	ContactEmail	string	`json:"contact_email"`
}

func main() {
	f, err := os.Open("mission3.json")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	decoder := json.NewDecoder(f)
	data := Data{}
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
	buf, _ := json.Marshal(data)

	secret := base32.StdEncoding.EncodeToString([]byte(data.ContactEmail + secretSuffix))
	passcode, _ := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Digits: 10,
		Period: 30,
		Algorithm: otp.AlgorithmSHA512,
	})

	client := &http.Client{}
	req, err := http.NewRequest("POST", submissionUrl, bytes.NewBuffer(buf))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(data.ContactEmail, passcode)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
	fmt.Println(res.Body)
}
