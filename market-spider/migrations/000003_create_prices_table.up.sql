create table prices (
  id serial primary key not null,
  asset_id integer not null,
  time	integer not null,
  listings	integer not null,
  price	integer not null,
  foreign key(asset_id) references assets(id)
);