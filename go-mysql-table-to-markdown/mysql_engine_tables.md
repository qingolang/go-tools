# 模块引擎

## 表名engine_cost

### 字段解释

| TABLE_CATALOG | TABLE_SCHEMA | TABLE_NAME | COLUMN_NAME | ORDINAL_POSITION | COLUMN_DEFAULT | IS_NULLABLE | DATA_TYPE | CHARACTER_MAXIMUM_LENGTH | CHARACTER_OCTET_LENGTH | NUMERIC_PRECISION | NUMERIC_SCALE | DATETIME_PRECISION | CHARACTER_SET_NAME | COLLATION_NAME | COLUMN_TYPE | COLUMN_KEY | EXTRA | PRIVILEGES | COLUMN_COMMENT | GENERATION_EXPRESSION | SRS_ID |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| def |mysql |engine_cost |comment |6 |<nil> |YES |varchar |1024 |3072 |<nil> |<nil> |<nil> |utf8 |utf8_general_ci |varchar(1024) | | |select,insert,update,references | | |<nil> |
| def |mysql |engine_cost |cost_name |3 |<nil> |NO |varchar |64 |192 |<nil> |<nil> |<nil> |utf8 |utf8_general_ci |varchar(64) |PRI | |select,insert,update,references | | |<nil> |
| def |mysql |engine_cost |cost_value |4 |<nil> |YES |float |<nil> |<nil> |12 |<nil> |<nil> |<nil> |<nil> |float | | |select,insert,update,references | | |<nil> |
| def |mysql |engine_cost |default_value |7 |<nil> |YES |float |<nil> |<nil> |12 |<nil> |<nil> |<nil> |<nil> |float | |VIRTUAL GENERATED |select,insert,update,references | |(case `cost_name` when _utf8mb4\'io_block_read_cost\' then 1.0 when _utf8mb4\'memory_block_read_cost\' then 0.25 else NULL end) |<nil> |
| def |mysql |engine_cost |device_type |2 |<nil> |NO |int |<nil> |<nil> |10 |0 |<nil> |<nil> |<nil> |int |PRI | |select,insert,update,references | | |<nil> |
| def |mysql |engine_cost |engine_name |1 |<nil> |NO |varchar |64 |192 |<nil> |<nil> |<nil> |utf8 |utf8_general_ci |varchar(64) |PRI | |select,insert,update,references | | |<nil> |
| def |mysql |engine_cost |last_update |5 |CURRENT_TIMESTAMP |NO |timestamp |<nil> |<nil> |<nil> |<nil> |0 |<nil> |<nil> |timestamp | |DEFAULT_GENERATED on update CURRENT_TIMESTAMP |select,insert,update,references | | |<nil> |
### 建表语句

```mysql
CREATE TABLE `engine_cost` (
  `engine_name` varchar(64) NOT NULL,
  `device_type` int NOT NULL,
  `cost_name` varchar(64) NOT NULL,
  `cost_value` float DEFAULT NULL,
  `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comment` varchar(1024) DEFAULT NULL,
  `default_value` float GENERATED ALWAYS AS ((case `cost_name` when _utf8mb3'io_block_read_cost' then 1.0 when _utf8mb3'memory_block_read_cost' then 0.25 else NULL end)) VIRTUAL,
  PRIMARY KEY (`cost_name`,`engine_name`,`device_type`)
) /*!50100 TABLESPACE `mysql` */ ENGINE=MyISAM DEFAULT CHARSET=utf8mb3 STATS_PERSISTENT=0 ROW_FORMAT=DYNAMIC```

