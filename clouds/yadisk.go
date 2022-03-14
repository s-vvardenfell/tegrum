package clouds

type YaDisk struct {
}

func NewYaDisk() *YaDisk {
	return &YaDisk{}
}

func (yd *YaDisk) DownLoadFile(fileId, dst string) {

}

func (yd *YaDisk) UploadFile(filename string) string {
	return ""
}
