package turn

func ChangePlayer(player int) int {
	if player == 1 {
		return 2
	}
	return 1
}
