package service

import (
	"fmt"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/postgres"

	"gorm.io/gorm"
)

type DictionaryService struct {
	DeliveryPostgresRepository postgres.PostgresRepository[entity.DeliveryDictionary]
	PaymentPostgresRepository  postgres.PostgresRepository[entity.PaymentDictionary]
	VATPostgresRepository      postgres.PostgresRepository[entity.VATPercentageDict]
	ProductCategoryRepository  postgres.PostgresRepository[entity.ProductCategoryDictionary]
}

func NewDictionaryService(conn *gorm.DB) DictionaryService {
	return DictionaryService{
		DeliveryPostgresRepository: postgres.NewPostgresRepository[entity.DeliveryDictionary](conn),
		PaymentPostgresRepository:  postgres.NewPostgresRepository[entity.PaymentDictionary](conn),
		VATPostgresRepository:      postgres.NewPostgresRepository[entity.VATPercentageDict](conn),
		ProductCategoryRepository:  postgres.NewPostgresRepository[entity.ProductCategoryDictionary](conn),
	}
}

//////////////////// DELIVERY DICTIONARY ////////////////////

func (s *DictionaryService) GetAllDeliveryDictionary() ([]entity.DeliveryDictionary, error) {
	allDeliveryDictionaries, err := s.DeliveryPostgresRepository.GetAll("Delivery_dict_ID")
	if err != nil {
		return nil, err
	}
	return allDeliveryDictionaries, nil
}

func (s *DictionaryService) GetDeliveryDictionary(deliveryDictionaryID uint) (*entity.DeliveryDictionary, error) {
	deliveryDictionary, err := s.DeliveryPostgresRepository.Get("Delivery_dict_ID", deliveryDictionaryID)
	fmt.Println("Am here (service)?")
	if err != nil {
		return nil, err
	}
	return deliveryDictionary, nil
}

func (s *DictionaryService) CreateDeliveryDictionary(deliveryDictionary *entity.DeliveryDictionary) error {
	err := s.DeliveryPostgresRepository.Create(deliveryDictionary)
	if err != nil {
		return err
	}
	return nil
}

func (s *DictionaryService) UpdateDeliveryDictionary(deliveryDictionary *entity.DeliveryDictionary) error {
	err := s.DeliveryPostgresRepository.Update(deliveryDictionary)
	if err != nil {
		return err
	}
	return nil
}

func (s *DictionaryService) DeleteDeliveryDictionary(deliveryDictionaryID uint) error {
	err := s.DeliveryPostgresRepository.Delete("Delivery_dict_ID", uint64(deliveryDictionaryID))
	if err != nil {
		return err
	}
	return nil
}

//////////////////// PAYMENT DICTIONARY ////////////////////

func (s *DictionaryService) GetAllPaymentDictionary() ([]entity.PaymentDictionary, error) {
	allPaymentDictionaries, err := s.PaymentPostgresRepository.GetAll("Payment_dict_ID")
	if err != nil {
		return nil, err
	}
	return allPaymentDictionaries, nil
}

func (s *DictionaryService) GetPaymentDictionary(paymentDictionaryID uint) (*entity.PaymentDictionary, error) {
	paymentDictionary, err := s.PaymentPostgresRepository.Get("Payment_dict_ID", paymentDictionaryID)
	if err != nil {
		return nil, err
	}
	return paymentDictionary, nil
}

func (s *DictionaryService) CreatePaymentDictionary(paymentDictionary *entity.PaymentDictionary) error {
	err := s.PaymentPostgresRepository.Create(paymentDictionary)
	if err != nil {
		return err
	}
	return nil
}

func (s *DictionaryService) UpdatePaymentDictionary(paymentDictionary *entity.PaymentDictionary) error {
	err := s.PaymentPostgresRepository.Update(paymentDictionary)
	if err != nil {
		return err
	}
	return nil
}

func (s *DictionaryService) DeletePaymentDictionary(paymentDictionaryID uint) error {
	err := s.PaymentPostgresRepository.Delete("Payment_dict_ID", uint64(paymentDictionaryID))
	if err != nil {
		return err
	}
	return nil
}

//////////////////// VAT PERCENTAGE DICTIONARY ////////////////////

func (s *DictionaryService) GetAllVATPercentageDictionary() ([]entity.VATPercentageDict, error) {
	allVATPercentageDicts, err := s.VATPostgresRepository.GetAll("VAT_percentage_dict_ID")
	if err != nil {
		return nil, err
	}
	return allVATPercentageDicts, nil
}

func (s *DictionaryService) GetVATPercentageDictionary(vatPercentageDictID uint) (*entity.VATPercentageDict, error) {
	vatPercentageDict, err := s.VATPostgresRepository.Get("VAT_percentage_dict_ID", vatPercentageDictID)
	if err != nil {
		return nil, err
	}
	return vatPercentageDict, nil
}

func (s *DictionaryService) CreateVATPercentageDictionary(vatPercentageDict *entity.VATPercentageDict) error {
	err := s.VATPostgresRepository.Create(vatPercentageDict)
	if err != nil {
		return err
	}
	return nil
}

func (s *DictionaryService) UpdateVATPercentageDictionary(vatPercentageDict *entity.VATPercentageDict) error {
	err := s.VATPostgresRepository.Update(vatPercentageDict)
	if err != nil {
		return err
	}
	return nil
}

func (s *DictionaryService) DeleteVATPercentageDictionary(vatPercentageDictID uint) error {
	err := s.VATPostgresRepository.Delete("VAT_percentage_dict_ID", uint64(vatPercentageDictID))
	if err != nil {
		return err
	}
	return nil
}

//////////////////// PRODUCT CATEGORY DICTIONARY ////////////////////

func (s *DictionaryService) GetAllProductCategoryDictionary() ([]entity.ProductCategoryDictionary, error) {
	allProductCategoryDictionaries, err := s.ProductCategoryRepository.GetAll("Product_dict_ID")
	if err != nil {
		return nil, err
	}
	return allProductCategoryDictionaries, nil
}

func (s *DictionaryService) GetProductCategoryDictionary(productCategoryDictionaryID uint) (*entity.ProductCategoryDictionary, error) {
	productCategoryDictionary, err := s.ProductCategoryRepository.Get("Product_dict_ID", productCategoryDictionaryID)
	if err != nil {
		return nil, err
	}
	return productCategoryDictionary, nil
}

func (s *DictionaryService) CreateProductCategoryDictionary(productCategoryDictionary *entity.ProductCategoryDictionary) error {
	err := s.ProductCategoryRepository.Create(productCategoryDictionary)
	if err != nil {
		return err
	}
	return nil
}

func (s *DictionaryService) UpdateProductCategoryDictionary(productCategoryDictionary *entity.ProductCategoryDictionary) error {
	err := s.ProductCategoryRepository.Update(productCategoryDictionary)
	if err != nil {
		return err
	}
	return nil
}

func (s *DictionaryService) DeleteProductCategoryDictionary(productCategoryDictionaryID uint) error {
	err := s.ProductCategoryRepository.Delete("Product_dict_ID", uint64(productCategoryDictionaryID))
	if err != nil {
		return err
	}
	return nil
}
