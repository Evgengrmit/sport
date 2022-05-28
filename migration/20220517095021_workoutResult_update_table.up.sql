ALTER TABLE workout_result ADD COLUMN user_name varchar(255) null;
ALTER TABLE workout_result ALTER created_at TYPE timestamptz;
ALTER TABLE workout_result ALTER COLUMN created_at SET DEFAULT current_timestamp;