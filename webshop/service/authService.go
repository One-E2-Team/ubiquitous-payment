package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/dto"
	"ubiquitous-payment/webshop/model"
)

func (service *Service) GetPrivileges(id uint) *[]string {
	privileges, err := service.WSRepository.GetPrivilegesByUserID(id)
	if err != nil {
		return nil
	}
	return privileges
}

func (service *Service) LogIn(dto dto.LogInDTO) (*model.User, error) {
	user, err := service.WSRepository.GetUserByEmail(dto.Email)

	if err != nil {
		return nil, fmt.Errorf("'" + dto.Email + "' " + err.Error())
	}
	if !user.IsValidated {
		return nil, fmt.Errorf(util.GetLoggingStringFromID(user.ProfileId) + " NOT VALIDATED")
	}
	if user.IsDeleted {
		return nil, fmt.Errorf(util.GetLoggingStringFromID(user.ProfileId) + " DELETED")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return nil, fmt.Errorf(util.GetLoggingStringFromID(user.ProfileId) + " " + err.Error())
	}
	return user, nil
}

func (service *Service) SetPSPAccessToken(accessToken string) error {
	webShop, err := service.WSRepository.GetFirstWebShop()
	if err != nil {
		return err
	}
	webShop.PSPAccessToken = accessToken
	return service.WSRepository.UpdateWebShop(webShop)
}
