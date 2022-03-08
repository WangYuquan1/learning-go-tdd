package main

import "testing"

/*
テスト作成ルール
1 xxx_test.goのような名前のファイルにある必要があります。
2 テスト関数はTestという単語で始まる必要があります。
3 テスト関数は1つの引数のみをとります。 t *testing.T
*/
func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q wang %q", got, want)
	}
}
