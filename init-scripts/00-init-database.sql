create table if not exists blacklist
(
    id varchar
);

create index idx_blacklist_id on blacklist (id);

insert into blacklist(id) values('1');
insert into blacklist(id) values('2');
insert into blacklist(id) values('3');
insert into blacklist(id) values('4');
insert into blacklist(id) values('5');
