package consts

// swagger:enum Filesystem
// ENUM(file,dir)
type Filesystem string

// swagger:enum FileStatus
// ENUM(normal="", uploading="u")
type FileStatus string

// swagger:enum FilesystemDriver
// ENUM(local, aliyun_oss, tencent_oss, huawei_oss)
type FilesystemDriver string
