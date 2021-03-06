package service

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	"ubiquitous-payment/bank/bankutil"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/model"
	"ubiquitous-payment/util"
)

func (service *Service) Register(request dto.RegistrationDTO, w http.ResponseWriter) error {
	var err error

	v := validator.New()
	checkCommonPass(v)
	checkWeakPass(v, err)
	checkUsername(v)
	err = v.Struct(request)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{\"message\":\"Invalid data.\",\n"))
		_, _ = w.Write([]byte("\"errors\":\""))
		for _, e := range err.(validator.ValidationErrors) {
			_, _ = w.Write([]byte(e.Field()))
			_, _ = w.Write([]byte(" "))
		}
		_, _ = w.Write([]byte("\"}"))
		return errors.New("validation error")
	}

	role, err := service.Repository.GetRoleByName("CLIENT")
	if err != nil {
		return err
	}

	client := model.Client{
		Username:  request.Username,
		Password:  hashAndSalt(request.Password),
		IsDeleted: false,
		Roles:     []model.Role{*role},
		Accounts:  []model.ClientAccount{*createNewClientAccount(request.Name, request.Surname)},
		LastActivityTimestamp: time.Now(),
	}

	return service.Repository.Create(&client)
}

func (service *Service) LogIn(credentials dto.LoginDTO) (*model.Client, error) {
	client, err := service.Repository.GetClientByUsername(credentials.Username)

	if err != nil {
		return nil, fmt.Errorf("'" + credentials.Username + "' " + err.Error())
	}
	if client.IsDeleted {
		return nil, fmt.Errorf(util.Uint2String(client.ID) + " DELETED")
	}
	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(credentials.Password))
	if err != nil {
		return nil, fmt.Errorf(util.Uint2String(client.ID) + " " + err.Error())
	}
	return client, nil
}

func (service *Service) ConfirmPassword(clientId uint, password string) (bool, error) {
	client, err := service.Repository.GetClientById(clientId)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (service *Service) GetPrivileges(clientId uint) *[]string {
	privileges, err := service.Repository.GetPrivilegesByClientId(clientId)
	if err != nil {
		return nil
	}
	return privileges
}

func checkCommonPass(v *validator.Validate) {
	_ = v.RegisterValidation("common_pass", func(fl validator.FieldLevel) bool {
		f, err := os.Open("common_pass.txt")
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(f)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if strings.Contains(fl.Field().String(), scanner.Text()) {
				return false
			}
		}
		return true
	})
}

func checkWeakPass(v *validator.Validate, err error) {
	_ = v.RegisterValidation("weak_pass", func(fl validator.FieldLevel) bool {
		if len(fl.Field().String()) < 8 {
			return false
		}
		ret, _ := regexp.MatchString("(.*[a-z].*)", fl.Field().String())
		if ret == false {
			return false
		}
		ret, _ = regexp.MatchString("(.*[A-Z].*)", fl.Field().String())
		if ret == false {
			return false
		}
		ret, _ = regexp.MatchString("(.*[0-9].*)", fl.Field().String())
		if ret == false {
			return false
		}
		ret, _ = regexp.MatchString("(.*[*!@#$%^&(){}\\[:;\\]<>,.?~_+\\-\\\\=|/].*)", fl.Field().String())
		if err != nil {
			fmt.Println(err)
			return false
		}
		return ret
	})
}

func checkUsername(v *validator.Validate) {
	_ = v.RegisterValidation("bad_username", func(fl validator.FieldLevel) bool {
		if len(fl.Field().String()) < 3 || len(fl.Field().String()) > 15 {
			return false
		}
		ret, _ := regexp.MatchString("([*!@#$%^&(){}\\[:;\\]<>,.?~+\\-\\\\=|/ ])", fl.Field().String())
		if ret {
			return false
		}
		return true
	})
}

func hashAndSalt(pass string) string {
	bytePass := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(hash)
}

func createNewClientAccount(name string, surname string) *model.ClientAccount {

	accountNumber := bankutil.PanPrefix + util.RandomString("0123456789", 10)
	pan := bankutil.PanPrefix + util.RandomString("0123456789", 10)
	validUntil := time.Now().AddDate(5, 0, 0).Format(util.MMyyDateFormat)

	creditCard := model.CreditCard{
		Pan:        util.GetEncryptedString(pan),
		Cvc:        util.GetEncryptedString(util.RandomString("0123456789", 3)),
		HolderName: name + " " + surname,
		ValidUntil: validUntil,
	}

	return &model.ClientAccount{
		AccountNumber: util.GetEncryptedString(accountNumber),
		Amount:        0,
		Secret:        util.GetEncryptedString(util.RandomString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_", 10)),
		IsActive:      true,
		CreditCards:   []model.CreditCard{creditCard},
	}
}
