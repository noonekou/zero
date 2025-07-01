CREATE TABLE `book`
(
  `book` varchar(255) NOT NULL COMMENT 'book name',
  `price` int NOT NULL DEFAULT 0 COMMENT 'book price',
  PRIMARY KEY(`book`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;