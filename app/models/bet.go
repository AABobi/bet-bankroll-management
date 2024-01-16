package models

import (
	"ap/db"

	"fmt"
	"time"
	//	"time"
)

type Bet struct {
	ID       int64
	UserID   int64
	BetValue int64
	Date     time.Time
	Tax      int64
	Win      bool
	Result   int64
}

func (b Bet) SaveBet() error {
	query := "INSERT INTO bets(userID, betValue, data, tax, win, result) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	fmt.Println("SaveBet", b.UserID, b.BetValue, b.Date, b.Tax, b.Win)
	_, err = stmt.Exec(b.UserID, b.BetValue, b.Date, b.Tax, b.Win, b.Result)

	if err != nil {
		return err
	}

	return nil
}

func GetAllBets() ([]Bet, error) {
	query := "SELECT * FROM bets"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bets []Bet
	for rows.Next() {
		var bet Bet

		err := rows.Scan(&bet.ID, &bet.UserID, &bet.BetValue, &bet.Date, &bet.Tax, &bet.Win, &bet.Result)

		if err != nil {
			return nil, err
		}

		bets = append(bets, bet)
	}

	return bets, nil
}
