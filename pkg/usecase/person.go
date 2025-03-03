package usecase

import (
	"github.com/google/uuid"
	"github.com/me/finance/pkg/entity"
	"github.com/me/finance/pkg/repository"
)

type PersonUseCase interface {
	CreatePerson(person entity.Person) error
	UpdatePerson(person entity.Person) error
	DeletePerson(id uuid.UUID) error
	FindPersonByID(id uuid.UUID) (entity.Person, error)
	FindAllPersons() ([]entity.Person, error)
}

type Person struct {
	repositoryPerson repository.RepositoryPerson
}

func NewPersonUseCase(r repository.RepositoryPerson) PersonUseCase {
	return &Person{
		repositoryPerson: r,
	}
}

func (p *Person) CreatePerson(person entity.Person) error {
	if err := p.repositoryPerson.Create(person); err != nil {
		return err
	}

	return nil
}

func (p *Person) UpdatePerson(person entity.Person) error {
	if err := p.repositoryPerson.Update(person); err != nil {
		return err
	}

	return nil
}

func (p *Person) DeletePerson(id uuid.UUID) error {
	if err := p.repositoryPerson.Delete(id); err != nil {
		return err
	}

	return nil
}

func (p *Person) FindPersonByID(id uuid.UUID) (entity.Person, error) {
	person, err := p.repositoryPerson.FindByID(id)
	if err != nil {
		return entity.Person{}, err
	}

	return person, nil
}

func (p *Person) FindAllPersons() ([]entity.Person, error) {
	persons, err := p.repositoryPerson.FindAll()
	if err != nil {
		return nil, err
	}
	return persons, nil
}
