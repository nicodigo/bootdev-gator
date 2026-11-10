-- +goose up
create table users (
	id integer primary key,
	created_at timestamp not null,
	updated_at timestamp not null,
	name text not null,
	unique(name)
);

-- +goose down
drop table users;
