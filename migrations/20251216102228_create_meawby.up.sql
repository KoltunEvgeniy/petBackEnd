CREATE TABLE accounts(
    id uuid PRIMARY KEY,
    phone text UNIQUE NOT NULL,
    role text NOT NULL DEFAULT 'client' CHECK (ROLE IN ('client', 'admin', 'master')),
    created_at timestamp NOT NULL DEFAULT now()
);

CREATE INDEX idx_accounts_phone ON accounts(phone);

CREATE TABLE IF NOT EXISTS clients(
    id uuid PRIMARY KEY,
    account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    name text NOT NULL,
    created_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS masters(
    id uuid PRIMARY KEY,
    account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    name text NOT NULL,
    is_active boolean NOT NULL DEFAULT TRUE,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS services(
    id uuid PRIMARY KEY,
    title text NOT NULL,
    duration_min int NOT NULL,
    price int NOT NULL,
    is_active boolean NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS master_services(
    master_id uuid NOT NULL REFERENCES masters(id) ON DELETE CASCADE,
    service_id uuid NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    PRIMARY KEY (master_id, service_id)
);

CREATE TABLE sms_codes(
    id uuid PRIMARY KEY,
    account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    code text NOT NULL,
    expires_at timestamp NOT NULL
);

CREATE INDEX idx_sms_codes_account_id ON sms_codes(account_id);

CREATE INDEX idx_sms_codes_expires_at ON sms_codes(expires_at);

CREATE TABLE refresh_tokens(
    id uuid PRIMARY KEY,
    account_id uuid NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    token text NOT NULL UNIQUE,
    expires_at timestamp NOT NULL
);

CREATE INDEX idx_refresh_tokens_account_id ON refresh_tokens(account_id);

CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);

CREATE TABLE IF NOT EXISTS appointments(
    id uuid PRIMARY KEY,
    client_id uuid NOT NULL REFERENCES clients(id),
    master_id uuid NOT NULL REFERENCES masters(id),
    service_id uuid NOT NULL REFERENCES services(id),
    start_at timestamp NOT NULL,
    end_at timestamp NOT NULL,
    status varchar(20) NOT NULL CHECK (status IN ('booked', 'canceled', 'completed', 'paid')),
    price int NOT NULL CHECK (price >= 0),
    created_at timestamp NOT NULL DEFAULT now(),
    CHECK (start_at < end_at)
);

CREATE INDEX idx_appointments_master_id ON appointments(master_id);

CREATE INDEX idx_appointments_client_id ON appointments(client_id);

CREATE INDEX idx_appointments_start_at ON appointments(start_at);

CREATE TABLE payments(
    id uuid PRIMARY KEY,
    appointment_id uuid NOT NULL REFERENCES appointments(id) ON DELETE CASCADE,
    amount int NOT NULL CHECK (amount >= 0),
    status varchar(20) NOT NULL CHECK (status IN ('pending', 'paid', 'failed')),
    created_at timestamp NOT NULL DEFAULT now()
);

CREATE INDEX idx_payments_appointment_id ON payments(appointment_id);

CREATE TABLE master_schedules(
    id uuid PRIMARY KEY,
    master_id uuid NOT NULL REFERENCES masters(id) ON DELETE CASCADE,
    day_of_week int NOT NULL CHECK (day_of_week BETWEEN 1 AND 7),
    start_time time NOT NULL,
    end_time time NOT NULL,
    UNIQUE (master_id, day_of_week)
);

CREATE TABLE master_days_off(
    id uuid PRIMARY KEY,
    master_id uuid NOT NULL REFERENCES masters(id) ON DELETE CASCADE,
    date date NOT NULL,
    reason varchar(255)
);

CREATE TABLE availability(
    id uuid PRIMARY KEY,
    master_id uuid NOT NULL REFERENCES masters(id) ON DELETE CASCADE,
    date date NOT NULL,
    is_available boolean NOT NULL DEFAULT TRUE,
    UNIQUE (master_id, date)
);

CREATE INDEX idx_master_schedules_master_id ON master_schedules(master_id);

CREATE INDEX idx_master_master_days_off_master_id ON master_days_off(master_id);

-- CREATE TABLE users(
--     id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
--     account_id uuid NOT NULL UNIQUE,
--     name varchar(20),
--     created_at timestamp NOT NULL DEFAULT now(),
--     CONSTRAINT fk_users_account FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
-- );
-- CREATE TABLE masters(
--     id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
--     account_id uuid NOT NULL UNIQUE,
--     name varchar(20) NOT NULL,
--     is_active boolean NOT NULL DEFAULT TRUE,
--     created_at timestamp NOT NULL DEFAULT now(),
--     CONSTRAINT fk_masters_account FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
-- );
-- CREATE TABLE services(
--     id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
--     title varchar(255) NOT NULL,
--     duration_min int NOT NULL CHECK (duration_min > 0),
--     price int NOT NULL CHECK (price >= 0),
--     is_active boolean NOT NULL DEFAULT TRUE,
--     created_at timestamp NOT NULL DEFAULT now()
-- );
-- CREATE TABLE master_services(
--     master_id uuid NOT NULL REFERENCES masters(id) ON DELETE CASCADE,
--     service_id uuid NOT NULL REFERENCES services(id) ON DELETE CASCADE,
--     PRIMARY KEY (master_id, service_id),
--     CONSTRAINT fk_ms_master FOREIGN KEY (master_id) REFERENCES masters(id) ON DELETE CASCADE,
--     CONSTRAINT fk_ms_service FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE CASCADE
-- );
-- CREATE TABLE refresh_tokens(
--     id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
--     account_id uuid NOT NULL UNIQUE,
--     token varchar(255) NOT NULL,
--     expires_at timestamp NOT NULL,
--     created_at timestamp DEFAULT now(),
--     CONSTRAINT fk_refresh_account FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
-- );
-- CREATE TABLE sms_codes(
--     id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
--     account_id uuid NOT NULL UNIQUE,
--     code varchar(255) NOT NULL,
--     expires_at timestamp NOT NULL,
--     created_at timestamp DEFAULT now(),
--     CONSTRAINT fk_sms_account FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
-- );
-- CREATE INDEX idx_appointments_master_time ON appointments(master_id, start_at, end_at);
-- CREATE INDEX idx_sms_account ON sms_codes(account_id);
-- CREATE INDEX idx_refresh_account ON refresh_tokens(account_id);
-- CREATE INDEX idx_master_services_master ON master_services(master_id);
-- migrate -path ./migrations -database 'postgres://localhost:5432/meawby?sslmode=disable' down
