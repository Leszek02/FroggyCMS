package handler

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"prestaprest/internal/application/schema"
	"prestaprest/internal/application/service"
	"prestaprest/internal/domain/entity"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service    *service.ProductService
	dictionary *service.DictionaryService
	media      *service.MediaService
}

func NewProductHandler(e *echo.Echo,
	service *service.ProductService,
	dictionary *service.DictionaryService,
	media *service.MediaService) *ProductHandler {
	Handler := &ProductHandler{
		service:    service,
		dictionary: dictionary,
		media:      media,
	}
	return Handler
}

func (ph *ProductHandler) GetAllProduct(c echo.Context) error {
	fmt.Println("Product: GetAllProduct")
	product, err := ph.service.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) GetProduct(c echo.Context) error {
	fmt.Println("Product: GetProduct")
	productID, err := strconv.ParseUint(c.Param("product_ID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	fmt.Println("ProductID: ", productID)
	product, err := ph.service.GetProduct(productID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) GetProductPhotos(c echo.Context) error {
	fmt.Println("Product: GetProductMedias")
	medias, err := ph.media.GetAllMedia()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, medias)
}

func (ph *ProductHandler) CreateProduct(c echo.Context) error {
	fmt.Println("Product: CreateProduct")
	fmt.Println("Form values:")
	fmt.Println("  name:", c.FormValue("name"))
	fmt.Println("  price:", c.FormValue("price"))
	fmt.Println("  product_code:", c.FormValue("product_code"))
	fmt.Println("  description:", c.FormValue("description"))
	fmt.Println("  is_animal:", c.FormValue("is_animal"))
	fmt.Println("  availability:", c.FormValue("availability"))
	fmt.Println("  vat_percentage:", c.FormValue("vat_percentage"))
	fmt.Println("  product_category:", c.FormValue("product_category"))

	product := new(entity.Product)
	product.Name = c.FormValue("name")
	product.ProductCode = c.FormValue("product_code")
	product.Description = c.FormValue("description")
	product.IsAnimal = c.FormValue("is_animal") == "on" || c.FormValue("is_animal") == "true"
	product.Availability = c.FormValue("availability") == "on" || c.FormValue("availability") == "true"

	if priceStr := c.FormValue("price"); priceStr != "" {
		price, _ := strconv.ParseFloat(priceStr, 32)
		product.Price = float32(price)
	}

	if vatPercentage := c.FormValue("vat_percentage"); vatPercentage != "" {
		vat, _ := strconv.ParseFloat(vatPercentage, 32)
		product.VatPercentage = float32(vat)
	}

	product.ProductCategory = c.FormValue("product_category")
	fmt.Println(product)

	err := ph.service.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to create product: %s", err),
		})
	}

	imagePaths := []struct {
		FieldName string
		Folder    string
	}{
		{"product_main_image", "../media/products/main"},
		{"product_icon_image", "../media/products/icons"},
	}

	for _, config := range imagePaths {
		file, err := c.FormFile(config.FieldName)

		existingMediaID := c.FormValue(config.FieldName + "_id")

		if err != nil && existingMediaID == "" {
			continue
		}

		var mediaID uint64

		if existingMediaID != "" {
			mediaID, _ = strconv.ParseUint(existingMediaID, 10, 64)
		} else {
			savedPath, err := ph.saveUploadedFile(file, config.Folder, file.Filename)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Failed to save " + config.FieldName + ": " + err.Error(),
				})
			}

			media := entity.NewMedia(file.Filename, savedPath)
			err = ph.media.CreateMedia(media)
			if err != nil {
				os.Remove(savedPath)
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Failed to save media record for " + config.FieldName,
				})
			}
			mediaID = media.Media_ID
		}

		productsMedia := entity.NewProductsMedias(product.ProductID, mediaID)
		err = ph.media.UpdateProductMedia(productsMedia)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to link photos",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Product created successfully",
		"product_id": product.ProductID,
	})
}

