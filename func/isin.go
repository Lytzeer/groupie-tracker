package groupie

func Isin(ele string, tab []string) bool {
	if tab == nil {
		return false
	} else {
		for _, element := range tab {
			if element == ele {
				return true
			}
		}
	}
	return false
}
