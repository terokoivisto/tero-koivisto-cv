package models

type Skill struct {
	Icon   string   `json:"icon"`
	Name   string   `json:"name"`
	Usages []string `json:"usages"`
}

type Experience struct {
	Company string `json:"company"`
	Title   string `json:"title"`
	From    string `json:"from"`
	To      string `json:"to"`
	Summary string `json:"summary"`
}

type CV struct {
	Name       string       `dynamodbav:"name" json:"name"`
	AboutMe    string       `dynamodbav:"aboutMe" json:"aboutMe"`
	PersonalMe string       `dynamodbav:"personalMe" json:"personalMe"`
	Location   string       `dynamodbav:"location" json:"location"`
	Title      string       `dynamodbav:"title" json:"title"`
	Skills     []Skill      `dynamodbav:"skills" json:"skills"`
	Experience []Experience `dynamodbav:"experience" json:"experience"`
}
