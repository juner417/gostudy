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
* ```Println```문에서 컴파일러는 메소드명과 수신자(perim []Path) 타입으로 어떤 함수를 호출할지를 결정한다. 
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

```golang
type IntList struct {
    Value int
    Tail  *IntList // 포인트 주소
}

//Sum 재귀함수
func (list *IntList) Sum() int {
    // 포인트 주소 list가 nil일 경우 처리, 빈 list
    if list == nil {
        return 0
    }
    return list.Value + list.Tail.Sum()
}
```

* net/url 패키지의 Values 타입의 정의
```golang
package url

// Values maps a string key to a list of values.
type Values map[string][]string

// Get returns the first value associated with the given key,
// or "" if there are none.
func (v Values) Get(key string) string {
    if vs := v[key]; len(vs) > 0 {
        return vs[0]
    }
    return ""
}

// Add adds the value to key.
// It appends to any existing values associated with key.
func (v Values) Add(key, value string) {
    v[key] = append(v[key], value)
}
```

```golang
    //urlvalues
    //net/url
    // value는 string slice타입의 map
    //type Values map[string][]string
    m := url.Values{"lang": {"en"}} //리터럴 방식으로 직접생성
    //이때는 m이 nil이 아니어서 append가능
    m.Add("item", "1")
    m.Add("item", "2") 

    fmt.Println(m.Get("lang"))   // "en"
    fmt.Println(m.Get("q"))      // "" 해당 키의 값은 len 0
    fmt.Println(m.Get("item"))   // "1" vs[0]
    fmt.Println(m["item"])       // "[1, 2]" 직접 맵 값 접근
    fmt.Printf("%v %#v %p\n", m, m, m) // map[item:[1 2] lang:[en]] url.Values{"item":[]string{"1", "2"}, "lang":[]string{"en"}} 0xc00007a030
    m = nil                      // nil로 초기화
    fmt.Println(m.Get("item"))   // "" nil map의 len 0 이다. 
    fmt.Printf("%v %#v %p\n", m, m, m) // map[] url.Values(nil) 0x0
    //m.Add("item", "3") // nil 맵(map[])을 변경하려고 해서 panic. nil 맵은 주소공간만 할당된 것으로 주소값이 없다.

    //아래처럼 해줘야함
    m = url.Values{"item": {"3"}} // 이렇게하면 nil이 아닌 새로운 객체가 들어감(리터럴을 이용한 초기화)
    //아래처럼 make 함수를 사용해도 됨 
    //m = make(url.Values))
    //m.Add("item", "3")
    fmt.Println(m.Get("item")) // "3"
```
* ```m = nil```에서 처럼 nil map으로 초기화 되면, 다시 선언해 주지 않으면 메소드 호출이 안됨, 
* 왜냐하면 url.Values가 map type인데 이게 nil map이 됐음(map type zero value)
* 그래서 실제로 메모리가 할당된 hashmap의 주소가 없기 때문에 메소드도 사용할 수 없음(Get이 왜 됐냐면 nil map의 len은 0이다.)
* 리터럴 혹은 make 함수로 새로운 map을 만들고 그것의 주소를 할당해 줘야함


## 6.3 내장 구조체를 통한 타입조합
[coloredpoint](./coloredpoint/coloredpoint.go)
```golang
type Point struct{ X, Y float64 }

//ColoredPoint ...
type ColoredPoint struct {
    Point //anonymous field, 익명필드
    Color color.RGBA
}
```
* 위처럼 Point 구조체를 ColoredPoint에 내장하여 Point의 필드를 사용가능하게 했다(내장 구조체)

```golang
    var cp coloredpoint.ColoredPoint // coloredpoint.ColoredPoint 선언
    cp.X = 1
    fmt.Println(cp.Point.X) // "1" 원래는 이렇게 접근해야 하지만
    cp.Point.Y = 2
    fmt.Println(cp.Y) // "2" 해당 구조체에 필드 타입이 annoymous field를 사용했다면 제거 가능함
    // 단축문법으로 Point의 모든 필드와 추가필드를 갖는 ColoredPoint를 정의할수 있다.
    // ref: http://golangtutorials.blogspot.com/2011/06/anonymous-fields-in-structs-like-object.html
```
* 원래는 ```cp.Point.X```로 지정해야 하나. 내장구조체를 anonymous field(익명필드)로 지정하면, 
* 단축문법으로 Point 언급없이 선택할수 있음

