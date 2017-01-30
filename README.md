# Bucket

Cli to get information of s3 buckets

```
$ bucket list

[{
  CreationDate: 2017-01-29 02:39:45 +0000 UTC,
  Name: "bronzdoc-bucket-0"
} {
  CreationDate: 2017-01-29 22:16:39 +0000 UTC,
  Name: "bronzdoc-bucket-1"
}]

```

Bucket will use the `default` profile in `~/.aws/credentials` to authenticate

# Install

```
curl -L "https://github.com/bronzdoc/bucket/releases/download/0.0.1/bucket-$(uname -s)-$(uname -m)" -o /usr/local/bin/bucket; chmod +x /usr/local/bin/bucket
```

# Usage

```
$ bucket

Get information of your AWS S3 buckets

Usage:
  bucket [command]

Available Commands:
  list        List AWS S3 buckets
  show        Show AWS S3 bucket objects
  version     Print the version number of Bucket

Flags:
      --config string   config file (default is $HOME/.bucket.yaml)

Use "bucket [command] --help" for more information about a command.

```
