package conn

// 应用调用主接口
type App interface {
	Request(ph *PackHead) (*PackHead, error)
	Response(ph *PackHead) error
	OnClose()
}
