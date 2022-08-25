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

ALTER TABLE entries
    ADD FOREIGN KEY (account_id) REFERENCES accounts (id);
ALTER TABLE transfers
    ADD FOREIGN KEY (from_account_id) REFERENCES accounts (id);
ALTER TABLE transfers
    ADD FOREIGN KEY (to_account_id) REFERENCES accounts (id);

CREATE INDEX ON accounts (owner);
CREATE INDEX ON entries (account_id);
CREATE INDEX ON transfers (from_account_id);
CREATE INDEX ON transfers (to_account_id);
CREATE INDEX ON transfers (from_account_id, to_account_id);

COMMENT ON TABLE accounts IS '账号表';
COMMENT ON COLUMN accounts.id IS '账号';
COMMENT ON COLUMN accounts.owner IS '所有者';