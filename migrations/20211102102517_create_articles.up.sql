create table articles (
                       id varchar not null primary key,
                       state varchar not null unique,
                       title varchar not null,
                       text varchar
);
