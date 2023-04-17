create table if not exists topics (
	id uuid primary key,
	category varchar,
	title varchar,
	capacity int,
	owner uuid,
	parties uuid[],
	created_at timestamptz,
	updated_at timestamptz,
	finished_at timestamptz
);
