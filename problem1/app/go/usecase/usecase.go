package usecase

import (
	"problem1/domain/model"
	"problem1/domain/repository"
)

type UserUsecase interface {
	GetUser(uID int) (*model.User, error)
	GetFriendList(uID int) ([]*model.User, error)
	GetFriendOfFriendList(uID int) ([]*model.User, error)
	GetFriendOfFriendListPaging(uID, limit, page int) ([]*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (uu *userUsecase) GetUser(uID int) (*model.User, error) {
	id, err := uu.userRepo.FindIDByUserID(uID)
	if err != nil {
		return nil, err
	}

	u, err := uu.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *userUsecase) GetFriendList(uID int) ([]*model.User, error) {
	id, err := uu.userRepo.FindIDByUserID(uID)
	if err != nil {
		return nil, err
	}

	fList, err := uu.findFriendListExceptBlock(id)
	if err != nil {
		return nil, err
	}

	return uu.getUniqueList(fList), nil
}

func (uu *userUsecase)  GetFriendOfFriendList(uID int) ([]*model.User, error) {
	id, err := uu.userRepo.FindIDByUserID(uID)
	if err != nil {
		return nil, err
	}

	fList, err := uu.findFriendListExceptBlock(id)
	if err != nil {
		return nil, err
	}

	ffList, err := uu.findFriendOfFriendListExcept1HopFriend(fList)
	if err != nil {
		return nil, err
	}

	return uu.getUniqueList(ffList), nil
}

func (uu *userUsecase) GetFriendOfFriendListPaging(uID, limit, page int) ([]*model.User, error) {
	ffList, err := uu.GetFriendOfFriendList(uID)
	if err != nil {
		return nil, err
	}

	end := limit * page
	start := end - (limit - 1)
	nffList := make([]*model.User, 0)
	for i, ff := range ffList {
		var u model.User
		if start <= i+1 && i+1 <= end {
			u.UserID, u.Name = ff.UserID, ff.Name
			nffList = append(nffList, &u)
		}
	}

	return nffList, nil
}

