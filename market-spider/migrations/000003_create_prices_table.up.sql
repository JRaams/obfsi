create table prices (
  id serial primary key not null,
  assets_id integer not null,
  time	integer not null,
  listings	integer not null,
  price	integer not null,
  foreign key(assets_id) references assets(id)
);