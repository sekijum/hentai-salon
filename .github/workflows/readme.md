# .secrets.exampleをコピー後適時編集する
cp .secrets.example .secrets

# secretを読み込ませて実行する
# act --secret-file .secretsは動かず。。main.ymlは.github/workflowsを見るのにsecretは見ないという謎い。。
act --secret-file .github/workflows/.secrets