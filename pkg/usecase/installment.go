package usecase

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/me/finance/pkg/dto"
	"github.com/me/finance/pkg/entity"
	"github.com/me/finance/pkg/repository"
)

type InstallmentUseCase interface {
	CreateInstallment(purchase entity.Purchase) error
	UpdateInstalment(id uuid.UUID) error
	DeleteInstallment(purchaseID uuid.UUID) error
	FindInstallmentByPurchaseID(id uuid.UUID) (dto.InstallmentResponse, error)
	FindInstallmentByMonth(month string) (dto.InstallmentResponse, error)
	FindInstallmentByNotPaid() (dto.InstallmentResponse, error)
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
		first       = true
		month       = purchase.Date
		err         error
	)

	slog.Info(fmt.Sprintf("purchaseID: %s", installment.PurchaseID.String()))

	for j := 1; j <= installment.Number; j++ {
		installment.ID = uuid.New()
		installment.Description = fmt.Sprintf("Parcela %d de %d", j, installment.Number)
		installment.Value = purchase.Amount / float64(installment.Number)

		month, err = calculeteDateNextInvoice(i, first, month, purchase.IDCreditCard)
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
func (i *Installment) UpdateInstalment(id uuid.UUID) error {
	if err := i.repositoryInstallment.All().Installment.Update(id); err != nil {
		return fmt.Errorf("error updating installment: %v", err)
	}

	return nil
}

func (i *Installment) DeleteInstallment(purchaseID uuid.UUID) error {
	if err := i.repositoryInstallment.All().Installment.Delete(purchaseID); err != nil {
		return fmt.Errorf("error deleting installment: %v", err)
	}

	return nil
}

func (i *Installment) FindInstallmentByPurchaseID(id uuid.UUID) (dto.InstallmentResponse, error) {
	installments, err := i.repositoryInstallment.All().Installment.FindByPurchaseID(id)
	if err != nil {
		return dto.InstallmentResponse{}, fmt.Errorf("error finding installment by purchaseID: %v", err)
	}

	response := processInstallmentResponse(installments)

	return response, nil
	
}

func (i *Installment) FindInstallmentByMonth(month string) (dto.InstallmentResponse, error) {
	installments, err := i.repositoryInstallment.All().Installment.FindByMonth(month)
	if err != nil {
		return dto.InstallmentResponse{}, fmt.Errorf("error finding installment by month: %v", err)
	}

	response := processInstallmentResponse(installments)

	return response, nil
}

func (i *Installment) FindInstallmentByNotPaid() (dto.InstallmentResponse, error) {
	installments, err := i.repositoryInstallment.All().Installment.FindByNotPaid()
	if err != nil {
		return dto.InstallmentResponse{}, fmt.Errorf("error finding installment by not paid: %v", err)
	}

	response := processInstallmentResponse(installments)

	return response, nil
}

func calculeteDateNextInvoice(i *Installment, first bool, date string, id uuid.UUID) (string, error) {
	var (
		cc  entity.CreditCard
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
			return "", err
		}

		if completeDate.Day() >= cc.InvoiceClosingDay {
			completeDate = completeDate.AddDate(0, 1, 0)

			return completeDate.Format("2006-01-02"), nil
		} else {
			return completeDate.Format("2006-01-02"), nil
		}
	}

	newDate, err2 := time.Parse("2006-01-02", date)
	if err1 != nil && err2 != nil {
		return "", msg
	}

	newDate = newDate.AddDate(0, 1, 0)

	return newDate.Format("2006-01-02"), nil
}

func processInstallmentResponse(installments []entity.Installment) dto.InstallmentResponse {
	paid, toPay := calculateTotal(installments)

	response := dto.InstallmentResponse{
		Response: installments,
		Paid:     paid,
		ToPay:    toPay,
		Total:    paid + toPay}

	return response
}

func calculateTotal(installments []entity.Installment) (float64, float64) {
	var toPay, paid float64

	for _, installment := range installments {
		if installment.Paid {
			paid += installment.Value
		} else {
			toPay += installment.Value
		}
	}

	return paid, toPay
}
