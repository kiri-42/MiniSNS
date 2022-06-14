package usecase

import (
	"problem1/domain/model"
	"problem1/domain/repository"
)

type UserUsecase interface {
	FindByID(id int) (*model.User, error)
	FindFriendsByID(id int) ([]*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (uu *userUsecase) FindByID(id int) (*model.User, error) {
	foundUser, err := uu.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (uu *userUsecase) FindFriendsByID(id int) ([]*model.User, error) {
	foundFriends, err := uu.userRepo.FindFriendsByID(id)
	if err != nil {
		return nil, err
	}

	userID, err := uu.userRepo.FindUserIDByID(id)
	if err != nil {
		return nil, err
	}

	idList := make([]int, 0)
	for _, friend := range foundFriends {
		if friend.User1ID != userID {
			idList = append(idList, friend.User1ID)
		} else {
			idList = append(idList, friend.User2ID)
		}
	}

	uList := make([]*model.User, 0)
	for _, id := range idList {
		user, err := uu.userRepo.FindByID(id)
		if err != nil {
			return nil, err
		}
		uList = append(uList, user)
	}

	return uList, nil
}
