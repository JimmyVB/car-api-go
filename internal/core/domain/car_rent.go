package domain

import (
	"time"
)

type CarRent struct {
	ID        int       `json:"id"`
	IdCar     int       `json:"idcar"`
	IdUser    int       `json:"iduser"`
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
}
