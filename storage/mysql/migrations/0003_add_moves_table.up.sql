create table if not exists `moves` (`Id` bigint not null primary key auto_increment, `GameId` varchar(16), `Version` bigint, `Name` varchar(16), `Blob` text)  engine=InnoDB charset=utf8;