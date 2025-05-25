import requests
import time
import concurrent.futures
import json
import os
import logging
from datetime import datetime

log_filename = f"security_test_log_{datetime.now().strftime('%Y%m%d_%H%M%S')}.log"
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s',
    handlers=[
        logging.FileHandler(log_filename, encoding='utf-8'),
        logging.StreamHandler()
    ]
)

BASE_URL = os.getenv('BASE_URL', 'http://localhost')

def test_authentication():
    logging.info("\n=== KIỂM TRA XÁC THỰC ===")
    
    logging.info("\nTest 1: Kiểm tra đăng ký với mật khẩu yếu")
    response = requests.post(f"{BASE_URL}/api/auth/register", json={
        "username": "testuser1",
        "password": "123"
    })
    logging.info(f"Trạng thái: {response.status_code}")
    logging.info(f"Phản hồi: {response.text}")

    logging.info("\nTest 2: Kiểm tra chống SQL Injection khi đăng ký")
    response = requests.post(f"{BASE_URL}/api/auth/register", json={
        "username": "admin' OR '1'='1",
        "password": "password123"
    })
    logging.info(f"Trạng thái: {response.status_code}")
    logging.info(f"Phản hồi: {response.text}")

    logging.info("\nTest 3: Kiểm tra đăng nhập với tài khoản không tồn tại")
    response = requests.post(f"{BASE_URL}/api/auth/login", json={
        "username": "nonexistent",
        "password": "password123"
    })
    logging.info(f"Trạng thái: {response.status_code}")
    logging.info(f"Phản hồi: {response.text}")

def test_authorization():
    logging.info("\n=== KIỂM TRA PHÂN QUYỀN ===")
    
    logging.info("\nTest 4: Kiểm tra truy cập khi không có token")
    response = requests.get(f"{BASE_URL}/api/profile")
    logging.info(f"Trạng thái: {response.status_code}")
    logging.info(f"Phản hồi: {response.text}")

    logging.info("\nTest 5: Kiểm tra truy cập với token không hợp lệ")
    headers = {"Authorization": "Bearer invalid_token"}
    response = requests.get(f"{BASE_URL}/api/profile", headers=headers)
    logging.info(f"Trạng thái: {response.status_code}")
    logging.info(f"Phản hồi: {response.text}")

    logging.info("\nTest 6: Kiểm tra truy cập route admin với tài khoản thường")
    requests.post(f"{BASE_URL}/api/auth/register", json={
        "username": "nguoidung_thuong",
        "password": "Password123!"
    })
    response = requests.post(f"{BASE_URL}/api/auth/login", json={
        "username": "nguoidung_thuong",
        "password": "Password123!"
    })
    if response.status_code == 200:
        token = response.json().get("token")
        headers = {"Authorization": f"Bearer {token}"}
        response = requests.post(f"{BASE_URL}/api/products", headers=headers, json={
            "name": "Test Product",
            "price": 99.99
        })
        logging.info(f"Trạng thái: {response.status_code}")
        logging.info(f"Phản hồi: {response.text}")

def test_input_validation():
    logging.info("\n=== KIỂM TRA VALIDATION DỮ LIỆU ĐẦU VÀO ===")
    
    logging.info("\nTest 7: Kiểm tra chống XSS trong username")
    response = requests.post(f"{BASE_URL}/api/auth/register", json={
        "username": "<script>alert('xss')</script>",
        "password": "Password123!"
    })
    logging.info(f"Trạng thái: {response.status_code}")
    logging.info(f"Phản hồi: {response.text}")

    logging.info("\nTest 8: Kiểm tra giới hạn kích thước dữ liệu")
    large_string = "A" * 1000000
    response = requests.post(f"{BASE_URL}/api/auth/register", json={
        "username": large_string,
        "password": "Password123!"
    })
    logging.info(f"Trạng thái: {response.status_code}")
    logging.info(f"Phản hồi: {response.text}")

def test_rate_limiting():
    logging.info("\n=== KIỂM TRA GIỚI HẠN TỐC ĐỘ REQUEST ===")
    
    logging.info("\nTest 9: Kiểm tra rate limiting")
    results = []
    
    def make_request():
        response = requests.post(f"{BASE_URL}/api/auth/login", json={
            "username": "test",
            "password": "test"
        })
        return response.status_code

    with concurrent.futures.ThreadPoolExecutor(max_workers=10) as executor:
        futures = [executor.submit(make_request) for _ in range(20)]
        for future in concurrent.futures.as_completed(futures):
            results.append(future.result())
    
    logging.info(f"Mã phản hồi: {results}")
    logging.info(f"Số lượng mã 429 (Quá nhiều request): {results.count(429)}")

def test_cors():
    logging.info("\n=== KIỂM TRA CẤU HÌNH CORS ===")
    
    logging.info("\nTest 10: Kiểm tra headers CORS")
    headers = {
        "Origin": "http://malicious-site.com",
        "Access-Control-Request-Method": "POST",
        "Access-Control-Request-Headers": "Content-Type"
    }
    response = requests.options(f"{BASE_URL}/api/auth/register", headers=headers)
    logging.info("Headers CORS:")
    for header, value in response.headers.items():
        if header.lower().startswith('access-control'):
            logging.info(f"{header}: {value}")

def log_test_summary(start_time):
    end_time = time.time()
    duration = end_time - start_time
    logging.info("\n=== KẾT QUẢ KIỂM TRA ===")
    logging.info(f"Thời gian chạy: {duration:.2f} giây")
    logging.info(f"File log: {log_filename}")

if __name__ == "__main__":
    try:
        start_time = time.time()
        logging.info("=== BẮT ĐẦU KIỂM TRA BẢO MẬT ===")
        logging.info(f"Thời gian bắt đầu: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        logging.info(f"URL API: {BASE_URL}")
        
        test_authentication()
        test_authorization()
        test_input_validation()
        test_rate_limiting()
        test_cors()
        
        log_test_summary(start_time)
        
    except requests.exceptions.RequestException as e:
        logging.error(f"Lỗi khi chạy test: {e}") 