let urladdress = ""
let qrsmallclick_toggle = false
let uploadareaclock_toggle = false
let popup_once = false
let file_context = 1
let user_isHost = false

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
    const isExpanded = uploadArea.classList.contains('show');
    uploadToggle.setAttribute('aria-expanded', isExpanded);
    if(isExpanded){
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
            user_isHost ? null : showPopup("호스트가 파일 변경 중...")
            loadPDF("/space/file")
        }
        if (jsonData.Sp_ws_type == 'chat') {
            floatingMessage(jsonData.Sp_c_content)
        }
        if (jsonData.Sp_ws_type == 'media_sync') {
            if (!user_isHost) {
                const mediaElement = document.querySelector('#media-viewer video, #media-viewer audio');
                if (mediaElement) {
                    if (Math.abs(mediaElement.currentTime - jsonData.CurrentTime) > 0.5) {
                        mediaElement.currentTime = jsonData.CurrentTime;
                    }
                    if (jsonData.IsPaused && !mediaElement.paused) {
                        mediaElement.pause();
                    } else if (!jsonData.IsPaused && mediaElement.paused) {
                        mediaElement.play();
                    }
                }
            }
        }
        if (jsonData.Sp_ws_type == 'emoji') {
            floatEmoji(jsonData.Emoji)
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
        xhr.setRequestHeader('X-CSRF-Token', document.querySelector('meta[name="csrf-token"]').getAttribute('content'));
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

function uploadFile() {
    const file = document.getElementById('fileInput').files[0];
    if (!file) return;

    const allowedExtensions = ['.pdf', '.png', '.jpg', '.jpeg', '.gif', '.webp', '.mp3', '.wav', '.ogg', '.mp4', '.webm'];
    const fileName = file.name.toLowerCase();
    const isAllowed = allowedExtensions.some(ext => fileName.endsWith(ext));

    if (!isAllowed) {
        showPopup('지원하지 않는 파일 형식입니다.');
        return;
    }

    const uploadBtn = document.getElementById('uploadButton');
    const originalText = uploadBtn.textContent;
    uploadBtn.disabled = true;
    uploadBtn.textContent = '업로드 중...';
    uploadBtn.setAttribute('aria-busy', 'true');

    const formData = new FormData();
    formData.append('file', file);
    fetch('/space/addfile', {
        method: 'POST',
        headers: {
            'X-CSRF-Token': document.querySelector('meta[name="csrf-token"]').getAttribute('content')
        },
        body: formData
    }).then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    }).then(data => {
        if (data.success) {
            // The websocket message will trigger loadPDF or loadMedia
            uploadToggle_onclick(); // Close panel only on success
        } else {
            showPopup('업로드에 실패했습니다.');
        }
    }).catch(error => {
        showPopup('업로드 중 오류가 발생했습니다.');
        console.error('Error:', error);
    }).finally(() => {
        uploadBtn.disabled = false;
        uploadBtn.textContent = originalText;
        uploadBtn.removeAttribute('aria-busy');
    });
}

function loadMedia(url, status) {
    const pdf_viewerDOM = document.getElementById('pdf-viewer');
    const media_viewerDOM = document.getElementById('media-viewer');
    pdf_viewerDOM.style.display = 'none';
    media_viewerDOM.style.display = 'block';

    let content = '';
    const cacheBuster = `?t=${new Date().getTime()}`;
    if (status == 3) {
        content = `<img src="${url}${cacheBuster}" style="max-width: 100%; height: auto;" />`;
    } else if (status == 2 || status == 4) {
        const tag = status == 2 ? 'audio' : 'video';
        // Remove native controls, add our custom ones if host
        content = `<${tag} ${user_isHost ? '' : 'autoplay'} src="${url}${cacheBuster}" style="max-width: 100%; width: 100%; height: auto;"></${tag}>`;

        if (user_isHost) {
            content += `
            <div id="custom-media-controls" class="custom-media-controls">
                <button id="play-pause-btn" class="media-btn" aria-label="재생">▶️</button>
                <input type="range" id="progress-bar" class="progress-bar" value="0" step="0.1" min="0" aria-label="재생 진행률">
                <span id="time-display">0:00 / 0:00</span>
            </div>`;
        }
    }

    media_viewerDOM.innerHTML = content;

    if (user_isHost && (status == 2 || status == 4)) {
        setupHostMediaControls();
    }
}

