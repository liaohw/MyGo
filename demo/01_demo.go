
// 变量
	var a int = 1
	var a2 = 2
	a3 := 3

	a , a2 = a2, a

// 函数
	func getName(firstName , lastName string){
		return "hw" , "liao"
	}
	_, xin := getName()

// 常量
	const zero = 0.0
	const (
		size int64 = 1024
		eof = -1			// 无类型整型常量
	)
	const u,v float32 = 0, 3	//=>0.0,3.0

// iota比较特殊，可以被认为是一个可被编译器修改的常量
	const (
		c0 = iota	//0,const第一次0
		c1 = iota	//1，每出现一次+1
	)
	// 也可以这么写
	const (
		c0 = iota	//0
		c1			//1
	)

// 枚举
	const (
		Sunday = iota	//0	包外可见
        Monday			//1 包外可见
        Tuesday
        numberOfDays	//3 包内私有 ,这个常量没有导出
	)

// 内置类型
	// 基础类型
		bool 
		int8,byte,int16,int,uint,uintptr
		float32,float64
		复数:complex64,complex128
		string
		字符 rune
		错误类型 error
	// 复合类型
	指针	pointer
	数组	array
	切片	slice
	字典	map
	通道	chan
	结构体	struct
	接口	interface







