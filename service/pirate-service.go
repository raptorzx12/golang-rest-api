package service

import (
	"fmt"
	"log"
	"programming/golang/rest-api/dto"
	"programming/golang/rest-api/entity"
	"programming/golang/rest-api/repository"

	"github.com/mashingan/smapping"
)

type PirateService interface {
	Insert(p dto.PirateCreateDTO) entity.Pirate
	Update(p dto.PirateUpdateDTO) entity.Pirate
	Delete(p entity.Pirate)
	All() []entity.Pirate
	FindByID(pirateID uint64) entity.Pirate
	IsAllowedToEdit(userID string, pirateID uint64) bool
}

type pirateService struct {
	pirateRepository repository.PirateRepository
}

//NewPirateService ...
func NewPirateService(pirateRepo repository.PirateRepository) PirateService {
	return &pirateService{
		pirateRepository: pirateRepo,
	}
}

func (service *pirateService) Insert(p dto.PirateCreateDTO) entity.Pirate {
	pirate := entity.Pirate{}
	err := smapping.FillStruct(&pirate, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.pirateRepository.InsertPirate(pirate)
	return res
}

func (service *pirateService) Update(p dto.PirateUpdateDTO) entity.Pirate {
	pirate := entity.Pirate{}
	err := smapping.FillStruct(&pirate, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.pirateRepository.UpdatePirate(pirate)
	return res
}

func (service *pirateService) Delete(p entity.Pirate) {
	service.pirateRepository.DeletePirate(p)
}

func (service *pirateService) All() []entity.Pirate {
	return service.pirateRepository.AllPirate()
}

func (service *pirateService) FindByID(pirateID uint64) entity.Pirate {
	return service.pirateRepository.FindPirateByID(pirateID)
}

func (service *pirateService) IsAllowedToEdit(userID string, pirateID uint64) bool {
	p := service.pirateRepository.FindPirateByID(pirateID)
	id := fmt.Sprintf("%v", p.UserID)
	return userID == id
}
