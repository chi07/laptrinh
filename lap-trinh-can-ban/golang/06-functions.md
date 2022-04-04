## Hàm là gì?

Hàm là một đoạn chương trình (block of code) để xử lý một nhiệm vụ cụ thể. 1 hàm gồm có tên (required), các tham số (optional), kiểu dữ liệu trả về (optional)

## Khai báo hàm

```go
func functionName (params Type) returnTypes {

}
```
Các bạn chú ý mình để tham số là `params` và `returnTypes` là thể hiện có thể 1 hoặc nhiều tham số hoặc nhiều kiểu giá trị trả về.

Ví dụ

**Kiểu khai báo đầy đủ**

```go
package main

import (
	"fmt"
)

func main() {
	var firstName, lastName = "Shinichi", "Pham"
	fullName := displayName(firstName, lastName)
	
	fmt.Printf("Hello, %s", fullName)
}

func displayName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

}
```
#### [Run in playground](https://play.golang.org/p/ltF4sqLT7vg)
```bash
Hello, Shinichi Pham
```  

Ở ví dụ trên mình khai báo 1 hàm (function) có tên là `displayUserInfo` gồm 2 tham số là `firstName` và `lastName` kiểu `string` và kiểu trả về của hàm là `string`.

Hàm trả về nhiều giá trị
```go
package main

import (
	"fmt"
)

func main() {
	a, b := 10, 2
	sum, sub, multiple, dive := Calculate(a, b)
	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(multiple)
	fmt.Println(dive)
}

func Calculate(a, b int)(int, int, int, int) {
	sum := a + b
	sub := a - b
	multiple := a * b
	dive := a/b // Giả sử b != 0
	return sum, sub, multiple, dive
}
```
#### [Run in playground](https://play.golang.org/p/nhqyAj9r0i4)

Chạy chương trình cho bạn kết quả.

```bash
12
8
20
5

```

**Kiểu khai báo không có tham số**
```go
package main

import (
	"fmt"
)

func main() {
	fullName := getName()
	
	fmt.Printf("Hello, %s", fullName)
}

func getName() string {
	return "Shinichi Pham"
}

```
#### [Run in playground](https://play.golang.org/p/qBRNMKrDTQB)

**Kiểu khai báo không có dữ liệu trả về**
```go
package main

import (
	"fmt"
)

func main() {
	firstName, lastName := "Shinichi", "Pham"
	printName(firstName, lastName)
}

func printName(firstName, lastName string) {
	fullName := firstName + " " + lastName
	fmt.Printf("Hello, %s", fullName)
}

```
#### [Run in playground](https://play.golang.org/p/pETd5qL3R9S)



## Một số lưu ý khi đặt tên cho hàm (function)
Các bạn có thể đặt tên tuỳ ý theo ý các bạn muốn, nhưng để code rõ ràng, clean các bạn nên tuân theo một số chuẩn mà các lập trình viên thường làm

> Đặt tên của hàm bằng tên tiếng anh, gắn gọn, mang tính gợi nhớ về chức năng của hàm
> Tên hàm nên là động từ. Ví dụ các bạn thường thấy như: `Create`, `Update`, `Save`, `Delete`, `placeOrder` đó đều là những động từ
> Các tham số, biến nên là các danh từ. Ví dụ: `fullName`, `firstName` ...
> tên hàm và tên biến nên đặt theo quy tắc camel (chữ cái đầu tiên của 1 từ viết thường, các từ sau viết hoa chữ cái đầu tiên. Ví dụ: `productName`, `productCode`...

## Blank Identifier
Mình ko biết dịch thế nào cho nó sát nghĩa và nó thuận 1 chút nên mình xin để nguyên như vậy. Bản chất Blank Identifier là `_`, nó cho phép bạn lược bỏ đi những giá trị mà bạn không cần sử dụng. Chúng ta hãy cùng xem lại 1 trong các ví dụ trên nhé
```go
sum, sub, multiple, dive := Calculate(a, b)
```

Theo như ví dụ bên trên chúng ta được 4 giá trị trả về là `sum` (tổng), `sub` (hiệu), `multiple` (tích), `dive` (thương). Nhưng giả sử trong bài toán chúng ta không cần đến 2 giá trị là tích và thương thì chúng ta sẽ dùng `Blank Identifier` để bỏ qua 2 giá trị này. Hãy thử xe nó như thế nào nhé

```go
package main

import (
	"fmt"
)

func main() {
	a, b := 10, 2
	sum, sub, _, _ := Calculate(a, b)
	fmt.Println(sum)
	fmt.Println(sub)
}

func Calculate(a, b int)(int, int, int, int) {
	sum := a + b
	sub := a - b
	multiple := a * b
	dive := a/b // Giả sử b != 0
	return sum, sub, multiple, dive
}
```
#### [Run in playground](https://play.golang.org/p/t4UNuWGXOYm)

Trong ví trên ta ko cần đến 2 giá trị là `multiple` và `dive`. Do đó mình dùng `_` để lược bỏ đi. Kết quả cho ra
```bash
12
8
```

Bài học hôm nay khá đơn giản và nhẹ nhàng. Bài học tiếp theo chúng ta sẽ học là về lệnh Packages.

#### Bài tiếp theo: [Packages](/bai-viet/packages.html)