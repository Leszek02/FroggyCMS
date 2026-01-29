package handler

import (
	"fmt"
	"net/http"
	"prestaprest/internal/application/schema"
	"prestaprest/internal/application/service"
	"prestaprest/internal/domain/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DictionaryHandler struct {
	service *service.DictionaryService
}

func NewDictionaryHandler(e *echo.Echo, service *service.DictionaryService) *DictionaryHandler {
	Handler := &DictionaryHandler{
		service: service,
	}
	return Handler
}

/////////////////// DELIVERY DICTIONARY ////////////////////

func (dh *DictionaryHandler) GetAllDeliveryDictionaries(c echo.Context) error {
	fmt.Println("Dictionary: GetAllDeliveryDictionaries")
	deliveryDictionary, err := dh.service.GetAllDeliveryDictionary()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, deliveryDictionary)
}

func (dh *DictionaryHandler) GetDeliveryDictionary(c echo.Context) error {
	fmt.Println("Dictionary: GetDeliveryDictionary")
	deliveryDictionaryIDuint16, err := strconv.ParseUint(c.Param("delivery_dictionary_ID"), 10, 16)
	deliveryDictionaryID := uint(deliveryDictionaryIDuint16)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	deliveryDictionary, err := dh.service.GetDeliveryDictionary(deliveryDictionaryID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, deliveryDictionary)
}

func (dh *DictionaryHandler) CreateDeliveryDictionary(c echo.Context) error {
	fmt.Println("Dictionary: CreateDeliveryDictionary")
	deliveryDictionary := new(entity.DeliveryDictionary)
	if err := c.Bind(deliveryDictionary); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	err := dh.service.CreateDeliveryDictionary(deliveryDictionary)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record created")
}

func (dh *DictionaryHandler) UpdateDeliveryDictionary(c echo.Context) error {
	fmt.Println("Dictionary: UpdateDeliveryDictionary")
	deliveryDictionaryIDuint16, err := strconv.ParseUint(c.Param("delivery_dictionary_ID"), 10, 16)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	deliveryDictionaryID := uint16(deliveryDictionaryIDuint16)
	deliveryDictionary := new(entity.DeliveryDictionary)
	if err := c.Bind(deliveryDictionary); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	deliveryDictionary.SetID(deliveryDictionaryID)
	err = dh.service.UpdateDeliveryDictionary(deliveryDictionary)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record updated")
}

