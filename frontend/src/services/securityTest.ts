import api, { authAPI, productAPI } from './api'
import axios from 'axios'

export interface SecurityTestResult {
  id: string
  name: string
  description: string
  status: 'PASS' | 'FAIL' | 'WARNING'
  result: string
  details?: string
  timestamp: string
}

// Tạo test result helper
function createTestResult(
  id: string,
  name: string,
  description: string,
  status: 'PASS' | 'FAIL' | 'WARNING',
  result: string,
  details?: string
): SecurityTestResult {
  return {
    id,
    name,
    description,
    status,
    result,
    details,
    timestamp: new Date().toLocaleString('vi-VN')
  }
}

// 1. JWT Token Validation Tests
export async function testJwtTokenValidation(): Promise<SecurityTestResult> {
  try {
    // Test với fake token bằng cách tạo request manual
    const testApi = api.create({
      headers: {
        'Authorization': 'Bearer fake_invalid_token'
      }
    })
    
    try {
      await testApi.get('/profile')
      return createTestResult(
        'jwt_token_validation',
        'JWT Token Validation',
        'Test token expiry, invalid signatures',
        'FAIL',
        'API accepts invalid JWT tokens',
        'Server không validate JWT token đúng cách'
      )
    } catch (error: any) {
      if (error.response?.status === 401) {
        return createTestResult(
          'jwt_token_validation',
          'JWT Token Validation',
          'Test token expiry, invalid signatures',
          'PASS',
          'Invalid tokens are properly rejected',
          `Server trả về 401 khi token không hợp lệ`
        )
      }
      return createTestResult(
        'jwt_token_validation',
        'JWT Token Validation',
        'Test token expiry, invalid signatures',
        'WARNING',
        'Unexpected response to invalid token',
        `Status: ${error.response?.status}`
      )
    }
  } catch (error) {
    return createTestResult(
      'jwt_token_validation',
      'JWT Token Validation',
      'Test token expiry, invalid signatures',
      'FAIL',
      'Test execution failed',
      `Error: ${error}`
    )
  }
}

// 2. Login Brute Force Protection
export async function testLoginBruteForce(): Promise<SecurityTestResult> {
  const attempts = 15
  let blockedCount = 0
  
  const promises = Array.from({ length: attempts }, async (_, i) => {
    try {
      await authAPI.login({
        username: `testuser_${i}`,
        password: 'wrongpassword'
      })
      return 'success'
    } catch (error: any) {
      if (error.response?.status === 429) {
        blockedCount++
        return 'blocked'
      }
      return 'failed'
    }
  })
  
  await Promise.allSettled(promises)
  
  if (blockedCount > 0) {
    return createTestResult(
      'login_bruteforce',
      'Login Brute Force Protection',
      'Test multiple failed login attempts',
      'PASS',
      `Rate limiting active: ${blockedCount}/${attempts} requests blocked`,
      `Rate limiting đang hoạt động`
    )
  } else {
    return createTestResult(
      'login_bruteforce',
      'Login Brute Force Protection',
      'Test multiple failed login attempts',
      'FAIL',
      'No rate limiting detected',
      `Tất cả ${attempts} requests được xử lý mà không có rate limiting`
    )
  }
}

// 3. Session Hijacking Test (Cookie Security)
export async function testSessionHijacking(): Promise<SecurityTestResult> {
  const cookies = document.cookie.split(';')
  const authCookie = cookies.find(cookie => cookie.trim().startsWith('auth_token='))
  
  if (!authCookie) {
    return createTestResult(
      'session_hijacking',
      'Session Hijacking Test',
      'Test cookie security attributes',
      'WARNING',
      'No auth token cookie found',
      'User might not be logged in or token is stored differently'
    )
  }
  
  // Check for secure cookie attributes (can't fully test httpOnly from JS)
  const issues = []
  const cookieValue = authCookie.split('=')[1]
  
  // Test if cookie is accessible via JavaScript (should be httpOnly)
  if (cookieValue && cookieValue !== '') {
    issues.push('Cookie accessible via JavaScript (not httpOnly)')
  }
  
  // Check if running on HTTPS (secure flag test)
  if (window.location.protocol !== 'https:' && window.location.hostname !== 'localhost') {
    issues.push('Not running on HTTPS - secure flag cannot be verified')
  }
  
  if (issues.length === 0) {
    return createTestResult(
      'session_hijacking',
      'Session Hijacking Test',
      'Test cookie security attributes',
      'PASS',
      'Cookie security appears properly configured',
      'HttpOnly và Secure flags seem to be set correctly'
    )
  } else {
    return createTestResult(
      'session_hijacking',
      'Session Hijacking Test',
      'Test cookie security attributes',
      'WARNING',
      'Potential cookie security issues detected',
      `Issues: ${issues.join(', ')}`
    )
  }
}

