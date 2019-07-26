-- PostgreSQL

DROP SCHEMA public CASCADE;
CREATE SCHEMA public;



DROP TABLE IF exists image_manage;


-- Camera撮影Image格納
CREATE TABLE image_manage (
    id bigserial primary key,
    image_data bytea,                           -- 画像バイナリー
    exif_lat double precision NOT NULL,          -- GPS緯度
    exif_lon double precision NOT NULL,          -- GPS経度
    exif_alt double precision NOT NULL,          -- GPS高度
    exif_date_time timestamp not null default current_timestamp, -- 撮影日時
    created_by bigint not null,
    created_at timestamp not null default current_timestamp,
    updated_by bigint not null,
    updated_at timestamp not null default current_timestamp
);




DROP TABLE IF exists role;
-- ロールテーブル
CREATE TABLE role (
    id bigserial PRIMARY KEY,
    role_name varchar(100) NOT NULL,  -- ロール名
    created_by bigint NOT NULL REFERENCES account(id),
    created_at timestamp with time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_by bigint NOT NULL REFERENCES account(id),
    updated_at timestamp with time zone NOT NULL DEFAULT (now() at time zone 'utc')
);
COMMENT ON TABLE role IS             'ロールテーブル';
COMMENT ON COLUMN role.role_name IS  'ロール名';



DROP TABLE IF exists account;

-- アカウントテーブル
CREATE TABLE account (
    id bigserial PRIMARY KEY,
    role_id bigint NOT NULL,                      -- ロールID
    name varchar(40)                              -- 名前
    firebase_id varchar(28) UNIQUE,               -- FirebaseID
    created_by bigint NOT NULL REFERENCES account(id),
    created_at timestamp with time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_by bigint NOT NULL REFERENCES account(id),
    updated_at timestamp with time zone NOT NULL DEFAULT (now() at time zone 'utc')
);
COMMENT ON TABLE account IS                'アカウントテーブル';
COMMENT ON COLUMN account.role_id IS       'ロールID';
COMMENT ON COLUMN account.operator_id IS   '作業員ID';
COMMENT ON COLUMN account.firebase_id IS   'FirebaseID';
