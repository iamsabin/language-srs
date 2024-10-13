package transliterate

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"language-srs/transliterate/jisho"

	"language-srs/model"
)

func Transliterate(input string) []model.Transliterate {
	input = strings.ReplaceAll(input, "'", "")
	cmd := exec.Command("/bin/sh", "-c",
		"docker exec ichiran-main-1 ichiran-cli --full "+input)
	// cmd := exec.Command("docker", "exec", "ichiran-master-main-1", "ichiran-cli", "--full", input)
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	o := unMarshalJSON(res)

	for i := range o {
		unquoteUnicodeFields(&o[i])
	}

	return Serialise(o)
}

func unquoteUnicodeFields(wordInfo *WordInfo) {
	var err error
	// wordInfo.Reading, err = strconv.Unquote(`"` + wordInfo.Reading + `"`)
	// if err != nil {
	// 	fmt.Println("Error unquoting Reading:", err)
	// 	return
	// }
	wordInfo.Text, err = strconv.Unquote(`"` + wordInfo.Text + `"`)
	if err != nil {
		fmt.Println("Error unquoting Text:", err)
		return
	}
	wordInfo.Kana, err = strconv.Unquote(`"` + wordInfo.Kana + `"`)
	if err != nil {
		fmt.Println("Error unquoting Kana:", err)
		return
	}
	for i := range wordInfo.Gloss {
		if strings.TrimSpace(wordInfo.Gloss[i].Gloss) == "" {
			continue
		}
		gloss, glossErr := strconv.Unquote(`"` + wordInfo.Gloss[i].Gloss + `"`)
		if glossErr != nil {
			continue
		}
		wordInfo.Gloss[i].Gloss = gloss

		// if wordInfo.Gloss[i].Info != "" {
		// 	wordInfo.Gloss[i].Info, err = strconv.Unquote(`"` + wordInfo.Gloss[i].Info + `"`)
		// 	if err != nil {
		// 		fmt.Println("Error unquoting Info:", err)
		// 		return
		// 	}
		// }
	}
}

func unMarshalJSON(input []byte) []WordInfo {
	// Parse the JSON string
	var data []interface{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}

	// Helper function to process the nested structure

	// Start processing the data
	return nestedUnmarshal(data)
}

func nestedUnmarshal(input []interface{}) []WordInfo {
	var allWordInfo []WordInfo
	for _, v := range input {
		switch vv := v.(type) {
		case []interface{}:
			allWordInfo = append(allWordInfo, nestedUnmarshal(vv)...)
		case map[string]interface{}:
			if nestedV, ok := vv["alternative"]; ok {
				sameWordInfo := nestedUnmarshal(nestedV.([]interface{}))

				wordInfo := sameWordInfo[0]
				for i, vw := range sameWordInfo {
					if i == 0 {
						continue
					}

					wordInfo.Kana += "," + vw.Kana
					wordInfo.Gloss = append(wordInfo.Gloss, vw.Gloss...)
				}

				allWordInfo = append(allWordInfo, wordInfo)
				break
			}
			data, _ := json.Marshal(vv)

			var wordInfo WordInfo
			_ = json.Unmarshal(data, &wordInfo)
			if len(wordInfo.Components) > 0 {
				comps := wordInfo.Components

				for i := range comps {
					if len(comps[i].Conj) == 0 {
						continue
					}
					comps[i].Gloss = comps[i].Conj[0].Gloss
				}
				allWordInfo = append(allWordInfo, comps...)
			} else {
				if len(wordInfo.Gloss) == 0 && len(wordInfo.Conj) > 0 {
					wordInfo.Gloss = wordInfo.Conj[0].Gloss
				}
				allWordInfo = append(allWordInfo, wordInfo)
			}
		}
	}

	return allWordInfo
}

func Serialise(input []WordInfo) []model.Transliterate {
	var result []model.Transliterate

	for i := range input {
		v := input[i]

		var meanings []string

		for _, val := range v.Gloss {
			meanings = append(meanings, val.GetMeanings()...)
		}

		if len(meanings) == 0 {
			meanings = jisho.SearchJisho(v.Text)
		}
		result = append(result, model.Transliterate{
			Kanji:    v.Text,
			Kana:     v.Kana,
			Meanings: meanings,
		})
	}

	return result
}
