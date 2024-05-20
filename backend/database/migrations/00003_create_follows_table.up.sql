create table follows (
	follower_user_id bigint,
	followed_user_id bigint,
  primary key (follower_user_id, followed_user_id),
	created_at timestamp default CURRENT_TIMESTAMP,
	updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
)
