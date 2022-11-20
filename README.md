# cubox-cli

![Go](https://github.com/jiyee/cubox-cli/workflows/Go/badge.svg)

[cubox](https://cubox.pro/) 命令行工具，支持添加 Link 和 Memo。

## 📥 安装

```bash
# 支持在不同平台运行
curl -sf https://gobinaries.com/jiyee/cubox-cli | sh
```

macOS 系统请下载使用 `cubox-cli-darwin-amd64.gz`

## 👉 使用

### 添加一条新的 Memo

```bash
$ cubox-cli new --api <CUSTOM_API> "a new memo from cubox-cli"
```

### 添加一条带标签的 Memo

```bash
$ cubox-cli new --api <CUSTOM_API> --tag "cubox-cli" "a new memo from cubox-cli"
```

### 添加一条带多个标签的 Memo

```bash
$ cubox-cli new --api <CUSTOM_API> --tag "cubox-cli" --tag "jiyee" "a new memo from cubox-cli"
```

### 将文本文件添加到 Memo

```bash
$ cat memo.txt | cubox-cli new --api <CUSTOM_API> --tag "cubox-cli"
```

### 使用环境变量来指定 API

```bash
$ export CUBOX_API=<CUSTOM_API>
$ cubox-cli new --tag "cubox-cli" "a new memo from cubox-cli"
```

## LICENCE

[MIT](./LICENSE)
