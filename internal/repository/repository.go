package repository

type UserRepo interface {
}

type MessageRepo interface {
}

type DatabaseRepo interface {
	UserRepo
	MessageRepo
}
