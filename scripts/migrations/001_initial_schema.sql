-- Create initial schema
CREATE TABLE users (
    id            CHAR(36) PRIMARY KEY, 
    email         VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE vaults (
    id         CHAR(36) PRIMARY KEY, 
    user_id    CHAR(36) NOT NULL,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE secrets (
    id         CHAR(36) PRIMARY KEY, 
    vault_id   CHAR(36) NOT NULL,
    name       VARCHAR(255) NOT NULL,
    data       TEXT NOT NULL,  -- Encrypted JSON (passwords, notes, etc.)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (vault_id) REFERENCES vaults(id) ON DELETE CASCADE
);

CREATE TABLE devices (
    id         CHAR(36) PRIMARY KEY, 
    user_id    CHAR(36) NOT NULL,
    device_id  VARCHAR(255) UNIQUE NOT NULL,
    last_sync  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE audit_logs (
    id         CHAR(36) PRIMARY KEY, 
    user_id    CHAR(36) NOT NULL,
    action     VARCHAR(255) NOT NULL,
    timestamp  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    metadata   JSON DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
