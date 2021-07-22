### 题目一

> 基于errgroup实现一个http server的启动和关闭，以及linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
>

代码：go_projects/src/toy_server/main.go

 ### 题目二

> 我们在数据库操作的时候，比如Data Access Object(DAO)层中当遇到一个sql.ErrNoRows的时候，是否应该wrap这个error,抛给上层。为什么？应该怎么做，请写出代码。

DAO层应该尽量向上层隐藏数据库细节，所以不应该抛出错误，而是应该处理这个错误，并且给上层提供一个方法来确定是否发生错误。这里我提供了额外的IsExist()方法用于判断是否出现ErrNoRows错误。

代码：go_projects/src/dao_simulation/main.go

