## 子命令设计
- help
  - usage: 列出命令说明
  - args: command string (= "all")
  - notes: 参数默认为all，打印所有命令的说明
- register
  - usage: 用户注册
  - args: username string, password string, email string, phone string
  - notes: None
- login
  - usage: 用户登录
  - args: username string, password string
  - notes: 若已登录，则先登出，无论是否能登录成功
- logout
  - usage: 用户登出
  - args: None
  - notes: 若未登录，则静默
- listUsers
  - usage: 列出所有用户
  - args: None
  - notes: 要求已登录
- removeUser
  - usage: 用户删除
  - args: username string, password string
  - notes: 若成功删除当前用户，登出
- createMeeting
  - usage: 创建会议
  - args: title string, startTime string, endTime string, participants []string
  - notes: 要求已登录,时间格式:"YYYY:MM:DD HH:mm"
- addParticipants
  - usage: 增加会议参与者
  - args: title string, participants []string
  - notes: 要求已登录,仅能操作当前用户为发起者的会议
- removeParticipants
  - usage: 删除会议参与者
  - args: title string, participants []string
  - notes: 要求已登录,仅能操作当前用户为发起者的会议，仅剩发起者的会议应删除
- listMeetings
  - usage: 列出所有与当前用户有关的会议
  - args: None
  - notes: 要求已登录,时间格式:"YYYY-MM-DD HH:mm"
- cancelMeeting
  - usage: 取消会议
  - args: title string
  - notes: 要求已登录,仅能操作当前用户为发起者的会议
- quitMeeting
  - usage: 退出会议
  - args: title string
  - notes: 要求已登录,仅能操作当前用户为参与者的会议
- clearMeetings
  - usage: 清空会议
  - args: None
  - notes: 要求已登录,清除当前用户为发起者的会议

## 编码约定
1. 获得参数并排除空值，使用utils.GetNonEmptyString(cmd, flag)
2. 调用service结束时，调用logger.Info("CMD called with ARGS", args...)
3. 错误集中写在err/err.go中，出错时调用logger.FatalIf(err.SomeErr)，不再将错误传递给上层函数
4. 可能接收到错误的地方（一般是别人写好的包），调用logger.FatalIf(err)
5. 关心是否已登录时：
```
	_, loggedIn := storage.LoadCurUser()
	if !loggedIn {
		logger.FatalIf(err.RequireLoggedIn)
	}
```
