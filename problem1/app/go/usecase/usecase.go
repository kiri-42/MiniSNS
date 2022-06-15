package usecase

import (
	"problem1/domain/model"
	"problem1/domain/repository"
)

type UserUsecaseI interface {
	GetUser(uID int) (*model.User, error)
	GetUserList() ([]*model.User, error)
	GetUserListPaging(limit, page int) ([]*model.User, error)
	GetFriendList(uID int) ([]*model.User, error)
	GetFriendOfFriendList(uID int) ([]*model.User, error)
	GetFriendOfFriendListPaging(uID, limit, page int) ([]*model.User, error)
}

type userUsecaseS struct {
	userRepo repository.UserRepositoryI
}

// NewUserUsecase はuserUsecaseのコンストラクタです。
func NewUserUsecase(userRepo repository.UserRepositoryI) UserUsecaseI {
	return &userUsecaseS{userRepo: userRepo}
}

// GetUser はUserをuser_idで取得します。
func (uu *userUsecaseS) GetUser(uID int) (*model.User, error) {
	id, err := uu.userRepo.FindID(uID)
	if err != nil {
		return nil, err
	}

	u, err := uu.userRepo.FindUser(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUserList はUser listを取得します。
func (uu *userUsecaseS) GetUserList() ([]*model.User, error) {
	uList, err := uu.userRepo.FindUserList()
	if err != nil {
		return nil, err
	}

	return uList, nil
}

// GetUserListPaging はpaging形式のUser listを取得します。
func (uu *userUsecaseS) GetUserListPaging(limit, page int) ([]*model.User, error) {
	uList, err := uu.userRepo.FindUserList()
	if err != nil {
		return nil, err
	}

	end := limit * page
	start := end - (limit - 1)
	nuList := make([]*model.User, 0)
	for i, u := range uList {
		var nu model.User
		if start <= i+1 && i+1 <= end {
			nu.UserID, nu.Name = u.UserID, u.Name
			nuList = append(nuList, &nu)
		}
	}

	return nuList, nil
}

// GetFriendList はfriend listをuser_idで取得します。
func (uu *userUsecaseS) GetFriendList(uID int) ([]*model.User, error) {
	id, err := uu.userRepo.FindID(uID)
	if err != nil {
		return nil, err
	}

	fList, err := uu.getFriendList(id)
	if err != nil {
		return nil, err
	}

	fList, err = uu.getFriendListExceptBlock(id, fList)
	if err != nil {
		return nil, err
	}

	return uu.getUniqueList(fList), nil
}

// GetFriendOfFriendList はfriendのfriend listをuser_idで取得します。
func (uu *userUsecaseS)  GetFriendOfFriendList(uID int) ([]*model.User, error) {
	fList, err := uu.GetFriendList(uID)
	if err != nil {
		return nil, err
	}

	ffList, err := uu.getFriendOfFriendList(fList)
	if err != nil {
		return nil, err
	}

	return uu.getUniqueList(ffList), nil
}

// GetFriendOfFriendListPaging はpaging形式のfriendのfriend listをuser_idで取得します。
func (uu *userUsecaseS) GetFriendOfFriendListPaging(uID, limit, page int) ([]*model.User, error) {
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

