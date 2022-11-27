package creational

// The Singleton design pattern is used to provide a single instance of an object
// and to guarantee that there are no duplicates. The object is created at the first call to use
// the instance and is then it is reused by all other parts of the application.
//
// Warning: this singleton counter implementation is not thread-safe and should not be used
// in concurrent execution. It shoulde be used in single-threaded applications only.
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
