package tus

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

type Uploader struct {
	files map[uuid.UUID]File
}

type File struct {
	fileName      string
	bytesTotal    int64
	bytesUploaded int64
}

func (u *Uploader) CreateRequest(c *gin.Context) {
	total := c.GetHeader("Upload-Length")
	totalLength, err := strconv.ParseInt(total, 10, 32)
	if err != nil {
		c.JSON(400, map[string]string{"message": "error, " + err.Error()})
		c.Abort()
		return
	}

	meta := c.GetHeader("Upload-Metadata")
	metaV := strings.Split(meta, ",")
	var metaKV map[string]string
	for _, va := range metaV {
		kv := strings.Split(va, " ")
		if len(kv) == 1 {
			metaKV[kv[0]] = ""
		} else {
			metaKV[kv[0]] = kv[1]
		}
	}
	filename, ok := metaKV["filename"]
	if !ok {
		c.JSON(400, map[string]string{"message": "error, no filename in Upload-Metadata header"})
		c.Abort()
		return
	}
	file := File{
		fileName:      filename,
		bytesTotal:    totalLength,
		bytesUploaded: 0,
	}
	id := uuid.New()
	u.files[id] = file
	c.Header("Location", "/"+id.String())
}

func (u *Uploader) HeadRequest(c *gin.Context) {
	strid := c.Param(":fileid")
	id, err := uuid.Parse(strid)
	if err != nil {
		c.JSON(400, map[string]string{"message": "error, " + err.Error()})
		c.Abort()
		return
	}
	f, ok := u.files[id]
	if !ok {
		c.JSON(404, map[string]string{"message": "error, " + err.Error()})
		c.Abort()
		return
	}
	c.Header("Upload-Offset", fmt.Sprintf("%v", f.bytesUploaded))
}
