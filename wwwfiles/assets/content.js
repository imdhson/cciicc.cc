let urladdress = ""
let qrsmallclick_toggle = false
let uploadareaclock_toggle = false
let popup_once = false
let file_context = 1

function showPopup(message) {
    const popup = document.getElementById('popup');
    const overlay = document.getElementById('overlay');
    
    popup.textContent = message;
    popup.style.display = 'block';
    overlay.style.display = 'block';
    
    setTimeout(() => {
        popup.style.opacity = '1';
        popup.style.transform = 'translate(-50%, -50%) scale(1)';
        overlay.style.opacity = '1';
    }, 10);
    
    setTimeout(() => {
        popup.style.opacity = '0';
        popup.style.transform = 'translate(-50%, -50%) scale(0.8)';
        overlay.style.opacity = '0';
        
        setTimeout(() => {
            popup.style.display = 'none';
            overlay.style.display = 'none';
        }, 400);
    }, 2000);
}

const uploadToggle = document.getElementById('uploadToggle');
const uploadArea = document.getElementById('uploadArea');
const fileInput = document.getElementById('fileInput');
const uploadButton = document.getElementById('uploadButton');

function qrsmallClick() {
    const qrsmall = document.getElementById("QRsmall");
    
    if (!qrsmallclick_toggle) {
        qrsmall.classList.add('expanded');
    } else {
        qrsmall.classList.remove('expanded');
    }
    
    qrsmallclick_toggle = !qrsmallclick_toggle;
}

function uploadToggle_onclick(){
    uploadArea.classList.toggle('show')
    if(uploadArea.classList.contains('show')){
        uploadToggle.textContent = '업로드 창 닫기'
    } else{
        uploadToggle.textContent = '업로드 창 보기'
    }
}

function space_content_onload(urladdress_i) {
    urladdress = urladdress_i

    const socket = new WebSocket("/ws");

    socket.onopen = function (event) {
        console.log("WebSocket 연결이 열렸습니다.");
        socket.send("클라이언트에서 보내는 메시지입니다!");
    };

    socket.onmessage = function (event) {
        console.log("서버로부터 메시지 수신:", event.data);
        let jsonData = JSON.parse(event.data)
        if (jsonData.Sp_file_status == 1 && jsonData.Sp_file_ext == '.pdf') {
            loadPDF("/space/file")
            file_context = jsonData.Sp_file_context == 0 ? 1 : jsonData.Sp_file_context
        }
        if (jsonData.Sp_ws_type == 'file_context' && jsonData.Sp_file_context != null) {
            file_context = parseInt(jsonData.Sp_file_context)
            loadPDF("/space/file")
        }
    };

    socket.onclose = function (event) {
        if (event.wasClean) {
            console.log(`연결이 정상적으로 종료되었습니다. 코드: ${event.code}, 이유: ${event.reason}`);
        } else {
            console.log('연결이 비정상적으로 종료되었습니다.');
        }
    };

    socket.onerror = function (error) {
        console.error(`WebSocket 에러 발생: ${error.message}`);
    };

    function sendMessage(message) {
        if (socket.readyState === WebSocket.OPEN) {
            socket.send(message);
        } else {
            console.log("WebSocket 연결이 열려있지 않습니다.");
        }
    }

    function closeConnection() {
        socket.close();
    }
}

function addComment_form(event) {
    if (event == -1 || event.key == "Enter") {
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
    const urladdress_copy = urladdress + "/" + sp_id;
    
    navigator.clipboard.writeText(urladdress_copy)
        .then(() => {
            showPopup("클립보드에 복사되었어요");
        })
        .catch(err => {
            console.error('클립보드 복사 실패:', err);
            showPopup("클립보드 복사에 실패했습니다");
        });
}

let pdfDoc = null
let pageRendering = false,
    pageNumPending = null,
    scale = 1.5;

function uploadPDF() {
    const file = document.getElementById('fileInput').files[0];
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
    const pdf_viewerDOM = document.getElementById('pdf-viewer')
    pdf_viewerDOM.style.display = 'block'
    pdfjsLib.getDocument(url).promise.then(function (pdf) {
        pdfDoc = pdf;
        document.getElementById('page-num').textContent = file_context + ' / ' + pdf.numPages;
        renderPage(file_context);
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
    let formData = new FormData();
    formData.append('file_context', num)
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
    if (file_context <= 1) {
        return;
    }
    file_context--;
    queueRenderPage(file_context);
}

function onNextPage() {
    if (file_context >= pdfDoc.numPages) {
        return;
    }
    file_context++;
    queueRenderPage(file_context);
}