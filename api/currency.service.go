package api

import (
	"eisen/api/data"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

func GetPage(UserId string, Direction string, Offset string, Limit string) data.CurrenciesDto {
	limit, err := strconv.Atoi(Limit)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(Offset)
	if err != nil {
		offset = 0
	}

	slice := data.Currencies
	if Direction == "sell" {
		user := data.GetUser(UserId)
		slice = lo.Keys(user.Portfolio)
		sort.Strings(slice)
	}

	return data.CurrenciesDto{Meta: data.PaginationMeta{Total: uint(len(slice)), Limit: uint(limit), Offset: offset}, Currencies: lo.Subset(slice, offset, uint(limit))}
}
