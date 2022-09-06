# go_todo_cli
Go製のTODOリストを使用できるCLIツールです。

# 使用技術
- go 1.19
- Docker

# コマンド
## Linux・Mac
```text
- ./todo -add [タスク名]       タスク登録
- ./todo -comp [タスクID]      タスク完了
- ./todo -del [タスクID]       タスク削除
- ./todo -list                タスクリスト表示
- ./todo -help                ヘルプ表示
```
## Windows
```text
- .\todo -add [タスク名]       タスク登録
- .\todo -comp [タスクID]      タスク完了
- .\todo -del [タスクID]       タスク削除
- .\todo -list                タスクリスト表示
- .\todo -help                ヘルプ表示
```

# 動作確認方法
## Linux
1. ターミナルなどを開き、任意のディレクトリで『 curl -OL https://github.com/ucho456/go_todo_cli/releases/download/v1.0.0/todo_1.0.0_linux_amd64.tar.gz 』を実行し、tar.gzファイルをダウンロード
2. 『 tar -zxvf todo_1.0.0_linux_amd64.tar.gz 』を実行し、tar.gzファイルを解凍
3. ./todo -help でヘルプを表示

## Mac
1. ターミナルなどを開き、任意のディレクトリで『 curl -OL https://github.com/ucho456/go_todo_cli/releases/download/v1.0.0/todo_1.0.0_darwin_amd64.tar.gz 』を実行し、tar.gzファイルをダウンロード
2. 『 tar -zxvf todo_1.0.0_darwin_amd64.tar.gz 』を実行し、tar.gzファイルを解凍
3. ./todo -help でヘルプを表示

## Windows
1. コマンドプロンプトなどを開き、任意のディレクトリで『 curl -OL https://github.com/ucho456/go_todo_cli/releases/download/v1.0.0/todo_1.0.0_windows_amd64.tar.gz 』を実行し、tar.gzファイルをダウンロード
2. 『 tar -xzf todo_1.0.0_windows_amd64.tar.gz 』を実行し、tar.gzファイルを解凍
3. .\todo -help でヘルプを表示

# ダウンロードページ
[こちら](https://github.com/ucho456/go_todo_cli/releases/tag/v1.0.0)からダウンロードページに遷移できます。