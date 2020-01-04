## 错误处理

参考：[清晰架构（Clean Architecture）的Go微服务: 日志管理](https://segmentfault.com/a/1190000021479989)

错误（error）处理:

错误处理与日志记录直接相关，所以我也在这里讨论一下。以下是我在处理错误时遵循的规则。

1.使用堆栈跟踪创建错误

错误消息本身需要包含堆栈跟踪信息。如果错误源自你的程序，你可以导入“github.com/pkg/errors”库来创建错误以包含堆栈跟踪。但是如果它是从另一个库生成的并且该库没有使用“pkg/errors”，你需要用“errors.Wrap（err，message）”语句包装该错误，以获取堆栈跟踪信息。由于我们无法控制第三方库，因此最好的解决方案是在我们的程序中对所有错误进行包装。详情请见[这里](https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package)³。

2.使用堆栈跟踪打印错误
你需要使用“logger.Log.Errorf（”％+vn“，err）”或“fmt.Printf（”％+vn“，err）”以便打印堆栈跟踪信息，关键是“+v”选项（当然你必须已经使用＃1）。

3.只有顶级函数才能处理错误
“处理”表示记录错误并将错误返回给调用者。因为只有顶级函数处理错误，所以错误只在程序中记录一次。顶层的调用者通常是面向用户的程序，它是用户界面程序（UI）或另一个微服务。你希望记录错误消息（因此你的程序中具有记录），然后将消息返回到UI或其他微服务，以便他们可以重试或对错误执行某些操作。

4.所有其他级别函数应只是将错误传播到较高级别
底层或中间层函数不要记录或处理错误，也不要丢弃错误。你可以向错误中添加更多数据，然后传播它。当出现错误时，你不希望停止整个应用程序。

恐慌（Panic）:

除了在本地的“main.go”之外，我从未使用过恐慌（Panic）。它更像是一个bug而不是一个功能。在让我们谈谈日志⁴中，Dave Cheney写道“人们普遍认为应用库不应该使用恐慌”。另一个错误是log.Fatal，它具有与恐慌相同的效果，也应该被禁止。 “log.Fatal”更糟糕，它看起来像一个日志，但是在输出日志后它“恐慌”，这违反了单一责任规则。

恐慌有两个问题。首先，它与错误的处理方式不同，但它实际上是一个错误，一个错误的子类型。现在，错误处理代码需要处理错误和恐慌，例如事务处理代码⁵中的错误处理代码。其次，它会停止应用程序，这非常糟糕。只有顶级主控制程序才能决定如何处理错误，所有其他被调用的函数应该只将错误传播到上层。特别是现在，服务网格层（Service Mesh）可以提供重试等功能，恐慌使其更加复杂。

如果你正在调用第三方库并且它在代码中产生恐慌，那么为了防止代码停止，你需要截获恐慌并从中恢复。以下是代码示例，你需要为每个可能发生恐慌的顶级函数执行此操作（在每个函数中放置“defer catchPanic（）”）。在下面的代码中，我们有一个函数“catchPanic”来捕获并从恐慌中恢复。函数“RegisterUser”在代码的第一行调用“defer catchPanic（）”。有关恐慌的详细讨论，请参阅此处⁶。

func catchPanic() {
    if p := recover(); p != nil {
        logger.Log.Errorf("%+v\n", p)
    }
}

func (uss *UserService) RegisterUser(ctx context.Context, req *uspb.RegisterUserReq)
    (*uspb.RegisterUserResp, error) {
    
     defer catchPanic()
    ruci, err := getRegistrationUseCase(uss.container)
    if err != nil {
        logger.Log.Errorf("%+v\n", err)
        return nil, errors.Wrap(err, "")
    }
    mu, err := userclient.GrpcToUser(req.User)
...
}

结论：

良好的日志记录可以使程序员更有效。你希望使用堆栈跟踪记录错误。 只有顶级函数才能处理错误，所有其他级别函数只应将错误传播到上一级。 不要使用恐慌。
