-- +goose Up
-- +goose StatementBegin
CREATE TABLE `persons` (
   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID of the record',
   `first_name` varchar(50) NOT NULL DEFAULT '' COMMENT 'First name of person',
   `middle_name` varchar(50) NOT NULL DEFAULT '' COMMENT 'Middle name of person',
   `last_name` varchar(50) NOT NULL DEFAULT '' COMMENT 'Last name of person',
   `email` varchar(100) NOT NULL COMMENT 'NOT unique email address for the person',
   `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'Time the record was created',
   `modified_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT 'Time the record was last modified',
   `is_deleted` tinyint(1) DEFAULT 0 COMMENT 'Flag for if the record is deleted',
   PRIMARY KEY `person_pkey` (`id`),
   KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Person model';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `persons`;
-- +goose StatementEnd