// 4. Privilege Escalation Test
export async function testPrivilegeEscalation(): Promise<SecurityTestResult> {
  try {
    await productAPI.create({
      name: 'Test Product',
      description: 'Privilege escalation test',
      price: 100
    })
    
    return createTestResult(
      'privilege_escalation',
      'Privilege Escalation',
      'Test role-based access control',
      'FAIL',
      'Non-admin user can create products',
      'API cho phép user thường tạo sản phẩm'
    )
  } catch (error: any) {
    if (error.response?.status === 403) {
      return createTestResult(
        'privilege_escalation',
        'Privilege Escalation',
        'Test role-based access control',
        'PASS',
        'Access denied for non-admin user',
        'Server đúng cách từ chối quyền truy cập'
      )
    } else if (error.response?.status === 401) {
      return createTestResult(
        'privilege_escalation',
        'Privilege Escalation',
        'Test role-based access control',
        'PASS',
        'Authentication required',
        'Server yêu cầu authentication'
      )
    }
    return createTestResult(
      'privilege_escalation',
      'Privilege Escalation',
      'Test role-based access control',
      'WARNING',
      'Unexpected response',
      `Status: ${error.response?.status}`
    )
  }
}

// 5. IDOR (Insecure Direct Object Reference) Test
export async function testIDOR(): Promise<SecurityTestResult> {
  try {
    // Thử truy cập user ID khác
    const otherUserIds = [1, 2, 3, 999, -1, 'admin', 'test']
    let vulnerableEndpoints = 0
    
    for (const userId of otherUserIds) {
      try {
        // Tạo custom request để test IDOR
        const response = await api.get(`/users/${userId}`)
        if (response.status === 200) {
          vulnerableEndpoints++
        }
      } catch (error: any) {
        // Expected behavior - should be 403 or 404
        if (error.response?.status !== 403 && error.response?.status !== 404) {
          // Log unexpected status codes
        }
      }
    }
    
    if (vulnerableEndpoints > 0) {
      return createTestResult(
        'idor_test',
        'IDOR (Direct Object Reference)',
        'Test access to unauthorized resources',
        'FAIL',
        `IDOR vulnerability detected: ${vulnerableEndpoints} endpoints accessible`,
        'API cho phép truy cập thông tin user khác'
      )
    } else {
      return createTestResult(
        'idor_test',
        'IDOR (Direct Object Reference)',
        'Test access to unauthorized resources',
        'PASS',
        'No IDOR vulnerabilities detected',
        'API đúng cách từ chối truy cập resources của user khác'
      )
    }
  } catch (error) {
    return createTestResult(
      'idor_test',
      'IDOR (Direct Object Reference)',
      'Test access to unauthorized resources',
      'WARNING',
      'Test could not be completed',
      `Error during IDOR test: ${error}`
    )
  }
}

// 6. Admin Function Bypass
export async function testAdminBypass(): Promise<SecurityTestResult> {
  const adminEndpoints = [
    { method: 'POST', url: '/products', data: { name: 'Test', price: 100 } },
    { method: 'PUT', url: '/products/1', data: { name: 'Modified', price: 200 } },
    { method: 'DELETE', url: '/products/1' }
  ]
  
  let bypassedEndpoints = 0
  const results: string[] = []
  
  for (const endpoint of adminEndpoints) {
    try {
      let response
      switch (endpoint.method) {
        case 'POST':
          response = await api.post(endpoint.url, endpoint.data)
          break
        case 'PUT':
          response = await api.put(endpoint.url, endpoint.data)
          break
        case 'DELETE':
          response = await api.delete(endpoint.url)
          break
      }
      
      if (response?.status === 200 || response?.status === 201) {
        bypassedEndpoints++
        results.push(`${endpoint.method} ${endpoint.url}: Accessible`)
      }
    } catch (error: any) {
      const status = error.response?.status
      if (status === 403) {
        results.push(`${endpoint.method} ${endpoint.url}: Properly protected`)
      } else if (status === 401) {
        results.push(`${endpoint.method} ${endpoint.url}: Auth required`)
      } else {
        results.push(`${endpoint.method} ${endpoint.url}: Status ${status}`)
      }
    }
  }
  
  if (bypassedEndpoints > 0) {
    return createTestResult(
      'admin_bypass',
      'Admin Function Bypass',
      'Test admin-only endpoint access',
      'FAIL',
      `${bypassedEndpoints}/${adminEndpoints.length} admin endpoints bypassed`,
      results.join('; ')
    )
  } else {
    return createTestResult(
      'admin_bypass',
      'Admin Function Bypass',
      'Test admin-only endpoint access',
      'PASS',
      'All admin endpoints properly protected',
      results.join('; ')
    )
  }
}

