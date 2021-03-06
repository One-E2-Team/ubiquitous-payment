package service

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/dto"
	"ubiquitous-payment/webshop/model"
	"ubiquitous-payment/webshop/wsutil/pspAuth"
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

func (service *Service) SetPSPAccessToken(accessUuid string) error {
	webShop, err := service.WSRepository.GetOurWebShop()
	if err != nil {
		return err
	}

	req := dto.PSPLoginDTO{Name: webShop.Name, AccessUuid: accessUuid}
	jsonReq, _ := json.Marshal(req)

	resp, err := pspAuth.PSPRequest(http.MethodPost, "/api/psp/web-shop-login",
		jsonReq, map[string]string{util.ContentType: util.ApplicationJson})
	if err != nil {
		return err
	}

	var accessToken string
	err = util.UnmarshalResponse(resp, &accessToken)
	if err != nil {
		return err
	}
	pspAuth.SetPspAccessToken(accessToken)
	webShop.PSPAccessToken = accessToken
	return service.WSRepository.UpdateWebShop(webShop)
}

func (service *Service) GetPSPAccessToken() (string, error) {
	webShop, err := service.WSRepository.GetOurWebShop()
	if err != nil {
		return "", err
	}
	return webShop.PSPAccessToken, nil
}

func (service *Service) ConfirmPassword(id uint, password string) (bool, error) {
	user, err := service.WSRepository.GetUserById(id)
	if err != nil{
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil{
		return false, err
	}
	return true, nil
}