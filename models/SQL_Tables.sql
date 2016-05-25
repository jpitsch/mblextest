CREATE TABLE `test_software`.`users` (
  `first_name` VARCHAR(75) NOT NULL COMMENT '',
  `last_name` VARCHAR(75) NOT NULL COMMENT '',
  `user_name` VARCHAR(50) NOT NULL COMMENT '',
  `password` VARCHAR(45) NOT NULL COMMENT '',
  `created` DATE NULL COMMENT '',
  `updated` DATE NULL COMMENT '',
  `email` VARCHAR(100) NOT NULL COMMENT '',
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '');


CREATE TABLE `test_software`.`tests` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
  `name` VARCHAR(75) NOT NULL COMMENT '',
  `type` INT NOT NULL COMMENT '',
  `published` TINYINT(1) NULL DEFAULT 0 COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '');


CREATE TABLE `test_software`.`questions` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
  `selected` INT NULL COMMENT '',
  `correct_answer` INT NOT NULL COMMENT '',
  `text` VARCHAR(300) NOT NULL COMMENT '',
  `number` INT NULL COMMENT '',
  `test_id` INT NOT NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '');


CREATE TABLE `test_software`.`answers` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
  `position` INT NOT NULL COMMENT '',
  `text` VARCHAR(500) NOT NULL COMMENT '',
  `question_id` INT NOT NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '');