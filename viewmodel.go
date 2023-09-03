package gmvvm

import (
	"sync"
	"unsafe"
)

type ValueTranslator[D any, V any] interface {
	ToModel(value V) (D, bool)
	ToView(value D) V
}

type IView[D any] interface {
	OnModelUpdate(value D)
	Bind(model *ViewModel[D])
	UpdateModel()
}

type View[D any, V any] struct {
	Value      V
	Translator ValueTranslator[D, V]
	viewModel  *ViewModel[D]
}

func NewView[D any]() *View[D, D] {
	return &View[D, D]{
		Translator: &TranslatorSameType[D]{},
	}
}

func (v *View[D, V]) Bind(model *ViewModel[D]) {
	v.viewModel = model
}

func (v *View[D, V]) OnModelUpdate(value D) {
	v.Value = v.Translator.ToView(value)
}

func (v *View[D, V]) UpdateModel() {
	data, _ := v.Translator.ToModel(v.Value)
	v.viewModel.UpdateModel(data)
}

func (v *View[D, V]) UpdateWith(update func(current V)) {

}

type IModel[D any] interface {
	ValuePointer() uintptr
	OnViewUpdate(value D)
	Bind(viewModel *ViewModel[D])
	UpdateViews()
}

type Model[D any] struct {
	Value     *D
	OnChange  func(src D, dst *D) bool
	viewModel *ViewModel[D]
}

func NewModel[D any](value *D) *Model[D] {
	return &Model[D]{
		Value: value,
		OnChange: func(src D, dst *D) bool {
			*dst = src
			return true
		},
	}
}

func (m *Model[D]) ValuePointer() uintptr {
	return uintptr(unsafe.Pointer(m.Value))
}

func (m *Model[D]) OnViewUpdate(value D) {
	if m.OnChange(value, m.Value) {
		m.UpdateViews()
	}
}

func (m *Model[D]) Bind(viewModel *ViewModel[D]) {
	m.viewModel = viewModel
}

func (m *Model[D]) UpdateViews() {
	m.viewModel.UpdateViews(*m.Value)
}

func (m *Model[D]) UpdateWith(update func(orig D) D) {
	*m.Value = update(*m.Value)
	m.UpdateViews()
}

type ViewModel[D any] struct {
	model IModel[D]
	views []IView[D]
	mux   sync.RWMutex
}

func (vm *ViewModel[D]) BindModel(model IModel[D]) {
	vm.mux.Lock()
	vm.model = model
	model.Bind(vm)
	vm.mux.Unlock()
}

func (vm *ViewModel[D]) BindView(view IView[D]) {
	vm.mux.Lock()
	view.Bind(vm)
	vm.views = append(vm.views, view)
	vm.mux.Unlock()
}

func (vm *ViewModel[D]) UpdateModel(value D) {
	vm.model.OnViewUpdate(value)
}

func (vm *ViewModel[D]) UpdateViews(value D) {
	for _, view := range vm.views {
		view.OnModelUpdate(value)
	}
}