func (dh *DictionaryHandler) DeleteDeliveryDictionary(c echo.Context) error {
	fmt.Println("Dictionary: DeleteDeliveryDictionary")
	deliveryDictionaryIDuint16, err := strconv.ParseUint(c.Param("delivery_dictionary_ID"), 10, 16)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	deliveryDictionaryID := uint(deliveryDictionaryIDuint16)
	err = dh.service.DeleteDeliveryDictionary(deliveryDictionaryID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (dh *DictionaryHandler) GetDeliveryDictionarySchema(c echo.Context) error {
	schema := schema.GenerateSchema(entity.DeliveryDictionary{})
	return c.JSON(http.StatusOK, schema)
}

/////////////////// PAYMENT DICTIONARY ////////////////////

func (dh *DictionaryHandler) GetAllPaymentDictionaries(c echo.Context) error {
	fmt.Println("Dictionary: GetAllPaymentDictionaries")
	paymentDictionary, err := dh.service.GetAllPaymentDictionary()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, paymentDictionary)
}

func (dh *DictionaryHandler) GetPaymentDictionary(c echo.Context) error {
	fmt.Println("Dictionary: GetPaymentDictionary")
	paymentDictionaryIDuint16, err := strconv.ParseUint(c.Param("payment_dictionary_ID"), 10, 16)
	paymentDictionaryID := uint(paymentDictionaryIDuint16)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	paymentDictionary, err := dh.service.GetPaymentDictionary(paymentDictionaryID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, paymentDictionary)
}

func (dh *DictionaryHandler) CreatePaymentDictionary(c echo.Context) error {
	fmt.Println("Dictionary: CreatePaymentDictionary")
	paymentDictionary := new(entity.PaymentDictionary)
	if err := c.Bind(paymentDictionary); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	err := dh.service.CreatePaymentDictionary(paymentDictionary)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record created")
}

func (dh *DictionaryHandler) UpdatePaymentDictionary(c echo.Context) error {
	fmt.Println("Dictionary: UpdatePaymentDictionary")
	paymentDictionaryIDuint16, err := strconv.ParseUint(c.Param("payment_dictionary_ID"), 10, 16)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	paymentDictionaryID := uint16(paymentDictionaryIDuint16)
	paymentDictionary := new(entity.PaymentDictionary)
	if err := c.Bind(paymentDictionary); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	paymentDictionary.SetID(paymentDictionaryID)
	err = dh.service.UpdatePaymentDictionary(paymentDictionary)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record updated")
}

func (dh *DictionaryHandler) DeletePaymentDictionary(c echo.Context) error {
	fmt.Println("Dictionary: DeletePaymentDictionary")
	paymentDictionaryIDuint16, err := strconv.ParseUint(c.Param("payment_dictionary_ID"), 10, 16)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	paymentDictionaryID := uint(paymentDictionaryIDuint16)
	err = dh.service.DeletePaymentDictionary(paymentDictionaryID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (dh *DictionaryHandler) GetPaymentDictionarySchema(c echo.Context) error {
	schema := schema.GenerateSchema(entity.PaymentDictionary{})
	return c.JSON(http.StatusOK, schema)
}

///////////////////// VAT PERCENTAGE DICTIONARY ////////////////////

func (dh *DictionaryHandler) GetAllVATPercentageDictionary(c echo.Context) error {
	fmt.Println("Dictionary: GetAllVATPercentageDicts")
	vatPercentageDicts, err := dh.service.GetAllVATPercentageDictionary()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, vatPercentageDicts)
}

func (dh *DictionaryHandler) GetVATPercentageDictionary(c echo.Context) error {
	fmt.Println("Dictionary: GetVATPercentageDict")
	vatPercentageDictIDuint8, err := strconv.ParseUint(c.Param("vat_percentage_dictionary_ID"), 10, 8)
	vatPercentageDictID := uint(vatPercentageDictIDuint8)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	vatPercentageDict, err := dh.service.GetVATPercentageDictionary(vatPercentageDictID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, vatPercentageDict)
}

func (dh *DictionaryHandler) CreateVATPercentageDictionary(c echo.Context) error {
	fmt.Println("Dictionary: CreateVATPercentageDict")
	vatPercentageDict := new(entity.VATPercentageDict)
	if err := c.Bind(vatPercentageDict); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	err := dh.service.CreateVATPercentageDictionary(vatPercentageDict)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record created")
}

func (dh *DictionaryHandler) UpdateVATPercentageDictionary(c echo.Context) error {
	fmt.Println("Dictionary: UpdateVATPercentageDict")
	vatPercentageDictIDuint8, err := strconv.ParseUint(c.Param("vat_percentage_dictionary_ID"), 10, 8)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	vatPercentageDictID := uint8(vatPercentageDictIDuint8)
	vatPercentageDict := new(entity.VATPercentageDict)
	if err := c.Bind(vatPercentageDict); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	vatPercentageDict.SetID(vatPercentageDictID)
	err = dh.service.UpdateVATPercentageDictionary(vatPercentageDict)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record updated")
}

func (dh *DictionaryHandler) DeleteVATPercentageDictionary(c echo.Context) error {
	fmt.Println("Dictionary: DeleteVATPercentageDict")
	vatPercentageDictIDuint8, err := strconv.ParseUint(c.Param("vat_percentage_dictionary_ID"), 10, 8)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	vatPercentageDictID := uint(vatPercentageDictIDuint8)
	err = dh.service.DeleteVATPercentageDictionary(vatPercentageDictID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (dh *DictionaryHandler) GetVATPercentageDictionarySchema(c echo.Context) error {
	schema := schema.GenerateSchema(entity.VATPercentageDict{})
	return c.JSON(http.StatusOK, schema)
}

/////////////////// PRODUCT CATEGORY DICTIONARY ////////////////////

func (dh *DictionaryHandler) GetAllProductCategoryDictionaries(c echo.Context) error {
	fmt.Println("Dictionary: GetAllProductCategoryDictionaries")
	productCategoryDictionaries, err := dh.service.GetAllProductCategoryDictionary()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, productCategoryDictionaries)
}

func (dh *DictionaryHandler) GetProductCategoryDictionary(c echo.Context) error {
	fmt.Println("Dictionary: GetProductCategoryDictionary")
	productCategoryDictionaryIDuint8, err := strconv.ParseUint(c.Param("product_category_dictionary_ID"), 10, 8)
	productCategoryDictionaryID := uint(productCategoryDictionaryIDuint8)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	productCategoryDictionary, err := dh.service.GetProductCategoryDictionary(productCategoryDictionaryID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, productCategoryDictionary)
}

func (dh *DictionaryHandler) CreateProductCategoryDictionary(c echo.Context) error {
	fmt.Println("Dictionary: CreateProductCategoryDictionary")
	productCategoryDictionary := new(entity.ProductCategoryDictionary)
	if err := c.Bind(productCategoryDictionary); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	err := dh.service.CreateProductCategoryDictionary(productCategoryDictionary)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record created")
}

func (dh *DictionaryHandler) UpdateProductCategoryDictionary(c echo.Context) error {
	fmt.Println("Dictionary: UpdateProductCategoryDictionary")
	productCategoryDictionaryIDuint8, err := strconv.ParseUint(c.Param("product_category_dictionary_ID"), 10, 8)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	productCategoryDictionaryID := uint8(productCategoryDictionaryIDuint8)
	productCategoryDictionary := new(entity.ProductCategoryDictionary)
	if err := c.Bind(productCategoryDictionary); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	productCategoryDictionary.SetID(productCategoryDictionaryID)
	err = dh.service.UpdateProductCategoryDictionary(productCategoryDictionary)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record updated")
}

func (dh *DictionaryHandler) DeleteProductCategoryDictionary(c echo.Context) error {
	fmt.Println("Dictionary: DeleteProductCategoryDictionary")
	productCategoryDictionaryIDuint8, err := strconv.ParseUint(c.Param("product_category_dictionary_ID"), 10, 8)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	productCategoryDictionaryID := uint(productCategoryDictionaryIDuint8)
	err = dh.service.DeleteProductCategoryDictionary(productCategoryDictionaryID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (dh *DictionaryHandler) GetProductCategorySchema(c echo.Context) error {
	schema := schema.GenerateSchema(entity.ProductCategoryDictionary{})
	return c.JSON(http.StatusOK, schema)
}
