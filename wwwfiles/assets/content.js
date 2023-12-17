let last_update = ""
let urladdress = ""
let last_comment_id = -1
let qrsmallclick_toggle = false
let sortbyRate_toggle = false
let last_interaction = new Date()
let popup_once = false

function startContent(urladdress_i) {
    urladdress = urladdress_i

    // mousemove 이벤트를 핸들링하여 마지막 인터랙션 시각 저장
    const element = document.querySelector('#comments');
    element.addEventListener('mousemove', (event) => {
        last_interaction = new Date()
    });

    //즉시 수행 초기에 한번
    onLoad(urladdress)
    //인터벌 설정 1초
    setInterval(() => onLoad(urladdress), 1000)
}

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
function content_orderClick() {
    const comments = document.getElementById("comments")
    comments.textContent = "" //comments 아래 내용 삭제
    last_comment_id = -1
    last_update = ""
    sortbyRate_toggle = !sortbyRate_toggle
    sortbyRate_toggle_imdt = true
    onLoad(urladdress)
}
function qrsmallClick() {
    var device_width = document.body.clientWidth;
    var big_device = device_width <= 850 ? false : true;

    if (!qrsmallclick_toggle && big_device) { // 토글 f , 큰화면
        let main = document.querySelector("main")
        main.style.gridTemplateColumns = "1fr"
        let qrbig = document.getElementById("QRbig")
        qrbig.style.display = "none"
    } else if (qrsmallclick_toggle && big_device) { //토글 true, 큰 화면
        let main = document.querySelector("main")
        main.style.gridTemplateColumns = "1fr 3fr"
        let qrbig = document.getElementById("QRbig")
        qrbig.style.display = "grid"
    } else if (!qrsmallclick_toggle && !big_device) { //토글 false, 작은 화면
        let main = document.querySelector("main")
        main.style.gridTemplateColumns = "1fr"
        let qrbig = document.getElementById("QRbig")
        qrbig.style.display = "none"

        let qrsmall = document.getElementById("QRsmall")
        qrsmall.style.height = "50%"
        qrsmall.style.top = "30px"
        qrsmall.style.right = "30px"
        qrsmall.style.border = "1px solid black"
        qrsmall.style.borderRadius = "20px"

    } else if (qrsmallclick_toggle && !big_device) { //토글t, 작은 화면
        let main = document.querySelector("main")
        main.style.gridTemplateColumns = "1fr"
        let qrbig = document.getElementById("QRbig")
        qrbig.style.display = "none"

        let qrsmall = document.getElementById("QRsmall")
        qrsmall.style.height = "inherit"
        qrsmall.style.top = "0"
        qrsmall.style.right = "0"
        qrsmall.style.border = "none"
        qrsmall.style.borderRadius = "0"
    }
    qrsmallclick_toggle = !qrsmallclick_toggle
}

function writeComment_c(element) {
    const comment_c = document.createElement("section")
    comment_c.id = "comment_c"
    comment_c.setAttribute("Sp_c_id", element.Sp_c_id)
    comment_c_color = element.Sp_c_color
    comment_c.style.backgroundColor = ""
    switch (comment_c_color) {
        case -1: comment_c.style.backgroundColor = "white"; break;
        case 0: comment_c.style.backgroundColor = "#ffcc99"; break;
        case 1: comment_c.style.backgroundColor = "#99e6ff"; break;
        case 2: comment_c.style.backgroundColor = "#80ffcc"; break;
        case 3: comment_c.style.backgroundColor = "#ff99ff"; break;
        case 4: comment_c.style.backgroundColor = "#ff8080"; break;
        case 5: comment_c.style.backgroundColor = "#4d94ff"; break;
    }

    var li_content = document.createElement("li")
    li_content.id = "comment_c_content"
    li_content.textContent = element.Sp_c_content
    if (element.Sp_c_content.length > 12) { //길이가 길면 글자 크기 작게 출력
        li_content.style.fontSize = "1em"
    }

    var li_name = document.createElement("li")
    li_name.id = "comment_c_name"
    li_name.style.textDecoration = "underline"
    li_name.textContent = element.Sp_c_guestname

    var li_rate = document.createElement("li")
    li_rate.id = "comment_c_rate"
    li_rate.textContent = "❤️ " + element.Sp_c_rate

    var li_rate_like = document.createElement("li")
    li_rate_like.textContent = "👍"
    li_rate_like.onclick = function () {
        send_c_rate(true, element.Sp_c_id)
        onLoad(urladdress) //클릭시 즉시 갱신
    }
    li_rate_like.style.fontSize = "2em"
    li_rate_like.style.display = "none"

    var li_rate_dislike = document.createElement("li")
    li_rate_dislike.textContent = "👎"
    li_rate_dislike.onclick = function () {
        send_c_rate(false, element.Sp_c_id)
        onLoad(urladdress) //클릭시 즉시 갱신
    }
    li_rate_dislike.style.fontSize = "2em"
    li_rate_dislike.style.display = "none"
    //하트 누를 때
    li_rate.onclick = function () {
        li_rate_like.style.display = "block"
        li_rate_dislike.style.display = "block"
        comment_c.style.height = "11em";
    };
    //댓글 영역에서 마우스 호버 해제시 
    comment_c.onmouseleave = function () {
        li_rate_like.style.display = "none"
        li_rate_dislike.style.display = "none"
        comment_c.style.height = "7em";
    }


    comment_c.append(li_content, li_name, li_rate, li_rate_like, li_rate_dislike)

    document.getElementById("comments").appendChild(comment_c)
    last_comment_id = element.Sp_c_id
}

