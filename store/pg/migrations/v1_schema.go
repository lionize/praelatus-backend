package migrations

const v1body = `
CREATE OR REPLACE FUNCTION update_date()	
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_date = now();
	RETURN NEW;	
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS users (
    id              SERIAL PRIMARY KEY,
    username        varchar(40) UNIQUE NOT NULL,
    password        varchar(250) NOT NULL,
    email           varchar(250) NOT NULL,
    full_name       varchar(250) NOT NULL,
    is_admin        boolean DEFAULT false,
    gravatar        varchar(250),
    profile_picture varchar(250)
);

CREATE TABLE IF NOT EXISTS teams (
    id           SERIAL PRIMARY KEY,
    name         varchar(40) NOT NULL,
    url_slug     varchar(80) NOT NULL,
    homepage     varchar(250),
    icon_url     varchar(250),

    lead_id integer REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS projects (
    id              SERIAL PRIMARY KEY,
    name            varchar(250) NOT NULL,
    key				varchar(40) NOT NULL,
    git_repo        varchar(250),
    homepage        varchar(250),
    icon_path       varchar(250),

    project_lead_id integer REFERENCES users (id),
    team_id         integer REFERENCES teams (id)
);

CREATE TABLE IF NOT EXISTS statuses (
    id   SERIAL PRIMARY KEY,
    name varchar(250) 
);

CREATE TABLE IF NOT EXISTS workflows (
    id   SERIAL PRIMARY KEY,
    name varchar(250),

    project_id integer REFERENCES projects (id)
);

CREATE TABLE IF NOT EXISTS transitions (
    id          SERIAL PRIMARY KEY,
	name		varchar(250),

    workflow_id integer REFERENCES workflows (id),
    status_id   integer REFERENCES statuses (id)
);

CREATE TABLE IF NOT EXISTS transitions_to_statuses (
    transition_id integer REFERENCES transitions (id),
    status_id     integer REFERENCES statuses (id)
);

CREATE TABLE IF NOT EXISTS hooks (
    id            SERIAL PRIMARY KEY,
    body          text,
    delivery      varchar(20),

    transition_id integer REFERENCES transitions (id)
);

CREATE TABLE IF NOT EXISTS fields (
    id        SERIAL PRIMARY KEY,
    name      varchar(250),
    data_type varchar(6)
);

CREATE TABLE IF NOT EXISTS ticket_types (
    id        SERIAL PRIMARY KEY,
    name      varchar(250),
    icon_path varchar(250)
);

CREATE TABLE IF NOT EXISTS tickets (
    id           SERIAL PRIMARY KEY,
	updated_date timestamp,
	created_date timestamp DEFAULT current_timestamp,
    key          varchar(250) NOT NULL,
    summary      varchar(250) NOT NULL,
    description  text NOT NULL,

    project_id     integer REFERENCES projects (id) NOT NULL,
    assignee_id    integer REFERENCES users (id),
    reporter_id    integer REFERENCES users (id) NOT NULL,
    ticket_type_id integer REFERENCES ticket_types (id) NOT NULL,
    status_id      integer REFERENCES statuses (id) NOT NULL
);

CREATE TRIGGER update_ticket_updated_date BEFORE 
UPDATE ON tickets FOR EACH ROW EXECUTE PROCEDURE update_date();

CREATE TABLE IF NOT EXISTS field_values (
    id		  SERIAL PRIMARY KEY,
    value     jsonb,

    ticket_id integer REFERENCES tickets (id),
    field_id  integer REFERENCES fields (id)
);
CREATE INDEX idxfv ON field_values (value);

CREATE TABLE IF NOT EXISTS field_tickettype_project (
    id             SERIAL PRIMARY KEY,

    field_id       integer REFERENCES fields (id),
    ticket_type_id integer REFERENCES ticket_types (id),
    project_id     integer REFERENCES projects (id)
);

CREATE TABLE IF NOT EXISTS database_information (
	schema_version integer	
);

CREATE TABLE IF NOT EXISTS comments (
	id SERIAL PRIMARY KEY,
	body text,
	author_id integer REFERENCES users (id),
	ticket_id integer REFERENCES tickets (id)
);

CREATE TABLE IF NOT EXISTS labels (
	id SERIAL PRIMARY KEY,
	name varchar(255)
);

CREATE TABLE IF NOT EXISTS labels_tickets (
	label_id integer REFERENCES labels (id),
	ticket_id integer REFERENCES tickets (id),
	PRIMARY KEY(label_id, ticket_id)
);

CREATE TABLE IF NOT EXISTS permissions (
	id			 SERIAL PRIMARY KEY,
	updated_date timestamp,
	created_date timestamp DEFAULT current_timestamp,
	level		 varchar(50),

	project_id	 integer REFERENCES projects (id),
	team_id		 integer REFERENCES teams(id),
	user_id		 integer REFERENCES users (id) NOT NULL
);
`

var v1schema = schema{1, v1body}
