COPY users
FROM
	'/tmp/data/users.csv' WITH (FORMAT csv, HEADER true);