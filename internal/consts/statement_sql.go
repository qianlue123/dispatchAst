package consts

/* 数据库命令, 两种显示形式 */

// mysql -D -e 'xxx;'
var E = map[int]string{
	1:  "mysql -e 'show databases ;' ",
	2:  "mysql -D asterisk -e 'show tables ;' ",
	3:  "mysql -D asterisk -e 'select id from devices ;' ",
	31: "mysql -D asterisk -e \"SELECT id,user,description FROM devices WHERE id='%s' LIMIT 1 ;\" ",

	// 反查已知列名所在的表, 可能不止一个
	11: "mysql -e \"SELECT table_name from information_schema.columns " +
		"where TABLE_SCHEMA='asterisk' and COLUMN_NAME='%s' ; \" ",
}

// mysql -D -e 'xxx \G'
var EG = map[int]string{
	1:  "mysql -e 'show databases \\G' ",
	2:  "mysql -D asterisk -e 'show tables \\G' ",
	3:  "mysql -D asterisk -e 'select id from devices \\G' ",
	31: "mysql -D asterisk -e \"SELECT id,user,description FROM devices WHERE id='%s' LIMIT 1 \\G \" ",

	// 反查已知列名所在的表, 可能不止一个
	11: "mysql -e \"SELECT table_name from information_schema.columns " +
		"where TABLE_SCHEMA='asterisk' and COLUMN_NAME='%s' \\G \" ",
}
