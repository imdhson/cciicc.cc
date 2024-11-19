let urladdress = ""
let qrsmallclick_toggle = false
let popup_once = false

function popup(toggle, text) {
    if (toggle && !popup_once) {
        //팝업 생성
        const popup = document.createElement("div")
        popup.id = "popup"
        popup.textContent = text
        const main = document.querySelector("main")
        main.appendChild(popup)
        popup_once = true
    } else if (!toggle) {
        const popup = document.getElementById("popup")
        if (popup != null || popup != undefined) {
            document.querySelector("main").removeChild(document.getElementById("popup"))
        }
        popup_once = false
    }

}

function qrsmallClick() {
    if (!qrsmallclick_toggle) { //토글 false, 작은 화면
        let main = document.querySelector("main")
        main.style.gridTemplateColumns = "1fr"

        let qrsmall = document.getElementById("QRsmall")
        qrsmall.style.height = "50%"
        qrsmall.style.top = "30px"
        qrsmall.style.right = "30px"
        qrsmall.style.border = "1px solid black"
        qrsmall.style.borderRadius = "20px"
    } else { //토글t, 작은 화면
        let main = document.querySelector("main")
        main.style.gridTemplateColumns = "1fr"

        let qrsmall = document.getElementById("QRsmall")
        qrsmall.style.height = "inherit"
        qrsmall.style.top = "0"
        qrsmall.style.right = "0"
        qrsmall.style.border = "none"
        qrsmall.style.borderRadius = "0"
    }
    qrsmallclick_toggle = !qrsmallclick_toggle
}

function space_content_onload(urladdress_i) {
    urladdress = urladdress_i

    const socket = new WebSocket("/ws");

    // 연결이 열리면 실행되는 이벤트 핸들러
    socket.onopen = function (event) {
        console.log("WebSocket 연결이 열렸습니다.");
        console.log(event.data)

        // 서버로 메시지 전송
        socket.send("클라이언트에서 보내는 메시지입니다!");
    };

    // 서버로부터 메시지를 받으면 실행되는 이벤트 핸들러 
    socket.onmessage = function (event) {
        console.log("서버로부터 메시지 수신:", event.data);
        if (event.data.Sp_file_status != 0) {
            loadPDF("/space/file")
        }
    };

    // 연결이 닫히면 실행되는 이벤트 핸들러
    socket.onclose = function (event) {
        if (event.wasClean) {
            console.log(`연결이 정상적으로 종료되었습니다. 코드: ${event.code}, 이유: ${event.reason}`);
        } else {
            console.log('연결이 비정상적으로 종료되었습니다.');
        }
    };

    // 에러가 발생하면 실행되는 이벤트 핸들러
    socket.onerror = function (error) {
        console.error(`WebSocket 에러 발생: ${error.message}`);
    };

    // 서버로 메시지를 보내는 함수
    function sendMessage(message) {
        // 연결 상태 확인
        if (socket.readyState === WebSocket.OPEN) {
            socket.send(message);
        } else {
            console.log("WebSocket 연결이 열려있지 않습니다.");
        }
    }

    // 연결 종료 함수
    function closeConnection() {
        socket.close();
    }
}

function addComment_form(event) { // 키보드의 모든 입력을 받고 엔터 혹은 터치(마우스) 클릭시에만 수행
    if (event == -1 || event.key == "Enter") { //터치 || 엔터
        const xhr = new XMLHttpRequest();
        const url = urladdress + "/space/addcomment";

        const form = document.getElementById("comment_form")
        const data = new FormData(form);
        xhr.open("POST", url, true);
        xhr.send(data);

        const form_text = document.getElementById("comment")
        form_text.value = ""
    }

}

function linkCopyToClipboard(sp_id) {
    // 새로운 텍스트 영역 요소 생성
    const textArea = document.createElement('textarea');
    textArea.value = urladdress + "/" + sp_id; // 복사하고자 하는 텍스트 설정

    // 스타일 설정하여 화면에 표시되지 않도록 함
    textArea.style.position = 'absolute';
    textArea.style.left = '-9999px';

    document.body.appendChild(textArea); // DOM에 추가
    textArea.focus(); // 텍스트 영역에 포커스
    textArea.select(); // 텍스트 영역의 텍스트 선택

    try {
        // 텍스트 복사 시도
        document.execCommand('copy');
        popup(true, "클립보드에 복사했어요.")
        setTimeout(function () {
            popup(false);
        }, 1000);
    } catch (err) {
        console.error('Failed to copy text', err);
    }

    document.body.removeChild(textArea); // DOM에서 텍스트 영역 제거
}

let pdfDoc = null,
    pageNum = 1,
    pageRendering = false,
    pageNumPending = null,
    scale = 1.5;

function uploadPDF() {
    const file = document.getElementById('pdf-file').files[0];
    if (file) {
        const formData = new FormData();
        formData.append('file', file);

        fetch('/space/addfile', {
            method: 'POST',
            body: formData
        }).then(response => response.json())
            .then(data => {
                if (data.success) {
                    loadPDF('/space/file');
                }
            });
    }
}

function loadPDF(url) {
    pdfjsLib.getDocument(url).promise.then(function (pdf) {
        pdfDoc = pdf;
        document.getElementById('page-num').textContent = pageNum + ' / ' + pdf.numPages;
        renderPage(pageNum);
    });
}

function renderPage(num) {
    pageRendering = true;
    pdfDoc.getPage(num).then(function (page) {
        const canvas = document.getElementById('pdf-render');
        const ctx = canvas.getContext('2d');
        const viewport = page.getViewport({ scale: scale });
        canvas.height = viewport.height;
        canvas.width = viewport.width;

        const renderContext = {
            canvasContext: ctx,
            viewport: viewport
        };
        page.render(renderContext);

        pageRendering = false;
        if (pageNumPending !== null) {
            renderPage(pageNumPending);
            pageNumPending = null;
        }
    });

    document.getElementById('page-num').textContent = num + ' / ' + pdfDoc.numPages;
}

function queueRenderPage(num) {

    //post file context로 페이지넘버 전송
    let formData = new FormData();
    formData.append('pageNum', num)
    fetch('/space/filecontext', {
        method: 'POST',
        body: formData,
    })
        .then(response => response.status)
        .then(data => {
            console.log('성공:', data);
        })
        .catch((error) => {
            console.error('에러:', error);
        });


    if (pageRendering) {
        pageNumPending = num;
    } else {
        renderPage(num);
    }
}

function onPrevPage() {
    if (pageNum <= 1) {
        return;
    }
    pageNum--;
    queueRenderPage(pageNum);
}

function onNextPage() {
    if (pageNum >= pdfDoc.numPages) {
        return;
    }
    pageNum++;
    queueRenderPage(pageNum);
}


