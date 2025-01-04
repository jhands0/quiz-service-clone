export interface Quiz {
    id: string;
    name: string
    questions: QuizQuestion[];
}

export interface Player {
    id: string;
    name: string;
}

export interface QuizQuestion {
    id: string;
    name: string;
    choices: QuizChoices[];
}

export interface QuizChoices {
    id: string;
    name: string;
    correct: boolean;
}