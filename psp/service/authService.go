package service

import (
	"bufio"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	"regexp"
	"strings"
	"ubiquitous-payment/psp/dto"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
	"ubiquitous-payment/util"
)

func (service *Service) Register(w http.ResponseWriter, dto dto.RegisterDTO) error {
	var err error
	v := validator.New() // TODO: move validator in some util method
	checkCommonPass(v)
	checkWeakPass(v, err)
	checkUsername(v)
	err = v.Struct(dto)

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

	supportedPaymentTypes, err := service.PSPRepository.GetAllPaymentTypes()
	if err != nil {
		return err
	}

	for _, selectedPaymentTypeName := range dto.PaymentTypes {
		if !service.paymentTypeListContains(supportedPaymentTypes, selectedPaymentTypeName) {
			return fmt.Errorf("'%s' payment type is not supported by PSP", selectedPaymentTypeName)
		}
	}

	var paymentTypes []model.PaymentType //TODO: improve to call repo once
	for _, selectedPaymentTypeName := range dto.PaymentTypes {
		paymentType, err := service.PSPRepository.GetPaymentTypeByName(selectedPaymentTypeName)
		if err != nil {
			return err
		}
		paymentTypes = append(paymentTypes, *paymentType)
	}

	webShop := model.WebShop{ID: primitive.NewObjectID(), Name: dto.WebShopName, PSPAccessToken: "", //TODO: generate ID in repo
		Accepted: false, PaymentTypes: paymentTypes, Accounts: nil}

	err = service.PSPRepository.CreateWebShop(&webShop)
	if err != nil {
		return err
	}

	webShopPrivilege := model.Privilege{Name: psputil.WebShopTokenPermissionName}
	webShopRole := model.Role{Name: psputil.WebShopRoleName, Privileges: []model.Privilege{webShopPrivilege}}
	webShopOwner := model.User{ID: primitive.NewObjectID(), Username: dto.Username, Password: hashAndSalt(dto.Password),
		IsDeleted: false, Roles: []model.Role{webShopRole}, WebShopId: util.MongoID2String(webShop.ID)}
	return service.PSPRepository.CreateUser(&webShopOwner)
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
