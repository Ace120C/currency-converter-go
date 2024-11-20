package main

import (
	"fmt"
	// "strings"
	// "net/http"
	// "encoding/json"
	"github.com/charmbracelet/huh"
	// "github.com/charmbracelet/lipgloss"
	// "os"
	// "errors"
	"strconv"
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

	var SelectedCurrency, SelectedCurrency2 string
  var WantedCurrency string

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
      huh.NewSelect[string]().
				Options(huh.NewOptions("USD $", "JPY ¥", "EUR €", "GBP £")...).
        Title("Which Currency to convert to? ").
				Value(&SelectedCurrency2),
      ),
    huh.NewGroup(
        huh.NewInput().
        Title("How much? ").
        Prompt(">").
        Value(&WantedCurrency),
	  ),
  )
   

 err := form.Run()

 Wanted_CurrencyInt, err := strconv.Atoi(WantedCurrency)

 if err != nil {
    panic("illegal character: it needs to be a number!")
 }

 	// Output the selected currency

  fmt.Println("Selected Currency:", SelectedCurrency, "Converted Currency:", SelectedCurrency2, "Wanted Currency:", Wanted_CurrencyInt)
}
