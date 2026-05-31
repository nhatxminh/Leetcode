package main

func rotateString(s string, goal string) bool {
	for range s {
		rotate(&s)
		if s == goal {
			return true
		}
	}
  return false
}

func rotate(s *string) {
	arr := []rune(*s)
	arr = append(arr, arr[0])
	arr = arr[1:]
	*s = string(arr)
}