package main

import (
	"fmt"
	"log"
	// "strings"
	// "net/http"
	// "encoding/json"
	"github.com/charmbracelet/huh"
	// "github.com/charmbracelet/lipgloss"
	// "os"
	// "errors"
	// "strconv"
	// xstrings "github.com/charmbracelet/x/exp/strings"
) 


type currencies struct {
  USD string
  JPY string
  EUR string
  GBP string
}

func main() {
	// accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	var SelectedCurrency string

	// Create the form
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Currency Converter").
				Description("Welcome to Currency Converter.\n").
				Next(true).
				NextLabel("Next"),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("USD $", "JPY ¥", "EUR €", "GBP £")...).
				Title("Choose Your Currency").
				Value(&SelectedCurrency),
		),
	)
  
 err := form.Run()

 if err != nil {
  log.Fatal(err)
 }
	// Output the selected currency
	fmt.Println("Selected Currency:", SelectedCurrency)
}
