package service

// Chat は、チャットの実装を表す。
type Chat interface {
	Post(args *PostArgs, reply *struct{}) error
}
