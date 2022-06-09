package helper

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/iitheo/theofetchrewards/pkg/app/dto/pointsdto"
	"github.com/iitheo/theofetchrewards/pkg/app/models/pointsmodel"
	"strings"
	"time"
)

func IsAddPointsValid(p *pointsdto.AddPointsDTO) (*pointsmodel.TransactionDB, error) {
	p.Payer = strings.TrimSpace(strings.ToUpper(p.Payer))
	p.Timestamp = strings.TrimSpace(p.Timestamp)

	if p.Points <= 0 {
		return nil, errors.New("points must be greater than 0")
	}

	if p.Payer == "" {
		return nil, errors.New("name cannot be empty")
	}

	tt, err := time.Parse(time.RFC3339, p.Timestamp)
	if err != nil {
		return nil, err
	}

	id := uuid.New()
	idStr := fmt.Sprintf("%s", id)

	return &pointsmodel.TransactionDB{
		ID:        idStr,
		Payer:     p.Payer,
		Points:    p.Points,
		Timestamp: tt,
	}, nil
}
