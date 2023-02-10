-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
   id INT NOT NULL AUTO_INCREMENT,
   name varchar(255) DEFAULT NULL COMMENT 'user name',
   age varchar(255) DEFAULT NULL COMMENT 'age',
   created_at datetime DEFAULT NULL COMMENT 'created at',
   updated_at datetime DEFAULT NULL COMMENT 'updated at',
   deleted_at timestamp NULL DEFAULT NULL COMMENT 'deleted at',
   INDEX user_id (id),
   PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
