package service

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"SampleIntelliq/app/model"
	"SampleIntelliq/app/repo"
)

//AddAddress add
func AddAddress(address *model.Address) (string, error) {
	if address == nil {
		return "Invalid Address", errors.New("Null address")
	}
	addrRepo := repo.NewAddressRepository()
	err := addrRepo.Save(address)
	if err != nil {
		return "Could not save", err
	}
	return "Saved Successfully !", nil
}

//UpdateAddress update
func UpdateAddress(address *model.Address) (string, error) {
	if address == nil {
		return "Invalid Address", errors.New("Null address")
	}
	addrRepo := repo.NewAddressRepository()
	err := addrRepo.Update(address)
	if err != nil {
		return "Could not update", err
	}
	return "Updated Successfully !", nil
}

//RemoveAddress remove
func RemoveAddress(addressID primitive.ObjectID) (string, error) {
	addrRepo := repo.NewAddressRepository()
	err := addrRepo.Delete(addressID)
	if err != nil {
		return "Could not delete", err
	}
	return "Removed Successfully !", nil
}

//FindAddressByCity find
func FindAddressByCity(city string) (*model.Address, error) {
	addrRepo := repo.NewAddressRepository()
	address, err := addrRepo.FindByCity(city)
	if err != nil {
		return nil, err
	}
	return address, nil
}

//FindAddressByID find
func FindAddressByID(addressID primitive.ObjectID) (*model.Address, error) {
	addrRepo := repo.NewAddressRepository()
	address, err := addrRepo.FindByID(addressID)
	if err != nil {
		return nil, err
	}
	return address, nil
}

//FindAllAddresses findAll
func FindAllAddresses() (model.Addresses, error) {
	addrRepo := repo.NewAddressRepository()
	addresses, err := addrRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

//FindAddressByState find
func FindAddressByState(state string) (model.Addresses, error) {
	addrRepo := repo.NewAddressRepository()
	addresses, err := addrRepo.FindByState(state)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}
