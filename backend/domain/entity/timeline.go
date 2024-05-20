package entity

type Timeline struct {
	Per    int64
	Page   int64
	Tweets []Tweet
}

func NewTimeline(
	per int64,
	page int64,
	tweets []Tweet,
) (*Timeline, error) {
	timeline := Timeline{
		Per:    per,
		Page:   page,
		Tweets: tweets,
	}

	return &timeline, nil
}
