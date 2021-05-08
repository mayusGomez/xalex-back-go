package services

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/mayusGomez/xalex/billing"
	"github.com/mayusGomez/xalex/billing/storage"
)

func CreateQuote(quote *billing.Quote, s storage.QuoteStorage) error {

	if quote.IDUser == "" {
		return errors.New("Wrong input values")
	}

	now := time.Now()
	hex_num := strconv.FormatInt(now.UnixNano(), 16)
	quote.RegisterDate = now
	quote.Professional = strings.ToUpper(quote.Professional)
	quote.Status = billing.PendingQuote
	quote.Code = hex_num

	for _, note := range quote.Notes {
		note.CreatedAt = time.Now()
	}

	err := s.Create(quote)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetQuotesByPage(IDUser string, quoteStatus []billing.QuoteStatus, filterField, fielterData string, pageNumber, pageSize int64, s storage.QuoteStorage) ([]billing.Quote, int64, error) {

	fielterData = strings.ToUpper(fielterData)
	quotes, total, err := s.GetQuotesByPage(IDUser, quoteStatus, filterField, fielterData, pageNumber, pageSize)

	for i, quote := range quotes {
		quote.SetMoneyToFloat()
		quotes[i] = quote
	}

	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	return quotes, total, nil
}

func UpdateQuote(quote *billing.Quote, s storage.QuoteStorage) error {

	if quote.IDUser == "" {
		return errors.New("Wrong input values")
	}

	quote.Professional = strings.ToUpper(quote.Professional)
	quote.SetMoneyToInt()

	log.Println("quote, update quote:", quote)

	err := s.Update(quote)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetQuote(userId, quoteId string, s storage.QuoteStorage) (billing.Quote, error) {

	quote, err := s.Get(quoteId)

	if quote.IDUser != userId {
		fmt.Println("Error, user not related", err)
		return billing.Quote{}, nil
	}

	quote.SetMoneyToFloat()
	if err != nil {
		fmt.Println("Error, Receive err from storage", err)
		return billing.Quote{}, err
	}
	return quote, nil
}

func AddQuoteNotes(userId, id string, note *billing.Notes, s storage.QuoteStorage) error {

	if userId == "" {
		return errors.New("Wrong input values")
	}

	note.CreatedAt = time.Now()

	err := s.AddNote(userId, id, note)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
