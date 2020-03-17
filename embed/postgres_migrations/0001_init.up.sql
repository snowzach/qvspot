CREATE TABLE IF NOT EXISTS vendor (
  id TEXT PRIMARY KEY NOT NULL,
  created timestamp with time zone,
  updated timestamp with time zone,
  name TEXT DEFAULT '',
  description TEXT DEFAULT ''
);

CREATE TABLE IF NOT EXISTS product (
  id TEXT PRIMARY KEY NOT NULL,
  created timestamp with time zone,
  updated timestamp with time zone,
  vendor_id TEXT DEFAULT '',
  name TEXT DEFAULT '',
  description TEXT DEFAULT '',
  pic_url TEXT DEFAULT '',
  attr jsonb DEFAULT '{}'::jsonb,
  attr_num jsonb DEFAULT '{}'::jsonb
);

CREATE TABLE IF NOT EXISTS location (
  id TEXT PRIMARY KEY NOT NULL,
  created timestamp with time zone,
  updated timestamp with time zone,
  vendor_id TEXT DEFAULT '',
  name TEXT DEFAULT '',
  description TEXT DEFAULT '',
  position jsonb DEFAULT '{}'::jsonb
);

CREATE TABLE IF NOT EXISTS item (
  id TEXT PRIMARY KEY NOT NULL,
  created timestamp with time zone,
  updated timestamp with time zone,
  vendor_id TEXT DEFAULT '',
  product_id TEXT DEFAULT '',
  location_id TEXT DEFAULT '',
  stock double precision DEFAULT 0.0,
  price double precision DEFAULT 0.0,
  unit boolean DEFAULT 'false',
  start_time timestamp with time zone,
  end_time timestamp with time zone
);
