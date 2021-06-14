CREATE TABLE IF NOT EXISTS contract_address (
    id serial primary key,
    address_hex varchar(42)  
);

CREATE TABLE IF NOT EXISTS regitered_users (
    id serial primary key,
    wallet_address varchar(42) not null,
    social_account_id text not null,
    social_source smallint not null default 0 check(social_source > 0),
    refresh_token text not null,
    reward_cupone_sign varchar,
    metadata json,
    added_at timestamp not null default now()
);

ALTER TABLE IF EXISTS regitered_users DROP CONSTRAINT IF EXISTS unique_source;
ALTER TABLE IF EXISTS regitered_users 
    ADD CONSTRAINT unique_source UNIQUE(social_account_id,social_source);