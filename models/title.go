package models

import (
	"fmt"
	"time"
)

// Title título de cobrança de entrada
type Title struct {
	CreateDate     time.Time
	ExpireDateTime time.Time
	ExpireDate     string
	AmountInCents  uint64
	OurNumber      uint
	Instructions   string
}

//ValidateInstructionsLength valida se texto das instruções possui quantidade de caracteres corretos
func (t Title) ValidateInstructionsLength() error {
	if len(t.Instructions) > 220 {
		return NewErrorResponse("MPInstructions", "Instruções não podem passar de 220 caracteres")
	}
	return nil
}

//IsExpireDateValid retorna um erro se a data de expiração for inválida
func (t *Title) IsExpireDateValid() error {
	d, err := parseDate(t.ExpireDate)
	if err != nil {
		return NewErrorResponse("MPExpireDate", fmt.Sprintf("Data em um formato inválido, esperamos AAAA-MM-DD e recebemos %s", t.ExpireDate))
	}
	n, _ := parseDate(time.Now().Format("2006-01-02"))
	t.CreateDate = n
	t.ExpireDateTime = d
	if t.CreateDate.After(t.ExpireDateTime) {
		return NewErrorResponse("MPExpireDate", "Data de expiração não pode ser menor que a data de hoje")
	}
	return nil
}

//IsAmountInCentsValid retorna um erro se o valor em centavos for inválido
func (t *Title) IsAmountInCentsValid() error {
	if t.AmountInCents < 1 {
		return NewErrorResponse("MPAmountInCents", "Valor não pode ser menor do que 1 centavo")
	}
	return nil
}

// GetCreateDate Retorna a data de crição do título
func (t *Title) GetCreateDate() time.Time {
	return t.CreateDate
}

func parseDate(t string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", t)
	if err != nil {
		return time.Now(), err
	}
	return date, nil
}
