package dao

// AddRelation userIdTo关注BeFollowed
func (a *Dao) AddRelation(userIdTo, BeFollowed int64) error {

}

func (a *Dao) DeleteRelation(userIdTo, BeFollowed int64) error {

}

// GetFollowListByUserId 登录用户关注的所有用户列表
func (a *Dao) GetFollowListByUserId(userId int64) ([]int, error) {

}

// GetFollowerListByUserId 当前用户的所有粉丝
func (a Dao) GetFollowerListByUserId(userId int64) ([]int, error) {

}
