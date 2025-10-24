# TodoListStudyDemo
简单的golang练习项目，备忘录的登陆注册以及CRUD
TodoList/
├── api/           # 接口处理层（接收请求、返回响应）
│   ├── tasks.go   # 任务相关接口实现
│   └── user.go    # 用户相关接口实现
├── cmd/           # 程序入口
│   └── main.go    # 服务启动入口
├── conf/          # 配置文件
│   ├── config.go  # 配置解析逻辑
│   └── config.ini # 配置参数
├── middleware/    # 中间件
│   └── jwt.go     # JWT认证中间件
├── model/         # 数据模型与数据库操作
│   ├── init.go    # 数据库初始化
│   ├── migrate.go # 数据表迁移
│   ├── task.go    # 任务模型
│   └── user.go    # 用户模型
├── pkg/           # 工具包
│   └── utils.go   # JWT生成与解析工具
├── routes/        # 路由配置
│   └── routes.go  # 路由注册
├── serializer/    # 数据序列化（模型转API响应格式）
│   ├── common.go  # 通用响应结构
│   ├── task.go    # 任务序列化
│   └── user.go    # 用户序列化
└── service/       # 业务逻辑层
    ├── task.go    # 任务相关业务逻辑
    └── user.go    # 用户相关业务逻辑
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
