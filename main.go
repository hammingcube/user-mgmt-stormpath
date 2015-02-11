package main

import (
	"github.com/darkhelmet/env"
	"github.com/jmcvetta/randutil"
	"github.com/jmcvetta/stormpath"
	"log"
)

func setupApplication() *stormpath.Application {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
	spApp := env.String("STORMPATH_APP")
	apiId := env.String("STORMPATH_API_ID")
	apiSecret := env.String("STORMPATH_API_SECRET")
	s := stormpath.Application{
		Href:      spApp,
		ApiId:     apiId,
		ApiSecret: apiSecret,
	}
	return &s
}

func createAccountTemplate() stormpath.Account {
	rnd, err := randutil.AlphaString(8)
	if err != nil {
		log.Fatal(err)
	}
	email := "jason.mcvetta+" + rnd + "@gmail.com"
	password := rnd + "Xy123" // Ensure we meet password requirements
	tmpl := stormpath.Account{
		Username:   rnd,
		Email:      email,
		Password:   password,
		GivenName:  "James",
		MiddleName: "T",
		Surname:    "Kirk",
	}
	return tmpl
}

func TestCreateAccount() {
	app := setupApplication()
	tmpl := createAccountTemplate()
	_, err := app.CreateAccount(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	//
	// Cleanup
	//
	//acct.Delete()
}

func main() {
	TestCreateAccount()
}