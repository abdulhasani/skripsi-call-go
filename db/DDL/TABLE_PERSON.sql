create table persons (
  id uuid default uuid_generate_v4() unique,
  name_first varchar,
  name_mid varchar,
  name_last varchar not null,
  id_number varchar not null,
  primary key (id, name_last)
);