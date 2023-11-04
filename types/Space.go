package types

type Spaces []Space

func (spaces *Spaces) isValid() bool {
	if len(*spaces) > 0 {
		return true
	} else {
		return false
	}
}

type Space struct {
	sp_id       string
	sp_name     string
	sp_view     int
	sp_comments []Sp_comment
}

type Sp_comment struct {
	sp_c_content   string
	sp_c_guestname string
	sp_c_color     int
	sp_c_comment   []Sp_comment
}
