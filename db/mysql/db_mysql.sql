CREATE DATABASE bank;
USE bank;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
    `customer_id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `birth_date` date NOT NULL,
    `city` varchar(100) NOT NULL,
    `zipcode` varchar(10) NOT NULL,
    `status` tinyint(1) NOT NULL DEFAULT '1',
    PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1006 DEFAULT CHARSET=latin1;

INSERT INTO `customers` VALUES
    (1000, 'Stepan', '1978-12-15', 'Dmitrov', '101', 1),
    (1001, 'Alex', '1983-05-11', 'Novgorod', '121', 1),
    (1002, 'Pavel', '1995-04-15', 'Elista', '131', 1),
    (1003, 'Vladislav', '1999-01-23', 'Moscow', '141', 1),
    (1004, 'Egor', '2002-12-08', 'Vladimir', '151', 1),
    (1005, 'Olga', '1977-12-07', 'Rostov', '161', 0);


DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
    `account_id` int(11) NOT NULL AUTO_INCREMENT,
    `customer_id` int(11) NOT NULL,
    `opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `account_type` varchar(10) NOT NULL,
    `pin` varchar(10) NOT NULL,
    `status` tinyint(4) NOT NULL DEFAULT '1',
    PRIMARY KEY (`account_id`),
    KEY `accounts_FK` (`customer_id`),
    CONSTRAINT `accounts_FK` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=956 DEFAULT CHARSET=latin1;;

INSERT INTO `accounts` VALUES
    (950, 1000, '2022-01-22 10:20:03', 'Saving', '1126', 1),
    (951, 1001, '2022-02-13 01:20:03', 'Checking', '6234', 1),
    (952, 1002, '2022-03-09 11:20:03', 'Checking', '1614', 1),
    (953, 1003, '2022-04-21 05:20:03', 'Saving', '6362', 1),
    (954, 1004, '2022-05-26 07:20:03', 'Checking', '8865', 1),
    (955, 1005, '2022-01-11 04:20:03', 'Saving', '1048', 0);

DROP TABLE IF EXISTS `transactions`;

CREATE TABLE `transactions` (
    `transaction_id` int(11) NOT NULL AUTO_INCREMENT,
    `account_id` int(11) NOT NULL,
    `amount` int(11) NOT NULL,
    `transaction_type` varchar(10) NOT NULL,
    `transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`transaction_id`),
    KEY `transactions_FK` (`account_id`),
    CONSTRAINT `transactions_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;;