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

# Các loại kiểu dữ liệu
Trong chương trình phần mềm cần phải định nghĩa trước những kiểu dữ liệu và lưu trữ nó ở dạng thích hợp nhất để chương trình chạy nhanh và tốn ít bộ nhớ.
Ví du: Để lưu tuổi của 1 người thì ta chỉ cần dùng kiểu `unit8` là đủ. Lý do, tuổi thọ của con người <= 120 tuổi và không thể là số âm

1. Kiểu dữ liệu 
> Kiểu dữ liệu boolean (Đúng sai)
- bool
> Kiểu dữ liệu số
* int8, int16, int32, int64, int
* uint8, uint16, uint32, uint64, uint
* float32, float64
* complex64, complex128
* byte


> Kiểu dữ liệu chuỗi
* rune
* string

# Số nguyên có dấu
`int8`: B
`size`: 8 bits
`range`: -128 to 127

`int16`: biểu diễn số nguyên có dấu 16 bit
size: 16 bits
range: -32768 to 32767

`int32`: biểu diễn số nguyên có dấu 32 bit
`size`: 32 bits
`range`: -2147483648 to 2147483647

`int64`: biểu diễn số nguyên có dấu 64 bit
`size`: 64 bits
`range`: -9223372036854775808 to 9223372036854775807

# Phép gán
> Để thiết lập giá trị của một biến, người ta thiết lập giá trị cho biến đó thông qua phép gán giá trị của biến.

Ví dụ :
height = 8;
length = 10;
width = 12;

Trước khi thiết lập giá trị cho biến thì biến đó cần phải được khai báo.

Ví dụ :
int height = 8; hoặc
int height;
height = 8;

Ngoài ra một biến có thể được gán giá trị bằng một biểu thức. Ví dụ :
int volume = height * length* width;

Bài tập: 
> Viết chương trình nhập vào tháng của một năm, cho biết số ngày của tháng đó. Nếu tháng nhập vào <1 hoặc >12 thì thông báo "Không tồn tại tháng này".

- Các tháng 1, 3, 5, 7, 8, 10, 12 có 31 ngày
- Các tháng 4, 6, 9, 11 có 30 ngày
- Nếu là tháng 2 thì yêu cầu nhập thêm năm, nếu là năm nhuận thì tháng 2 có 29 ngày, còn lại là 28 ngày. Năm nhuận là năm chia hết cho 4.