## go代码生成助手 是给yii-gii提供 类型信息的 

## 已有的实现
- https://github.com/go-xorm/cmd
- https://github.com/smallnest/gen
- https://github.com/Shelnutt2/db2struct
- https://github.com/xxjwxc/gormt

https://github.com/gohouse/converter

https://github.com/fraenky8/tables-to-go

https://github.com/xxjwxc/gormt

https://github.com/xo/xo

### 如何开发web程序
https://www.alexedwards.net/blog/a-recap-of-request-handling


Generally you shouldn't use the DefaultServeMux because it poses a security risk.

Because the DefaultServeMux is stored in a global variable, any package is able to access it and register a route – including any third-party packages that your application imports. If one of those third-party packages is compromised, they could use the DefaultServeMux to expose a malicious handler to the web.

So as a rule of thumb it's a good idea to avoid the DefaultServeMux, and instead use your own locally-scoped ServeMux, like we have been so far. But if you did decide to use it...
