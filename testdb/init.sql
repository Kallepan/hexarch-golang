CREATE TABLE IF NOT EXISTS arith_history (
    id SERIAL PRIMARY KEY,
    x INT NOT NULL,
    y INT NOT NULL,
    result INT NOT NULL,
    operation VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS arith_history_operation_idx ON arith_history (operation);
