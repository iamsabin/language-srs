package main

import (
	"strconv"

	"language-srs/wanikani"
)

var recentMistakes = []string{
	"画面",
	"活気",
	"句",
	"以後",
	"事典",
	"返",
	"浴",
	"証明",
	"植える",
	"文句",
	"禁止",
	"妥協",
	"欠席",
	"保つ",
	"平",
	"頑丈",
	"丈夫",
	"苦しい",
	"映像",
	"老人",
	"伝える",
	"器",
	"参る",
	"近々",
	"工業",
	"人形",
	"首",
	"返る",
	"巳",
	"反",
	"企画",
	"昔",
}

func getRecentMistakes() []wanikani.Subject {
	var allSubjects []wanikani.Subject

	for i, v := range recentMistakes {
		allSubjects = append(allSubjects, wanikani.Subject{
			ID:   strconv.Itoa(i),
			Text: v,
		})
	}

	return allSubjects
}
