# Gin Helloworld 애플리케이션

이 애플리케이션은 Go 언어와 Gin 프레임워크를 사용하여 작성된 간단한 웹 애플리케이션입니다. 지정된 포트에서 요청을 수신하고 환영 메시지를 반환합니다. 환영 메시지는 `TARGET` 환경 변수의 값에 따라 동적으로 구성됩니다.

## 목차

- [개요](#개요)
- [필요 사항](#필요-사항)
- [설치 방법](#설치-방법)
- [사용 방법](#사용-방법)
- [환경 변수](#환경-변수)
- [라이선스](#라이선스)

---

## 개요

이 애플리케이션은 Gin 프레임워크를 사용한 최소한의 웹 서버 예제입니다. 기본적으로 서버는 `8080` 포트에서 실행되며, `/` 경로로 들어오는 HTTP GET 요청에 대해 "Hello [TARGET 환경 변수 값]" 메시지를 반환합니다. `TARGET` 환경 변수가 설정되지 않은 경우 기본 메시지는 "Hello World!"로 설정됩니다.

---

## 필요 사항

이 애플리케이션을 실행하려면 아래와 같은 요구 사항이 필요합니다:

- Go 1.18 이상
- Git (소스 코드 클론 용)
- 인터넷 연결 (의존성 설치 시 필요)

---

## 설치 방법

1. **리포지토리 클론**  
먼저 프로젝트를 클론합니다:
   ```bash
   git clone https://github.com/your-repo/helloworld.git
   cd helloworld
   ```

2. **모듈 정리**  
   필요 라이브러리를 설치/정리합니다:
   ```bash
   go mod tidy
   ```

---

## 사용 방법

1. **애플리케이션 실행**  
아래 명령어를 실행하여 애플리케이션을 시작합니다:
   ```bash
   go run helloworld.go
   ```

2. **웹 브라우저에서 테스트**  
웹 브라우저로 접속하거나 cURL로 테스트합니다:
   ```bash
   curl http://localhost:8080/
   ```

   기본적으로 "Hello World!" 메시지가 표시됩니다.

---

## 환경 변수

애플리케이션은 다음 환경 변수를 사용할 수 있습니다:

- `SERVER_PORT`  
  HTTP 서버가 수신 대기할 포트를 설정합니다. 설정하지 않으면 기본값 `8080`이 사용됩니다.

- `TARGET`  
  응답 메시지에 표시될 이름을 설정합니다. 설정하지 않은 경우 "World"가 기본값으로 사용됩니다.

**예시:**
```bash
export TARGET=GinFramework
export SERVER_PORT=9090
go run helloworld.go
```

실행 후 `http://localhost:9090`에 접근하면 "Hello GinFramework!" 메시지가 표시됩니다.

---
