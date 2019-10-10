# sqlBuilder

用来方便 递进式构造sql语句的
不然就要手动拼接sql了 又丑陋又容易出错

其实满足一般查询条件即可 主要集中在构造Where 子句部分

## 局限
- godb
    + 对于构造like条件支持不够 Where部分的构造能力太弱：
    相较于[ozzo-dbx building-query-conditions](https://github.com/go-ozzo/ozzo-dbx#building-query-conditions)
    
- ozzo-dbx
    + 不能脱离db对象 需要一起使用 如果想跟其他orm库配合比较困难    
    
- goqu
    http://doug-martin.github.io/goqu/docs/expressions.html 看来都需要互斥使用 不能只用片段了！    
