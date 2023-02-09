package main

import "testing"

func Test_lengthOfNoRepeatingSubString(t *testing.T) {

	tests := []struct {
		s    string
		want int
	}{
		{"123321", 3},
		{"12344321", 4},
		{"黑化肥挥发发灰会花飞化肥挥发发黑会飞花", 8},
	}
	for _, tt := range tests {

		if got := lengthOfNoRepeatingSubString(tt.s); got != tt.want {
			t.Errorf("lengthOfNoRepeatingSubString() = %v, want %v", got, tt.want)
		}
	}
}

// 性能测试
func BenchmarkSubstr(b *testing.B) {
	test := struct {
		s    string
		want int
	}{
		"黑化肥挥发发灰会花飞化肥挥发发黑会飞花", 8,
	}
	for i := 0; i < b.N; i++ {
		if got := lengthOfNoRepeatingSubString(test.s); got != test.want {
			b.Errorf("lengthOfNoRepeatingSubString() = %v, want %v", got, test.want)
		}
	}

}
