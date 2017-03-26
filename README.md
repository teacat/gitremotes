# Gitremotes

Gitremotes 能夠幫你建立多個 Origin 遠端，而不需要手動逐一輸入遠端路徑。

# 說明

在目前目錄下新增一個 `.gitremotes` 檔案，在裡面存放多個路徑，第一行的遠端路徑會成為主要的 `push` 和 `fetch` 遠端。

```bash
git@github.com:TeaMeow/Gitremotes.git
git@bitbucket.org:TeaMeow/Gitremotes.git
```

接著在同個目錄下執行 `gitremotes` 就能夠在目前目錄下建立多個 `origin` 遠端。

# 實際用法

```bash
# 將希望的多個 origin 遠端存放入 `.gitremotes` 檔案中。
# 第一行的遠端路徑會成為主要的 push 和 fetch 遠端。
$ echo "git@github.com:TeaMeow/Gitremotes.git" >> .gitremotes
$ echo "git@bitbucket.org:TeaMeow/Gitremotes.git" >> .gitremotes

# 查看路徑是否正確。
$ cat .gitremotes
git@github.com:TeaMeow/Gitremotes.git
git@bitbucket.org:TeaMeow/Gitremotes.git

# 接著透過 `gitremotes` 就能夠在目前目錄下建立多個 `origin` 遠端。
$ gitremotes
>> /Users/YamiOdymel/go/src/github.com/TeaMeow/Gitremotes
origin  git@github.com:TeaMeow/Gitremotes.git (fetch)
origin  git@github.com:TeaMeow/Gitremotes.git (push)
origin  git@github.com:TeaMeow/Filary.git (push)
```