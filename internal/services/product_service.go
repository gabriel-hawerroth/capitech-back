package services

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
	"github.com/gabriel-hawerroth/capitech-back/internal/entity"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"
	awsclients "github.com/gabriel-hawerroth/capitech-back/third_party/aws"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
	SearchLogService  SearchLogService
	S3Client          awsclients.S3Client
}

func NewProductService(productRepository repositories.ProductRepository, s3Client awsclients.S3Client, searchLogService SearchLogService) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
		SearchLogService:  searchLogService,
		S3Client:          s3Client,
	}
}

func (s *ProductService) GetById(id *int) (*entity.Product, error) {
	product, err := s.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	go func() {
		s.SearchLogService.Save(dto.SaveSearchLogDTO{
			FieldKey:   "id",
			FieldValue: strconv.Itoa(*id),
		})
	}()

	return product, nil
}

func (s *ProductService) GetFilteredProducts(params dto.ProductQueryParams) (*dto.PaginationResponse[entity.Product], error) {
	content, err := s.ProductRepository.GetFilteredProducts(params)
	if err != nil {
		return nil, err
	}

	go s.SaveSearchLog(params.Filters)

	totalItems, err := s.ProductRepository.GetFilteredProductsCount(params)
	if err != nil {
		return nil, err
	}

	return &dto.PaginationResponse[entity.Product]{
		Content:    content,
		TotalItems: totalItems,
	}, nil
}

func (s *ProductService) Create(dto dto.SaveProductDTO) (*entity.Product, error) {
	return s.ProductRepository.Create(dto)
}

func (s *ProductService) Update(id int, dto dto.SaveProductDTO) (*entity.Product, error) {
	return s.ProductRepository.Update(id, dto)
}

func (s *ProductService) ChangeImage(productId *int, w http.ResponseWriter, r *http.Request) error {
	const maxUploadSize = 5 << 20

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return err
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return err
	}
	defer file.Close()

	if handler.Size > maxUploadSize {
		http.Error(w, "File size exceeds 5MB", http.StatusBadRequest)
		return err
	}

	fileType := handler.Header.Get("Content-Type")
	if !strings.HasPrefix(fileType, "image/") {
		http.Error(w, "The uploaded file is not an image", http.StatusBadRequest)
		return err
	}

	fileName := s.S3Client.GetS3ProductFileName(productId)

	product, err := s.GetById(productId)
	if err != nil {
		return fmt.Errorf("failed to get product: %v", err)
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	// Upload image to S3
	wg.Add(1)
	go func() {
		defer wg.Done()

		if product.Image != nil && *product.Image != "" {
			err = s.S3Client.UpdateS3File(product.Image, &fileName, &file)
		} else {
			err = s.S3Client.UploadS3File(&fileName, &file)
		}

		if err != nil {
			errChan <- fmt.Errorf("failed to upload image to s3: %v", err)
		}
	}()

	// Update image on database
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := s.ProductRepository.ChangeImage(productId, &fileName); err != nil {
			errChan <- fmt.Errorf("failed to update image on database: %v", err)
		}
	}()

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return <-errChan
	}

	return nil
}

func (s *ProductService) RemoveImage(productId *int) error {
	product, err := s.GetById(productId)
	if err != nil {
		return fmt.Errorf("failed to get product: %v", err)
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	// Delete image from S3
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := s.S3Client.DeleteS3File(product.Image); err != nil {
			errChan <- fmt.Errorf("failed to delete image from s3: %v", err)
		}
	}()

	// Delete image on database
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := s.ProductRepository.RemoveImage(productId); err != nil {
			errChan <- fmt.Errorf("failed to remove image on database: %v", err)
		}
	}()

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return <-errChan
	}

	return nil
}

func (s *ProductService) ChangePrice(productId int, newPrice float64) error {
	return s.ProductRepository.ChangePrice(productId, newPrice)
}

func (s *ProductService) ChangeStockQuantity(productId int, newStockQuantity int) error {
	return s.ProductRepository.ChangeStockQuantity(productId, newStockQuantity)
}

func (s *ProductService) SaveSearchLog(filters dto.ProductFilter) {
	if isNotEmpty(filters.Name) {
		s.saveLog("name", *filters.Name)
	}

	if filters.MinPrice > 0 {
		s.saveLog("price", strconv.Itoa(int(filters.MinPrice)))
	}

	if filters.MaxPrice < 50000 {
		s.saveLog("price", strconv.Itoa(int(filters.MaxPrice)))
	}

	if len(filters.Categories) > 0 {
		for index := range filters.Categories {
			s.saveLog("category", strconv.Itoa(filters.Categories[index]))
		}
	}
}

func (s *ProductService) saveLog(fieldKey, fieldValue string) {
	s.SearchLogService.Save(dto.SaveSearchLogDTO{
		FieldKey:   fieldKey,
		FieldValue: fieldValue,
	})
}

func isNotEmpty(s *string) bool {
	return s != nil && *s != ""
}
