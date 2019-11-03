# Ch6 method
* golang에서는 객체를 단순히 메소드를 가진 값이나 변수로 정의함
* 메소드는 특정 타입과 관련된 함수로 정의한다. 

* 객체지향 프로그램은 메소드를 통해 데이터 구조의 특성과 동작을 표현하므로 , 사용자는 객체의 구현에 직접 접근할 필요가 없다. 

## 6.1 메소드 선언
* 메소드는 일반 함수 선언을 변형해 함수명 앞에 부가적인 파라미터(reciever, reciever parameter)를 추가한 형태로 선언한다. 

[geometry](./geometry/geometry.go)
```golang
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

```golang
//...
func main() {
	p := Point{1, 2}
	q := Point{4, 6}

    //"5" call Distance func
    //패키지 수준 함수 geometry.Distance 를 선언
	fmt.Println(Distance(p, q)) 
	//"5" call Distance method for Point struct
	//Point type의 메소드를 선언 Point.Distanc
	fmt.Println(p.Distance(q))

}
```
* ```p.Distance``` 표현식은 Point 타입의 수신자 p에 대응하는 Distance 메소드를 선택하므로 셀렉터라(selector)한다. 
* 이 셀렉터는 구조체 타입의 필드를 선택할때도 사용한다. 
  * 이것은 메소드와 필드가 같은 네임스페이스(geometry.Point)에 있으므로 
* 각 타입에는 메소드가 속한 자신만의 네임스페이스가 있으므로 다른 타입에서는 같은 이름인 Distance를 메소드명으로 사용할수 있다. 

```golang {.line-numbers}
//Path is path to connect between points
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			// point type Distance 메소드를 호출한다.
            // (*path[i-1]).Distance(*path[i])
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
```
* point type Distance 메소드를 호출한다.

```golang
	perim := Path{
		Point{5, 1},
		Point{1, 1},
		Point{5, 4},
		Point{1, 1},
	}
    // "12", path type의 Distance 메소드 호출
	fmt.Println(perim.Distance()) 
```
* 
