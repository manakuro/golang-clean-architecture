
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE credit_cards (
   id INT NOT NULL AUTO_INCREMENT,
   user_id INT NOT NULL,
   number varchar(255) DEFAULT NULL COMMENT 'credit number',
   created_at datetime DEFAULT NULL COMMENT 'created at',
   updated_at datetime DEFAULT NULL COMMENT 'updated at',
   deleted_at timestamp NULL DEFAULT NULL COMMENT 'deleted at',
   PRIMARY KEY(id),
   CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='credit card';

-- +goose Down
DROP TABLE credit_cards
-- SQL section 'Down' is executed when this migration is rolled back

