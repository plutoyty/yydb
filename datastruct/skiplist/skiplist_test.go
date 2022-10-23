package skiplist

import (
	"testing"
)

//测试获取排名
func TestSkipList(t *testing.T) {
	s := makeSkiplist()
	//print(s.level)
	s.insert("12", 12)
	print(s.getRank("12", 12))
	print(s.getByRank(1).Member)
}

//测试删除
func TestSkiplist_RemoveRangeByRank(t *testing.T) {
	s := makeSkiplist()
	s.insert("lyy", 38)
	s.insert("lll", 40)
	//fmt.Printf("%s", s)
	println(s.getRank("lll", 40))
	println(s.getByRank(2).Member)
	s.remove("lyy", 38)

	println(s.getRank("lll", 40))
	print(s.getByRank(1).Member)
}
