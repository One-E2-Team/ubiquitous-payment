package service

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

	type AccessTokenData struct {
		Name        string `json:"name"`
		AccessToken string `json:"accessToken"` //TODO: change name to uuid
	}

	req := AccessTokenData{
		Name:        webShop.Name,
		AccessToken: accessToken,
	}
	jsonReq, _ := json.Marshal(req)

	resp, err := util.PSPRequest(http.MethodPost, "/api/psp/web-shop-login",
		jsonReq, map[string]string{util.ContentType: util.ApplicationJson})
	if err != nil {
		return err
	}

	var realAccessToken string
	err = util.UnmarshalResponse(resp, &realAccessToken)
	if err != nil {
		return err
	}
	util.SetPspJwt(realAccessToken)
	webShop.PSPAccessToken = realAccessToken
	return service.WSRepository.UpdateWebShop(webShop)
}