function setupHostMediaControls() {
    const mediaElement = document.querySelector('#media-viewer video, #media-viewer audio');
    const playPauseBtn = document.getElementById('play-pause-btn');
    const progressBar = document.getElementById('progress-bar');
    const timeDisplay = document.getElementById('time-display');

    if (!mediaElement) return;

    // Sync to server function
    let lastSyncTime = 0;
    const syncMediaState = (force = false) => {
        const now = Date.now();
        if (!force && now - lastSyncTime < 500) return; // limit sync rate
        lastSyncTime = now;

        let formData = new FormData();
        formData.append('currentTime', mediaElement.currentTime);
        formData.append('isPaused', mediaElement.paused);
        fetch('/space/mediasync', {
            method: 'POST',
            headers: {
                'X-CSRF-Token': document.querySelector('meta[name="csrf-token"]').getAttribute('content')
            },
            body: formData
        }).catch(e => console.error("Media sync error", e));
    };

    const formatTime = (time) => {
        const minutes = Math.floor(time / 60);
        const seconds = Math.floor(time % 60);
        return `${minutes}:${seconds < 10 ? '0' : ''}${seconds}`;
    };

    mediaElement.addEventListener('loadedmetadata', () => {
        progressBar.max = mediaElement.duration;
        timeDisplay.textContent = `${formatTime(mediaElement.currentTime)} / ${formatTime(mediaElement.duration)}`;
    });

    mediaElement.addEventListener('timeupdate', () => {
        progressBar.value = mediaElement.currentTime;
        timeDisplay.textContent = `${formatTime(mediaElement.currentTime)} / ${formatTime(mediaElement.duration)}`;
        if (!mediaElement.paused) syncMediaState();
    });

    playPauseBtn.addEventListener('click', () => {
        if (mediaElement.paused) {
            mediaElement.play();
            playPauseBtn.textContent = '⏸️';
            playPauseBtn.setAttribute('aria-label', '일시정지');
        } else {
            mediaElement.pause();
            playPauseBtn.textContent = '▶️';
            playPauseBtn.setAttribute('aria-label', '재생');
        }
        syncMediaState(true);
    });

    progressBar.addEventListener('input', () => {
        mediaElement.currentTime = progressBar.value;
        syncMediaState(true);
    });

    mediaElement.addEventListener('play', () => {
        playPauseBtn.textContent = '⏸️';
        playPauseBtn.setAttribute('aria-label', '일시정지');
        syncMediaState(true);
    });

    mediaElement.addEventListener('pause', () => {
        playPauseBtn.textContent = '▶️';
        playPauseBtn.setAttribute('aria-label', '재생');
        syncMediaState(true);
    });
}


