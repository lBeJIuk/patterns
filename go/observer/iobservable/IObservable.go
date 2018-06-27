package iobservable

type IObservable interface {
	RegisterObserver(name string, observable IObservable)
	RemoveObserver(name string)
	NotifyObservers()
}