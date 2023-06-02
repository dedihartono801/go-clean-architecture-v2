CREATE TABLE transaction (
    id VARCHAR(255) NOT NULL,
    admin_id VARCHAR(255) NOT NULL,
    total_quantity int NOT NULL,
    total_transaction int NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    PRIMARY KEY(id)
);