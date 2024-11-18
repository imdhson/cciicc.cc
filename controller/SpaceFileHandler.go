package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"cciicc/types"

	"cciicc/service"
)

func SpaceFileHandler(w http.ResponseWriter, r *http.Request) {
	var user types.User
	var user_success bool
	session, getcookie_err := r.Cookie("ub_session")
	if getcookie_err != nil {
		service.ErrHandler(getcookie_err, "getcookie_err")
		w.WriteHeader(http.StatusForbidden)
		return
	} else { //오류 없을 때
		user, user_success = service.GetUserFromSession(session.Value)
	}

	if !user_success { //user success false일 때
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var space_filepath string
	space, _ := service.GetSpaceFrom_space_id(user.User_related_spaceid)
	if space.Sp_file_status == types.SP_FILESTATUS_PDF {
		space_filepath = filepath.Clean("wwwfiles/host_file/" + user.User_related_spaceid + space.Sp_file_ext)
	}

	// 파일 열기
	file, err := os.Open(space_filepath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	defer file.Close()

	//파일 정보 가져오기
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// Content-Disposition 헤더 설정 (선택사항)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name()))
	// Content-Type 설정 (필요에 따라 적절히 변경)
	w.Header().Set("Content-Type", "application/octet-stream")
	// Content-Length 설정
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
	// 파일 내용 전송
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}
