
qrsmallclick_toggle = false
function qrsmallClick() {
    var device_width = document.body.clientWidth;
    if ( device_width <= 850 && qrsmallclick_toggle){
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