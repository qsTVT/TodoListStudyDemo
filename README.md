# TodoListStudyDemo
简单的golang练习项目，备忘录的登陆注册以及CRUD

#注意事项
##数据库会在程序启动时自动迁移（创建表结构）
##JWT 令牌有效期为 24 小时
##分页默认每页 15 条数据
##任务状态：0 表示未完成，1 表示已完成

#技术栈
后端框架：Gin
ORM：GORM
数据库：MySQL
身份认证：JWT
配置管理：ini
日志：logrus

#环境要求
Go 1.16+
MySQL 5.7+
（可选）Redis（配置文件中包含，但当前功能未充分使用）
