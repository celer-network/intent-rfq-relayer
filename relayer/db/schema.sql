CREATE DATABASE IF NOT EXISTS rfq_relayer;
USE rfq_relayer;

CREATE TABLE IF NOT EXISTS relayer (
    quote_hash TEXT PRIMARY KEY NOT NULL,
    src_chain_id INT NOT NULL,
    src_token_symbol TEXT NOT NULL,
    dst_chain_id INT NOT NULL,
    dst_token_symbol TEXT NOT NULL,
    base_fee TEXT NOT NULL,
    base_fee_usd DECIMAL NOT NULL,
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    INDEX(create_time)
);

