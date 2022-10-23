drop table posts cascade;
drop table comments;

create table posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);
