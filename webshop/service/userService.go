package service

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	"ubiquitous-payment/webshop/dto"
	"ubiquitous-payment/webshop/model"
	"github.com/google/uuid"
)

func (service *Service) Register(w http.ResponseWriter, dto dto.RegistrationDto) error {
	var err error

	v := validator.New()
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
		return fmt.Errorf("validation error")
	}

	if dto.Role == "ADMIN" {
		return fmt.Errorf("admin must not be registered")
	}
	role, err := service.WSRepository.GetRoleByName(dto.Role)
	if err != nil {
		return err
	}
	profile := model.Profile{Name: dto.Name}
	err = service.WSRepository.CreateProfile(&profile)
	if err != nil {
		return err
	}
	user := model.User{ProfileId: profile.ID, Email: dto.Email, Username: dto.Username,
		Password: hashAndSalt(dto.Password),IsDeleted: false, IsValidated: false, ValidationUuid: uuid.NewString(),
		ValidationExpire: time.Now().Add(1 * time.Hour), Roles: []model.Role{*role}}

	err = service.WSRepository.CreateUser(&user)
	if err != nil {
		return err
	}
	return nil
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