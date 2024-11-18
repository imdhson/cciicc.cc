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
