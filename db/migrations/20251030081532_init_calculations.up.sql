CREATE TABLE calculations (
                              id SERIAL PRIMARY KEY,
                              expression VARCHAR(255) NOT NULL,
                              result FLOAT NOT NULL,
                              created_at TIMESTAMP DEFAULT NOW()
);