* 메소드도 동일한 방식으로 적용된다. 
```golang
    red := color.RGBA{255, 0, 0, 255}
    blue := color.RGBA{0, 0, 255, 255}
    var cp1 = coloredpoint.ColoredPoint{coloredpoint.Point{1, 1}, red}
    var cp2 = coloredpoint.ColoredPoint{coloredpoint.Point{5, 4}, blue}
    fmt.Println(cp1.Distance(cp2.Point)) // "5"
    cp1.ScaleBy(2)
    cp2.ScaleBy(2)
    fmt.Println(cp1.Point.Distance(cp2.Point)) // "10" 이것도 Point 제거 가능, 하지만 cp2.Point는 제거 불가능. is-a 관계가 아니다.
    //cp1.Distance(cp2) // 불가능
    //이러한 방식으로 많은 메소드가 있는 복잡한 타입을,
    //소수의 메소드를 갖는 여러 필드의 조합으로 만들수 있게 한다.
    //ColoredPoint는 Point가 아니지만, Point를 갖고(has-a)있고,
    //Point에서 승격된 두개의 추가 메소드(Distance, ScaleBy)를 갖고있다.
```
* ```cp1.Point.Distance(cp2.Point)``` 의 Point도 익명필드로 지정되었기 때문에 제거가 가능하다. 
* 이러한 방식으로 많은 메소드가 있는 복잡한 타입을, 소수의 메소드를 갖는 여러 필드의 조합으로 만들수 있게 한다.

* 익명필드의 타입은 명명된 타입의 포인터가 될수 있다. 
* 아래의 경우도 Point struct 필드와 메소드가 승격됨 refered(간접적으로)
```golang
//PColoredPoint ...
type PColoredPoint struct {
    *Point //anonymous field, 이것도 Point struct 필드와 메소드가 승격됨 refered(간접적으로)
    Color  color.RGBA
}
...
    // *Point
    cp3 := coloredpoint.PColoredPoint{&coloredpoint.Point{1, 1}, red}
    cp4 := coloredpoint.PColoredPoint{&coloredpoint.Point{5, 4}, blue}
    fmt.Println(cp3.Distance(*cp4.Point)) //*cp4.Point{5,4} 이 필요함.
    cp4.Point = cp3.Point                 // cp3의 Point주소(&cp3.Point)를 넘겨줌
    cp3.ScaleBy(2)                        // "{2, 2}" cp3의 값을 변경하면, cp4도 변경됨. 윗줄에서 cp4.Point에 cp3.Point 주소값을 주었기 때문에.
    fmt.Println(*cp3.Point, *cp4.Point)   //"{2,2}" "{2,2}"
```

* 구조체의 타입은 두개 이상의 익명필드를 가질수 있다. 
* 아래와 같이 선언된 구조체에서는 아래의 필드와 메소드가 승격된다. 
  * 이 타입의 값은 Point의 모든 메소드와
  * RGBA의 모든 메소드
  * DColoredPoint에 정의된 추가 메소드들을 직접갖게 된다.
```golang
    // 두개 이상의 익명필드를 가질수 있고, 아래와 같이 선언 됐다면
    //DColoredPoint ...
    //type DColoredPoint struct {
    //    Point
    //    color.RGBA
    //}

    dp := coloredpoint.DColoredPoint{coloredpoint.Point{1, 1}, red}
    dp.ScaleBy(2)
```

* ```dp.ScaleBy``` 와 같은 셀렉터를 메소드로 연결할때 컴파일러가 승격된 필드와 메소드를 찾는 순서
  * 직접 ScaleBy로 선언된 메소드
  * 그다음 DColoredPoint에 내장된 필드에서 승격된 메소드
  * 그 다음에 내장된 Point와 RGBA에서 두번 승격된 메소드 순으로 찾는다.
* 만약 같은 단계에 두개 이상의 메소드가 있어서 셀렉터로 선택할 수 없으면 오류


## 6.4 메소드 값과 메소드 표현식
[geometry](./geometry/geometry.go)
* 아래의 코드에서 셀렉터 ```pp4.Distance``` 메소드(Point.Distance)를 특정 수신자값 pp4와 결합하는 함수인 메소드 값을 내보낸다(주소값). 
```golang
    pp4 := geometry.Point{1, 2}
    qq4 := geometry.Point{4, 6}

    distanceFromP := pp4.Distance
    fmt.Printf("method value pp4.Distance -> %p\n", pp4.Distance)        // 메소드 값
    fmt.Printf("distance between pp4 and qq4 %v \n", distanceFromP(qq4)) //"5"
    var origin geometry.Point                                            // origin 선언 zero 값 {0,0}
    fmt.Println(distanceFromP(origin))                                   // "2.23606797749979"

    scaleP := pp4.ScaleBy // 메소드 값
    scaleP(2)
    fmt.Println(pp4) //{2,4} ScaleBy는 수신자 인수가 *Point 타입이다.
    scaleP(3)
    fmt.Println(pp4) //{6, 12}
    scaleP(10)
    fmt.Println(pp4) //{60, 120}
    fmt.Printf("method value distanceFromP -> %p\n", distanceFromP)
```
* 이런 메소드 값은 
  * 패키지 api가 함수 값을 호출하거나, time.AfterFunc(10 * time.second, r.Launch)
  * 해당함수에서 특정 수신자의 메소드를 호출하려고 할때 유용하다.

