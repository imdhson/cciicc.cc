package types

type Sp_ws_type_chat struct {
	Sp_ws_type     string
	Sp_c_content   string
	Sp_c_guestname string
}

type Sp_ws_type_file_new struct {
	Sp_ws_type   string
	Sp_file_type Sp_file_status
}

type Sp_ws_type_file_context struct {
	Sp_ws_type      string
	Sp_file_context string
	Sp_file_type    Sp_file_status
}

func New_Sp_ws_type_chat(sp_c_guestname string, sp_c_content string) Sp_ws_type_chat {
	return Sp_ws_type_chat{
		Sp_ws_type:     "chat",
		Sp_c_guestname: sp_c_guestname,
		Sp_c_content:   sp_c_content,
	}
}

func New_Sp_ws_type_file_new(sp_file_type Sp_file_status) Sp_ws_type_file_new {
	return Sp_ws_type_file_new{
		Sp_ws_type:   "file_new",
		Sp_file_type: sp_file_type,
	}
}

func New_Sp_ws_type_file_context(sp_file_type Sp_file_status, sp_file_context string) Sp_ws_type_file_context {
	return Sp_ws_type_file_context{
		Sp_ws_type:      "file_context",
		Sp_file_context: sp_file_context,
		Sp_file_type:    sp_file_type,
	}
}
