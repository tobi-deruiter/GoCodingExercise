package controller

import (
	"GoCodingExercise/models"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Post Requests

func ProcessReceipt(receipt models.Receipt) models.ProcessResponse {
	receipt = receipt.SetId(strconv.Itoa(len(models.GetReceipts()) + 1)) // set ID to index in receipts for simplicity

	// Set receipt points based on data
	var points int

	// One point for every alphanumeric character in the retailer name.
	re_alphanum := regexp.MustCompile(`[[:alnum:]]`)
	for i, _ := range receipt.Retailer {
		if re_alphanum.FindString(string(receipt.Retailer[i])) != "" {
			points++
		}
	}

	// 50 points if the total is a round dollar amount with no cents.
	if strings.Split(receipt.Total, ".")[1] == "00" {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil && int(total*100)%25 == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt.
	points += 5 * (len(receipt.Items) / 2)

	// If the trimmed length of the item description is a multiple of 3, multiply the price by
	// 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		if len(item.ShortDescription)%3 == 0 {
			p, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(p * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd.
	if day, _ := strconv.Atoi(strings.Split(receipt.PurchaseDate, "-")[2]); day%2 == 1 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if hour, _ := strconv.Atoi(strings.Split(receipt.PurchaseTime, ":")[0]); hour >= 14 && hour <= 16 {
		points += 10
	}

	receipt = receipt.SetPoints(points)

	// add processed receipt to list
	models.AddReceipt(receipt)

	return models.NewProcessResponse(receipt.Id())
}

// Get Requests

func GetPoints(pointsRequest models.PointsRequest) models.PointsResponse {
	return models.NewPointsResponse(models.GetReceipts()[pointsRequest.Idx()].Points())
}
