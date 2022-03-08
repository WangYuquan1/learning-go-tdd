package main

import "testing"

/*
テスト作成ルール
1 xxx_test.goのような名前のファイルにある必要があります。
2 テスト関数はTestという単語で始まる必要があります。
3 テスト関数は1つの引数のみをとります。 t *testing.T
*/
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
