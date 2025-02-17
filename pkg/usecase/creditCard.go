package usecase

import (
	"github.com/google/uuid"
	"github.com/me/financial/pkg/entity"
	"github.com/me/financial/pkg/repository"
)

type CreditCardUseCase interface {
	CreateCreditCard(cc entity.CreditCard) error
	UpdateCreditCard(cc entity.CreditCard) error
	DeleteCreditCard(id uuid.UUID) error
	FindCreditCardByID(id uuid.UUID) (entity.CreditCard, error)
	FindAllCreditCards() ([]entity.CreditCard, error)
}

type CreditCard struct {
	repositoryCreditCard repository.RepositoryCreditCard
}

func NewCreditCardUseCase(r repository.RepositoryCreditCard) CreditCardUseCase {
	return &CreditCard{
		repositoryCreditCard: r,
	}
}

func (c *CreditCard) CreateCreditCard(cc entity.CreditCard) error {
	if err := c.repositoryCreditCard.Create(cc); err != nil {
		return err
	}

	return nil
}

func (c *CreditCard) UpdateCreditCard(cc entity.CreditCard) error {
	if err := c.repositoryCreditCard.Update(cc); err != nil {
		return err
	}

	return nil
}

func (c *CreditCard) DeleteCreditCard(id uuid.UUID) error {
	if err := c.repositoryCreditCard.Delete(id); err != nil {
		return err
	}

	return nil
}

func (c *CreditCard) FindCreditCardByID(id uuid.UUID) (entity.CreditCard, error) {
	cc, err := c.repositoryCreditCard.FindByID(id)
	if err != nil {
		return entity.CreditCard{}, err
	}

	return cc, nil
}

func (c *CreditCard) FindAllCreditCards() ([]entity.CreditCard, error) {
	cc, err := c.repositoryCreditCard.FindAll()
	if err != nil {
		return []entity.CreditCard{}, err
	}

	return cc, nil
}