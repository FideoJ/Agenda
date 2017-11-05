# Agenda-go
![travisCI](https://travis-ci.org/FideoJ/Agenda.svg?branch=master)

## Introdution
Agenda是一个基于命令行界面的会议管理程序，使用golang实现。
## Install
```
go get github.com/FideoJ/Agenda
```

## Usage
```
go build
./Agenda [subcommand] [flags]
```

## Subcommands
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
  - notes: 若已有登录用户，则先登出，无论后续是否能登录成功
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
  - notes: 要求已登录,时间格式:"YYYY-MM-DD HH:mm"
- addParticipants
  - usage: 增加会议参与者
  - args: title string, participant string
  - notes: 要求已登录,仅能操作当前用户为发起者的会议
- removeParticipants
  - usage: 删除会议参与者
  - args: title string, participant string
  - notes: 要求已登录,仅能操作当前用户为发起者的会议，仅剩发起者的会议应删除
- listMeetings
  - usage: 列出所有与当前用户有关的会议
  - args: None
  - notes: 要求已登录,时间格式:"YYYY:MM:DD HH:mm"
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

## Test Subcommands
Data persistence: data of users and meetings will be stored as `*.json` files.

All the related `.json` file is stored in `$homedir/.agenda/` of current user. (eg. `$HOME/.agenda/` in *nix)

```shell
# register
$ ./Agenda register -utest -ptest -etest@test.com -t12345678910
[INFO] 2017/11/05 13:56:01 Register called with username:[test], password:[test], email:[test@test.com], phone:[12345678910]
$ ./Agenda register -utest2 -ptest -etest2@test2.com -t12345678911
[INFO] 2017/11/05 13:57:32 Register called with username:[test2], password:[test], email:[test2@test2.com], phone:[12345678911]

# login
$ ./Agenda login -utest --password test
[INFO] 2017/11/05 13:58:08 Login called with username:[test], password:[test]

# list all users. some users are pre-registered.
$ ./Agenda listUsers
USERNAME             EMAIL                PHONE
abc                  asdf                 3432
abe                  asdf                 3432
marsh                abc@123.com          1234
test                 test@test.com        12345678910
test2                test2@test2.com      12345678911
abd                  asdf                 3432
[INFO] 2017/11/05 13:59:16 ListUsers called

# removeUser: test2
$ ./Agenda removeUser --username test2 --password=test
[INFO] 2017/11/05 14:03:13 RemoveUser called with username:[test2], password:[test]

# createMeeting
$ ./Agenda createMeeting --title=testMeeting -s"2012:12:12 12:12" -e"2012:12:13 13:13" -pabc -pabe
[ERROR] 2017/11/05 14:05:38 parsing time "2012:12:12 12:12" as "2006-01-02 15:04": cannot parse ":12:12 12:12" as "-"
$ ./Agenda createMeeting --title=testMeeting -s"2012-12-12 12:12" -e"2012-12-13 13:13" -pabc -pabe
[INFO] 2017/11/05 14:06:19 CreateMeeting called with title: [testMeeting], startTime: [2012-12-12 12:12], endTime: [2012-12-13 13:13], participants: [[abc abe]]

# addParticipant: user[abd] to meeting[testMeeting]
$ ./Agenda addParticipant --title=testMeeting -pabd
[INFO] 2017/11/05 14:07:26 addParticipant called with title: [testMeeting], participants: [[abd]]

# removeParticipant: user[abe] from meeting[testMeeting]
$ ./Agenda removeParticipant -ttestMeeting -pabe
[INFO] 2017/11/05 14:08:17 removeParticipant called with title: [testMeeting], participants: [[abe]]

# listMeetings: (created another meeting)
$ ./Agenda listMeetings
TITLE           SPONSOR         START-TIME           END-TIME             PARTICIPANTS
testMeeting     test            2012-12-12 12:12     2012-12-13 13:13     abc            abd
testMeeting2    test            2011-12-12 12:12     2011-12-13 13:13     abc
[INFO] 2017/11/05 14:09:22 listMeetings called

# cancelMeeting
$ ./Agenda cancelMeeting --title=testMeeting2
[INFO] 2017/11/05 14:10:07 CancelMeeting called with title: [testMeeting2]
$ ./Agenda listMeetings
TITLE           SPONSOR         START-TIME           END-TIME             PARTICIPANTS
testMeeting     test            2012-12-12 12:12     2012-12-13 13:13     abc            abd
[INFO] 2017/11/05 14:10:12 listMeetings called

# quitMeeting
$ ./Agenda logout
[INFO] 2017/11/05 14:16:38 Logout called
$ ./Agenda login --username=abc -p111
[ERROR] 2017/11/05 14:16:49 Wrong username or password
$ ./Agenda login --username=abc -p123
[INFO] 2017/11/05 14:17:01 Login called with username:[abc], password:[123]
$ ./Agenda quitMeeting --title=testMeeting
[INFO] 2017/11/05 14:17:53 QuitMeeting called with title: [testMeeting]
$ ./Agenda login -utest -ptest
[INFO] 2017/11/05 14:18:43 Login called with username:[test], password:[test]
$ ./Agenda listMeetings
TITLE           SPONSOR         START-TIME           END-TIME             PARTICIPANTS
testMeeting     test            2012-12-12 12:12     2012-12-13 13:13     abd
[INFO] 2017/11/05 14:18:47 listMeetings called

# clearMeetings
$ ./Agenda clearMeetings
[INFO] 2017/11/05 14:19:20 ClearMeetings called
$ ./Agenda listMeetings
TITLE           SPONSOR         START-TIME           END-TIME             PARTICIPANTS
[INFO] 2017/11/05 14:19:26 listMeetings called
```