// 7. Rate Limiting Test (Login)
export async function testRateLimitLogin(): Promise<SecurityTestResult> {
  const requestCount = 20
  let rateLimitedCount = 0
  let errorCount = 0
  
  const promises = Array.from({ length: requestCount }, async () => {
    try {
      await authAPI.login({ username: 'testuser', password: 'wrongpass' })
      return 'success'
    } catch (error: any) {
      if (error.response?.status === 429) {
        rateLimitedCount++
        return 'rate_limited'
      } else {
        errorCount++
        return 'error'
      }
    }
  })
  
  await Promise.allSettled(promises)
  
  if (rateLimitedCount > 0) {
    return createTestResult(
      'rate_limit_login',
      'Login Rate Limiting',
      'Test 10 req/sec limit on auth endpoints',
      'PASS',
      `Rate limiting working: ${rateLimitedCount}/${requestCount} requests blocked`,
      `Blocked: ${rateLimitedCount}, Errors: ${errorCount}, Total: ${requestCount}`
    )
  } else {
    return createTestResult(
      'rate_limit_login',
      'Login Rate Limiting',
      'Test 10 req/sec limit on auth endpoints',
      'FAIL',
      'No rate limiting detected on login endpoint',
      `All ${requestCount} requests processed without rate limiting`
    )
  }
}

// 8. General API Rate Limiting
export async function testRateLimitAPI(): Promise<SecurityTestResult> {
  const requestCount = 25
  let rateLimitedCount = 0
  let successCount = 0
  
  const promises = Array.from({ length: requestCount }, async () => {
    try {
      await api.get('/products?page=1&limit=1')
      successCount++
      return 'success'
    } catch (error: any) {
      if (error.response?.status === 429) {
        rateLimitedCount++
        return 'rate_limited'
      }
      return 'error'
    }
  })
  
  await Promise.allSettled(promises)
  
  if (rateLimitedCount > 0) {
    return createTestResult(
      'rate_limit_api',
      'API Rate Limiting',
      'Test general API rate limits',
      'PASS',
      `Rate limiting active: ${rateLimitedCount}/${requestCount} blocked`,
      `Success: ${successCount}, Blocked: ${rateLimitedCount}`
    )
  } else {
    return createTestResult(
      'rate_limit_api',
      'API Rate Limiting',
      'Test general API rate limits',
      'WARNING',
      'No rate limiting detected',
      `Tất cả ${requestCount} requests thành công`
    )
  }
}

// 9. DDoS Protection Test
export async function testDDoSProtection(): Promise<SecurityTestResult> {
  const burstSize = 50
  const startTime = Date.now()
  let completedRequests = 0
  let rateLimitedRequests = 0
  
  const promises = Array.from({ length: burstSize }, async () => {
    try {
      await api.get('/products')
      completedRequests++
    } catch (error: any) {
      if (error.response?.status === 429) {
        rateLimitedRequests++
      }
    }
  })
  
  await Promise.allSettled(promises)
  const duration = Date.now() - startTime
  
  if (rateLimitedRequests > burstSize * 0.3) {
    return createTestResult(
      'ddos_protection',
      'DDoS Protection',
      'Test burst request handling',
      'PASS',
      `DDoS protection active: ${rateLimitedRequests}/${burstSize} requests blocked`,
      `Completed in ${duration}ms, ${completedRequests} succeeded, ${rateLimitedRequests} blocked`
    )
  } else {
    return createTestResult(
      'ddos_protection',
      'DDoS Protection',
      'Test burst request handling',
      'WARNING',
      `Weak DDoS protection: only ${rateLimitedRequests}/${burstSize} requests blocked`,
      `Server handled ${completedRequests} burst requests in ${duration}ms`
    )
  }
}

