package usecase

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/me/financial/pkg/entity"
	"github.com/me/financial/pkg/repository"
)

type InstallmentUseCase interface {
	CreateInstallment(purchase entity.Purchase) error
}

type Installment struct {
	repositoryInstallment repository.RepositoryAll
}

func NewInstallmentUseCase(r repository.RepositoryAll) InstallmentUseCase {
	return &Installment{
		repositoryInstallment: r,
	}
}

func (i *Installment) CreateInstallment(purchase entity.Purchase) error {
	var (
		installment = purchase.Installment
		first = true
		month = purchase.Date
		err   error
	)

	slog.Info(fmt.Sprintf("purchaseID: %s", installment.PurchaseID.String()))

	for j := 1; j <= installment.Number; j++ {
		installment.ID = uuid.New()
		installment.Description = fmt.Sprintf("Parcela %d de %d", j, installment.Number)
		installment.Value = purchase.Amount / float64(installment.Number)

		month, first, err = calculeteDate(i, first, month, purchase.IDCreditCard)
		if err != nil {
			return err
		}
		installment.Month = month
		installment.Paid = false

		if err := i.repositoryInstallment.All().Installment.Create(installment); err != nil {
			return err
		}
	}

	return nil
}

func calculeteDate(i *Installment, first bool, date string, id uuid.UUID) (string, bool, error) {
	var (
		cc entity.CreditCard
		err error
		msg error
	)

	completeDate, err1 := time.Parse("2006-01-02", date)
	if err1 != nil {
		msg = fmt.Errorf("Error parsing date: %v", err1)
	}

	if first {
		cc, err = i.repositoryInstallment.All().CreditCard.FindByID(id)
		if err != nil {
			return "", false, err
		}

		if completeDate.Day() >= cc.InvoiceClosingDay {
			completeDate = completeDate.AddDate(0, 1, 0)

			return completeDate.Format("2006-01-02"), false, nil
		} else {
			return completeDate.Format("2006-01-02"), false, nil
		}
	}

	newDate, err2 := time.Parse("2006-01-02", date)
	if err1 != nil && err2 != nil {
		return "", false, msg 
	}

	newDate = newDate.AddDate(0, 1, 0)

	return newDate.Format("2006-01-02"), false, nil			
}
