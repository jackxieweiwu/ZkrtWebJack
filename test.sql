create database zkrt_db;

USE zkrt_db;

show variables like 'character_set_%';
show variables like 'collation_%';
SET GLOBAL character_set_server=utf8;
alter database zkrt_db character set utf8;

CREATE TABLE IF NOT EXISTS `zkrt_db`.`ZkrtUser` (
  `Zkrt_id` INT NOT NULL AUTO_INCREMENT,
  `Zkrt_user` VARCHAR(45) NOT NULL,
  `Zkrt_password` VARCHAR(45) NOT NULL,
  `Zkrt_fastName` VARCHAR(45) NOT NULL,
  `Zkrt_lastName` VARCHAR(45) NOT NULL,
  `Zkrt_Jurisdiction` INT NOT NULL,
  `Zkrt_time` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`Zkrt_id`))
ENGINE = InnoDB;


CREATE TABLE IF NOT EXISTS `zkrt_db`.`DroneMsg` (
  `DroneMsg_id` INT NOT NULL AUTO_INCREMENT,
  `DroneMsg_droneGPS` VARCHAR(45) NOT NULL,
  `DroneMsg_peoGPS` VARCHAR(45) NULL,
  `DroneMsg_videoPath` VARCHAR(500) NULL,
  `DroneMsg_photoPath` VARCHAR(500) NULL,
  `DroneMsg_alt` VARCHAR(45) NULL,
  `DroneMsg_y` VARCHAR(45) NULL,
  `DroneMsg_r` VARCHAR(45) NULL,
  `DroneMsg_p` VARCHAR(45) NULL,
  `DroneMsg_time` VARCHAR(45) NULL,
  PRIMARY KEY (`DroneMsg_id`))
ENGINE = InnoDB;



insert into ZkrtUser(Zkrt_user,Zkrt_password,Zkrt_fastName,Zkrt_lastName,
Zkrt_Jurisdiction,Zkrt_time )
values('admin','admin','xie','weiwu',0,'2017-06-14 18:14');

insert into ZkrtUser(Zkrt_user,Zkrt_password,Zkrt_fastName,Zkrt_lastName,
Zkrt_Jurisdiction,Zkrt_time )
values('jack','jack','jack_xie','jack_weiwu',0,'2017-06-14 18:29');