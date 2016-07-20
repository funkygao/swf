DROP TABLE IF EXISTS `history`;
CREATE TABLE `history` (
    `run_id` bigint(20) NOT NULL,
    `events` text NOT NULL DEFAULT '',
    PRIMARY KEY (`run_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
