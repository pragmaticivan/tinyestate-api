package schema

import "github.com/GuiaBolso/darwin"

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Creating table states",
		Script: `CREATE TABLE states (
					id UUID,
					name 			VARCHAR(255) NOT NULL,
					acronym 		VARCHAR(255) NOT NULL,
					created_at 		TIMESTAMP NOT NULL,
					updated_at 		TIMESTAMP NOT NULL,

					PRIMARY KEY (id)
				);`,
	},
}
