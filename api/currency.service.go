package api

import (
	"eisen/api/data"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

func getCurrenciesPage(list []string, Offset string, Limit string) data.CurrenciesDto {
	limit, err := strconv.Atoi(Limit)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(Offset)
	if err != nil {
		offset = 0
	}
	return data.CurrenciesDto{Meta: data.PaginationMeta{Total: uint(len(list)), Limit: uint(limit), Offset: offset}, Currencies: lo.Subset(list, offset, uint(limit))}
}

func GetPage(UserId string, Direction string, Offset string, Limit string) data.CurrenciesDto {

	slice := data.Currencies
	if Direction == "sell" {
		user := data.GetUser(UserId)
		slice = lo.Keys(user.Portfolio)
		sort.Strings(slice)
	}

	return getCurrenciesPage(slice, Offset, Limit)
}

func GetBalances(UserId string, Offset string, Limit string) data.BalancesDto {
	user := data.GetUser(UserId)
	slice := lo.Keys(user.Portfolio)
	sort.Strings(slice)
	currencies := getCurrenciesPage(slice, Offset, Limit)

	return data.BalancesDto{Meta: currencies.Meta, Balances: lo.Map(currencies.Currencies, func(symbol string, _ int) data.Balance {
		return data.Balance{Currency: symbol, Balance: user.Portfolio[symbol]}
	})}
}

func Exchange(UserId string, Direction string, exchangeDto data.ExchangeDto) data.ExchangeResultDto {
	user := data.GetUser(UserId)
	balance, ok := user.Portfolio[exchangeDto.Currency]
	if ok {
		if Direction == "sell" {
			user.Portfolio[exchangeDto.Currency] = balance - exchangeDto.Amount
		} else {
			user.Portfolio[exchangeDto.Currency] = balance + exchangeDto.Amount
		}
	}
	return data.ExchangeResultDto{Success: ok, Balance: data.Balance{Currency: exchangeDto.Currency, Balance: user.Portfolio[exchangeDto.Currency]}}
}
