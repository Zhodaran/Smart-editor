package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type word struct {
	kind     []string
	factur   []string
	name     []string
	color    []string
	size     []string
	quantity []string
}

func firstKind(line string) string {
	parts := strings.SplitN(line, " ", 2)
	kinds := []string{"мозаика", "мозайка", "ступень"}
	for _, kid := range kinds {
		if strings.Contains(strings.ToLower(parts[0]), strings.ToLower(kid)) || strings.Contains(strings.ToLower(parts[0]), "керамогранит") || strings.Contains(strings.ToLower(parts[0]), "плитка") || strings.Contains(strings.ToLower(parts[0]), "плита") || strings.Contains(strings.ToLower(parts[0]), "керамическая") || strings.Contains(strings.ToLower(parts[0]), "настенная") || strings.Contains(strings.ToLower(parts[0]), "напольная") || strings.Contains(strings.ToLower(parts[0]), "керамическая") || strings.Contains(strings.ToLower(parts[0]), "керамический") {

			if len(parts) == 2 {
				result := strings.Join([]string{parts[1]}, " ")
				if strings.Contains(strings.ToLower(result), "плитка") || strings.Contains(strings.ToLower(result), "плита") || strings.Contains(strings.ToLower(result), "керамическая") || strings.Contains(strings.ToLower(result), "настенная") || strings.Contains(strings.ToLower(result), "напольная") || strings.Contains(strings.ToLower(result), "керамическая") || strings.Contains(strings.ToLower(result), "керамический") {
					add := strings.SplitN(result, " ", 2)
					kind := parts[0] + " " + add[0]
					return kind
				} else {
					return strings.Join([]string{parts[0]}, " ")
				}
			}
		}
	}
	return ""
}

func firstFactur(line string) string {
	keyWord := []string{" altacera ", "cersanit", " alma ceramica ", "globaltile", "global tile", "maimoon", "laparet", "ceradim", " ab ", "ab ceramic", " lcm ", "staro", "керамин "}
	line = strings.ToLower(line)
	foundSupplier := "" // Переменная для хранения найденного поставщика

	for _, key := range keyWord {
		matched, err := regexp.MatchString(key, line)
		if err != nil {
			continue
		}
		if matched {
			if key == "globaltile" {
				return " Global Tile "
			}
			words := strings.Fields(key)
			var upperText []rune
			if len(words) > 1 {
				for _, word := range words {
					runes := []rune(word)
					upperText = append(upperText, unicode.ToUpper(runes[0]))
					upperText = append(upperText, runes[1:]...)
					upperText = append(upperText, ' ')
				}
				foundSupplier = string(upperText)
			} else {
				runes := []rune(words[0])
				if len(runes) <= 3 {
					for _, r := range runes {
						upperText = append(upperText, unicode.ToUpper(r))
					}
					foundSupplier = string(upperText)
				} else {
					upperText = append(upperText, unicode.ToUpper(runes[0])) // Заглавная первая буква
					upperText = append(upperText, runes[1:]...)              // Остальные буквы
					foundSupplier = string(upperText)                        // Сохраняем найденного поставщика
				}
			}
			// Если мы уже нашли поставщика, выходим из цикла
			if foundSupplier != "" {
				return foundSupplier
			}
		}
	}
	return "Поставщик не найден"
}

func ColorFound(line string) string {
	colors := []string{" light ", "sugar white ", "red ", " green ", " blue ", " yellow ", " black ", " white ", " gray ", " orange ", " purple ", " pink ", " brown ", " cyan ", " beige ", " lavender ", " mint ", " peach ", " gold ", " silver ", " burgundy ", " olive ", " sand ", " violet ", " emerald ", " aquamarine ", " malachite ", " graphite ", " magenta ", " cream ", " salmon ", " dark blue ", " dark green ", " dark red ", " dark purple ", " pastel ", " light brown ", " light blue ", " light green ", " light pink ", " light yellow ", " light gray ", " dark beige ", " light beige ", " laurel ", " copper ", " garnet ", " light violet ", " light purple ", " light olive ", " grey ", " dark ", " красный ", " зеленый ", " синий ", " желтый ", " черный ", " белый ", " серый ", " оранжевый ", " фиолетовый ", " розовый ", " коричневый ", " голубой ", " бежевый ", " лавандовый ", " мятный ", " персиковый ", " золотой ", " серебряный ", " бордовый ", " оливковый ", "бежевая", " песочный ", " сиреневый ", " изумрудный ", " аквамариновый ", " малахитовый ", " графитовый ", " пурпурный ", " кремовый ", " салатовый ", " темно-синий ", " темно-зеленый ", " темно-красный ", " темно-фиолетовый ", " пастельный ", " светло-коричневый ", " светло-голубой ", " светло-зеленый ", " светло-розовый ", " светло-желтый ", " светло-серый ", " темно-бежевый ", " светло-бежевый ", " лавровый ", " медный ", " гранатовый ", " светло-сиреневый ", " светло-фиолетовый ", " светло-оливковый ", " сахарная белая ", " красная ", " зеленая ", " синяя ", " желтая ", " черная ", " белая ", " серая ", " оранжевая ", " фиолетовая ", " розовая ", " коричневая ", " голубая ", " бежевую ", " лавандовая ", " мятная ", " персиковая ", " золотая ", " серебряная ", " бордовая ", " оливковая ", " песочная ", " сиреневая ", " изумрудная ", " аквамариновая ", " малахитовая ", " графитовая ", " пурпурная ", " кремовая ", " салатовая ", " темно-синяя ", " темно-зеленая ", " темно-красная ", " темно-фиолетовая ", " пастельная ", " светло-коричневая ", " светло-голубая ", " светло-зеленая ", " светло-розовая ", " светло-желтая ", " светло-серая ", " темно-бежевую ", " светло-бежевую ", " лавровая ", " медная ", " гранатовая ", " светло-сиреневая ", " светло-фиолетовая ", " светло-оливковая "}

	for _, color := range colors {
		if strings.Contains(strings.ToLower(line), strings.ToLower(color)) {

			return color

		}
	}
	return ""
}

