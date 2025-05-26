# 🚀 VHU_BMWA: Triển khai và đánh giá bảo mật RESTful API

## 📚 Giới thiệu
Chào mừng đến với bài tập siêu to khổng lồ của môn học "Bảo mật Web & Ứng dụng"! 
Đây không phải là một API bình thường - đây là một API với full-stack security logging như NASA mới code bằng AI.

### 🎯 Mục tiêu
- Xây dựng một RESTful API không phải cái API "hello world"
- Triển khai hệ thống logging ngon-bổ-rẻ như buffet
- Bảo mật chắc chắn như két sắt ngân hàng
- Làm cho giảng viên phải trầm trồ: "Ồ, hay đấy!, sài AI đúng không?" 😎

## 🛠 Công nghệ sử dụng
- **Backend**: Go (Gin)
- **Frontend**: Vue.js (React quá mainstream 😏)
- **Database**: MySQL
- **Container**: Docker
- **Logging**: Logrus

## 🏗 Kiến trúc hệ thống
```
┌─────────────┐      ┌──────────┐      ┌─────────┐
│   Frontend  │ <──> │  Nginx   │ <──> │ Backend │
└─────────────┘      └──────────┘      └─────────┘
                                           │
                                      ┌─────────┐
                                      │ MySQL   │
                                      └─────────┘
```

## 📝 Tính năng logging siêu cấp
### 1. Access Log 📊
- Request method, path, status code
- Thời gian xử lý (để biết API nào chậm như rùa 🐢)
- IP address (để biết ai đang truy cập)
- User agent (để biết ai còn xài IE 🤦‍♂️)

### 2. Security Log 🔒
- Các sự kiện xác thực (login/register)
- Các nỗ lực tấn công SQL injection
- Các thử nghiệm XSS
- Rate limiting violations (spam là cook)

### 3. Error Log ⚠️
- Stack trace đầy đủ (để debug không phải đau đầu)
- Component gây lỗi (để biết đứa nào code bug)
- User context (để biết ai là nạn nhân của bug)

## 🚀 Hướng dẫn cài đặt

### Yêu cầu hệ thống
- Docker
- Docker Compose
- Não 🧠 (quan trọng nhất)

### Các bước cài đặt
1. Clone repository:
```bash
git clone https://github.com/phuongaz/VHU_BMWA.git
cd VHU_BMWA
```

2. Khởi động project:
```bash
docker-compose up --build -d
```
(Đi pha cà phê ☕, chờ Docker build)

3. Kiểm tra các container:
```bash
docker-compose ps
```

## 🔍 API Endpoints

### 👤 Authentication
- `POST /api/auth/register` - Đăng ký
- `POST /api/auth/login` - Đăng nhập

### 🛍 Products
- `GET /api/products` - Lấy danh sách sản phẩm
- `POST /api/products` - Thêm sản phẩm `required Role admin`
- `PUT /api/products/:id` - Cập nhật sản phẩm `required Role admin`
- `DELETE /api/products/:id` - Xóa sản phẩm `required Role admin`

## 📊 Monitoring

### Log Files
- `access.log` - Ghi lại mọi request 
- `security.log` - Ghi lại các vấn đề bảo mật
- `error.log` - Ghi lại các lỗi

### Rate Limiting
- 10 request/giây
- Áp dụng cho tất cả các endpoint

## 🔒 Bảo mật
- JWT Authentication 
- Password hashing với bcrypt
- Rate limiting
- Input validation 

## 🎯 Mục tiêu đạt được
- [x] Xây dựng RESTful API hoàn chỉnh
- [x] Triển khai hệ thống logging chi tiết
- [ ] Điểm A môn học.....

## 🙏 Special Thanks
- Cảm ơn Stack Overflow (vì đã giúp copy-paste code)
- Cảm ơn ChatGPT, grok, claude, Gemini, ...... (vì đã giúp code, 'gần như toàn bộ frontend :v')
- Cảm ơn Github Copilot (vì đã giúp code)
---