# Các khái niệm cơ bản trong lập trình

## 1. Biến (Variables)

Một chương trình luôn luôn phải có các biến, biến là một thành phần quan trọng nhất của chương trình. Muốn sử dụng 1 biến nào đó thì chúng ta cần phải khai báo trước.

> Định nghĩa: Biến là vùng bộ nhớ được cấp phát dùng để lưu trữ giá trị cho một kiểu dữ liệu nào đó. Tại mỗi thời điểm biến số có 1 giá trị. Một biến gồm 3 thành phần:

- Tên biến
- Kiểu dữ liệu
- Giá trị

### 1.1 Cú pháp khai báo
> Có 2 cách khai báo biến:
- Cách 1:
```
var tenBien kieuGiaTri
```

Ví dụ ta có biến: 

```go
var address string 
```
Trong Ví dụ trên ta có:
* Tên biến: address
* Kiểu dữ liệu: string (chuỗi)
* giá trị: ""

Cách 2: Khai báo biến có khởi tạo giá trị ban đầu.

Nếu muốn khởi tạo giá trị cho biến ta dùng:
```go
var address string = "Dong Do, Hung Ha, Thai Binh"
```

Hiểu nôm na như: address là địa chỉ nhà ông Thoại

Ví dụ: Giải phương trình bậc 2
Solution: 
- Tính delta: `delta = b*b - 4*a*c`
- Nếu delta < 0 => Phương trình vô nghiệm
- Nếu delta = 0 => Phương trình có nghiệm duy nhất
- Nếu delta > 0 => Phương trình có 2 nghiệm phân biệt

Phân tích bài toán:
- Đầu vào (Input): Hệ số: a, b, c
- Đầu ra: Đưa ra các nghiệm của phương trình

Vậy những điều cần làm:
- Tạo ra các biến: a, b, c, delta
- Biến kết quả có thể có.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1
	var b float64 = 2
	var c float64 = -3
	var delta float64

	delta = b*b - 4*a*c

	if delta < 0 {
		fmt.Println("Phương trình vô nghiệm")
	} else if delta == 0 {
		x := -b / (2 * a)
		fmt.Println("Phương trình có nghiệm duy nhất: ", x)
	} else {
		x1 := (-b + math.Sqrt(delta)) / 2 * a
		x2 := (-b - math.Sqrt(delta)) / 2 * a
		fmt.Println("Phương trình có 2 nghiệm phân biệt x1 =", x1, " và x2 =", x2)
	}
}

```