func FindFuncX(line string) string {
	re := regexp.MustCompile(`\b(\d+[,\.]?\d*)\s*[xXхХ]\s*(\d+[,\.]?\d*)\b`)
	matches := re.FindString(line)
	if matches == "" {
		return "Нет размера"
	}
	return matches
}

func FindQuantity(line string) string {
	re := regexp.MustCompile(`\b[0-4][.,]\d+\b`)
	matches := re.FindString(line)
	if matches == "" {
		return "Нет квадратных метров или 1"
	}
	return matches
}

func FindName(line string, found []string) []string {
	var name []string
	var res []string
	var result []string
	parts := strings.SplitN(line, " ", -1)
	found2 := strings.Join(found, " ")
	found1 := strings.SplitN(found2, " ", -1)

	for _, str := range found1 {
		foundCor := strings.TrimSpace(str)
		if foundCor != "" {
			res = append(res, foundCor)
		}
	}
	for _, part := range parts {
		foundCor := strings.TrimSpace(part)
		if foundCor != "" {
			result = append(result, foundCor)
		}
	}
	name = result
	//fmt.Println(name)
	var filteredName []string
	for _, n := range name {
		//fmt.Println(n)
		found := false
		for _, f := range res {
			if strings.EqualFold(n, f) {
				found = true
				break
			}

		}
		if !found {
			if strings.ToLower(n) == "globaltile" {

			} else {
				filteredName = append(filteredName, n)
			}

		}

	}
	//fmt.Println(filteredName)
	return filteredName
}

func removeElement(slice []string, index int) []string { //Удаление уже найденных элементов из наименования
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func CharToLowwer(words string) string {
	var upperText []rune
	words = strings.ToLower(words)
	for _, word := range strings.Fields(words) {
		runes := []rune(word)
		upperText = append(upperText, unicode.ToUpper(runes[0]))
		upperText = append(upperText, runes[1:]...)
		upperText = append(upperText, ' ')
	}

	return string(upperText)
}

func main() {
	id := 0
	items := make(map[int]string)
	address := "0.0.0.0:8085"

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var item []string
	//var words []string
	w := word{}
	for scanner.Scan() {
		originLine := scanner.Text()
		par := strings.SplitN(originLine, " ", 2)
		var line string
		if len(par) > 1{
			line = string(par[1])
		}else{
			line = originLine
		}
		
		kind := firstKind(line)
		kind = CharToLowwer(kind)
		w.kind = append(w.kind, kind, "\n")
		factur := firstFactur(line)
		w.factur = append(w.factur, factur, "\n")
		colorFound := ColorFound(line)
		w.color = append(w.color, colorFound)
		size := FindFuncX(line)
		w.size = append(w.size, size, "\n")
		quantity := FindQuantity(line)
		w.quantity = append(w.quantity, quantity, "\n")

		/*if factur != "" {
			combined := fmt.Sprintf(" %s %s", kind, factur)
			if colorFound != "" {
				combined = fmt.Sprintf("%s %s", combined, colorFound)
			}
			if size != "" {
				combined = fmt.Sprintf("%s %s", combined, size)
			}
			if quantity != "" {
				combined = fmt.Sprintf("%s %s", combined, quantity)
			}
			item = append(item, combined)

			name := FindName(line, item)
			w.name = append(w.name, name...)
		} else {
			combined := fmt.Sprintf("%s %s", kind, " ")
			if colorFound != "" {
				combined = fmt.Sprintf("%s %s", combined, colorFound)
			}
			if size != "" {
				combined = fmt.Sprintf("%s %s", combined, size)
			}
			if quantity != "" {
				combined = fmt.Sprintf("%s %s", combined, quantity)
			}
			item = append(item, combined)




		}*/
		item = append(item, kind, factur, colorFound, size, quantity)
		w.name = FindName(line, item)
		wordes := strings.Join(w.name, " ")
		wordes = CharToLowwer(wordes)
		//var wordes string
		fmt.Println(wordes)
		//wordes = strings.Join(name, " ")
		//fmt.Println(wordes)
		quantity += " м²"
		size += " см "
		var position []string
		position = append(position, kind, factur, wordes, colorFound, size, quantity)
		positions := strings.Join(position, " ")
		items[id] = positions
		id++
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html; charset=utf-8")
		for _, worl := range items {
			fmt.Fprintf(w, "%s<br>", worl)
		}
	})

	err = http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println("Error")
	}

}
