CREATE TABLE users
(
    username            text PRIMARY KEY,
    hashed_password     text        NOT NULL,
    full_name           text        NOT NULL,
    email               text UNIQUE NOT NULL,
    password_changed_at timestamptz NOT NULL DEFAULT '0001-01-01 00:0:00',
    created_at          timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE accounts
(
    id         bigserial PRIMARY KEY,
    owner      text        NOT NULL,
    balance    bigint      NOT NULL,
    currency   text        NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (NOW())
);

CREATE TABLE entries
(
    id         bigserial PRIMARY KEY,
    account_id bigint      NOT NULL,
    amount     bigint      NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (NOW())
);

CREATE TABLE transfers
(
    id              bigserial PRIMARY KEY,
    from_account_id bigint      NOT NULL,
    to_account_id   bigint      NOT NULL,
    amount          bigint      NOT NULL,
    created_at      timestamptz NOT NULL DEFAULT (NOW())
);

ALTER TABLE accounts
    ADD FOREIGN KEY (owner) REFERENCES users (username);
ALTER TABLE entries
    ADD FOREIGN KEY (account_id) REFERENCES accounts (id);
ALTER TABLE transfers
    ADD FOREIGN KEY (from_account_id) REFERENCES accounts (id);
ALTER TABLE transfers
    ADD FOREIGN KEY (to_account_id) REFERENCES accounts (id);

CREATE INDEX ON accounts (owner);
CREATE UNIQUE INDEX  ON accounts (owner, currency);
CREATE INDEX ON entries (account_id);
CREATE INDEX ON transfers (from_account_id);
CREATE INDEX ON transfers (to_account_id);
CREATE INDEX ON transfers (from_account_id, to_account_id);
