package usecase

import (
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
	repositoryPurchase repository.RepositoryPurchase
}

func NewPurchaseUseCase(r repository.RepositoryPurchase) PurchaseUseCase {
	return &Purchase{
		repositoryPurchase: r,
	}
}

func (p *Purchase) CreatePurchase(purchase entity.Purchase) error {
	purchase.Installment = purchase.Amount / float64(purchase.InstallmentNumber)

	if err := p.repositoryPurchase.Create(purchase); err != nil {
		return err
	}

	return nil
}

func (p *Purchase) UpdatePurchase(purchase entity.Purchase) error {
	if err := p.repositoryPurchase.Update(purchase); err != nil {
		return err
	}

	return nil
}

func (p *Purchase) DeletePurchase(id uuid.UUID) error {
	if err := p.repositoryPurchase.Delete(id); err != nil {
		return err
	}

	return nil
}

func (p *Purchase) FindPurchaseByID(id uuid.UUID) (dto.PurchaseResponse, error) {
	purchase, err := p.repositoryPurchase.FindByID(id)
	if err != nil {
		return dto.PurchaseResponse{}, err
	}

	return purchase, err
}

func (p *Purchase) FindPurchaseByDate(date string) (dto.PurchaseResponseTotal, error) {
	purchases, err := p.repositoryPurchase.FindByDate(date)
	if err != nil {
		return dto.PurchaseResponseTotal{}, err
	}

	response := processResponse(purchases)

	return response, err
}

func (p *Purchase) FindPurchaseByMonth(date string) (dto.PurchaseResponseTotal, error) {
	purchases, err := p.repositoryPurchase.FindByMonth(date)
	if err != nil {
		return dto.PurchaseResponseTotal{}, err
	}

	response := processResponse(purchases)

	return response, err
}

func (p *Purchase) FindPurchaseByPerson(personID uuid.UUID) (dto.PurchaseResponseTotal, error) {
	purchases, err := p.repositoryPurchase.FindByPerson(personID)
	if err != nil {
		return dto.PurchaseResponseTotal{}, err
	}

	response := processResponse(purchases)

	return response, err
}

func (p *Purchase) FindAllPurchases() ([]dto.PurchaseResponse, error) {
	purchases, err := p.repositoryPurchase.FindAll()
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
