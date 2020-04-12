package proxy

/**
代理模式用于延迟处理操作或者在进行实际操作前后进行其它处理。

	代理模式的常见用法有
	1. 虚代理
	2. COW代理
	3. 远程代理
	4. 保护代理
	5. Cache 代理
	6. 防火墙代理
	7. 同步代理
	8. 智能指引
*/

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string
	res += "pre:"
	res += p.real.Do()
	res += ":after"
	return res
}