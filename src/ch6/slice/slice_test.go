package slice

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2,1)
	t.Log(s2[0], s2[1], s2[2],s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan","Feb","Mar","Apr","May","Jun","Jul",
		"Aug","Sep","Oct","Nov","Dec"}
	q2 := year[3:6]
	t.Log(q2,len(q2),cap(q2)) //[Apr May Jun] 3 9

	sumer := year[5:8]
	t.Log(sumer,len(sumer),cap(sumer)) //[Jun Jul Aug] 3 7

	sumer[0] = "Unknow"
	t.Log(q2) //[Apr May Unknow]
	t.Log(year) //[Jan Feb Mar Apr May Unknow Jul Aug Sep Oct Nov Dec]

}
