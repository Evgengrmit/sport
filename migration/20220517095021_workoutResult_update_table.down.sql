ALTER TABLE workout_result DROP COLUMN user_name;
ALTER TABLE workout_result ALTER created_at TYPE timestamp;
ALTER TABLE workout_result ALTER COLUMN created_at SET DEFAULT current_date;