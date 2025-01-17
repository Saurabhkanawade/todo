CREATE TABLE users
(id UUID PRIMARY KEY, name varchar(255) NOT NUll,
 email varchar(255) unique NOT NUll,
 password_hash varchar(255) NOT NULL,
 created_at timestamp default current_timestamp,
 updated_at timestamp default current_timestamp
)

select * from task_tags

CREATE TABLE todo_list(
 id uuid primary key, name varchar(255) not null,
 user_id uuid not null,	
 created_at timestamp default current_timestamp,
 updated_at timestamp default current_timestamp,
 foreign key (user_id) references users(id)	
)

CREATE TYPE task_priority AS ENUM ('low', 'medium', 'high');
CREATE TYPE task_status AS ENUM ('pending', 'in progress', 'completed');

-- to show all the enums
SELECT n.nspname AS schema_name,
       t.typname AS enum_name
FROM pg_type t
JOIN pg_catalog.pg_namespace n ON n.oid = t.typnamespace
WHERE t.typtype = 'e';


CREATE TABLE tasks(
    id uuid primary key,
    todo_list_id uuid not null,
    title varchar(255) not null,
	description text null,
	due_date timestamp null,
	ppriority task_priority DEFAULT 'low',
    status task_status DEFAULT 'pending',
	created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
	foreign key (todo_list_id) references todo_list(id)
)

CREATE TABLE tags
( id uuid primary key, name varchar(50) unique not null,
  created_at timestamp default current_timestamp
)

CREATE TABLE task_tags (
	task_id uuid not null,
	tag_id uuid not null,
    created_at timestamp default current_timestamp,
	foreign key (task_id) references tasks(id),
	foreign key (tag_id) references tags(id)
)

ALTER TABLE task_tags add column task_tag_id uuid primary key