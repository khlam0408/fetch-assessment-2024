package internal

import (
	"strconv"
	"math"
	"strings"
	"fetch-assessment/models"
)

func CalculatePoints(points int, currReceipt models.Receipt) int{

	// This calculates all the rules for the points. 

	totalPoints := retailerPoints(points, currReceipt) 
	totalPoints += purchaseDateOdd(points, currReceipt)
	totalPoints += timePurchaseBefore(points, currReceipt)
	totalPoints += totalRound(points, currReceipt)
	totalPoints += totalMultiple(points, currReceipt)
	totalPoints += itemMultiplePrice(points, currReceipt)
	totalPoints += pairItems(points, currReceipt)

	return totalPoints
}
func retailerPoints(points int, currReceipt models.Receipt) int{
	// Find retailer name points. Only the alphabetical characters, no numbers or symbols

	// ASCII Nums: 65 - 90 and 97 - 122 
	for i := 0; i <len(currReceipt.Retailer); i++{
		if int(currReceipt.Retailer[i]) >= 65 && int(currReceipt.Retailer[i]) <= 90{
			points++
		}
		if int(currReceipt.Retailer[i]) >= 97 && int(currReceipt.Retailer[i]) <= 122{
			points++
		}
	}
	return points
}

func purchaseDateOdd(points int, currReceipt models.Receipt) int{
		// Find if purchase date is odd (so the day not the month)

		start := len(currReceipt.PurchaseDate)-2
		dayAsStr := string(currReceipt.PurchaseDate[start:]) // This only gets the day 
	
		dayAsInt, err := strconv.Atoi(dayAsStr)
		if err != nil{
			return 0
		}
		if dayAsInt % 2 == 1{
			points += 6
		}
	return points
}

func timePurchaseBefore(points int, currReceipt models.Receipt) int{
	// Find of time of purchase is between 2:00pm and 4:00pm
	timePurchased := string(currReceipt.PurchaseTime[:2]) // This gets the hour 
	timeAsInt, err := strconv.Atoi(timePurchased)
	if err != nil{
		return 0
	}
	if timeAsInt >= 14 && timeAsInt < 16{
		points += 10
	}
	return points
}

func totalRound(points int, currReceipt models.Receipt) int{
	// Find if the total is a rounded number already (so no cents)
	start := len(currReceipt.Total) - 2
	cents := string(currReceipt.Total[start:]) // This gets the cents
	centsAsInt, err := strconv.Atoi(cents)
	if err != nil{
		return 0
	}
	if centsAsInt == 0{
		points += 50
	}
	return points
}

func totalMultiple(points int, currReceipt models.Receipt) int{
	// Find if the total is multiple of 0.25
	totalPurchased, err := strconv.ParseFloat(currReceipt.Total, 8) // Converts the total to a float
	if err != nil{
		return 0
	}
	if math.Mod(totalPurchased, 0.25) == 0{
		points += 25
	}
	return points
}

func itemMultiplePrice(points int, currReceipt models.Receipt) int{
	// If the length of the items description is a multiple of 3, multiply its price by 0.2 and round to the nearest integer
	for i := 0; i<len(currReceipt.Items); i++{
		currItem := currReceipt.Items[i]
		result := strings.TrimSpace(currItem["shortDescription"]) // trims off any leading or trailing spaces
		lenOfItem := len(result)
		if lenOfItem % 3 == 0{
			priceAsFloat, err := strconv.ParseFloat(currItem["price"], 8)
			if err != nil{
				return 0
			}
			total := math.Ceil(priceAsFloat * 0.2) // Rounds it up, even if there is just 0.01
			points += int(total)
		}

	}
	return points
}

func pairItems(points int, currReceipt models.Receipt) int{
	// For every two items, add 5 points. So if there is an odd amount of items, only count the items that have a pair. If no pair, it does not count
	numOfItems := len(currReceipt.Items)
	points += 5 * (numOfItems / 2)
	return points
}