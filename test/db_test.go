package test

import (
	"coreGo/database"
	"github.com/stretchr/testify/assert"
	"testing"
)
/*
	CREATE TABLE `users` (
	  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	  `name` varchar(25) NOT NULL,
	  `age` tinyint(3) unsigned NOT NULL DEFAULT '0',
	  PRIMARY KEY (`id`)
	) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

	--
	-- 转存表中的数据 `users`
	--

	INSERT INTO `users` (`id`, `name`, `age`) VALUES
	(1, '张三', 25),
	(2, '李四', 22),
	(3, '田七', 25);
*/

func TestStructQueryField(t *testing.T) {
	user := database.User{1, "张三", 25}
	assert.Equal(t, database.StructQueryField(1), &user)

	//database.StructTx()
	//database.RawQueryField()
	//database.RawQueryAllField()
}

func TestStructQueryAllField(t *testing.T) {
	users := make([]database.User, 0)
	users = append(users, database.User{1, "张三", 25})
	users = append(users, database.User{2, "李四", 22})
	users = append(users, database.User{3, "田七", 25})
	assert.Equal(t, database.StructQueryAllField(), &users)
}

func TestStructInsert(t *testing.T) {
	assert.Equal(t, database.StructInsert(), int64(1))
}

func TestStructUpdate(t *testing.T) {
	assert.Equal(t, database.StructUpdate(), int64(1))
}

func TestStructDel(t *testing.T) {
	assert.Equal(t, database.StructDel(), int64(1))
}

func TestRawQueryField(t *testing.T) {
	database.RawQueryField()
}

func TestRawQueryAllField(t *testing.T) {
	database.RawQueryAllField()
}