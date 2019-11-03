# Ch6 method
* golang에서는 객체를 단순히 메소드를 가진 값이나 변수로 정의함
* 메소드는 특정 타입과 관련된 함수로 정의한다. 

* 객체지향 프로그램은 메소드를 통해 데이터 구조의 특성과 동작을 표현하므로 , 사용자는 객체의 구현에 직접 접근할 필요가 없다. 

## 6.1 메소드 선언
* 메소드는 일반 함수 선언을 변형해 함수명 앞에 부가적인 파라미터(reciever)를 추가한 형태로 선언한다. 

[geometry](./geometry/geometry.go)
```golang
package geometry

import (
	"math"
)

//Point struct for method
type Point struct{ X, Y float64 }

//Distance legacy
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Distance method for Point type
// p is reciever
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```
* p를 메소드의 수신자(reciever)라 부른다. 
  * 초기 객체지향 언어에서 메소드 호출을 객체에 메시지를 전송한다 라는 전통이 있다나 뭐래나...

```golang
//...
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

}
```
* ```p.Distance``` 표현식은 Point 타입의 수신자 p에 대응하는 Distance 메소드를 선택하므로 셀렉터라(selector)한다. 
* 이 셀렉터는 구조체 타입의 필드를 선택할때도 사용한다. 
  * 이것은 메소드와 필드가 같은 네임스페이스(geometry.Point)에 있으므로 
* 각 타입에는 메소드가 속한 자신만의 네임스페이스가 있으므로 다른 타입에서는 같은 이름인 Distance를 메소드명으로 사용할수 있다. 

```golang {.line-numbers}
//Path is path to connect between points
type Path []Point

//Distance is method for Path type
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			// point type Distance를 호출한다.
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
```
* sum라인은 point type Distance 메소드를 호출한다.

```golang
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // 12

```
* ```Println``문에서 컴파일러는 메소드명과 수신자(perim []Path) 타입으로 어떤 함수를 호출할지를 결정한다. 
* 그래서 위의 코드에서 path[i-1]는 Point 타입이므로 Point.Distance가 호출되고, 
* perim의 경우 Path 타입이므로 Path.Distance가 호출된다. 
* 특정 타입의 모든 메소드명을 유일해야 한다, 하지만 서로 다른 타입에서는 같은 메소드명을 사용할수 있다(메소드의 장점)

## 6.2 포인터 수신자가 있는 메소드
* 함수를 호출하면 각 인수(변수) 값의 복사본이 생성되어 전달됨(golang 변수는 call-by-value가 기본)
* 아래의 목적이 필요할 경우 포인터를 이용해 변수의 주소를 전달해야 함
  * 함수에서 인수(변수) 값을 변경해야 할때
  * 인수(변수)가 커서 가급적 복사하고 싶지 않을 때

```golang
//ScaleBy 주어진 값만큼 좌표 수정
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

//LocalScaleBy not point reciever
func (p Point) LocalScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
```

* 이 메소드의 이름은 ```(*Point).Scaleby```: 포인터 수신자(pointer receiver)값은 Point 타입 구조체를 가리키고(*Point), 그것의 ScaleBy 메소드를 호출한다. 
* 관행은 Point 구조체의 메소드 중 포인터 수신자가 있다면, 모든 메소드에도 포인터 수신자로 해야 한다...
* 수신자(receiver)선언은 명명된 타입(Point)과 이 타입의 포인터(*Point)만 사용할수 있다. 

* (*Point).ScaleBy 메소드는 *Point 타입 수신자(포인터 수신자)를 통해서 호출할수 있다. 
* 하지만 메소드의 수신자 인수(변수)가(아래의 경우 r)가 *Point지만, 선언된 수신자가 *Point 타입이 아닐경우(Point), 컴파일러가 묵시적으로 수신자의 주소값을 가지고 참조한다. 

```golang
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
```

* 헷갈리기 쉬우니 한번 더 정리
  * golang 컴파일러가 수신자 인수(변수)의 값, 주소값 종류별로 알아서 참조한다. 하지만 모든 경우는 아니므로 주의해야 함

1. 수신자 인수(변수)의 타입과 수신자의 파라미터 타입이 갈을경우,
```golang
	geometry.Point{1, 2}.Distance(q) // Point type 메소드 Distance는 값으로 수신자에게 파라미터를 보냄
	pptr.ScaleBy(2)                  // *Point type 메소드 ScaleBy는 포인터로 수신자에게 파라메터를 보냄 *pptr
```

2. 수신자 인수(변수)가 일반타입(T)이고, 파라미터가 포인터 수신자 타입일때(*T) 컴파일러가 묵시적으로 변수의 주소를 취함
```golang
	p.ScaleBy(2) // 묵시적으로 (&p)을 취함 (&p).ScaleBy(2)
```

3. 수신자 인수(변수)가 포인터 타입(*T)이고, 파라미터가 일반타입(T)일때 컴파일러가 묵시적으로 변수의 값을 역참조해 값을 읽음
```golang
    pptr.Distance(q) // 묵시적으로 *pptr(pp1.Point{1,2})을 취함.
```

### 6.2.1 nil은 유효한 수신자 값
* 일부 메소드도, 맵과 슬라이스와 같이 nil이 유의미한 제로값인 경우 인수로 nil 포인터를 사용한다. 
