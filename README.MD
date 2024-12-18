# 실시간 파일, 의견 공유 플랫폼 | cciicc

## 소개

cciicc는 웹 기반 실시간 연동 파일, 텍스트 공유 플랫폼입니다. 강의실, 회의실과 같이 단체로 대화를 하며 공통의 정보를 열람해야 할 때 유용하게 사용할 수 있습니다. 웹 표준을 준수하여 다양한 디바이스에서 동일한 기능을 제공합니다.

## 개발 배경
- 대면 상황에서 온라인 자료 접근의 불편함 해소
- 교육 현장에서의 정보 공유 효율성 증대
- 온라인 회의 시 자료 공유의 품질 향상

## 설치 및 구성

### 1. requirements
- Go언어 1.23 이상

### 2. 환경 변수 수정하기
- 
### 3. 서버 실행하기
아래 명령어를 사용하여 프로젝트를 설치하고 실행합니다:
```bash
git clone https://github.com/imdhson/cciicc.cc.git
cd cciicc.cc
go run main.go
```
### 4. 접속하기
웹브라우저를 열고 [http://localhost:8000](http://localhost:8000) 에 접속하세요.

## 주요 기능

- 다수의 사용자들이 호스트가 주최한 웹 공간에 참여
- 호스트는 텍스트, 비디오, 이미지, 오디오 등 파일을 업로드
- 게스트는 호스트와 파일을 자동으로 실시간 열람 가능
- 게스트는 텍스트 형태로 채팅 전송 가능
- QR 코드로 호스트가 만든 공간에 빠른 방문
- 호스트가 만든 공간을 위한 쉬운 URL 주소 자동 생성

## 발표, 구동 영상
### Youtube:
[![Video Title](https://img.youtube.com/vi/6xaHcNrjRQ8/0.jpg)](https://www.youtube.com/watch?v=6xaHcNrjRQ8)

## 라이선스

이 프로젝트는 [MIT LICENSE](LICENSE.md) 하에 배포됩니다.

## 연락처

개발자: 손동휘

손동휘 이메일: mail@imdhson.com

Instagram: @imdhson

GitHub: [https://github.com/imdhson](https://github.com/imdhson)

## memo
```
<h1>cciicc server source code</h1>
<br>
import qrcode "github.com/skip2/go-qrcode"    https://github.com/skip2/go-qrcode
<br>
systemctl 사용시 cli 비활성화해야함. 안그럼 무한 입력리턴 발생

<br><br>
https://certbot.eff.org/instructions?ws=other&os=debianbuster
<br>
ssl은 listenandservetls에서 키 연결하면됨.<br> 현재 certbot으로 저장된 위치에서 수행중
<br>
Certificate is saved at: /etc/letsencrypt/live/cciicc.cc/fullchain.pem
Key is saved at:         /etc/letsencrypt/live/cciicc.cc/privkey.pem
```