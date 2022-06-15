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
	findByID(id int) (*model.User, error)
	findIDByUserID(userID int) (int, error)
	findFriendList(id int) ([]*model.User, error)
	findFriendOfFriendList(fList []*model.User) ([]*model.User, error)
	findFriendListExceptBlock(id int) ([]*model.User, error)
	findFriendOfFriendListExcept1HopFriend(fList []*model.User) ([]*model.User, error)
	rmBlockUser(fList []*model.User, bList []*model.Link, id int) ([]*model.User, error)
	rm1HopFriend(ffList []*model.User, fList []*model.User) ([]*model.User)
	getUniqueList(fList []*model.User) ([]*model.User)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (uu *userUsecase) GetUser(uID int) (*model.User, error) {
	id, err := uu.findIDByUserID(uID)
	if err != nil {
		return nil, err
	}

	u, err := uu.findByID(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *userUsecase) GetFriendList(uID int) ([]*model.User, error) {
	id, err := uu.findIDByUserID(uID)
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
	id, err := uu.findIDByUserID(uID)
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

func (uu *userUsecase) findByID(id int) (*model.User, error) {
	foundUser, err := uu.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (uu *userUsecase) findIDByUserID(userID int) (int, error) {
	id, err := uu.userRepo.FindIDByUserID(userID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uu *userUsecase) findFriendList(id int) ([]*model.User, error) {
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

func (uu *userUsecase) findFriendOfFriendList(fList []*model.User) ([]*model.User, error) {
	ffList := make([]*model.User, 0)
	for _, f := range fList {
		nfList, err := uu.findFriendListExceptBlock(f.ID)
		if err != nil {
			return nil, err
		}
		ffList = append(ffList, nfList...)
	}

	return ffList, nil
}

func (uu *userUsecase) findFriendListExceptBlock(id int) ([]*model.User, error) {
	fList, err := uu.findFriendList(id)
	if err != nil {
		return nil, err
	}

	bList, err := uu.userRepo.FindBlockList(id)
	if err != nil {
		return nil, err
	}

	return uu.rmBlockUser(fList, bList, id)
}

func (uu *userUsecase) rmBlockUser(fList []*model.User, bList []*model.Link, id int) ([]*model.User, error) {
	nList := make([]*model.User, 0)

	userID, err := uu.userRepo.FindUserIDByID(id)
	if err != nil {
		return nil, err
	}

	for _, f := range fList {
		isBlock := false

		for _, b := range bList {
			if b.User1ID == f.UserID && b.User2ID == userID || b.User1ID == userID && b.User2ID == f.UserID {
				isBlock = true
				break
			}
		}

		if !isBlock {
			nList = append(nList, f)
		}
	}

	return nList, nil
}

func (uu *userUsecase) rm1HopFriend(ffList []*model.User, fList []*model.User) ([]*model.User) {
	nffList := make([]*model.User, 0)

	for _, ff := range ffList {
		isFriend := false

		for _, f := range fList {
			if f.UserID == ff.UserID {
				isFriend = true
				break
			}
		}

		if !isFriend {
			nffList = append(nffList, ff)
		}
	}

	return nffList
}

func (uu *userUsecase) findFriendOfFriendListExcept1HopFriend(fList []*model.User) ([]*model.User, error) {
	ffList, err := uu.findFriendOfFriendList(fList)
	if err != nil {
		return nil, err
	}

	ffList = uu.rm1HopFriend(ffList, fList)

	return ffList, nil
}

func (uu *userUsecase) getUniqueList(fList []*model.User) ([]*model.User) {
	nfList := make([]*model.User, 0)

	for _, f := range fList {
		isUnique := true

		for _, nf := range nfList {
			if f.UserID == nf.UserID {
				isUnique = false
			}
		}

		if isUnique {
			nfList = append(nfList, f)
		}
	}

	return nfList
}
