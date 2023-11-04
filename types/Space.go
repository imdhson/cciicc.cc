package types

type Spaces []Space

type Space struct {
	Sp_id       string
	Sp_name     string
	Sp_view     int
	Sp_comments []Sp_comment
}

type Sp_comment struct {
	Sp_c_content   string
	Sp_c_guestname string
	Sp_c_color     int
	Sp_c_comment   []Sp_comment
}
