package components

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

/**  设置opt
opt := &cos.MultiUploadOptions{
	OptIni: &cos.InitiateMultipartUploadOptions{
		//ContentType
		ObjectPutHeaderOptions:&cos.ObjectPutHeaderOptions{
			ContentType: "text/jpeg",
		},
	},
}
*/
type Cosupload struct {
}

var domain string = "https://zhiji-1253509153.cos.ap-shanghai.myqcloud.com/"
var u, _ = url.Parse(domain)
var b = &cos.BaseURL{BucketURL: u}
var c = cos.NewClient(b, &http.Client{
	Transport: &cos.AuthorizationTransport{
		//如实填写账号和密钥，也可以设置为环境变量
		SecretID:  os.Getenv("AKIDyl13pNrliYBZjofLl00lfMp31MVXmr9B"),
		SecretKey: os.Getenv("vbNuJTNAhy0A5sQjSxDEAsSntWv5Cfui"),
	},
})

/**
 * @return string 绝对地址
 * @author lin
 * @param name 上传的服务器地址 /images/test.png
 * @param file 本地图片地址
 * 上传 文件
 */
func (co Cosupload) Uploadfile(name, file string) (string, error) {
	_, _, err := c.Object.Upload(
		context.Background(), name, file, nil,
	)
	if err != nil {
		return "", err
	}
	return domain + name, err
}

/**
 * @return bool true为成功
 * @author lin
 * @param key 服务器地址  /images/test.png
 * @param file 本地图片地址
 * 删除 文件
 */
func (co Cosupload) Delfile(key string) (bool, error) {
	//key := "/images/alert-failed.png"
	_, err := c.Object.Delete(context.Background(), key)
	if err != nil {
		return false, err
	}
	return true, err
}

/**
 * @return string 本地路劲
 * @author lin
 * @param key 服务器地址  /images/test.png
 * @param localpath 本地路劲
 * 获取 文件
 */
func (co Cosupload) Getfile(key, localpath string) string {
	resp, err := c.Object.Get(context.Background(), key, nil)
	if err != nil {
		panic(err)
	}
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	// 2. 下载对象到本地文件
	_, err = c.Object.GetToFile(context.Background(), key, localpath, nil)
	if err != nil {
		panic(err)
	}
	return localpath
}
