package usecase

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/me/financial/pkg/dto"
	"github.com/me/financial/pkg/entity"
	"github.com/me/financial/pkg/repository"
)

type PurchaseUseCase interface {
	CreatePurchase(purchase entity.Purchase) error
	UpdatePurchase(purchase entity.Purchase) error
	DeletePurchase(id uuid.UUID) error
	FindPurchaseByID(id uuid.UUID) (dto.PurchaseResponse, error)
	FindPurchaseByDate(date string) (dto.PurchaseResponseTotal, error)
	FindPurchaseByMonth(date string) (dto.PurchaseResponseTotal, error)
	FindPurchaseByPerson(id uuid.UUID) (dto.PurchaseResponseTotal, error)
	FindAllPurchases() ([]dto.PurchaseResponse, error)
}

type Purchase struct {
	repositoryPurchase repository.RepositoryAll
}

func NewPurchaseUseCase(r repository.RepositoryAll) PurchaseUseCase {
	return &Purchase{
		repositoryPurchase: r,
	}
}

func (p *Purchase) CreatePurchase(purchase entity.Purchase) error {
	var (
		savedID uuid.UUID
		err     error
	)

	if savedID, err = p.repositoryPurchase.All().Purchase.Create(purchase); err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("savedID: %s", savedID.String()))

	purchase.Installment.PurchaseID = savedID

	var ir InstallmentUseCase

	ir = NewInstallmentUseCase(p.repositoryPurchase)

	if err := ir.CreateInstallment(purchase); err != nil {
		return err
	}

	return nil
}

func (p *Purchase) UpdatePurchase(purchase entity.Purchase) error {
	if err := p.repositoryPurchase.All().Purchase.Update(purchase); err != nil {
		return err
	}

	return nil
}

func (p *Purchase) DeletePurchase(id uuid.UUID) error {
	if err := p.repositoryPurchase.All().Purchase.Delete(id); err != nil {
		return err
	}

	return nil
}

func (p *Purchase) FindPurchaseByID(id uuid.UUID) (dto.PurchaseResponse, error) {
	purchase, err := p.repositoryPurchase.All().Purchase.FindByID(id)
	if err != nil {
		return dto.PurchaseResponse{}, err
	}

	return purchase, err
}

func (p *Purchase) FindPurchaseByDate(date string) (dto.PurchaseResponseTotal, error) {
	purchases, err := p.repositoryPurchase.All().Purchase.FindByDate(date)
	if err != nil {
		return dto.PurchaseResponseTotal{}, err
	}

	response := processResponse(purchases)

	return response, err
}

func (p *Purchase) FindPurchaseByMonth(date string) (dto.PurchaseResponseTotal, error) {
	purchases, err := p.repositoryPurchase.All().Purchase.FindByMonth(date)
	if err != nil {
		return dto.PurchaseResponseTotal{}, err
	}

	response := processResponse(purchases)

	return response, err
}

func (p *Purchase) FindPurchaseByPerson(personID uuid.UUID) (dto.PurchaseResponseTotal, error) {
	purchases, err := p.repositoryPurchase.All().Purchase.FindByPerson(personID)
	if err != nil {
		return dto.PurchaseResponseTotal{}, err
	}

	response := processResponse(purchases)

	return response, err
}

func (p *Purchase) FindAllPurchases() ([]dto.PurchaseResponse, error) {
	purchases, err := p.repositoryPurchase.All().Purchase.FindAll()
	if err != nil {
		return nil, err
	}
	
	return purchases, err
}

func processResponse(purchases []dto.PurchaseResponse) dto.PurchaseResponseTotal {
	total := 0.0

	for _, purchase := range purchases {
		total += purchase.Amount
	}

	response := dto.PurchaseResponseTotal{
		Responses: purchases,
		Quantity:  len(purchases),
		Total:     total,
	}

	return response
}
