package main

import (
	pb "com.fengberlin/shippy/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {

	var users []*pb.User
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {

	var user *pb.User
	user.Id = id
	err := repo.db.First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {

	err := repo.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {

	user := &pb.User{}
	if err := repo.db.Where("email = ? ", email).Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}