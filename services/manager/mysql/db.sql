DROP TABLE IF EXISTS `workflow_type`;
CREATE TABLE `workflow_type` (
    `name` varchar(128) NOT NULL,
    `ver` char(64) NOT NULL DEFAULT '',
    `domain` char(64) NOT NULL DEFAULT '',
    `cluster` char(64) NOT NULL DEFAULT '',
    `desc` varchar(1024) NOT NULL DEFAULT '',
    `child_policy` varchar(128) NOT NULL DEFAULT '',
    PRIMARY KEY (`name`, `ver`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `activity_type`;
CREATE TABLE `activity_type` (
    `name` varchar(128) NOT NULL,
    `ver` char(64) NOT NULL DEFAULT '',
    `domain` char(64) NOT NULL DEFAULT '',
    `cluster` char(64) NOT NULL DEFAULT '',
    `desc` varchar(1024) NOT NULL DEFAULT '',
    PRIMARY KEY (`name`, `ver`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

