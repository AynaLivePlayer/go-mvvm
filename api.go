package gmvvm

var _viewmodels = make(map[uintptr]interface{})

//func GetModel[D any](value *D) IModel[D] {
//	model := IModel[D](NewModel[D](value))
//	if viewModel, ok := _viewmodels[model.ValuePointer()]; !ok {
//		vm := &ViewModel[D]{
//			views: make([]IView[D], 0),
//		}
//		vm.BindModel(model)
//		_viewmodels[model.ValuePointer()] = vm
//	} else {
//		model = viewModel.(*ViewModel[D]).model
//	}
//	return model
//}

func WatchModel[D any](model IModel[D], view IView[D]) {
	if _, ok := _viewmodels[model.ValuePointer()]; !ok {
		vm := &ViewModel[D]{
			views: make([]IView[D], 0),
		}
		vm.BindModel(model)
		_viewmodels[model.ValuePointer()] = vm
	}
	vm := _viewmodels[model.ValuePointer()].(*ViewModel[D])
	vm.BindView(view)
}

func WatchValue[D any](value *D, view IView[D]) {
	WatchModel[D](NewModel[D](value), view)
}
