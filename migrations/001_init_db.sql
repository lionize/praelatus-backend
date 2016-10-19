CREATE TABLE IF NOT EXISTS users (
    id        SERIAL PRIMARY KEY,
    username  varchar(40) UNIQUE NOT NULL,
    password  varchar(250) NOT NULL,
    email     varchar(250) NOT NULL,
    full_name varchar(250) NOT NULL,
    is_admin  boolean DEFAULT false
)

CREATE TABLE IF NOT EXISTS teams (
    id           SERIAL PRIMARY KEY,
    name         varchar(40) NOT NULL,
    url_slug     varchar(80) NOT NULL,
    homepage     varchar(250),
    team_lead_id integer REFERENCES users (id)
)

CREATE TABLE IF NOT EXISTS projects (
    id              SERIAL PRIMARY KEY,
    name            varchar(250) NOT NULL,
    project_key     varchar(40) NOT NULL,
    git_repo        varchar(250),
    homepage        varchar(250),

    project_lead_id integer REFERENCES users (id),
    team_id         integer REFERENCES teams (id)
)

CREATE TABLE IF NOT EXISTS statuses (
    id   SERIAL PRIMARY KEY,
    name varchar(250) 
)

CREATE TABLE IF NOT EXISTS workflows (
    id         SERIAL PRIMARY KEY,
    name       varchar(250),
)

CREATE TYPE data_types AS ENUM ('FLOAT', 'STRING', 'INT', 'DATE')
CREATE TABLE IF NOT EXISTS fields (
    id        SERIAL PRIMARY KEY,
    name      varchar(250),
    data_type data_types
)

CREATE TABLE IF NOT EXISTS tickets (
    id          SERIAL PRIMARY KEY,
    ticket_key  varchar(250) NOT NULL,
    summary     varchar(250) NOT NULL,
    description text NOT NULL,

    project_id  integer REFERENCES projects (id),
    assignee_id integer REFERENCES users (id),
    reporter_id integer REFERENCES users (id),
    status_id   integer REFERENCES status (id)

)


