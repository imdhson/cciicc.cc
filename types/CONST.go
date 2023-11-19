package types

type Sp_c_color int

const (
	WHILE  Sp_c_color = -1
	ORANGE Sp_c_color = iota
	SKYBLUE
	GREEN
	PINK
	RED
	BLUE
)

const (
	SESSION_EXPIRE_DURATION_HOURS        = 24
	URL_ADDESS                    string = "http://localhost"
	//use for QR generation, html linking
)

const ( //KR
	SERVICE_NAME   string = "이름테스트"
	SERVICE_DETAIL string = "서비스설명ㅁㅁㅁㅁㅁㅁ"
	FOOTER_TERMS   string = "서비스 이용시 약관에 동의하고 쿠키 사용에 동의하는 것으로 간주됨."

	MAIN_NEXT_UP string = "이어서 시작하기"
	MAIN_HOST    string = "새로운 space 생성하기"
	MAIN_GUEST   string = "코드 입력하여 시작하기"

	HOST_DETAIL          string = "Host로 시작하기"
	HOST_SPACENAME       string = "Space 이름"
	HOST_SPACENAME_INPUT string = "Space 이름을 지정하세요. 예) 창의문제해결 아이디어"
	HOST_USERNAME        string = "당신의 닉네임"
	HOST_USERNAME_INPUT  string = "당신의 닉네임을 입력하세요"
	HOST_FORM_BUTTON     string = "space 생성하기"

	GUEST_DETAIL         string = "Guest로 시작하기"
	GUEST_SPACEID        string = "Space ID 입력"
	GUEST_SPACEID_INPUT  string = "Host가 제공한 Space ID를 입력하세요. 예) "
	GUEST_USERNAME       string = "당신의 닉네임"
	GUEST_USERNAME_INPUT string = "당신의 닉네임을 입력하세요"
	GUEST_FORM_BUTTON    string = "space 입장하기"
)
