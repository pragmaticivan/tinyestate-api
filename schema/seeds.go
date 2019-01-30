package schema

// seeds is a string constant containing all of the queries needed to get the
// db seeded to a useful state for development.
//
// Using a constant in a .go file is an easy way to ensure the queries are part
// of the compiled executable and avoids pathing issues with the working
// directory. It has the downside that it lacks syntax highlighting and may be
// harder to read for some cases compared to using .sql files. You may also
// consider a combined approach using a tool like packr or go-bindata.
//
// Note that database servers besides PostgreSQL may not support running
// multiple queries as part of the same execution so this single large constant
// may need to be broken up.

const seeds = `
INSERT INTO states (id, name, acronym, created_at, updated_at) VALUES
	('a2b0639f-2cc6-44b8-b97b-15d69dbb511e', 'California', 'CA', '2017-01-01 00:00:10.2-07', '2017-01-01 00:00:10.2-07'),
	('72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 'Texas', 'TX', '2017-01-01 00:00:10.2-07', '2017-01-01 00:00:10.2-07');
`
