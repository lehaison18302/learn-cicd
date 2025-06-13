# Dự án Quản lý Người dùng với Go, MariaDB, Redis, JWT

## Mô tả

Đây là dự án mẫu xây dựng RESTful API quản lý người dùng sử dụng Golang, MariaDB, Redis và xác thực bằng JWT.  
Dự án hỗ trợ các chức năng: đăng nhập, thêm/sửa/xóa/lấy danh sách người dùng, ghi log thao tác qua Redis và worker ghi log vào database.

---

## Kiến trúc

- **Backend:** Golang (Gin framework)
- **Database:** MariaDB
- **Cache/Queue:** Redis (ghi log thao tác)
- **Xác thực:** JWT
- **Triển khai:** Docker Compose

---

## Cài đặt & Chạy dự án

### 1. Yêu cầu

- [Docker Desktop](https://www.docker.com/products/docker-desktop/) (Windows/Mac/Linux)
- Git

### 2. Clone dự án

```bash
git clone https://github.com/<your-username>/<repo-name>.git
cd <tên-thư-mục-dự-án>
```

### 3. Chạy bằng Docker Compose

```bash
cd my_project
docker compose up --build
```

- MariaDB sẽ chạy ở cổng `3307` trên máy bạn.
- Redis chạy ở cổng `6379`.
- Ứng dụng Go chạy ở cổng `8080`.

### 4. Truy cập ứng dụng
- Chạy project: 
```bash
cd my_project
go mod tidy 
go run ./cmd/api
```

- API: `http://localhost:8080/users`, `http://localhost:8080/login`, ...
- Giao diện web: `http://localhost:8080/web/users.html`

---

## Cấu hình biến môi trường (nếu cần)

Bạn có thể chỉnh sửa file `.env` hoặc phần `environment` trong `docker-compose.yml` để thay đổi thông tin kết nối DB, Redis, JWT secret...

---

## Các API chính

- `POST /login` — Đăng nhập, trả về JWT
- `GET /users` — Lấy danh sách người dùng (có phân trang)
- `POST /users/create` — Thêm người dùng mới
- `PUT /users/:id` — Cập nhật thông tin người dùng
- `DELETE /users/:id` — Xóa người dùng

**Lưu ý:**  
Các API (trừ `/login`) yêu cầu gửi JWT qua header:  
```
Authorization: Bearer <token>
```

---

## Ghi log thao tác

- Mỗi thao tác thêm/sửa/xóa sẽ ghi log vào Redis.
- Worker tự động lấy log từ Redis và ghi vào bảng `user_action_logs` trong MariaDB.

---

## Chạy kiểm thử (CI)



---
