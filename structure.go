package goProject01

import (
	"math/rand"
)

// Example 人脸实例
type Example struct {
	picture string
	name string
	id string
	department string
}

// ExampleSet 人脸实例集合
type ExampleSet struct{
	examples []Example
	indexes []int
	length int
}

// Maintain 人脸实例集合：维护
func (es ExampleSet) Maintain(){

}

// AddExample 人脸实例集合：添加实例
func (es ExampleSet) AddExample (example Example)  {
	es.examples[es.length] = example
	es.indexes[es.length] = es.length
	es.length++
}

// GetExample 人脸实例集合：取得指定实例
func (es ExampleSet) GetExample(index int) (e Example,info int){
	//异常：人脸实例集合为空
	if es.length == 0 {
		return Example{},1
	}
	//异常：下标越界
	if index >= es.length || index < 0{
		return Example{},2
	}
	return es.examples[index],0
}

// GetRandomExample 人脸实例集合：随机取得一个实例
func (es ExampleSet) GetRandomExample() (e Example,info int){
	//异常：人脸实例集合为空
	if es.length == 0 {
		return Example{},1
	}
	e,_ = es.GetExample(rand.Intn(es.length))
	info = 0
	return
}