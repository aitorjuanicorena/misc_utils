package misc_utils

func excelPos2Col(pos int, initial_pos rune) string {

	// initial_pos: Columna en la cual quieres empezar
	// pos: "Posición numérica". Empieza en 1, no en 0
	var int_column int = int(initial_pos) + (pos - 1)

	str_col := string(int_column)

	if int_column > 'Z' {

		var base rune = ('A') + rune((int_column-'Z'-1)/26)

		str_col = string(base) + string('A'+((int_column-'Z'-1)%26))

	}

	return str_col
}
