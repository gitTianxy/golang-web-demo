/* user */
CREATE TABLE IF NOT EXISTS `golang_web_demo`.`user` (
 `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
 `name` VARCHAR(45) NOT NULL COMMENT '用户名称',
 `age` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户年龄',
 `sex` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别',
 PRIMARY KEY (`id`))
 ENGINE = InnoDB
 AUTO_INCREMENT = 1
 DEFAULT CHARACTER SET = utf8
 COLLATE = utf8_general_ci
 COMMENT = '用户表'

/* item */
CREATE TABLE `golang_web_demo`.`item` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(45) NOT NULL COMMENT '名称',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='item 表'

