CREATE TABLE users (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	username STRING NOT NULL,
	node STRING NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	UNIQUE INDEX users_username_key (username ASC),
	FAMILY "primary" (id, username, node)
)