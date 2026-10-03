package controller

import (
	"cciicc/service"
	"cciicc/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getUserFromRequest(r *http.Request) (types.User, error) {
	var user types.User
	session, err := r.Cookie("ub_session")
	if err != nil {
		return user, err
	}

	user, success := service.GetUserFromSession(session.Value)
	if !success || !user.User_isHost {
		return user, fmt.Errorf("user not found or is not host")
	}
	return user, nil
}

func saveUploadedFile(r *http.Request, spaceID string) (string, error) {
	// 최대 2GB 파일 크기 제한
	r.ParseMultipartForm(2 << 30)

	file, filehandlerFormFile, err := r.FormFile("file")
	if err != nil {
		service.ErrHandler(err, "post handler file ")
		return "", err
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(filepath.Base(filehandlerFormFile.Filename)))
	if ext != ".pdf" {
		return "", fmt.Errorf("invalid file extension")
	}

	// 서버에 파일 생성 wwwfiles/host_file/ space_id.확장자
	dst, err := os.Create("wwwfiles/host_file/" + spaceID + ext)
	if err != nil {
		service.ErrHandler(err, "post hanlder os create")
		return "", err
	}
	defer dst.Close()

	//파일시스템에 복사
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return ext, nil
}

func updateSpaceAndBroadcast(spaceID string, ext string) error {
	space, success := service.GetSpaceFrom_space_id(spaceID)
	if !success {
		return fmt.Errorf("unable to find space")
	}

	//space.Sp_filestatus 변경
	if ext == ".pdf" { //파일 확장자가 pdf일 경우
		space.Sp_file_status = types.SP_FILESTATUS_PDF
		space.Sp_file_ext = ext
	}

	//같은 ws_space에 websocket broadcast 시도
	ws_hub := types.GetInstance_ws_hub()
	ws_space := ws_hub.Ws_GetOrCreateSpace(spaceID)
	ws_file_context := types.New_Sp_ws_type_file_context(space.Sp_file_status, "1")
	ws_file_context_encoded, err := json.MarshalIndent(ws_file_context, " ", "\t")
	if err != nil {
		service.ErrHandler(err, "posthandler file context json")
		return err
	}

	ws_hub.Broadcast([]byte(ws_file_context_encoded), ws_space)
	return nil
}

func PostHandler_file(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user, err := getUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ext, err := saveUploadedFile(r, user.User_related_spaceid)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "invalid file extension" || strings.HasPrefix(err.Error(), "http: no such file") {
			status = http.StatusBadRequest
		}
		http.Error(w, err.Error(), status)
		return
	}

	// Content-Type 헤더 설정
	w.Header().Set("Content-Type", "application/json")
	// HTTP 상태 코드 설정 (선택사항)
	w.WriteHeader(http.StatusOK)
	// JSON 인코딩 및 응답 작성
	response := struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}
	json.NewEncoder(w).Encode(response)

	err = updateSpaceAndBroadcast(user.User_related_spaceid, ext)
	if err != nil {
		// Log the error but don't change the HTTP response as we already sent StatusOK
		service.ErrHandler(err, "update space and broadcast failed")
	}
}
