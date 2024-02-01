package arithmetic

import "testing"

func TestAddition(t *testing.T) {
	adapter := NewAdapter()
	result, err := adapter.Addition(1, 2)
	if err != nil {
		t.Errorf("Addition() error = %v", err)
		return
	}
	if result != 3 {
		t.Errorf("Addition() result = %v, want 3", result)
	}
}

func TestSubtraction(t *testing.T) {
	adapter := NewAdapter()
	result, err := adapter.Subtraction(1, 2)
	if err != nil {
		t.Errorf("Subtraction() error = %v", err)
		return
	}
	if result != -1 {
		t.Errorf("Subtraction() result = %v, want -1", result)
	}
}

func TestMultiplication(t *testing.T) {
	adapter := NewAdapter()
	result, err := adapter.Multiplication(1, 2)
	if err != nil {
		t.Errorf("Multiplication() error = %v", err)
		return
	}
	if result != 2 {
		t.Errorf("Multiplication() result = %v, want 2", result)
	}
}

func TestDivision(t *testing.T) {
	adapter := NewAdapter()
	result, err := adapter.Division(1, 2)
	if err != nil {
		t.Errorf("Division() error = %v", err)
		return
	}
	if result != 0 {
		t.Errorf("Division() result = %v, want 0", result)
	}
}
