CREATE TABLE IF NOT EXISTS manufacturers (
    manufacturer_id VARCHAR(21) PRIMARY KEY,       -- ランダム生成 ID（NanoID など）
    name VARCHAR(512) NOT NULL UNIQUE, -- メーカー名（前方一致検索用に INDEX 追加）
    deleted_at TIMESTAMP NULL,        -- 論理削除（NULL: 有効, 日付: 削除）
    official_website TEXT,            -- 公式サイト URL
    description TEXT,                 -- メーカーの説明
    logo_url TEXT,                     -- メーカーのロゴ画像 URL
    contact_email VARCHAR(255),       -- 問い合わせメールアドレス
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- 更新日時
);

CREATE INDEX IF NOT EXISTS idx_manufacturers_name ON manufacturers(name);