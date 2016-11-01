package migrations

const v2Schema = `
CREATE TABLE IF NOT EXISTS labels (
	id SERIAL PRIMARY KEY,
	name varchar(255),
)

CREATE TABLE IF NOT EXISTS labels_tickets (
	label_id integer REFERENCES labels (id),
	ticket_id integer REFERENCES tickets (id),
	PRIMARY KEY(label_id, ticket_id)
);
`
