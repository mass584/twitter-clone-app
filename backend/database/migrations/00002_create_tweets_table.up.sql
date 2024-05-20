create table tweets (
	id bigint AUTO_INCREMENT primary key,
	user_id bigint,
	text_contents varchar(255),
	tweet_at datetime,
	created_at timestamp default CURRENT_TIMESTAMP,
	updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
)
