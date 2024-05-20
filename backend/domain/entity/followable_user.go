package entity

type FollowableUser struct {
	User       User
	IsFollowed bool
}

func NewFollowableUser(
	User User,
	IsFollowed bool,
) (*FollowableUser, error) {
	followable_user := FollowableUser{
		User:       User,
		IsFollowed: IsFollowed,
	}

	return &followable_user, nil
}
