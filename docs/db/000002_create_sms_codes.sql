-- 创建短信验证码表
CREATE TABLE IF NOT EXISTS t_sms_codes (
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    deleted_at  TIMESTAMPTZ,
    phone       VARCHAR(80),
    code        VARCHAR(20),
    expire_time TIMESTAMPTZ,
    is_used     BOOLEAN DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS idx_sms_codes_deleted_at ON t_sms_codes(deleted_at);
