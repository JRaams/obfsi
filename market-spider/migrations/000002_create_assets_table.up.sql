create table assets (
  id serial primary key not null,
  name text not null,
  assetgroup_id integer not null,
  nameColor	text not null,
  iconUrl	text not null,
  type	text not null,
  foreign key(assetgroup_id) references assetgroups(id)
);