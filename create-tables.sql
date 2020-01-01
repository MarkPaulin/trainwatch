CREATE TABLE stations (
  station_code text not null,
  station_name text,
  primary key (station_code)
)

CREATE TABLE trains (
  station_code text not null,
  timestamp text,
  headcode text,
  uid text not null,
  arrive_status text,
  arrive_timestamp text,
  depart_timestamp text,
  date text,
  platform integer,
  status text,
  expected_arrive_time text,
  expected_depart_time text,
  delay integer,
  origin text,
  destination text,
  foreign key (station_code) references stations(station_code),
  primary key (station_code, uid, date)
)
