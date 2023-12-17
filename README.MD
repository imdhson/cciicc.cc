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
Successfully received certificate.
Certificate is saved at: /etc/letsencrypt/live/cciicc.cc/fullchain.pem
Key is saved at:         /etc/letsencrypt/live/cciicc.cc/privkey.pem