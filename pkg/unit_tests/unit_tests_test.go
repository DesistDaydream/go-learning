package main

import "testing"

func TestUnitTests(t *testing.T) {
	// 这是测试时需要传递给 UnitTests() 的参数
	type args struct {
		needArgs string
	}
	// 可以创建多个测试
	tests := []struct {
		// 测试名称
		name string
		// 需要传递给 UnitTests() 的参数
		args args
		// 需要判断 UnitTests() 的返回值
		want bool
	}{
		// TODO: 在这里写具体的测试用例，也就是执行 UnitTests() 时想要传递的参数和想要获取到的返回值
		// 这是一个 struct 类型的数组，注意书写格式。
		{
			name: "这是第一条测试在下面填写测试想要传递的参数以及想要获取到的返回值",
			args: args{"unittests"},
			want: true,
		},
		{
			name: "这里是第二条测试用例中需要用到的信息",
			args: args{"这里的参数会导致返回值为 false,进而会导致本次测试失败"},
			want: true,
		},
	}
	// 执行我们提供的每一条测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnitTests(tt.args.needArgs); got != tt.want {
				// 如果 UnitTests() 的返回值与我们填写的 want(即想要获得的返回值) 不一致，那么将会报错
				t.Errorf("UnitTests() = %v, want %v", got, tt.want)
			}
		})
	}
}
