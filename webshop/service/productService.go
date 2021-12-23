package service

import (
	"errors"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"ubiquitous-payment/webshop/model"
)

func (service *Service) CreateProduct(product model.Product, r *http.Request) error {
	picture, picHeader, err := r.FormFile("file")
	if err != nil {
		return err
	}
	fileName, err := saveFile(picture, picHeader)
	if err != nil {
		return err
	}
	product.MediaPath = fileName
	return service.WSRepository.CreateProduct(&product)
}

func (service *Service) UpdateProduct(productID uint, updatedProduct model.Product) error {
	product, err := service.WSRepository.GetProduct(productID)
	if err != nil {
		return err
	}
	if !product.IsActive {
		return errors.New("cannot update deactivated product")
	}
	if product.Price != updatedProduct.Price {
		product.Deactivate()
		err = service.WSRepository.UpdateProduct(product)
		if err != nil {
			return err
		}
		return service.WSRepository.CreateProduct(&updatedProduct)
	}
	product.Update(updatedProduct)
	return service.WSRepository.UpdateProduct(product)
}

func (service *Service) GetActiveProducts() ([]model.Product, error) {
	return service.WSRepository.GetActiveProducts()
}

func saveFile(picture multipart.File, picHeader *multipart.FileHeader) (string, error) {
	uuid := uuid.NewString()
	fileSplit := strings.Split(picHeader.Filename, ".")
	fileName := uuid + "." + fileSplit[1]
	f, err := os.OpenFile("../staticdata/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	defer func(picture multipart.File) {
		_ = picture.Close()
	}(picture)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(f, picture)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
