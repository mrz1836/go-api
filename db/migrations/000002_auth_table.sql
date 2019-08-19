-- +goose Up
-- +goose StatementBegin
CREATE TABLE `auths` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID of the record',
    `person_id` bigint(20) unsigned NOT NULL COMMENT 'ID of the related person record',
    `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'Time the record was created',
    `modified_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'Time the record was last modified',
    `password_digest` varchar(512) DEFAULT NULL COMMENT 'Password digest for authentication',
    `yubikey_digest` varchar(512) DEFAULT NULL COMMENT 'Secure Yubikey ID',
    `yubikey_backup_digest` varchar(512) DEFAULT NULL COMMENT 'Secure Yubikey Backup ID',
    `email` varchar(100) NOT NULL COMMENT 'Unique email address for the person',
    `email_confirm_token` varchar(512) DEFAULT NULL COMMENT 'Confirm email token (encrypted email address)',
    `email_confirmed` tinyint(1) DEFAULT 0 COMMENT 'Flag for if email is confirmed',
    `email_confirm_time` timestamp NULL DEFAULT NULL COMMENT 'Last confirm email time',
    `last_ip_address` varchar(40) DEFAULT NULL COMMENT 'Last IP address used to login',
    `last_login_at` timestamp NULL DEFAULT NULL COMMENT 'Last login time',
    `last_user_agent` varchar(255) DEFAULT NULL COMMENT 'Last user agent from login',
    `login_count` int(5) unsigned DEFAULT 0 COMMENT 'User incremental login count',
    `reset_force` tinyint(1) DEFAULT 0 COMMENT 'Flag for if the person must reset password before login',
    `reset_password_time` timestamp NULL DEFAULT NULL COMMENT 'Last reset password time',
    `reset_password_token` char(32) DEFAULT NULL COMMENT 'Reset password token',
    `reset_token_expires_at` timestamp NULL DEFAULT NULL COMMENT 'Password reset expiration time',
    `locked` tinyint(1) DEFAULT 0 COMMENT 'Locking out from login',
    `locked_time` timestamp NULL DEFAULT NULL COMMENT 'Time when lock out occurred',
    `locked_by_user_id` bigint(20) unsigned DEFAULT NULL COMMENT 'User ID who created the record',
    `is_deleted` tinyint(1) DEFAULT 0 COMMENT 'Flag for if the record is deleted',
    PRIMARY KEY `auths_pkey` (`id`),
    UNIQUE KEY `email` (`email`),
    UNIQUE KEY `person_id` (`person_id`),
    KEY `is_deleted` (`is_deleted`),
    KEY `reset_password_token` (`reset_password_token`),
    KEY `email_confirm_token` (`email_confirm_token`),
    KEY `email_auth` (`email`,`password_digest`(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Auth record for person model';
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE auths ADD CONSTRAINT auths_fk_1 FOREIGN KEY (person_id) REFERENCES persons(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `auths`;
-- +goose StatementEnd