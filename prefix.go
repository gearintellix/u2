package u2

// ScanPrefix to get all u2 binding with a prefix
func ScanPrefix(q string, prefixs []string) (items map[string][]string) {
	items = make(map[string][]string)

	for _, v := range prefixs {
		items[v] = []string{}
	}

	qr := []rune(q)
	for _, v := range prefixs {
		i := 0
		for {
			i = index(q, "__"+v, i)
			if i >= 0 {
				if i > 0 && string(qr[i-1]) == "\\" {
					continue
				}

				i += 2 + len(v)
				for {
					ii := index(q, "__", i)
					if ii > i {
						if ii > 0 && string(qr[ii-1]) == "\\" {
							i = ii + 2
							continue
						}

						items[v] = append(items[v], subStr(q, i, ii-i))
						i = ii + 2
					}
					break
				}
				continue
			}
			break
		}
	}

	return items
}
