package police

func fontSelected(args []string) string {
	if len(args) < 2 {
		return "standard"
	} else {
		return args[1]
	}
}

func Choice(args []string) string {
	font := fontSelected(args)

	switch font {
	case "standard":
		return "./standard.txt"
	case "shadow":
		return "./shadow.txt"
	case "thinkertoy":
		return "./thinkertoy.txt"
	default:
		return "./standard.txt"
	}
}
