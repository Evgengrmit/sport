ALTER TABLE users ADD COLUMN phone varchar(11) null;
ALTER TABLE users ALTER COLUMN email DROP not null;