package numeric

import (
	"fmt"
	"strconv"
)

func RoundingRating(rating float64) float64 {
	ratingStr := fmt.Sprintf("%.2f", rating)

	ratingRounded, err := strconv.ParseFloat(ratingStr, 64)

	if err != nil {
		return -1
	}

	return ratingRounded
}
