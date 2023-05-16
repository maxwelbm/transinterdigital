CREATE DATABASE transinterdigital;

CREATE TABLE IF NOT EXISTS account (
     id SERIAL PRIMARY KEY,
     name VARCHAR(100),
     cpf VARCHAR(14),
     secret VARCHAR(100),
    balance NUMERIC(38, 10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

INSERT INTO account (name, cpf, secret, balance, created_at) VALUES
    ('João Silva', '123.456.789-00', 'hash_1', 100.50, NOW()),
    ('Maria Santos', '987.654.321-00', 'hash_2', 250.75, NOW()),
    ('Pedro Oliveira', '456.789.123-00', 'hash_3', 500.00, NOW()),
    ('Ana Pereira', '321.654.987-00', 'hash_4', 1000.25, NOW()),
    ('Lucas Mendes', '789.123.456-00', 'hash_5', 750.80, NOW());


CREATE TABLE transfers (
    id SERIAL PRIMARY KEY,
    account_origin_id INTEGER,
    account_destination_id INTEGER,
    amount NUMERIC(38, 10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);