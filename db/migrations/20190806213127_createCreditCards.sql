
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE credit_cards (
   id INT NOT NULL AUTO_INCREMENT,
   user_id INT NOT NULL,
   number varchar(255) DEFAULT NULL COMMENT 'ユーザ名',
   created_at datetime DEFAULT NULL COMMENT '作成日時',
   updated_at datetime DEFAULT NULL COMMENT '更新日時',
   deleted_at timestamp NULL DEFAULT NULL COMMENT '削除日時',
   PRIMARY KEY(id),
   CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='クレジットカード';

-- +goose Down
DROP TABLE credit_cards
-- SQL section 'Down' is executed when this migration is rolled back

