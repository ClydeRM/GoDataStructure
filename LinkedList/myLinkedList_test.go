package LinkedList

import (
	"reflect"
	"testing"
)

func TestLinkedList_Prepend(t *testing.T) {
	type args struct {
		dataList []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// test cases
		{
			"Prepend_EmptyList",
			args{
				dataList: []string{},
			},
			[]string{},
		},
		{
			"Prepend_SingleNode",
			args{
				dataList: []string{"Node1"},
			},
			[]string{"Node1"},
		},
		{
			"Prepend_MultipleList",
			args{
				dataList: []string{"Node1", "Node2", "Node3"},
			},
			[]string{"Node3", "Node2", "Node1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewLinkedList[string]()
			for _, data := range tt.args.dataList {
				newNode := &Node[string]{Data: data}
				list.Prepend(newNode)
			}

			gotValues := listToSlice(list)
			if !reflect.DeepEqual(gotValues, tt.want) {
				t.Errorf("LinkedList %s: got = %v, want = %v", "Prepend()", gotValues, tt.want)
			}
		})
	}
}

func TestLinkedList_Append(t *testing.T) {
	type args struct {
		dataList []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// test cases
		{
			"Append_EmptyList",
			args{
				dataList: []string{},
			},
			[]string{},
		},
		{
			"Append_SingleNode",
			args{
				dataList: []string{"Node1"},
			},
			[]string{"Node1"},
		},
		{
			"Append_MultipleList",
			args{
				dataList: []string{"Node1", "Node2", "Node3"},
			},
			[]string{"Node1", "Node2", "Node3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewLinkedList[string]()
			for _, data := range tt.args.dataList {
				newNode := &Node[string]{Data: data}
				list.Append(newNode)
			}

			gotValues := listToSlice(list)
			if !reflect.DeepEqual(gotValues, tt.want) {
				t.Errorf("LinkedList %s: got = %v, want = %v", "Append()", gotValues, tt.want)
			}
		})
	}
}

func TestLinkedList_Pop(t *testing.T) {
	type args struct {
		dataList []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// test cases
		// test cases
		{
			"Pop_EmptyList",
			args{
				dataList: []string{},
			},
			[]string{},
		},
		{
			"Pop_SingleNode",
			args{
				dataList: []string{"Node1"},
			},
			[]string{},
		},
		{
			"Pop_MultipleList",
			args{
				dataList: []string{"Node1", "Node2", "Node3"},
			},
			[]string{"Node1", "Node2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewLinkedList[string]()
			for _, data := range tt.args.dataList {
				newNode := &Node[string]{Data: data}

				list.Append(newNode)
			}

			list.Pop()

			gotValues := listToSlice(list)
			if !reflect.DeepEqual(gotValues, tt.want) {
				t.Errorf("LinkedList %s: got = %v, want = %v", "Pop()", gotValues, tt.want)
			}

			if list.length != len(tt.want) {
				t.Errorf("LinkedList %s: length = %d, want = %d", "Pop()", list.length, len(tt.want))
			}

		})
	}
}

// 輔助函數
func listToSlice[T any](list *LinkedList[T]) []T {
	var result []T
	if list.head == nil {
		return []T{}
	}

	current := list.head
	for current != nil {
		result = append(result, current.Data)
		current = current.next
	}
	return result
}
