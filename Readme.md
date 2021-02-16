# package_list_watcher_client

パッケージマネージャの情報をライブラリ経由で定期的に掠め取って、

サーバへ送信するツール。

## 設定ファイルを用意する。

* path: ~/.config/plwc/config.tml

内容例は以下の通り

```
[Base]
URL = "http://127.0.0.1:3000"
```
