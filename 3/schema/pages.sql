CREATE DATABASE cms;
USE cms;

CREATE TABLE `pages` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `page_guid` varchar(256) NOT NULL DEFAULT '',
  `page_title` varchar(256) DEFAULT NULL,
  `page_content` mediumtext,
  `page_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `page_guid` (`page_guid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

INSERT INTO `pages` (`id`, `page_guid`, `page_title`, `page_content`, `page_date`) VALUES (NULL, 'hello-world', 'Hello, World!', 'Hey! How are you doin?', CURRENT_TIMESTAMP);
INSERT INTO `pages` (`id`, `page_guid`, `page_title`, `page_content`, `page_date`) VALUES (NULL, 'cool', 'Cool!', 'Really cool!', CURRENT_TIMESTAMP);
INSERT INTO `pages` (`id`, `page_guid`, `page_title`, `page_content`, `page_date`) VALUES (NULL, 'long-long', 'Long-long long long long!', 'Really cool and long, very-very-very long. Too long! Very-very long! Longer then everything else!', CURRENT_TIMESTAMP);
