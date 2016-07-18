DROP TABLE IF EXISTS `history`;
CREATE TABLE `history` (
    `name` varchar(128) NOT NULL,
    `ver` char(64) NOT NULL DEFAULT '',
    `domain` char(64) NOT NULL DEFAULT '',
    `cluster` char(64) NOT NULL DEFAULT '',
    `desc` varchar(1024) NOT NULL DEFAULT '',
    `child_policy` varchar(128) NOT NULL DEFAULT '',
    PRIMARY KEY (`name`, `ver`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
