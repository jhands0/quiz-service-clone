package service

import (
	"backend/internal/collection"
	"backend/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizService struct {
	quizCollection *collection.QuizCollection
}

func Quiz(quizCollection *collection.QuizCollection) *QuizService {
	return &QuizService{
		quizCollection: quizCollection,
	}
}

func (s QuizService) GetQuizById(id primitive.ObjectID) (*entity.Quiz, error) {
	return s.quizCollection.GetQuizById(id)
}

func (s QuizService) GetQuizzes() (*[]entity.Quiz, error) {
	return s.quizCollection.GetQuizzes()
}
