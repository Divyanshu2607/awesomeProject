package matrix

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	rows     int
	columns  int
	elements [][]int
}

func NewMatrix(rows, columns int) Matrix {
	elements := make([][]int, rows)
	for i := range elements {
		elements[i] = make(
			[]int,
			columns,
		)
	}
	return Matrix{
		rows:     rows,
		columns:  columns,
		elements: elements,
	}
}

func (m *Matrix) GetRows() int {
	return m.rows
}

func (m *Matrix) GetColumns() int {
	return m.columns
}

func (m *Matrix) SetElement(i, j, value int) {
	if i >= 0 && i < m.rows && j >= 0 && j < m.columns {
		m.elements[i][j] = value
	}
}

func (m *Matrix) Add(other Matrix) (Matrix, error) {
	if m.rows != other.rows || m.columns != other.columns {
		return Matrix{}, fmt.Errorf("matrices must have the same dimensions to be added")
	}

	result := NewMatrix(m.rows, m.columns)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			result.elements[i][j] = m.elements[i][j] + other.elements[i][j]
		}
	}

	return result, nil
}

func (m *Matrix) PrintMatrix() {
	jsonData, err := json.MarshalIndent(m.elements, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling matrix:", err)
		return
	}
	fmt.Println(string(jsonData))
}
