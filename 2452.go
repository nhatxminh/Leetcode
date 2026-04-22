package main

func twoEditWords(queries []string, dictionary []string) []string {
	result := make([]string, 0)
	correct := 0
	incorrect := 0
	skip := 0

	for _, query := range queries {
        if len(query) <= 2 {
            result = append(result, query)
            continue
        }
		for _, value := range dictionary {
            if skip == 1 {
				break
			}
			queryRune := []rune(query)
			valueRune := []rune(value)
			correct = 0
			incorrect = 0

			for i := range len(queryRune) {
				if queryRune[i] == valueRune[i] {
					correct++
					
					if len(queryRune) - correct <= 2 {
						result = append(result, query)
						skip = 1
						break
					}
				} else {
					incorrect++
					if incorrect > 2 {
						break
					}
				}
			}
		}
        skip = 0
	}

  return result
}