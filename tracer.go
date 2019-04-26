package trace

// Traverはコード内での出来事を記録できるオブジェクトを表すインターフェースです
type Tracer interface {
	Tracer(...interface{})
}