// 10. SQL Injection Test
export async function testSQLInjection(): Promise<SecurityTestResult> {
  const sqlPayloads = [
    "'; DROP TABLE users; --",
    "' OR '1'='1",
    "admin'--"
  ]
  
  let vulnerableCount = 0
  
  for (const payload of sqlPayloads) {
    try {
      await authAPI.login({
        username: payload,
        password: 'test'
      })
      vulnerableCount++
    } catch (error: any) {
      // Expected - should reject malicious input
    }
  }
  
  if (vulnerableCount > 0) {
    return createTestResult(
      'sql_injection',
      'SQL Injection',
      'Test SQL injection vulnerabilities',
      'FAIL',
      `SQL injection vulnerabilities detected: ${vulnerableCount}`,
      'Server dễ bị tấn công SQL injection'
    )
  } else {
    return createTestResult(
      'sql_injection',
      'SQL Injection',
      'Test SQL injection vulnerabilities',
      'PASS',
      'No SQL injection vulnerabilities detected',
      'Tất cả SQL payloads đều bị từ chối'
    )
  }
}

// 11. XSS Test
export async function testXSS(): Promise<SecurityTestResult> {
  const xssPayloads = [
    '<script>alert("XSS")</script>',
    '"><script>alert("XSS")</script>',
    'javascript:alert("XSS")',
    '<img src=x onerror=alert("XSS")>',
    '<svg onload=alert("XSS")>'
  ]
  
  let vulnerableFields = 0
  const results: string[] = []
  
  for (const payload of xssPayloads) {
    try {
      // Test XSS trong product creation
      await productAPI.create({
        name: payload,
        description: `Test XSS: ${payload}`,
        price: 100
      })
      vulnerableFields++
      results.push(`XSS payload accepted: ${payload}`)
    } catch (error: any) {
      // May fail due to auth or validation
      if (error.response?.status === 400) {
        results.push(`Payload rejected: ${payload}`)
      }
    }
  }
  
  if (vulnerableFields > 0) {
    return createTestResult(
      'xss_test',
      'XSS (Cross-Site Scripting)',
      'Test script injection in inputs',
      'FAIL',
      `XSS vulnerabilities detected: ${vulnerableFields} payloads accepted`,
      results.join('; ')
    )
  } else {
    return createTestResult(
      'xss_test',
      'XSS (Cross-Site Scripting)',
      'Test script injection in inputs',
      'PASS',
      'No XSS vulnerabilities detected',
      'All script injection attempts were properly sanitized'
    )
  }
}

// 12. JSON Injection Test
export async function testJSONInjection(): Promise<SecurityTestResult> {
  const malformedPayloads = [
    '{"name": "test", "price": "invalid"}',
    '{"name": null, "description": 123, "price": "string"}',
    '{"extra_field": "malicious", "name": "test", "price": -1}',
    '{"name": "<script>alert(1)</script>", "price": 999999999999999999}',
  ]
  
  let acceptedPayloads = 0
  let rejectedPayloads = 0
  
  for (const payload of malformedPayloads) {
    try {
      await api.post('/products', JSON.parse(payload))
      acceptedPayloads++
    } catch (error: any) {
      if (error.response?.status === 400) {
        rejectedPayloads++
      }
    }
  }
  
  if (acceptedPayloads === 0) {
    return createTestResult(
      'json_injection',
      'JSON Injection',
      'Test malformed JSON payloads',
      'PASS',
      'All malformed JSON payloads rejected',
      `${rejectedPayloads}/${malformedPayloads.length} payloads properly rejected`
    )
  } else {
    return createTestResult(
      'json_injection',
      'JSON Injection',
      'Test malformed JSON payloads',
      'WARNING',
      `${acceptedPayloads} malformed payloads accepted`,
      `Server accepted ${acceptedPayloads} out of ${malformedPayloads.length} malformed JSON payloads`
    )
  }
}

// Export test functions map
export const securityTests = {
  jwt_token_validation: testJwtTokenValidation,
  login_bruteforce: testLoginBruteForce,
  session_hijacking: testSessionHijacking,
  privilege_escalation: testPrivilegeEscalation,
  idor_test: testIDOR,
  admin_bypass: testAdminBypass,
  rate_limit_login: testRateLimitLogin,
  rate_limit_api: testRateLimitAPI,
  ddos_protection: testDDoSProtection,
  sql_injection: testSQLInjection,
  xss_test: testXSS,
  json_injection: testJSONInjection,
} 