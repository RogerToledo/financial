package usecase

import (
	"github.com/google/uuid"
	"github.com/me/financial/pkg/entity"
	"github.com/me/financial/pkg/repository"
)

type PurchaseTypeUseCase interface {
	CreatePurchaseType(pt entity.PurchaseType) error
	UpdatePurchaseType(pt entity.PurchaseType) error
	DeletePurchaseType(id uuid.UUID) error
	FindPurchaseTypeByID(id uuid.UUID) (entity.PurchaseType, error)
	FindAllPurchaseTypes() ([]entity.PurchaseType, error)	
}

type PurchaseType struct {
	repository repository.RepositoryPurchaseType
}

func NewPurchaseTypeUseCase(r repository.RepositoryPurchaseType) PurchaseTypeUseCase {
	return &PurchaseType{
		repository: r,
	}
}

func (p *PurchaseType) CreatePurchaseType(pt entity.PurchaseType) error {
	if err := p.repository.Create(pt); err != nil {
		return err
	}

	return nil
}

func (p *PurchaseType) UpdatePurchaseType(pt entity.PurchaseType) error {
	if err := p.repository.Update(pt); err != nil {
		return err
	}

	return nil
}

func (p *PurchaseType) DeletePurchaseType(id uuid.UUID) error {
	if err := p.repository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (p *PurchaseType) FindPurchaseTypeByID(id uuid.UUID) (entity.PurchaseType, error) {
	purchaseType, err := p.repository.FindByID(id)
	if err != nil {
		return entity.PurchaseType{}, err
	}
	
	return purchaseType, nil
}

func (p *PurchaseType) FindAllPurchaseTypes() ([]entity.PurchaseType, error) {
	purchaseTypes, err := p.repository.FindAll()
	if err != nil {
		return []entity.PurchaseType{}, err
	}

	return purchaseTypes, nil
}
