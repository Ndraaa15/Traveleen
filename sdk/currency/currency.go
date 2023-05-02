package currency

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leekchan/accounting"
)

func FormatRupiah(price float64) string {
	ac := accounting.Accounting{Symbol: "Rp. ", Precision: 0, Thousand: ".", Decimal: ""}
	formattedPrice := ac.FormatMoney(price)

	return formattedPrice
}

func ConvertRupiahIntoFloat(price string) float64 {
	str := strings.Replace(price, ".", "", -1)

	convertPrice, err := strconv.ParseFloat(str[3:], 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	return convertPrice
}
