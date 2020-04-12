package command

import "fmt"

/**
命令模式本质是把某个对象的方法调用封装到对象中，方便传递、存储、调用。

示例中把主板单中的启动(start)方法和重启(reboot)方法封装为命令对象，再传递到主机(box)对象中。于两个按钮进行绑定：

第一个机箱(box1)设置按钮1(buttion1) 为开机按钮2(buttion2)为重启。
第二个机箱(box1)设置按钮2(buttion2) 为开机按钮1(buttion1)为重启。
从而得到配置灵活性。

除了配置灵活外，使用命令模式还可以用作：
1. 批处理
2. 任务队列
3. undo, redo
等把具体命令封装到对象中使用的场合
*/

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

type MotherBoard struct{}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

type Box struct {
	buttion1 Command
	buttion2 Command
}

func NewBox(buttion1, buttion2 Command) *Box {
	return &Box{
		buttion1: buttion1,
		buttion2: buttion2,
	}
}

func (b *Box) PressButtion1() {
	b.buttion1.Execute()
}

func (b *Box) PressButtion2() {
	b.buttion2.Execute()
}