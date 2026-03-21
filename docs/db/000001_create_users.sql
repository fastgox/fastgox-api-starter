-- 创建用户表
CREATE TABLE IF NOT EXISTS t_user (
    id           BIGSERIAL PRIMARY KEY,
    phone        VARCHAR(80),
    channel_code VARCHAR(80),
    status       SMALLINT DEFAULT 1,
    create_time  TIMESTAMPTZ,
    update_time  TIMESTAMPTZ,
    platform     VARCHAR(80),
    is_auth      SMALLINT DEFAULT 0
);

COMMENT ON COLUMN t_user.channel_code IS '渠道代码';
COMMENT ON COLUMN t_user.is_auth IS '是否已实名认证';