function loadPDF(url) {
    const pdf_viewerDOM = document.getElementById('pdf-viewer')
    pdf_viewerDOM.style.display = 'block'

    // 최신 PDF.js 라이브러리 버전 사용
    pdfjsLib.GlobalWorkerOptions.workerSrc = '/assets/lib/pdf.worker.min.js';
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
        const ctx = canvas.getContext('2d', { willReadFrequently: true });  // willReadFrequently 속성 추가
        const viewport = page.getViewport({ scale: 2 });  // scale 값 2로 변경
        canvas.height = viewport.height;
        canvas.width = viewport.width;

        const renderContext = {
            canvasContext: ctx,
            viewport: viewport,
            background: "rgba(0,0,0,0)",
            annotationLayer: {},
            optionalContentConfigPromise: undefined,
            annotationCanvasMap: undefined,
            renderInteractiveForms: false,
            enableWebGL: false,
            useOnlyCssZoom: true,
            maxCanvasPixels: 16777216,
            pageColors: undefined,
            eventBus: undefined,
            renderingQueue: undefined,
            textLayerMode: 0,  // DISABLE
            removePageBorders: false,  
            renderer: "canvas",
            disableFontFace: false,
            useSystemFonts: false,
            rotatePages: false  // rotatePages 옵션 설정
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
        headers: {
            'X-CSRF-Token': document.querySelector('meta[name="csrf-token"]').getAttribute('content')
        },
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

function ImHost(){
    user_isHost = true
}


function exportData() {
    const exportBtn = document.getElementById('exportButton');
    const originalText = exportBtn ? exportBtn.textContent : '';
    if (exportBtn) {
        exportBtn.disabled = true;
        exportBtn.textContent = '다운로드 중...';
        exportBtn.setAttribute('aria-busy', 'true');
    }

    fetch('/space/export', {
        method: 'GET',
        headers: {
            'X-CSRF-Token': document.querySelector('meta[name="csrf-token"]').getAttribute('content')
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Export failed');
        }
        return response.blob();
    })
    .then(blob => {
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.style.display = 'none';
        a.href = url;
        a.download = 'space_record.json';
        document.body.appendChild(a);
        a.click();
        window.URL.revokeObjectURL(url);
        a.remove();
        showPopup("기록 다운로드가 완료되었습니다.");
    })
    .catch(error => {
        console.error('Error:', error);
        showPopup("기록 다운로드에 실패했습니다.");
    })
    .finally(() => {
        if (exportBtn) {
            exportBtn.disabled = false;
            exportBtn.textContent = originalText;
            exportBtn.removeAttribute('aria-busy');
        }
    });
}

function sendEmoji(emojiChar) {
    let formData = new FormData();
    formData.append('emoji', emojiChar);
    fetch('/space/addemoji', {
        method: 'POST',
        headers: {
            'X-CSRF-Token': document.querySelector('meta[name="csrf-token"]').getAttribute('content')
        },
        body: formData
    })
    .catch(error => {
        console.error('Error sending emoji:', error);
    });
}

function floatEmoji(emojiChar) {
    const emojiEl = document.createElement('div');
    emojiEl.textContent = emojiChar;
    emojiEl.className = 'floating-emoji';

    // 화면 우측 하단에서 랜덤한 위치 지정
    const rightOffset = Math.random() * 20 + 5; // 5% ~ 25% from right
    emojiEl.style.right = `${rightOffset}%`;

    // 흔들림 효과를 위한 변수 설정
    const xOffset = (Math.random() - 0.5) * 50; // -25px ~ 25px
    emojiEl.style.setProperty('--x-offset', `${xOffset}px`);

    document.body.appendChild(emojiEl);

    setTimeout(() => {
        emojiEl.remove();
    }, 3000); // CSS 애니메이션 시간과 동일하게 설정
}

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

function floatingMessage(text) {
    const messageElement = document.createElement('div');
    messageElement.textContent = text;
    messageElement.className = 'floating-message';
    
    document.getElementById('message-container').appendChild(messageElement);
    
    // 랜덤한 수평 및 수직 위치 오프셋 적용
    const horizontalOffset = Math.random() * window.innerWidth * 0.8; // 화면 너비의 최대 80%까지 랜덤 오프셋
    messageElement.style.right = `${horizontalOffset}px`;
    
    const verticalOffset = Math.random() * window.innerHeight * 0.8; // 화면 높이의 최대 80%까지 랜덤 오프셋
    messageElement.style.bottom = `${verticalOffset}px`;
    
    setTimeout(() => {
        messageElement.remove();
    }, 15000); // 애니메이션 시간과 동일하게 설정
}