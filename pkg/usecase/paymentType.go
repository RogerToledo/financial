package usecase

import (
	"github.com/google/uuid"
	"github.com/me/financial/pkg/entity"
	"github.com/me/financial/pkg/repository"
)

type PaymentTypeUseCase interface {
	CreatePaymentType(paymentType entity.PaymentType) error
	UpdatePaymentType(paymentType entity.PaymentType) error
	DeletePaymentType(id uuid.UUID) error
	FindPaymentTypeByID(id uuid.UUID) (entity.PaymentType, error)
	FindAllPaymentTypes() ([]entity.PaymentType, error)
}

type PaymentType struct {
	repositoryPaymentType repository.RepositoryPaymentType
}

func NewPaymentTypeUseCase(r repository.RepositoryPaymentType) PaymentTypeUseCase {
	return &PaymentType{
		repositoryPaymentType: r,
	}
}

func (p *PaymentType) CreatePaymentType(paymentType entity.PaymentType) error {
	if err := p.repositoryPaymentType.Create(paymentType); err != nil {
		return err
	}

	return nil
}

func (p *PaymentType) UpdatePaymentType(paymentType entity.PaymentType) error {
	if err := p.repositoryPaymentType.Update(paymentType); err != nil {
		return err
	}

	return nil
}

func (p *PaymentType) DeletePaymentType(id uuid.UUID) error {
	if err := p.repositoryPaymentType.Delete(id); err != nil {
		return err
	}

	return nil
}

func (p *PaymentType) FindPaymentTypeByID(id uuid.UUID) (entity.PaymentType, error) {
	paymentType, err := p.repositoryPaymentType.FindByID(id)
	if err != nil {
		return entity.PaymentType{}, err
	}

	return paymentType, nil
}

func (p *PaymentType) FindAllPaymentTypes() ([]entity.PaymentType, error) {
	paymentTypes, err := p.repositoryPaymentType.FindAll()
	if err != nil {
		return nil, err
	}
	
	return paymentTypes, nil
}