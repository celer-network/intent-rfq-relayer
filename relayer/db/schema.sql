CREATE DATABASE IF NOT EXISTS rfq_relayer;
USE rfq_relayer;

CREATE TABLE IF NOT EXISTS relayer (
    quote_hash TEXT PRIMARY KEY NOT NULL,
    src_chain_id INT NOT NULL,
    src_token_symbol TEXT NOT NULL,
    src_token_addr TEXT NOT NULL,
    src_amount TEXT NOT NULL,
    src_fix_amount DECIMAL NOT NULL,
    src_fix_amount_usd DECIMAL NOT NULL,
    src_release_amount TEXT NOT NULL,
    src_release_fix_amount DECIMAL NOT NULL,
    src_release_fix_amount_usd DECIMAL NOT NULL,
    sender TEXT NOT NULL,
    dst_chain_id INT NOT NULL,
    dst_token_symbol TEXT NOT NULL,
    dst_token_addr TEXT NOT NULL,
    dst_amount TEXT NOT NULL,
    dst_fix_amount DECIMAL NOT NULL,
    dst_fix_amount_usd DECIMAL NOT NULL,
    receiver TEXT NOT NULL,
    deadline TIMESTAMPTZ NOT NULL,
    nonce INT NOT NULL,
    refund_to TEXT NOT NULL,
    liq_provider TEXT NOT NULL,
    status INT NOT NULL,
    source_server_type INT NOT NULL,
    latency INT NOT NULL DEFAULT 0,
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    update_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    INDEX(latency, status),
    INDEX(create_time),
    INDEX(liq_provider, create_time)
);

