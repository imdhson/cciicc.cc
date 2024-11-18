package controller

import (
	"cciicc/service"
	"cciicc/types"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func PostHandler_file(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//세션 가져오기
	var user types.User
	var user_success bool
	session, getcookie_err := r.Cookie("ub_session")
	if getcookie_err != nil {
		http.Error(w, getcookie_err.Error(), http.StatusBadRequest)
		return
	} else {
		user, user_success = service.GetUserFromSession(session.Value)
	}

	if !user_success { //user success false일 때 or space_id 불일치시
		http.Error(w, "user을 찾을 수 없음", http.StatusBadRequest)
		return
	}

	// 최대 2GB 파일 크기 제한
	r.ParseMultipartForm(2 << 30)

	file, filehandlerFormFile, err := r.FormFile("file")
	if err != nil {
		service.ErrHandler(err, "post handler file ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 서버에 파일 생성 wwwfiles/host_file/ space_id.확장자
	dst, err := os.Create("wwwfiles/host_file/" + user.User_related_spaceid + filepath.Ext(filehandlerFormFile.Filename))
	if err != nil {
		service.ErrHandler(err, "post hanlder os create")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	//파일시스템에 복사
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Content-Type 헤더 설정
	w.Header().Set("Content-Type", "application/json")
	// HTTP 상태 코드 설정 (선택사항)
	w.WriteHeader(http.StatusOK)
	// JSON 인코딩 및 응답 작성
	// 응답 데이터 구조체 정의
	response := struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}

	space, space_success := service.GetSpaceFrom_space_id(user.User_related_spaceid)
	if !space_success {
		http.Error(w, "unable to find user", http.StatusInternalServerError)
		return
	}

	//space.Sp_filestatus 변경
	log.Println(filepath.Ext(filehandlerFormFile.Filename))
	if filepath.Ext(filehandlerFormFile.Filename) == ".pdf" { //파일 확장자가 pdf일 경우
		space.Sp_file_status = types.SP_FILESTATUS_PDF
		space.Sp_file_ext = filepath.Ext(filehandlerFormFile.Filename)
	}

	//성공 쓰기
	json.NewEncoder(w).Encode(response)

}
