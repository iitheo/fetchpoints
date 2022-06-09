package pointsrepo

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/iitheo/theofetchrewards/pkg/app/models/pointsmodel"
	"sort"
	"time"
)

var listOfPointsFromRepo = make([]pointsmodel.TransactionDB, 0)

func init() {
	id1 := uuid.New()
	id1Str := fmt.Sprintf("%s", id1)
	point1 := pointsmodel.TransactionDB{
		ID:        id1Str,
		Payer:     "DANNON",
		Points:    1000,
		Timestamp: time.Date(2020, time.Month(11), 02, 14, 00, 00, 000, time.Local),
	}

	id2 := uuid.New()
	id2Str := fmt.Sprintf("%s", id2)
	point2 := pointsmodel.TransactionDB{
		ID:        id2Str,
		Payer:     "UNILEVER",
		Points:    200,
		Timestamp: time.Date(2020, time.Month(10), 31, 11, 00, 00, 000, time.Local),
	}

	id3 := uuid.New()
	id3Str := fmt.Sprintf("%s", id3)
	point3 := pointsmodel.TransactionDB{
		ID:        id3Str,
		Payer:     "DANNON",
		Points:    250,
		Timestamp: time.Date(2020, time.Month(10), 31, 15, 00, 00, 000, time.Local),
	}

	id4 := uuid.New()
	id4Str := fmt.Sprintf("%s", id4)
	point4 := pointsmodel.TransactionDB{
		ID:        id4Str,
		Payer:     "MILLER COORS",
		Points:    10000,
		Timestamp: time.Date(2020, time.Month(11), 01, 14, 00, 00, 000, time.Local),
	}

	id5 := uuid.New()
	id5Str := fmt.Sprintf("%s", id5)
	point5 := pointsmodel.TransactionDB{
		ID:        id5Str,
		Payer:     "DANNON",
		Points:    300,
		Timestamp: time.Date(2020, time.Month(10), 31, 10, 00, 00, 000, time.Local),
	}
	listOfPointsFromRepo = append(listOfPointsFromRepo, point1, point2, point3, point4, point5)
}
func SavePoints(newPoint *pointsmodel.TransactionDB) {
	listOfPointsFromRepo = append(listOfPointsFromRepo, *newPoint)
}

func GetAllPointsFromRepo() []pointsmodel.TransactionDB {
	sort.Slice(listOfPointsFromRepo, func(i, j int) bool {
		return listOfPointsFromRepo[i].Timestamp.Unix() > listOfPointsFromRepo[j].Timestamp.Unix()
	})
	return listOfPointsFromRepo
}

func SpendPoints(pointsToSpend int) []pointsmodel.TransactionDB {

	sort.Slice(listOfPointsFromRepo, func(i, j int) bool {
		return listOfPointsFromRepo[i].Timestamp.Unix() > listOfPointsFromRepo[j].Timestamp.Unix()
	})
	var amountWithDrawn int

	amountWithDrawn = pointsToSpend
	for k := range listOfPointsFromRepo {
		if amountWithDrawn == 0 {
			break
		}

		if listOfPointsFromRepo[k].Points >= amountWithDrawn {
			listOfPointsFromRepo[k].Points -= amountWithDrawn
			amountWithDrawn = 0

		} else if listOfPointsFromRepo[k].Points < amountWithDrawn && listOfPointsFromRepo[k].Points > 0 {
			amountWithDrawn -= listOfPointsFromRepo[k].Points
			listOfPointsFromRepo[k].Points = 0
		}
	}

	return listOfPointsFromRepo
}
