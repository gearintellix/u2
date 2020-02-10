package u2

import (
	"fmt"
	"strings"

	"github.com/gearintellix/serr"
)

// TagInfo object
type TagInfo struct {
	Tag   string
	Key   string
	Index string
	Value string
	Meta  map[string]string
}

// ScanTags to get all u2tag binding
func ScanTags(q string, tag string) (nq string, tags []TagInfo, errx serr.SErr) {
	tags = []TagInfo{}

	nq = ""
	i, io, it, qr := 0, 0, -1, []rune(q)

	for {
		it++

		qkey := fmt.Sprintf("<%s:", tag)
		i = index(q, qkey, i)
		if i >= 0 {
			nq += subStr(q, io, i-io)

			ctag := TagInfo{
				Tag:  tag,
				Meta: make(map[string]string),
			}

			fns := []string{"[", "{", "/", ">"}
			cstx := ""

			ii := i + len(qkey)
			ic := ii

			meta := ""
			for {
				if cstx == "!" {
					break
				}
				ii++

				switch cstx {
				case "[":
					ctag.Key = trim(subStr(q, ic, ii-ic-1))

					iii := index(q, "]", ii)
					if iii > ii {
						ctag.Index = trim(subStr(q, ii, iii-ii))
						cstx = ""
						continue
					}

					errx = serr.Newc(fmt.Sprintf("no ending tag %s:%s on syntax '['", ctag.Tag, ctag.Key), "while parsing script")
					return nq, tags, errx

				case "{\"":
					if string(qr[ii]) == "\"" {
						ii++
						cstx = "{"
						continue
					}

					iii := index(q, "\"", ii)
					if iii > ii {
						if len(q) >= iii+1 && string(qr[iii+1]) == "\"" {
							ii = iii + 1
							continue
						}
						ii = iii + 1
						cstx = "{"
						continue
					}

					errx = serr.Newc(fmt.Sprintf("no ending meta tag %s:%s on syntax '\"'", ctag.Tag, ctag.Key), "while parsing script")
					return nq, tags, errx

				case "{":
					if meta == "" {
						fns = []string{"/", ">"}
						if ctag.Key == "" {
							ctag.Key = trim(subStr(q, ic, ii-ic-1))
						}
						meta = ">"
						ic = ii
					}

					switch string(qr[ii]) {
					case "\"":
						cstx = "{\""

					case "}":
						meta = subStr(q, ic, ii-ic-1)
						cstx = ""
					}

				case "/":
					if string(qr[ii]) == ">" {
						if ctag.Key == "" {
							ctag.Key = trim(subStr(q, ic, ii-ic-1))
						}

						ii++
						cstx = "!"
						continue
					}

					errx = serr.Newc(fmt.Sprintf("unknown syntax %s on %s:%s", string(qr[ii]), ctag.Tag, ctag.Key), "while parsing script")
					return nq, tags, errx

				case ">":
					if ctag.Key == "" {
						ctag.Key = trim(subStr(q, ic, ii-ic-1))
					}

					qkey := fmt.Sprintf("</%s:%s>", ctag.Tag, ctag.Key)
					iii := index(q, qkey, ii)
					if iii > ii {
						ctag.Value = subStr(q, ii, iii-ii)
						ii = iii + len(qkey)
						cstx = "!"
						continue
					}

					errx = serr.Newc(fmt.Sprintf("no ending of tag %s:%s", ctag.Tag, ctag.Key), "while parsing script")
					return nq, tags, errx

				default:
					if len(qr) <= ii {
						cstx = "!"
						continue
					}
					if isArrExists(string(qr[ii]), fns) {
						cstx = string(qr[ii])
					}
				}
			}

			if meta != "" {
				metas := strings.Split(meta, "\n")

				done := false
				quote := false
				temp := []string{"", ""}

				for _, v := range metas {
					vr := []rune(v)
					iii := 0

					skip := false
					for {
						if skip || iii == -1 || (!done && len(v) <= (iii+1)) {
							break
						}

						if !done && !quote {
							iiy := index(v, ":", iii)
							if iiy > iii {
								temp[0] = trim(subStr(v, iii, iiy-iii))
								iii = iiy + 1

								iiii := iii
								for {
									if quote || done || len(v) <= iiii {
										break
									}

									switch string(vr[iiii]) {
									case "\"":
										iii = iiii + 1
										quote = true

									case ";":
										temp[1] = trim(subStr(v, iii, iiii-iii))
										iii = iiii + 1
										done = true

									default:
										iiii++
									}
								}

								if !quote && !done {
									temp[1] = trim(subStr(v, iii, iiii-iii))
									iii = iiii + 1
									done = true
								}

							} else {
								iii = -1
							}
						}

						if !done && quote {
							iix := iii

							for {
								if done {
									break
								}

								iiii := index(v, "\"", iii)
								if iiii >= iii {
									if len(v) > iiii+1 && string(vr[iiii+1]) == "\"" {
										iii = iiii + 2
										continue
									}

									temp[1] += subStr(v, iix, iiii-iix)
									quote, done = false, true
									iii = iiii + 1

									iix = index(v, ";", iii)
									if iix >= iii {
										iii = iix + 1
									}
									break
								}
								temp[1] += subStr(v, iix, 0)
								skip = true
								break
							}
							continue
						}

						if done {
							isMeta := false
							if strings.HasPrefix(temp[0], "~~") {
								temp[0] = subStr(temp[0], 1, 0)
							} else if strings.HasPrefix(temp[0], "~") {
								temp[0] = subStr(temp[0], 1, 0)
								isMeta = true
							}

							switch true {
							case !isMeta && temp[0] == "key":
								ctag.Key = temp[1]

							case !isMeta && temp[0] == "index":
								ctag.Index = temp[1]

							case !isMeta && temp[0] == "value":
								if ctag.Value == "" {
									ctag.Value = temp[1]
								}

							default:
								ctag.Meta[temp[0]] = temp[1]
							}

							temp = []string{"", ""}
							done = false
							continue
						}
					}
				}
			}

			tags = append(tags, ctag)

			nq += fmt.Sprintf("__#%s:%d__", ctag.Tag, it)
			i, io = ii, ii

			continue
		}
		nq += subStr(q, io, i-io+1)
		break
	}

	return nq, tags, errx
}
