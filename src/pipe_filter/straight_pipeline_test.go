/*
pipe filter 是一种软件设计架构模式

------------------------------------------ Ads SSP ------------------------------------------
|                                                                                            |
|   request parsing -> params processing -> verification -> filtering -> ranking -> filling  |
|                                                                                            |
---------------------------------------------------------------------------------------------

- 非常适合与数据处理以及数据分析系统
- Filter 封装数据处理的功能
- 松耦合：Filter 只跟数据（格式）耦合
- Pipe 用于连接 Filter 传递数据或者在异步处理过程中缓冲数据流，进程内同步调用时，pipe 演变为数据在方法调用间传递

Demo
                       ["1", "2", "3"]     [1,2,3]
"1,2,3"  -->  SplitFilter  -->  ToIntFilter  -->  SumFilter  --> 6
*/
package pipe_filter

import "testing"

func TestStraightPipeline(t *testing.T) {
	// 构造 filter
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()

	sp := NewStraightPipeline("pipe demo",spliter,converter,sum)

	ret ,err := sp.Process("1,2,3")

	if err != nil {
		t.Fatal(err)
	}

	if ret != 6 {
		t.Fatalf("The expected is 6,but the actuak is %d",ret)
	}
}
