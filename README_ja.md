サンプルサーバー
-----------

## 目的
  + GOとJavaの性能比較。
  + ついでにLL言語(Ruby)とも性能比較

## 対象
  + HTTPサーバー
    + GETとPOSTリクエストを行う。
    + POSTで値を送信し、入力値のチェックと計算結果野表示を行う。
    + レスポンスはファイルから取得してクライアントに送る。
    + なるべく、3言語で実装内容をそろえる。

## 結果
  go > java > ruby