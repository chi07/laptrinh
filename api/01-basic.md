# API
- Nếu coi ngôi nhà là 1 ví dụ thì ta có thể coi 
- Frontend là các thành phần có thể nhìn thấy được của ngôi nhà (hình dáng (nhà ống, nhà tầng, nhà thái), cách phân chia các tầng, các phòng gọi là layout).
- Backend (API) có thể coi là phần ko nhìn thấy được như săt thép được ẩn trong các khối bê tông.
- 
- Nói 1 cách khác, trong các ứng dụng web, ứng dụng mobile. Frontend là các màn hình được hiển thị cho người dùng có thể nhìn thấy được
- Backend(API) chính là một phần tiếp nhận các dữ liệu từ Web hoặc mobile gửi lên, sau đó thực hiện các nghiêpj vụ cần thiết nhưng tương tác với Cơ sở dữ liệu và trả về kết quả cho Frontend hiển thị ra cho người dùng nhinf thấy được.

- 1 API có các thành phần sau
```
- Request: là các thông tin cần thiết để thực hiện 1 nghiệp vụ
- Host: Địa chỉ của server
- Method: GET, POST, PUT, DELETE
- URL: đường dẫn đến API
- Headers: định dạng dữ liệu trả về
- Body: dữ liệu gửi lên
- Response: dữ liệu trả về
```
Hãy tưởng tượng 1 API như sau
Khi cần lấy thông tin của một bài viết có ID là 1 ta có thể tạo 1 API như sau

http://24h.com.vn/api/articles/1
```
- Request: yêu cầu lấy thông tin của bài viết có mã là 1
- Host: 24h.com.vn
- Method: GET, POST, PUT, DELETE
- URL: api/user/1
- Headers: json
- Body: {}
- Response: json
```

Nói 1 cách nôm na ở đây là: Tôi có yêu cầu thông tin số nhân viên của phòng kế toán ở VMO
```
Host: 18 tôn thất thuyết
Method: GET
URL: accounting/employees
Headers: json
Body: {}
Response: json
```


