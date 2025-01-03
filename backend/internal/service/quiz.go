package service

import (
	"github.com/jhands0/kahoot-clone/internal/collection"
	"github.com/jhands0/kahoot-clone/internal/entity"
)

type QuizService struct {
	quizCollection *collection.QuizCollection
}

func Quiz(quizCollection *collection.QuizCollection) *QuizService {
	return &QuizService{
		quizCollection: quizCollection,
	}
}

func (s QuizService) GetQuizzes() (*[]entity.Quiz, error) {
	return s.quizCollection.GetQuizzes()
}