* **메소드값**은 **메소드 표현식**과 연관된다. 
* 메소드를 호출할 때는 통상의 함수와 달리 특별히 셀렉터 문법 ```x.add``` 을 이용해 수신자를 지정해야 한다.
* 타입 T에서 ```T.f```, ```(*T).f``` 로 작성하는 메소드 표현식은 
* 통상적인 첫 번째 파라미터를 수신자로 받는 함수 값을 산출하므로 일반적인 방법으로 호출할 수 있다. 
```golang
    pp5 := geometry.Point{1, 2}
    qq5 := geometry.Point{4, 6}

    distance := geometry.Point.Distance                                                      // 메소드 표현식 - 수신자(Point)를 지정해야 한다.
    fmt.Printf("distance -> %p, Point.Distance -> %p \n", distance, geometry.Point.Distance) // 주소 같음
    fmt.Println(distance(pp5, qq5))                                                          //"5"
    fmt.Printf("%T\n", distance)                                                             //func(geometry.Point, geometry.Point) float64

    scale := (*geometry.Point).ScaleBy // 메소드 표현식
    scale(&pp5, 2)                     // 수신자가 *Point 타입이라 주소값으로 첫번째 파라미터에 줌
    fmt.Println(pp5)                   // "{2 4}"
    fmt.Printf("%T\n", scale)          //func(*geometry.Point, float64)
```

* **메소드 표현식**은 한타입의 여러 메소드 중 하나를 선택하고
* 선택한 메소드를 여러 수신자에서 호출할 때 도움이 된다.   
[geometry](./geometry/geometry.go#L68)

```golang
type Point struct{ X, Y float64 }
type Path []Point
//...
//Add ...
func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

//Sub ...
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

//TranslateBy ...
func (path Path) TranslateBy(offset Point, add bool) Path {
    var op func(p, q Point) Point // op 변수에 익명 함수 선언
    if add {
        op = Point.Add // 메소드 표현식
    } else {
        op = Point.Sub // 메소드 표현식
    }
    for i := range path {
        // path[i].Add(offset) 또는 path[i].Sub(offset)을 호출한다.
        path[i] = op(path[i], offset)
    }
    // fmt.Printf("%#v, %v \n", op, path) //for debuging
    return path
}
```

## 6.5 비트벡터 타입
[intset](./intset/intset.go)
```golang
    var x, y intset.IntSet
    x.Add(1) //
    x.Add(144)
    x.Add(9)
    fmt.Println(x.String()) // "{1 9 144}"

    y.Add(9)
    y.Add(42)
    fmt.Println(y.String()) // "{9 42}"

    x.UnionWith(&y)
    fmt.Println(x.String())           // "{1 9 42 144}"
    fmt.Println(x.Has(9), x.Has(123)) // "true false"
```

## 6.6 캡슐화
* 객체의 사용자가 객체의 변수나 메소드에 접근할 수 없는 경우 -> 캡슐화되어 있다.(캡슐화 -> 정보은닉)
* golang에서는 패키지 레벨의 캡슐화가 된다.(타입이 아님)
* 대문자로 시작하는 메소드는 정의된 패키지에서 노출된다.
* 대문자가 아니면 노출되지 않는다.(구조체의 필드, 타입의 메소드 모두 동일한 방식으로 적용)
* 패키지의 멤버에 대한 접근 제한은 구조체의 필드나 타입의 메소드에도 적용된다. 
* 객체를 캡슐화 하려면 구조체로 만들어서 필드명을 노출시키지 않아야 한다. 
```golang
    // 아래의 코드는 불가능...
    // intset package 의 구조체의 word 필드는 외부로 노출되지 않았다.(소문자)
    // Add, Has, UniWith 등의 메소드로만 이 필드는 접근,변경 가능하다(Encapsulation)
    zz := intset.IntSet{
        []uint64{1, 2, 3, 4},
    }
    fmt.Println(zz) 
    //./main.go:282:11: implicit assignment of unexported field 'words' in intset.IntSet literal

    // Encapsulation 할때 구조체로 안할경우, 직접 객체의 데이터에 접근할수 있으므로
    // 아래의 코드가 가능... Encapsulation 안됨
    z := intset.IntsetN{1, 2, 3, 4}
    fmt.Printf("%v\n", z) //[1,2,3,4]
    z = append(z, 100)
    fmt.Printf("%v\n", z) //[1 2 3 4 100] 이러면 뭐...
```

### 6.6.1 캡슐화의 장점
1. 사용자가 직접 객체 변수를 수정할 수 없다. 
2. 설계자가 api 호환성을 유지하면서 구현을 자유롭게 진화 시킬수 있다. 
3. 사용자가 객체의 변수를 임의로 변경할 수 없다