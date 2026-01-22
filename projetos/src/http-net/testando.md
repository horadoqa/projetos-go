# ROTAS

Testando:

```bash
curl -I http://localhost:8080/
```

Resultado:

```bash
Accept-Ranges: bytes
Content-Length: 1496
Content-Type: text/html; charset=utf-8
Last-Modified: Thu, 22 Jan 2026 14:07:32 GMT
Date: Thu, 22 Jan 2026 14:52:41 GMT
```

LOG

2026/01/22 11:52:41 Request: HEAD /
2026/01/22 11:52:41 Headers: map[Accept:[*/*] User-Agent:[curl/7.81.0]]
2026/01/22 11:52:41 Response status: 200

---

Testando o healthcheck

```bash
curl -I http://localhost:8080/healthcheck
```

Resultado:

```bash
HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 1493
Content-Type: text/html; charset=utf-8
Last-Modified: Thu, 22 Jan 2026 13:51:59 GMT
Date: Thu, 22 Jan 2026 14:52:56 GMT
```

LOG

2026/01/22 11:52:56 Request: HEAD /healthcheck
2026/01/22 11:52:56 Headers: map[Accept:[*/*] User-Agent:[curl/7.81.0]]
2026/01/22 11:52:56 Response status: 200

---

Testando o POST

```bash
curl -X POST http://localhost:8080/api/data -H "Content-Type: application/json" -d '{"teste":123}'
```

Resultado:

```bash
{"message":"Dados recebidos","body":"{"teste":123}"}%
```

LOG:

2026/01/22 11:51:49 Request: POST /api/data
2026/01/22 11:51:49 Headers: map[Accept:[*/*] Content-Length:[13] Content-Type:[application/json] User-Agent:[curl/7.81.0]]
2026/01/22 11:51:49 Response status: 200