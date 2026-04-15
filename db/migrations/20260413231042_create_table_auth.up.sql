-- USER TABLE
CREATE TABLE users (
  id VARCHAR(36) NOT NULL,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  password VARCHAR(255) NOT NULL,
  google_id VARCHAR(255) NOT NULL,
  email_verified BOOLEAN DEFAULT FALSE,
  image VARCHAR(255) NULL ,
  role VARCHAR(50) DEFAULT 'USER',
  created_at  BIGINT NOT NULL,
  updated_at  BIGINT NOT NULL,
  PRIMARY KEY (id)
);

-- SESSION TABLE
CREATE TABLE sessions (
  id VARCHAR(36) NOT NULL,
  token TEXT NOT NULL,
  ip_address VARCHAR(50),
  user_agent TEXT,
  user_id VARCHAR(36) NOT NULL,
  expired_at BIGINT NOT NULL,
  created_at  BIGINT NOT NULL,
  updated_at  BIGINT NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_session_user_id 
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE email_otps (
    id VARCHAR(36) NOT NULL,
    email VARCHAR(255) NOT NULL,
    otp_code VARCHAR(255) NOT NULL,
    expired_at BIGINT NOT NULL,
    is_used BOOLEAN DEFAULT FALSE,
    attempt_count INT DEFAULT 0,
    max_attempt INT DEFAULT 5,
    created_at  BIGINT NOT NULL,
    updated_at  BIGINT NOT NULL
);

CREATE INDEX idx_email_otps ON email_otps (email);