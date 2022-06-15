package usecase

import (
	"problem1/domain/model"
)

func (uu *userUsecase) getFriendList(id int) ([]*model.User, error) {
	flList, err := uu.userRepo.FindFriendLinkList(id)
	if err != nil {
		return nil, err
	}

	userID, err := uu.userRepo.FindUserID(id)
	if err != nil {
		return nil, err
	}

	idList := make([]int, 0)
	for _, fl := range flList {
		if fl.User1ID != userID {
			idList = append(idList, fl.User1ID)
		} else {
			idList = append(idList, fl.User2ID)
		}
	}

	uList := make([]*model.User, 0)
	for _, id := range idList {
		user, err := uu.userRepo.FindUser(id)
		if err != nil {
			return nil, err
		}
		uList = append(uList, user)
	}

	return uList, nil
}

func (uu *userUsecase) getFriendListExceptBlock(id int, fList []*model.User) ([]*model.User, error) {
	bList, err := uu.userRepo.FindBlockList(id)
	if err != nil {
		return nil, err
	}

	return uu.rmBlockUser(fList, bList, id)
}

func (uu *userUsecase) getFriendOfFriendList(fList []*model.User) ([]*model.User, error) {
	ffList := make([]*model.User, 0)

	for _, f := range fList {
		nfList, err := uu.GetFriendList(f.UserID)
		if err != nil {
			return nil, err
		}

		ffList = append(ffList, nfList...)
	}

	return ffList, nil
}

func (uu *userUsecase) getFriendOfFriendListExcept1HopFriend(fList []*model.User) ([]*model.User, error) {
	ffList, err := uu.getFriendOfFriendList(fList)
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

func (uu *userUsecase) rmBlockUser(fList []*model.User, bList []*model.Link, id int) ([]*model.User, error) {
	nuList := make([]*model.User, 0)

	userID, err := uu.userRepo.FindUserID(id)
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
			nuList = append(nuList, f)
		}
	}

	return nuList, nil
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
