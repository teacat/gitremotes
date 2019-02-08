# Gitremotes

Gitremotes 能夠幫你建立多個 Origin 遠端，而不需要手動逐一輸入遠端路徑。

## 說明

在目前目錄下新增一個 `.gitremotes` 檔案，在裡面存放多個路徑，第一行的遠端路徑會成為主要的 `push` 和 `fetch` 遠端。

```bash
git@github.com:teacat/gitremotes.git
git@bitbucket.org:teacat/gitremotes.git
```

接著在同個目錄下執行 `gitremotes` 就能夠在目前目錄下建立多個 `origin` 遠端。

## 實際用法

使用前，請先確定你的本地 Git 專案有先透過 `git init` 初始化，否則可能會發生 `exit status 128` 錯誤。

```bash
# 下載並安裝 Gitremotes。
$ go get github.com/teacat/gitremotes

# 將希望的多個 origin 遠端存放入 `.gitremotes` 檔案中。
# 第一行的遠端路徑會成為主要的 push 和 fetch 遠端。
$ echo "git@github.com:teacat/gitremotes.git" >> .gitremotes
$ echo "git@bitbucket.org:teacat/gitremotes.git" >> .gitremotes

# 查看路徑是否正確。
$ cat .gitremotes
git@github.com:teacat/gitremotes.git
git@bitbucket.org:teacat/gitremotes.git

# 確定本地 Git 專案有先初始化。
$ git init

# 接著透過 `gitremotes` 就能夠在目前目錄下建立多個 `origin` 遠端。
$ gitremotes
>> /Users/YamiOdymel/go/src/github.com/teacat/gitremotes
origin  git@github.com:teacat/gitremotes.git (fetch)
origin  git@github.com:teacat/gitremotes.git (push)
origin  git@bitbucket.org:teacat/gitremotes.git (push)
```