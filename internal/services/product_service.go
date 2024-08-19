package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
	"github.com/gabriel-hawerroth/capitech-back/internal/entity"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"
	awsclients "github.com/gabriel-hawerroth/capitech-back/third_party/aws"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
	S3Client          awsclients.S3Client
}

func NewProductService(productRepository repositories.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (s *ProductService) GetById(id int) (*entity.Product, error) {
	return s.ProductRepository.GetById(id)
}

func (s *ProductService) GetFilteredProducts(params dto.ProductQueryParams) (*dto.PaginationResponse[entity.Product], error) {
	content, err := s.ProductRepository.GetFilteredProducts(params)
	if err != nil {
		return nil, err
	}

	totalItems, err := s.ProductRepository.GetFilteredProductsCount(params)
	if err != nil {
		return nil, err
	}

	return &dto.PaginationResponse[entity.Product]{
		Content:    content,
		TotalItems: totalItems,
	}, nil
}

func (s *ProductService) Create(dto dto.SaveProductDto) (*entity.Product, error) {
	return s.ProductRepository.Create(dto)
}

func (s *ProductService) Update(id int, dto dto.SaveProductDto) (*entity.Product, error) {
	return s.ProductRepository.Update(id, dto)
}

func (s *ProductService) ChangeImage(productId int, w http.ResponseWriter, r *http.Request) error {
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

	if product.Image != nil && *product.Image != "" {
		err = s.S3Client.UpdateS3File(*product.Image, fileName, file)
	} else {
		err = s.S3Client.UploadS3File(fileName, file)
	}

	if err != nil {
		return fmt.Errorf("failed to upload image: %v", err)
	}

	// product.Image = &fileName

	// TODO update product image in database

	return nil
}