func (ph *ProductHandler) UpdateProduct(c echo.Context) error {
	fmt.Println("Product: UpdateProduct")
	productID, err := strconv.ParseUint(c.Param("product_ID"), 10, 64)
	fmt.Println("Product: CreateProduct")
	fmt.Println("Form values:")
	fmt.Println("  name:", c.FormValue("name"))
	fmt.Println("  price:", c.FormValue("price"))
	fmt.Println("  product_code:", c.FormValue("product_code"))
	fmt.Println("  description:", c.FormValue("description"))
	fmt.Println("  is_animal:", c.FormValue("is_animal"))
	fmt.Println("  availability:", c.FormValue("availability"))
	fmt.Println("  vat_percentage:", c.FormValue("vat_percentage"))
	fmt.Println("  product_category:", c.FormValue("product_category"))

	product := new(entity.Product)
	product.Name = c.FormValue("name")
	product.ProductCode = c.FormValue("product_code")
	product.Description = c.FormValue("description")
	product.IsAnimal = c.FormValue("is_animal") == "on" || c.FormValue("is_animal") == "true"
	product.Availability = c.FormValue("availability") == "on" || c.FormValue("availability") == "true"

	if priceStr := c.FormValue("price"); priceStr != "" {
		price, _ := strconv.ParseFloat(priceStr, 32)
		product.Price = float32(price)
	}

	if vatPercentage := c.FormValue("vat_percentage"); vatPercentage != "" {
		vat, _ := strconv.ParseFloat(vatPercentage, 32)
		product.VatPercentage = float32(vat)
	}

	product.ProductCategory = c.FormValue("product_category")
	fmt.Println(product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to create product: %s", err),
		})
	}

	imagePaths := []struct {
		FieldName string
		Folder    string
	}{
		{"product_main_image", "../media/products/main"},
		{"product_icon_image", "../media/products/icons"},
	}

	for _, config := range imagePaths {
		file, err := c.FormFile(config.FieldName)

		existingMediaID := c.FormValue(config.FieldName + "_id")

		if err != nil && existingMediaID == "" {
			continue
		}

		var mediaID uint64

		if existingMediaID != "" {
			mediaID, _ = strconv.ParseUint(existingMediaID, 10, 64)
		} else {
			savedPath, err := ph.saveUploadedFile(file, config.Folder, file.Filename)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Failed to save " + config.FieldName + ": " + err.Error(),
				})
			}

			media := entity.NewMedia(file.Filename, savedPath)
			err = ph.media.CreateMedia(media)
			if err != nil {
				os.Remove(savedPath)
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Failed to save media record for " + config.FieldName,
				})
			}
			mediaID = media.Media_ID
		}

		productsMedia := entity.NewProductsMedias(productID, mediaID)
		err = ph.media.UpdateProductMedia(productsMedia)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to link photos",
			})
		}
	}
	product.ProductID = productID
	err = ph.service.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record updated")
}

func (ph *ProductHandler) DeleteProduct(c echo.Context) error {
	fmt.Println("Product: DeleteProduct")
	productID, err := strconv.ParseUint(c.Param("product_ID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	err = ph.service.DeleteProduct(productID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (ph *ProductHandler) ValidateProduct(c echo.Context) error {
	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return errors.New("Product name is required")
	}

	priceStr := strings.TrimSpace(c.FormValue("price"))
	if priceStr == "" {
		return errors.New("Price is required")
	}
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		return errors.New("Invalid price format")
	}
	if price <= 0 {
		return errors.New("Price must be greater than 0")
	}

	vatStr := strings.TrimSpace(c.FormValue("vat_percentage"))
	if vatStr == "" {
		return errors.New("VAT percentage is required")
	}
	vat, err := strconv.ParseFloat(vatStr, 32)
	if err != nil {
		return errors.New("Invalid VAT percentage format")
	}
	if vat < 0 || vat > 100 {
		return errors.New("VAT percentage must be between 0 and 100")
	}

	productCode := strings.TrimSpace(c.FormValue("product_code"))
	if productCode == "" {
		return errors.New("Product code is required")
	}

	description := strings.TrimSpace(c.FormValue("description"))
	if description == "" {
		return errors.New("Description is required")
	}

	productCategory := strings.TrimSpace(c.FormValue("product_category"))
	if productCategory == "" {
		return errors.New("Product category is required")
	}

	return nil
}

func (ph *ProductHandler) saveUploadedFile(file *multipart.FileHeader, folderPath string, fileName string) (string, error) {
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	filePath := filepath.Join(folderPath, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}

	return filePath, nil
}

func (ph *ProductHandler) GetProductSchema(c echo.Context) error {
	schema := schema.GenerateSchema(entity.Product{})
	return c.JSON(http.StatusOK, schema)
}