function sortByID(Sp_comments) {
    Sp_comments.forEach(element => {
        if (element.Sp_c_id > last_comment_id) {
            writeComment_c(element)
        }
    });
}

function sortByRate(Sp_comments) {
    const comments = document.getElementById("comments")
    comments.textContent = "" //comments 아래 내용 삭제
    Sp_comments.sort((a, b) => {
        if (a.Sp_c_rate < b.Sp_c_rate) {
            return 1;
        } else if (a.Sp_c_rate > b.Sp_c_rate) {
            return -1;
        } else {
            return 0;
        }
    });
    Sp_comments.forEach(element => {
        writeComment_c(element)
    });

}

function send_c_rate(like, comment_c_id) {
    const xhr = new XMLHttpRequest();
    const url = urladdress + "/space/comment_rate";
    const data = new FormData();

    data.append("comment_c_id", comment_c_id);

    if (like) {
        data.append("like_dislike", "like");
    } else {
        data.append("like_dislike", "dislike");
    }

    xhr.open("POST", url, true);
    xhr.send(data);
}

function rateRefresh_on_DOM(Sp_comments) {
    const comments = document.getElementById("comments")
    for (comments_element of comments.children) {
        comments_element_c_id = comments_element.getAttribute("Sp_c_id")
        Sp_comments.forEach(Sp_comments_element => {
            if (comments_element_c_id == Sp_comments_element.Sp_c_id) {
                comments_element.querySelector("#comment_c_rate").textContent = "❤️ " + Sp_comments_element.Sp_c_rate
            }
        });
    }

}

function onLoad() {
    //json 받아서 쓰기
    const xhr = new XMLHttpRequest();
    xhr.open("GET", urladdress + "/space/json");
    xhr.onload = function () {
        if (xhr.status == 200) {
            const jsonData = JSON.parse(xhr.responseText)

            //조회수 갱신
            p_content_view_count = document.getElementById("content_view_count")
            p_content_view_count.textContent = jsonData.Sp_view

            //rate 갱신
            rateRefresh_on_DOM(jsonData.Sp_comments)

            //마지막 인터랙션 수행 후 1이하면 리턴 1초 이상이어야 아래 수행 //실험 중인 기능 
            if (last_interaction.getTime() + 1000 > Date.now()) {
                return
            }

            //time에 변동이 있을 때만 그리기
            if (last_update == jsonData.Sp_lastupdate) {
                return
            }

            if (jsonData.Sp_comments != null) {//댓글이 존재한다는 뜻
                popup(false)

                if (sortbyRate_toggle) { // 좋아요순대로 정렬
                    sortByRate(jsonData.Sp_comments)
                } else { //ID순 대로 정렬
                    //부드럽게 맨 아래로 이동
                    window.scrollTo({ top: document.body.scrollHeight, behavior: "smooth", duration: 5000 })

                    sortByID(jsonData.Sp_comments)
                }

                last_update = jsonData.Sp_lastupdate
            } else { // //Sp_comments 널 일 때
                if (last_comment_id == -1) {//댓글이 없을 때 
                    //첫 생성시 데이터 안불러와서 기본값 넣기
                    p_content_view_count = document.getElementById("content_view_count")
                    p_content_view_count.textContent = 0
                    //팝업 생성
                    popup(true, "empty") // no comment 
                }

            }

        }
    };
    xhr.send()
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
        popup(true, "Copied to clipboard")
        setTimeout(function () {
            popup(false);
        }, 1000);
    } catch (err) {
        console.error('Failed to copy text', err);
    }

    document.body.removeChild(textArea); // DOM에서 텍스트 영역 제거
}
