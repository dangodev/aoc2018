package main

import "fmt"

func offByOne(str1 string, str2 string) bool {
	if str1 == str2 {
		return false // if same string, don’t even bother
	}

	diffs := 0
	for i := range str1 {
		if str1[i] != str2[i] {
			diffs++
		}
	}

	return diffs == 1
}

func main() {
	boxIDs := []string{"bpacnmelhhzpygfsjoxtvkwuor", "biacnmelnizqygfsjoctvkwudr", "bpaccmllhizyygfsjoxtvkwudr", "rpacnmelhizqsufsjoxtvkwudr", "bfacnmelhizqygfsjoxtvwwudp", "bpacnmelhizqynfsjodtvkyudr", "bpafnmelhizqpgfsjjxtvkwudr", "bpackmelhizcygfsjoxtvkwudo", "bmacnmilhizqygfsjoltvkwudr", "bpafnmelhizuygfsjoxtvkwsdr", "boacnmylhizqygfsjoxtvxwudr", "bpbcjmelhizqygfsjoxtgkwudr", "bpacnmglhizqygfsjixtlkwudr", "bpacnmclhizqygfsjoxtvkwtqr", "bpacnmelhczqygtsjoptvkwudr", "bpacnmelhizqywfsaoxtvkbudr", "apacnmelhizqygcsjoxtvkwhdr", "bpacnmelrizqygfsbpxtvkwudr", "tpkcnmelpizqygfsjoxtvkwudr", "bpacnmelhizqlgfsjobtmkwudr", "npacnmelhizqygffjoxtvkwudf", "bpacnmeehqzqygqsjoxtvkwudr", "bpecnmelhizqigfsjvxtvkwudr", "bpacnmelhizqysfsjoxtvkdfdr", "bpacnfelhkzqygfsjoxtvkwfdr", "bpacnbelvizqygfsjoxthkwudr", "bpacnoelhizqygfejoxtvkwudn", "bpacnmelhizqygfzpkxtvkwudr", "bpahnmelhizqyufsjoxmvkwudr", "bpacnmelhizqygfsnoxtvkwmmr", "bpacnmelhizqygfsjoatvkludf", "bpacnmylhizqygfsjlxtvksudr", "bpacnmekhpzqygysjoxtvkwudr", "bpacnselhizqogfswoxtvkwudr", "bpacnmelhizqprfsjoxwvkwudr", "bpatnmelhinqygfsjoctvkwudr", "bpacnqelhqzqygfsxoxtvkwudr", "bpabnmelhiyqygfsjoxtykwudr", "bpacnivlhizqygfsjoxtviwudr", "bpkcnmylhizqygfsjoxtvkwcdr", "bpafnmflhizqygtsjoxtvkwudr", "bpachmelhizqygfsjixtvkwudg", "bpacymelhizqygfsjoxtykwuar", "bpacnkelhizqdgfsjoxtskwudr", "bpacnmezhizqggbsjoxtvkwudr", "bpacnmqlhizqygrsjoxzvkwudr", "bpaczmelhizqyhfsjoxfvkwudr", "bdacnmelhyzqygusjoxtvkwudr", "bpacbmelhizqywfsjostvkwudr", "bpacnmelhihzygfstoxtvkwudr", "bpactmelhizqygfsjcxtvkwydr", "bkacnmethizqytfsjoxtvkwudr", "bpacnmalhizqydfskoxtvkwudr", "spacnmelbizqygfsjoxdvkwudr", "lpalnmelhizoygfsjoxtvkwudr", "bpacjmeghizqygfsjoxtviwudr", "bpacnmeqhizxygfsjoxgvkwudr", "bpacnmelhizqygosjoxtvkkuhr", "bpacnmelhiznbxfsjoxtvkwudr", "bgacnmelhizqygfsjbxivkwudr", "bpacnmelhizqygfjjowtvswudr", "bpacnmelhizqygfsjovtgkmudr", "bpacnmelcmzqygfspoxtvkwudr", "bpvcnmelhizqyvfcjoxtvkwudr", "bpacnmeahizqjgfsjoxtvkwukr", "bpacnoelwizqygfsjoxtvkaudr", "xpacnmelhizqygfsjoxdvkwedr", "mpacnmelqizqygfsjoxtvkwudx", "bppcnmelhizqygfsjfxtvkhudr", "bpacnmclhizqyhfsjaxtvkwudr", "opacsmelhizqygfsjmxtvkwudr", "bpafnmelhizqjgfsjoxtvkrudr", "bpdcnmilhizqygfsjoxtvkludr", "bpainmelhizqygfsjtntvkwudr", "bradnmelhizqygfsjextvkwudr", "bpacnmelhizqygfmsoxtvkwudg", "bpacneelhizqygvrjoxtvkwudr", "bpacnpelhizqygfsjoxyvkwudf", "bpacnmelhizqygfsqoqtvkwodr", "bpacnmelhizjyghsjoxcvkwudr", "bpacnmelmibqygfsjoxtvnwudr", "jpacnmelaizqygfwjoxtvkwudr", "zpachmelhizqygfsjsxtvkwudr", "bpacnmelfizqykfsjomtvkwudr", "bpacnmllwizqygfsjoxtvkwusr", "bpaynmelhizqygfsjoxtvowcdr", "jpacnmqlhizqygfsjoxtvknudr", "bpacxmelhizqyffsjoxtvkwugr", "apawnmelhizqygfsjtxtvkwudr", "mpacnmelhitqigfsjoxtvkwudr", "bpacnmelhhzqygfsjoxtvkyzdr", "gpacnmelhizqynfsjoxtvkwudm", "bnacnkelhizqygfsjoxtpkwudr", "bpacnmelfizqygfsumxtvkwudr", "bpacnmelhisqygfsjohtvowudr", "bpacnmelhimqygxsjoxtvkwudn", "bpscnmeliizqygfsjoxtvkwunr", "qpacnmelhizqycfsjoxtvkwndr", "bpacnmelhijqygfsjohtvkyudr", "bpacnmelhizqykfsjkxtvknudr", "bpacnqilhizqygfsjoxtvkoudr", "bpacnmelhizqzgmsjoxtvkwurr", "bpdcnmelhizqygfsjoutukwudr", "bpecnmeghizqygfsjoxgvkwudr", "bpicnmelhizqygfrjoxtvlwudr", "bpacnmelhizfygfsroxtvkwodr", "buacnmelhizqygjsjoxtvkvudr", "bpacnmelhixqykfsjoxtvrwudr", "bpacnmelhizqygvejcxtvkwudr", "bpacnmjlhizqylfsjoxtvkwuor", "qpacnmelhizqygfsjoxfdkwudr", "bpfcnmemhizqygfsjoxtvknudr", "bpacnmelhizqoffsjqxtvkwudr", "hpacnielhiqqygfsjoxtvkwudr", "gpacnmelhizqygfsewxtvkwudr", "bpacnmellizqylxsjoxtvkwudr", "bpacnmenhizqymfsjoxtvkmudr", "bpacnfelhizqygcsjoltvkwudr", "bpacnmelhqqqygfsjoxtvkuudr", "bplgnmelhiqqygfsjoxtvkwudr", "bpacnzelhizqygfgjoxtvnwudr", "bpacnmelhizqygfsjoktvknunr", "bpacnmdlhioqygfnjoxtvkwudr", "epacnmelwizqyjfsjoxtvkwudr", "bpacxmelhazfygfsjoxtvkwudr", "bpacnmejhezqygfsjoxtskwudr", "bpacnqelhihqyzfsjoxtvkwudr", "bpacnbelhizqyrfsjoxtvkmudr", "bpacnmelhizqygfsjoxtylwzdr", "bpacnmelwizqygfsjodtvkhudr", "bpacnnelhizqygfsjoxtwkwadr", "bpacimelhizqygfsnoxtvkwuor", "bpacnmelhizqyaasjoxtlkwudr", "bpacnmelhizqyeffjoxtvkwuds", "bpacnmenhizqygxscoxtvkwudr", "bpacnmelhidqygfsjowtskwudr", "bpacnmeliizqygfsjoxhvkwucr", "bpacimelhizqygfsjoxtvktuwr", "bpainmelhhzqygfsjzxtvkwudr", "bpacamelhizqygfsjogtvkwbdr", "bpccnmelgizqygfsjoxtykwudr", "bpacnmelhizwegfsjoxtvkwadr", "bpackmelhbzqygqsjoxtvkwudr", "bpacymeihizqyffsjoxtvkwudr", "bpacnielhczqygfsjoxtvkwudk", "bpacnmejhizqygffjoxjvkwudr", "ppacnmelhizqygfsjoxtigwudr", "bpjcnmolhizqygfsjoxtvkwndr", "bpacnmelcizqygrsjoxtakwudr", "cpawnmelhizqygfsjoxmvkwudr", "bwacnmelhizqygesjoxtakwudr", "bpacnmelhizqygfsjexsvkwddr", "bpaunmelhiuqygfsjoxtvkwtdr", "bpacnmellimqygfsjextvkwudr", "bpacnmerhizqygfsaoxvvkwudr", "bpacnmglhizqygfsjixtukwudr", "ppacnmelhizqygfsjoxtvkdudp", "bpacnmedhizqygukjoxtvkwudr", "bpccnmelhizqngfsjoxtvkwadr", "bgacnmeldizqygfscoxtvkwudr", "bpacngelhizsygfsjoxtvkwkdr", "bpacnpelhizqygfsjoxctkwudr", "bpacnmylhizqygfcjoxtvkwmdr", "npacnmelhizqygfsjoxtwkwuds", "bpaxnmelhizqydfsjoxyvkwudr", "bpacnhelhizjygfsjoxtvkmudr", "bpacnkelhczqygfnjoxtvkwudr", "bfacnmelhizrygfsjoxtvkwodr", "bpycnmelhizqygfofoxtvkwudr", "qpacpselhizqygfsjoxtvkwudr", "bpvcnmelhezqygfsjoxttkwudr", "bpacnmwlhizqygfijoxtmkwudr", "bsacnmelhikqygfsjoxttkwudr", "bpccnxelhizqyafsjoxtvkwudr", "bpacnmelhizqygfswhxtvewudr", "vpacnmzlhizqygfsvoxtvkwudr", "bpacnmelhihqygfsjoxtvkqurr", "bpacnmelhixqygazjoxtvkwudr", "bpavnmelhizqygfsjozpvkwudr", "bpacnmclhizuygfsjoxmvkwudr", "bpacnmelhizryufsjoxtkkwudr", "bpacnmelhtzqygfsjobtvkwufr", "bpacnmelhizqmlfsjoxtvkwudq", "bpaaneelhizqygfsjlxtvkwudr", "bpacnmelhxzqygfsjoxthkwuhr", "bpacnmeshizqygfcjoxtvkwude", "bpacnzqlhizqygfsxoxtvkwudr", "bgaanmelhizqycfsjoxtvkwudr", "bpacnmexhizqygfsroxtvkwudn", "bpmmnmelhizqygfajoxtvkwudr", "bpacnmelhizqylfsjoxtckwhdr", "bpicnmelhizqyrfsjoxtvkwudi", "zpacnmelhizvycfsjoxtvkwudr", "bpamnmkllizqygfsjoxtvkwudr", "bpacnmelhrzqyrfsjoxgvkwudr", "bpadnmelhczqygfsjoxtlkwudr", "bpacrmelhizqygrsjoxtvkiudr", "lpacnmelhizqygfsjoxtgkwxdr", "fpacnmalhiuqygfsjoxtvkwudr", "bpacnmelhizqygfsjixtvfwcdr", "bpccnmelhxzqygfkjoxtvkwudr", "bpacnmepaizqygfsjoctvkwudr", "tpacnmelhivqygfsxoxtvkwudr", "kpacnfelhitqygfsjoxtvkwudr", "baacnzelhizqygfsjoxtvkwudx", "bcycnmeghizqygfsjoxtvkwudr", "wpacotelhizqygfsjoxtvkwudr", "bpacnmsshizqygrsjoxtvkwudr", "blacnmelhizqygfsjoxtykwvdr", "bkacnmelhizqygfsjoxuvkludr", "bpacnmelhizaugfsjoxtvhwudr", "fpavnmelhizqygfsgoxtvkwudr", "bpachmelnizqygfsjextvkwudr", "bpacnmelhizqpgfsjoxtvkwldu", "bpacnmelhizqygfsloftvywudr", "bpacntelhvzqygfejoxtvkwudr", "bpacnmeldizqygfsjmxtvkdudr", "byacnmelhizqygfsjsxtvkwudh", "bpacnmellizqygssxoxtvkwudr", "bpacnmelhizqygfsjootvknuir", "bpacnmelhitqjgfsjoxivkwudr", "bpacnmelhazaygfsjoxtvfwudr", "bpacnzenhizqygfsjzxtvkwudr", "bpacnmelhizqypfsdoxtvkwuar", "bpannmelhizqygnsjoxtvkwndr", "bracnmeldizsygfsjoxtvkwudr", "bpacnmelhizwygfsjugtvkwudr", "bpatnmelhizqygfsjoytvkwulr", "upacnmelhizqygfsjurtvkwudr", "bpaenmezhizqygfsjostvkwudr", "bpacnmelhizpygfsjodhvkwudr", "bpacnmelhizqygfsjogtvkguwr", "bpacnmelhisqygfsjoxtpkuudr", "bxacnmelhizqygfsjdxtvkfudr", "bpacnmelhizqygfsjohqvkwudu", "bzacnmtlhizqygfsjoxsvkwudr", "bpacnmplhixrygfsjoxtvkwudr", "bpacnmelhizqhgfsjomtvkwudg", "bpacnmezhizqygfsjxxtykwudr", "bpacnmwlhizqygfujoxtzkwudr", "tpacnmelhizqygfsjoxkvpwudr", "bpawsmenhizqygfsjoxtvkwudr", "bpacnmelhizqtgfsjoxttkwuqr", "bpkcbmelhizqygfsjoxtvkwucr", "bpacfmekhizqygfsjoxtvkwuds", "bpacnmethizqynfajoxtvkwudr", "bpocnmclhizqygfsjoxtvkwukr", "zpacnmwlhizqygfsjoxzvkwudr", "bpacpoelhqzqygfsjoxtvkwudr", "bpacnlelhizqyzfsjoxtvkwukr"}

	// part 1
	twice := map[string]string{}
	thrice := map[string]string{}

	for _, id := range boxIDs {
		letterCount := make(map[string]int)
		for _, rune := range id {
			l := string(rune)
			if letterCount[l] > 0 {
				letterCount[l]++
			} else {
				letterCount[l] = 1
			}
		}

		for _, v := range letterCount {
			if v == 2 {
				twice[id] = id
			}
			if v == 3 {
				thrice[id] = id
			}
		}
	}

	checksum := len(twice) * len(thrice)
	fmt.Printf("Day 02, Part 1: %v", checksum)
	fmt.Println()

	// part 2
	var box1 string
	var box2 string

Loop:
	for _, id1 := range boxIDs {
		for _, id2 := range boxIDs {
			if offByOne(id1, id2) {
				box1 = id1
				box2 = id2
				break Loop
			}
		}
	}

	sameChars := ""
	for i := range box1 {
		if box1[i] == box2[i] {
			sameChars = sameChars + string(box1[i])
		}
	}

	fmt.Printf("Day 02, Part 2: %v", sameChars)
	fmt.Println()
}
