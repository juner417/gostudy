package main

import (
	"fmt"
	"net/url"

	"github.com/juner417/gostudy/ch6/coloredpoint"
	customurl "github.com/juner417/gostudy/ch6/customurl"
	"github.com/juner417/gostudy/ch6/geometry"
)

//IntList 는 정수의 링크드 리스트
// nil *IntList는 빈 목록을 표시함 <-- 요렇게 주석으로 적어줘라
type IntList struct {
	Value int
	Tail  *IntList // 포인트 주소
}

//Sum 재귀함수
func (list *IntList) Sum() int {
	//
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	//########### 6.1 ###########
	fmt.Println("## 6.1 result ##")
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	//"5" call Distance func
	//패키지 수준 함수 geometry.Distance func를 선언(p,q를 인자로 받음)
	fmt.Println(geometry.Distance(p, q))
	//fmt.Printf("type: %T\n %#v\n", geometry.Distance, geometry.Distance)

	//"5" call Distance method for Point struct
	//Point type의 메소드를 선언 Point.Distanc(p객체의 slector로 호출)
	fmt.Println(p.Distance(q))
	//fmt.Printf("type: %T\n %#v\n", q.Distance, q.Distance)

	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // 12

	//########### 6.2 ###########
	fmt.Println("## 6.2 result ##")
	r := &geometry.Point{1, 2} //r은 Point 구조체의 포인터변수(*Point 수신자)
	//r이 아래같이 포인터변수가 아닐때도 컴파일러는 r.ScaleBy(3)을 (&r).ScaleBy로 변경한다.
	//r := geometry.Point{1, 2}
	//r.ScaleBy(3) //(&r).ScaleBy(3) 과 동일

	fmt.Println(*r) //{1,2}
	r.ScaleBy(3)
	//(&r).ScaleBy(3) // Error 이미 r이 포인터기 때문에 **r을 참조함 이렇게 하려면 위 주석처럼 선언해야 함

	fmt.Println(r)    //&{3,6}
	r.LocalScaleBy(2) //call by value로 변경 안됨
	fmt.Println(r)    //&{3,6}
	// r  value, addr: &geometry.Point{X:3, Y:6}, 0xc0000b0030
	fmt.Printf("r values, addr: %v,%+v,%#v,%p\n", *r, *r, *r, r)

	//case1
	pp1 := geometry.Point{1, 2}
	pptr := &pp1
	pptr.ScaleBy(2)
	fmt.Println(pp1) //{2, 4}

	//case2
	pp2 := geometry.Point{1, 2}
	(&pp2).ScaleBy(2)
	fmt.Println(pp2) // {2, 4}

	// 위의 case1,2 처럼 하지 않아도 되는 것은
	// 수신자 p가 Point 타입의 변수지만, 메소드에 *Point 타입의 변수가 필요할때
	// 컴파일러가 묵시적으로 &pp1, &pp2로 변경함

	pp3 := geometry.Point{1, 2}
	pp3.ScaleBy(2) // 이게 된다고.. 컴파일러가 해주니까... 그말을 풀어쓰느라 어렵게 쓴것임 쫄지마셈
	fmt.Println(pp3)

	//geometry.Point{1, 2}.ScaleBy(2) //주소가 할 당되지 않은 Point수신자의 *Point 메소드는 임시 값의 주소를 얻을수 없으므로 호출이 안됨
	//./main.go:49:22: cannot call pointer method on geometry.Point literal // ScaleBy는 포인터수신자를 인자로같는 메소드라 주소가 없어서 호출이 안됨
	//./main.go:49:22: cannot take the address of geometry.Point literal // 지금의 선언은 리터럴 타입은 주소를 가질수 없다. 받는 변수가 없어서

	// 헷갈리기 쉬우니 한번더 정리
	// #1. 수신자 인수(변수)의 타입과 수신자의 파라미터 타입이 갈을경우,
	geometry.Point{1, 2}.Distance(q) // Point type 메소드 Distance는 값으로 수신자에게 파라미터를 보냄
	pptr.ScaleBy(2)                  // *Point type 메소드 ScaleBy는 포인터로 수신자에게 파라메터를 보냄 *pptr

	// #2. 수신자 인수(변수)가 일반타입(T)이고, 파라미터가 포인터 수신자 타입일때(*T) 컴파일러가 묵시적으로 변수의 주소를 취함
	p.ScaleBy(2) // 묵시적으로 (&p)을 취함 (&p).ScaleBy(2)

	// #3. 수신자 인수(변수)가 포인터 타입(*T)이고, 파라미터가 일반타입(T)일때 컴파일러가 묵시적으로 변수의 값을 역참조해 값을 읽음
	pptr.Distance(q) // 묵시적으로 *pptr(pp1.Point{1,2})을 취함.

	//########### 6.2.1 ###########
	fmt.Println("## 6.2.1 result ##")
	//urlvalues
	//net/url
	// value는 string slice타입의 map
	//type Values map[string][]string
	m := url.Values{"lang": {"en"}} //직접생성
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println("m lang key has ->", m.Get("lang"))  // "en"
	fmt.Println("m lang q has ->", m.Get("q"))       // "" 해당 키의 값은 len 0
	fmt.Println("m lang item has ->", m.Get("item")) // "1" vs[0]
	fmt.Println("m[item] ->", m["item"])             // "[1, 2]" 직접 맵 값 접근
	fmt.Printf("m map info: %v %#v %p\n", m, m, m)   // map[item:[1 2] lang:[en]] url.Values{"item":[]string{"1", "2"}, "lang":[]string{"en"}} 0xc00007a030
	m = nil                                          // nil로 초기화
	fmt.Println(m.Get("item"))                       // "" nil map의 len 0 이다.
	fmt.Printf("nil map info: %v %#v %p\n", m, m, m) // map[] url.Values(nil) 0x0
	//m.Add("item", "3") // nil 맵(map[])을 변경하려고 해서 panic. nil 맵은 주소공간만 할당된 것으로 주소값이 없다.
	//아래처럼 해줘야함
	m = url.Values{"item": {"3"}}                        // 이렇게하면 nil이 아닌 새로운 객체가 들어감
	fmt.Println("new m item key has -> ", m.Get("item")) // "3"

	// customurl에 Add 메소드에 nil 일때 처리 하려고 했는데 그럴려면 values를 포인터로 받아야 함...
	var t customurl.Values                                       // 이렇게하면 nil 값이 제로값으로 들어감
	fmt.Printf("t customurl.Values info : %v %#v %p\n", t, t, t) // nil 값
	t = make(customurl.Values)                                   // make 함수로 nil map 초기화 리터럴 방식도 가능
	t.Add("lang", "en")                                          // nil 처리는 됐는데 values가 일반 변수라 call-by-value로 들어가서 지역변수로 처리
	t.Add("item", "1")                                           // nil 처리는 됐는데 values가 일반 변수라 call-by-value로 들어가서 지역변수로 처리
	//t = customurl.Values{"lang": {"ko", "jp"}} // 리터럴 방식
	fmt.Println("key lang value: ", t.Get("lang"))
	fmt.Println("key item value: ", t.Get("item"))

	//########### 6.3 ###########
	fmt.Println("## 6.3 result ##")

	var cp coloredpoint.ColoredPoint
	cp.X = 1

}
