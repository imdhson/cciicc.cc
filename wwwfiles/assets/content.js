let last_update = ""
let last_comment_id = -1
qrsmallclick_toggle = false


function qrsmallClick() {
    var device_width = document.body.clientWidth;
    if (device_width <= 850 && qrsmallclick_toggle) {
        //디바이스가 850px 미만이고 
        // 버튼이 눌러져서 트루 상태일때는 원복을 위해 한번작동함
        return
    }
    if (!qrsmallclick_toggle) {
        let main = document.querySelector("main")
        main.style.gridTemplateColumns = "1fr"
        let qrbig = document.getElementById("QRbig")
        qrbig.style.display = "none"
    } else {
        let main = document.querySelector("main")
        main.style.gridTemplateColumns = "1fr 2fr"
        let qrbig = document.getElementById("QRbig")
        qrbig.style.display = "grid"
    }
    qrsmallclick_toggle = !qrsmallclick_toggle
}

function onLoad(urladdress) {
    //json 받아서 쓰기
    const xhr = new XMLHttpRequest();
    xhr.open("GET", urladdress + "/space/json");
    xhr.onload = function () {
        if (xhr.status == 200) {
            const jsonData = JSON.parse(xhr.responseText)

            //time에 변동이 있을 때만 그리기
            if (last_update == jsonData.Sp_lastupdate) {
                return
            }
            if (jsonData.Sp_comments != null) {
                jsonData.Sp_comments.forEach(element => {
                    if (element.Sp_c_id > last_comment_id) {
                        const comment_c = document.createElement("section")
                        comment_c.id = "comment_c"
                        var li_content = document.createElement("li")
                        li_content.textContent = element.Sp_c_content
                        var li_name = document.createElement("li")
                        li_name.textContent = element.Sp_c_guestname
                        var li_rate = document.createElement("li")
                        li_rate.textContent = element.Sp_c_rate
                        comment_c.append(li_content, li_name, li_rate)
                        document.getElementById("comments").appendChild(comment_c)

                        last_comment_id = element.Sp_c_id
                        last_update = jsonData.Sp_lastupdate
                    }
                });
            } else{ // 댓글이 없을 때 
                if (last_comment_id == -1) {
                    //첫 생성시 데이터 안불러와서 기본값 넣기
                    p_content_view_count = document.getElementById("content_view_count")
                    p_content_view_count.textContent = 0
                    const popup = document.createElement("div")
                    popup.id = "popup"
                    popup.style.position = "fixed"
                    popup.style.display = "block"
                    popup.textContent = "no comment"
                    popup.style.top = "20%"
                    popup.style.left = "20%"
                    popup.style.fontSize = "10em"
                    popup.style.background = "white"
                    const main = document.querySelector("header")
                    main.appendChild(popup)
                } else {
                    const popup = document.getElementById("popup")
                    popup.style.display = "none"
                }

            }
            //조회수 갱신
            p_content_view_count = document.getElementById("content_view_count")
            p_content_view_count.textContent = jsonData.Sp_view
        }
    };
    xhr.send()
}

function intervel(urladdress) {
    //인터벌 설정 5초
    onLoad(urladdress)
    setInterval(() => onLoad(urladdress), 1000)
}


/*
const xhr = new XMLHttpRequest();
const url = "https://example.com";
const data = new FormData();
data.append("key1", "value1");
data.append("key2", "value2");

xhr.open("POST", url, true);
xhr.send(data);

*/