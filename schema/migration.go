package schema

import "github.com/GuiaBolso/darwin"

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Creating table states",
		Script: `CREATE TABLE states (
					id serial PRIMARY KEY,
					name 			VARCHAR(255) NOT NULL,
					abbreviation 	VARCHAR(255) NOT NULL,
					created_at 		TIMESTAMP NOT NULL,
					updated_at 		TIMESTAMP NOT NULL,
					UNIQUE(name),
					UNIQUE(abbreviation)
				);`,
	},
	{
		Version:     2,
		Description: "Creating table cities",
		Script: `CREATE TABLE cities (
					id serial PRIMARY KEY,
					name 			VARCHAR(255) NOT NULL,
					allows_on_wheels  BOOLEAN NOT NULL,
					allows_on_foundation  BOOLEAN NOT NULL,
					requires_care_giver  BOOLEAN NOT NULL,
					state_id INTEGER NOT NULL,
					latitude NUMERIC(14, 11)  NOT NULL,
					longitude NUMERIC(14, 11)  NOT NULL,
					created_at 		TIMESTAMP NOT NULL,
					updated_at 		TIMESTAMP NOT NULL,
					FOREIGN KEY (state_id) REFERENCES states (id)
				);`,
	},
	{
		Version:     3,
		Description: "Creating table canonicals",
		Script: `CREATE TABLE canonicals (
					id serial PRIMARY KEY,
					name 			VARCHAR(255),
					canonical 			VARCHAR(255) NOT NULL,
					allows_on_wheels  BOOLEAN NOT NULL,
					allows_on_foundation  BOOLEAN NOT NULL,
					requires_care_giver  BOOLEAN NOT NULL,
					state_id INTEGER NOT NULL,
					latitude NUMERIC(14, 11)  NOT NULL,
					longitude NUMERIC(14, 11)  NOT NULL,
					created_at 		TIMESTAMP NOT NULL,
					updated_at 		TIMESTAMP NOT NULL,
					FOREIGN KEY (state_id) REFERENCES states (id)
				);`,
	},
	{
		Version:     4,
		Description: "Drop state_id fk from canonicals",
		Script:      `ALTER TABLE canonicals DROP COLUMN IF EXISTS state_id`,
	},
	{
		Version:     5,
		Description: "Change latitude and longitude to default null",
		Script: `ALTER TABLE canonicals ALTER COLUMN latitude DROP NOT NULL;
					  ALTER TABLE canonicals ALTER COLUMN longitude DROP NOT NULL;`,
	},
}
