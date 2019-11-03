# Ch6 method
* golang에서는 객체를 단순히 메소드를 가진 값이나 변수로 정의함
* 메소드는 특정 타입과 관련된 함수로 정의한다. 

* 객체지향 프로그램은 메소드를 통해 데이터 구조의 특성과 동작을 표현하므로 , 사용자는 객체의 구현에 직접 접근할 필요가 없다. 

## 6.1 메소드 선언
* 메소드는 일반 함수 선언을 변형해 함수명 앞에 부가적인 파라미터(reciever, reciever parameter)를 추가한 형태로 선언한다. 

[geometry](./geometry/geometry.go)
```
package geometry

import "math"

type Point struct{ X, Y float64 }

// legacy
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Distance.. method for Point type
// p is reciever
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

* p를 메소드의 수신자(reciever)라 부른다. 
  * 초기 객체지향 언어에서 메소드 호출을 객체에 메시지를 전송한다 라는 전통이 있다나 뭐래나...

```
//...
func main() {
	p := Point{1, 2}
	q := Point{4, 6}

    //"5" call Distance func
    //패키지 수준 함수 geometry.Distance 를 선언
	fmt.Println(Distance(p, q)) 
	fmt.Println(p.Distance(q))  
    //"5" call Distance method for Point struct
}
```

