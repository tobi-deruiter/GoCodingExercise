package middleware

import (
	"GoCodingExercise/models"
	"encoding/json"
	"net/http"
	"reflect"
	"regexp"
)

/*
	PRIVATE
*/

// Remove leading/trailing spaces and make sure field is not blank/null

/*
	PUBLIC
*/

// Check for bad requests, returns status and Receipt if ok; BadRequest if not
func IsValidReceipt(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	// get receipt from request body
	receipt := models.Receipt{}
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		return 400, models.NewBadRequest()
	}

	// check each element that it follows the required format
	re_removeSpaces := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_retailer := regexp.MustCompile(`^[\w\s\-&]+$`)
	re_pDate := regexp.MustCompile(`^\d{4}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$`)
	re_pTime := regexp.MustCompile(`^([01][0-9]|2[0-4]):([0-5][0-9])$`)
	re_iDesc := regexp.MustCompile(`^[\w\s\-]+$`)
	re_price := regexp.MustCompile(`^\d+\.\d{2}$`)

	receipt.Retailer = re_retailer.FindString(receipt.Retailer)
	receipt.PurchaseDate = re_pDate.FindString(receipt.PurchaseDate)
	receipt.PurchaseTime = re_pTime.FindString(receipt.PurchaseTime)
	receipt.Total = re_price.FindString(receipt.Total)

	for i, item := range receipt.Items {
		item.ShortDescription = re_removeSpaces.ReplaceAllString(re_iDesc.FindString(item.ShortDescription), "")
		item.Price = re_price.FindString(item.Price)
		receipt.Items[i] = item
	}

	// if any field didnt follow the format, it will be left blank
	// if blank, it as a bad request
	var s string
	v := reflect.ValueOf(receipt)
	for i := range v.NumField() {
		f := v.Field(i)
		name := v.Type().Field(i).Name
		if name != "id" && name != "points" {
			if f.Type() == reflect.TypeOf(s) {
				if f.String() == "" {
					return 400, models.NewBadRequest()
				}
			} else {
				items, _ := f.Interface().([]models.Item)
				if len(items) == 0 {
					return 400, models.NewBadRequest()
				}
				for j := range items {
					w := reflect.ValueOf(items[j])
					for k := range w.NumField() {
						g := w.Field(k)
						if g.String() == "" {
							return 400, models.NewBadRequest()
						}
					}
				}
			}
		}
	}

	return 200, receipt
}

// Check if ID is valid, returns status and PointsRequest if ok; NotFound if not
func IsValidId(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	id := r.PathValue("id")
	for i, r := range models.GetReceipts() {
		if r.Id() == id {
			return 200, models.NewPointsRequest(r.Id(), i)
		}
	}
	return 404, models.NewNotFound()
}
