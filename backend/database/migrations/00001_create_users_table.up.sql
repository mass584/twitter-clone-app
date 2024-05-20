create table users (
	id bigint AUTO_INCREMENT primary key,
	display_name varchar(255) not null,
	email varchar(255) not null,
	created_at timestamp default CURRENT_TIMESTAMP,
	updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
	unique idx_email (email)
)
