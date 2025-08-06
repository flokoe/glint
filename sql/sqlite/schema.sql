CREATE TABLE providers (
    id         integer  PRIMARY KEY AUTOINCREMENT,
    type       text     NOT NULL,
    url        text     NOT NULL,
    created_at integer  NOT NULL DEFAULT (unixepoch('now'))
);

CREATE TABLE llms (
    id           integer PRIMARY KEY AUTOINCREMENT,
    string       text    NOT NULL,
    name         text    NOT NULL,
    provider_id  integer NOT NULL,
    context_size integer NOT NULL,
    capabilities text,
    is_enabled   integer NOT NULL DEFAULT TRUE,
    created_at   integer NOT NULL DEFAULT (unixepoch('now')),
    updated_at   integer NOT NULL DEFAULT (unixepoch('now')),

    CONSTRAINT fk_providers_llm FOREIGN KEY (provider_id) REFERENCES providers (id)
);


CREATE TABLE conversations (
    id         integer PRIMARY KEY AUTOINCREMENT,
    user_id    integer NOT NULL,
    uuid       text    NOT NULL,
    title      text    DEFAULT 'New chat',
    is_pinned  integer NOT NULL DEFAULT FALSE,
    created_at integer NOT NULL DEFAULT (unixepoch('now')),
    updated_at integer NOT NULL DEFAULT (unixepoch('now')),

    CONSTRAINT fk_users_conversations FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT uni_conversations_uuid UNIQUE (uuid)
);

CREATE TABLE messages (
    id               integer PRIMARY KEY AUTOINCREMENT,
    conversation_id  integer NOT NULL,
    parent_id        integer,
    role             text    NOT NULL,
    content          text    NOT NULL,
    content_rendered text,
    llm_id           integer NOT NULL,
    created_at       integer NOT NULL DEFAULT (unixepoch('now')),

    CONSTRAINT fk_messages_llms FOREIGN KEY (llm_id) REFERENCES llms (id),
    CONSTRAINT fk_conversations_messages FOREIGN KEY (conversation_id) REFERENCES conversations (id)
);
