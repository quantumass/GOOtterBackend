package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

const (
	PayPalEndpoint = "https://api.sandbox.paypal.com/v1/payments/payment"
)

type PaymentInfo struct {
	PaymentId  string `json:"paymentId"`
	CreateTime string `json:"create_time"`
	ID         string `json:"id"`
	Intent     string `json:"intent"`
	Links      []struct {
		Href   string `json:"href"`
		Method string `json:"method"`
		Rel    string `json:"rel"`
	} `json:"links"`
	Payer struct {
		Address struct {
			AddressLine1 string `json:"address_line_1"`
			AddressLine2 string `json:"address_line_2"`
			AdminArea1   string `json:"admin_area_1"`
			AdminArea2   string `json:"admin_area_2"`
			CountryCode  string `json:"country_code"`
			PostalCode   string `json:"postal_code"`
		} `json:"address"`
		EmailAddress string `json:"email_address"`
		Name         struct {
			GivenName string `json:"given_name"`
			Surname   string `json:"surname"`
		} `json:"name"`
		PayerID string `json:"payer_id"`
		Phone   struct {
			PhoneNumber struct {
				NationalNumber string `json:"national_number"`
			} `json:"phone_number"`
		} `json:"phone"`
	} `json:"payer"`
	PurchaseUnits []struct {
		Amount struct {
			CurrencyCode string `json:"currency_code"`
			Value        string `json:"value"`
		} `json:"amount"`
		Description string `json:"description"`
		Payee       struct {
			EmailAddress string `json:"email_address"`
			MerchantID   string `json:"merchant_id"`
		} `json:"payee"`
		Payments struct {
			Captures []struct {
				Amount struct {
					CurrencyCode string `json:"currency_code"`
					Value        string `json:"value"`
				} `json:"amount"`
				CreateTime       string `json:"create_time"`
				FinalCapture     bool   `json:"final_capture"`
				ID               string `json:"id"`
				SellerProtection struct {
					DisputeCategories []string `json:"dispute_categories"`
					Status            string   `json:"status"`
				} `json:"seller_protection"`
				Status     string `json:"status"`
				UpdateTime string `json:"update_time"`
			} `json:"captures"`
		} `json:"payments"`
		ReferenceID string `json:"reference_id"`
		Shipping    struct {
			Address struct {
				AddressLine1 string `json:"address_line_1"`
				AddressLine2 string `json:"address_line_2"`
				AdminArea1   string `json:"admin_area_1"`
				AdminArea2   string `json:"admin_area_2"`
				CountryCode  string `json:"country_code"`
				PostalCode   string `json:"postal_code"`
			} `json:"address"`
			Name struct {
				FullName string `json:"full_name"`
			} `json:"name"`
		} `json:"shipping"`
		SoftDescriptor string `json:"soft_descriptor"`
	} `json:"purchase_units"`
	Status     string `json:"status"`
	UpdateTime string `json:"update_time"`
}

func validatePayment(paymentID string) bool {
	// Set up your PayPal REST API credentials
	clientID := "AfYNudfz7iM75muJHFSMqy96LYDYqocNH7LBmoH5Tn4C8aMHVc5siQaZTIRUEDxVp01AtfX2juq4zIdR"
	secret := "EPemTyCa_yWtJdgWGlM3z_xuURgKv7XUHm7U5ouQGJLUKeeswXjkOs2Fot8OpA-gEY3_kGvInJ5qVPgj"

	// Set up the request headers and data
	auth := fmt.Sprintf("%s:%s", clientID, secret)
	authHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth)))
	data := url.Values{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", PayPalEndpoint, paymentID), strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/json")

	// Make the API call to verify the payment
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check the payment status in the response
	if resp.StatusCode == 200 {
		var payment map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&payment)
		if err != nil {
			panic(err)
		}

		if payment["state"] == "approved" {
			fmt.Println("Payment was successful!")
			return true
		} else {
			fmt.Println("Payment was not successful.")
			return false
		}

	} else {
		fmt.Printf("There was an error verifying the payment: %s\n", resp.Status)
		return false
	}
}

func main() {
	app := pocketbase.New()

	app.OnBeforeBootstrap().Add(func(e *core.BootstrapEvent) error {
		log.Println("Hello world, the server is starting, ready ? start !")
		return nil
	})

	app.OnModelBeforeUpdate().Add(func(e *core.ModelEvent) error {
		if record, ok := e.Model.(*models.Record); ok && record.Collection().Name == "users" {
			var payment PaymentInfo
			paymentAsJson := record.Get("paymentInfo").(types.JsonRaw)
			err := json.Unmarshal(paymentAsJson, &payment)
			if err != nil {
				fmt.Println("Error unmarshalling JSON:", err)
				return nil
			}
			if paymentAsJson != nil {
				record.Set("credits", 300)
				record.Set("paymentInfo", "")
				record.Set("previousPaymentInfo", paymentAsJson)
			}
		}
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
