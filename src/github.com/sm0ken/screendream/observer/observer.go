package observer


type PlayerObserver interface{
	PlayerNotify(reason bool)
}

