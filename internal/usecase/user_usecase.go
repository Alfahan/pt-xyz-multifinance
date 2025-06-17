package usecase

import (
	"context"
	"errors"
	"pt-xyz-multifinance/internal/domain"
	"pt-xyz-multifinance/internal/handler/dto"
	"pt-xyz-multifinance/internal/repository"
	"pt-xyz-multifinance/pkg"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

// Register registers a new user
func (u *userUsecase) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// Validasi: Pastikan email unik
	existingUser, err := u.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Buat domain object
	user := &domain.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Simpan ke repository
	if err := u.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	// Buat response
	resp := &dto.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return resp, nil
}

// Login authenticates a user and returns their details
func (u *userUsecase) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// Cari user berdasarkan email
	user, err := u.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate token
	token := pkg.GenerateToken(user.ID) // Pastikan user.ID bertipe int64

	// Buat response
	resp := &dto.LoginResponse{
		Token: token,
	}

	return resp, nil
}
