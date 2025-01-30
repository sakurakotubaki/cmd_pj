package test

import (
	"errors"
	"testing"
)

// 未実装の関数のスタブ
func unimplementedFunction() (string, error) {
	// TODO: 実装する
	return "", errors.New("未実装の機能")
}

// 失敗するテストケース
func TestUnimplementedFunction(t *testing.T) {
	// このテストは意図的に失敗するように設計されています
	result, err := unimplementedFunction()

	// エラーがない場合は失敗
	if err == nil {
		t.Error("エラーが発生するはずですが、エラーがありません")
	}

	// 期待される結果と異なる場合に失敗
	expectedResult := "期待される結果"
	if result != expectedResult {
		t.Errorf("期待される結果 '%s' と異なります。実際の結果: '%s'", expectedResult, result)
	}
}

// 別の失敗するテストケース
func TestAlwaysFail(t *testing.T) {
	// 常に失敗するテスト
	t.Error("このテストは常に失敗します")
}

// パニックを発生させるテスト
func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("パニックが発生するはずでしたが、発生しませんでした")
		}
	}()

	// 意図的にパニックを発生させる
	panic("意図的なパニック")
}