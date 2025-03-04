package usecase

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/me/finance/pkg/dto"
	"github.com/me/finance/pkg/entity"
	"github.com/me/finance/pkg/repository"
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
	tx, err := p.repositoryPurchase.All().Purchase.BeginTransaction()
	if err != nil {
		return fmt.Errorf("error on begin transaction: %v", err)
	}

	var	savedID uuid.UUID

	if savedID, err = p.repositoryPurchase.All().Purchase.Create(tx, purchase); err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("savedID: %s", savedID.String()))

	purchase.Installment.PurchaseID = savedID

	ir := NewInstallmentUseCase(p.repositoryPurchase)

	if err := ir.CreateInstallment(purchase); err != nil {
		p.repositoryPurchase.All().Purchase.Rollback(tx)

		return err
	}

	p.repositoryPurchase.All().Purchase.Commit(tx)

	return nil
}

func (p *Purchase) UpdatePurchase(purchase entity.Purchase) error {
	tx, err := p.repositoryPurchase.All().Purchase.BeginTransaction()
	if err != nil {
		return fmt.Errorf("error on begin transaction: %v", err)
	}

	if err := p.repositoryPurchase.All().Purchase.Update(tx, purchase); err != nil {
		return err
	}

	purchase.Installment.PurchaseID = purchase.ID

	ir := NewInstallmentUseCase(p.repositoryPurchase)
	if err := ir.DeleteInstallment(purchase.ID); err != nil {
		p.repositoryPurchase.All().Purchase.Rollback(tx)

		return err
	}	

	if err := ir.CreateInstallment(purchase); err != nil {
		p.repositoryPurchase.All().Purchase.Rollback(tx)

		return err
	}

	p.repositoryPurchase.All().Purchase.Commit(tx)

	return nil
}

func (p *Purchase) DeletePurchase(id uuid.UUID) error {
	tx, err := p.repositoryPurchase.All().Purchase.BeginTransaction()
	if err != nil {
		return fmt.Errorf("error on begin transaction: %v", err)
	}

 	if err := p.repositoryPurchase.All().Purchase.Delete(tx, id); err != nil {
		return err
	}

	ir := NewInstallmentUseCase(p.repositoryPurchase)
	if err := ir.DeleteInstallment(id) ; err != nil {
		p.repositoryPurchase.All().Purchase.Rollback(tx)
		return err
	}

	p.repositoryPurchase.All().Purchase.Commit(tx)

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

	response := processPurchaseResponse(purchases)

	return response, err
}

func (p *Purchase) FindPurchaseByMonth(date string) (dto.PurchaseResponseTotal, error) {
	purchases, err := p.repositoryPurchase.All().Purchase.FindByMonth(date)
	if err != nil {
		return dto.PurchaseResponseTotal{}, err
	}

	response := processPurchaseResponse(purchases)

	return response, err
}

func (p *Purchase) FindPurchaseByPerson(personID uuid.UUID) (dto.PurchaseResponseTotal, error) {
	purchases, err := p.repositoryPurchase.All().Purchase.FindByPerson(personID)
	if err != nil {
		return dto.PurchaseResponseTotal{}, err
	}

	response := processPurchaseResponse(purchases)

	return response, err
}

func (p *Purchase) FindAllPurchases() ([]dto.PurchaseResponse, error) {
	purchases, err := p.repositoryPurchase.All().Purchase.FindAll()
	if err != nil {
		return nil, err
	}
	
	return purchases, err
}

func processPurchaseResponse(purchases []dto.PurchaseResponse) dto.PurchaseResponseTotal {
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
