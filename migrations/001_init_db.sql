CREATE TABLE IF NOT EXISTS users (
    id        SERIAL PRIMARY KEY,
    username  varchar(40) UNIQUE NOT NULL,
    password  varchar(250) NOT NULL,
    email     varchar(250) NOT NULL,
    full_name varchar(250) NOT NULL,
    is_admin  boolean DEFAULT false
);

CREATE TABLE IF NOT EXISTS teams (
    id           SERIAL PRIMARY KEY,
    name         varchar(40) NOT NULL,
    url_slug     varchar(80) NOT NULL,
    homepage     varchar(250),
    team_lead_id integer REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS projects (
    id              SERIAL PRIMARY KEY,
    name            varchar(250) NOT NULL,
    project_key     varchar(40) NOT NULL,
    git_repo        varchar(250),
    homepage        varchar(250),

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

CREATE TABLE IF NOT EXISTS workflow_transitions (
    id             SERIAL PRIMARY KEY,

    workflow_id    integer REFERENCES workflows (id),
    from_status_id integer REFERENCES statuses (id),
    to_status_id   integer REFERENCES statuses (id)
);

CREATE TYPE method AS ENUM('EMAIL', 'POST', 'PUT', 'GET', 'DELETE');
CREATE TABLE IF NOT EXISTS hooks (
    id            SERIAL PRIMARY KEY,
    body          text,
    delivery      method,

    transition_id integer REFERENCES workflow_transitions
)

CREATE TYPE data_types AS ENUM ('FLOAT', 'STRING', 'INT', 'DATE');
CREATE TABLE IF NOT EXISTS fields (
    id        SERIAL PRIMARY KEY,
    name      varchar(250),
    data_type data_types
);

-- TODO: Create ticket_types table

CREATE TABLE IF NOT EXISTS tickets (
    id          SERIAL PRIMARY KEY,
    ticket_key  varchar(250) NOT NULL,
    summary     varchar(250) NOT NULL,
    description text NOT NULL,

    project_id     integer REFERENCES projects (id),
    assignee_id    integer REFERENCES users (id),
    reporter_id    integer REFERENCES users (id),
    ticket_type_id integer REFERENCES ticket_types (id),
    status_id      integer REFERENCES status (id)
);

CREATE TABLE IF NOT EXISTS field_values (
    id SERIAL PRIMARY KEY,
    value     jsonb,

    ticket_id integer REFERENCES tickets (id),
    field_id  integer REFERENCES fields (id)
);
CREATE INDEX idxfv ON field_values (value);

-- TODO: finish this table
CREATE TABLE IF NOT EXISTS field_tickettype_project (

);
