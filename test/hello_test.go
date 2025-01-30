package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	// テストケース1: 基本的な "hello" の表示
	hello := "hello"
	fmt.Println(hello)

	// テストケース2: アサーション
	if hello != "hello" {
		// 左が期待される値、右が実際の値
		t.Errorf("Expected 'hello', but got %s", hello)
	}
}

// １の数字渡されているかを確認
func TestOne(t *testing.T) {
	// テストケース1: 基本的な 1 の表示
	one := 1
	fmt.Println(one)

	// テストケース2: アサーション
	if one != 1 {
		// 左が期待される値、右が実際の値
		t.Errorf("Expected 1, but got %d", one)
	}
}

// 追加のテストケース: 1の型と値を確認
func TestOneType(t *testing.T) {
	one := 1
	
	// 型の確認
	if reflect.TypeOf(one).Kind() != reflect.Int {
		t.Errorf("Expected type int, but got %T", one)
	}
	
	// 値の範囲確認
	if one < 0 || one > 1 {
		t.Errorf("Expected value 1, but got %d", one)
	}
}

// Benchmarkも追加
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello := "hello"
		_ = hello
	}
}
