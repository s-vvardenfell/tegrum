package archiver

type ArchiverExtracter interface {
	Archive(source, target string) error
	Extract(archive, target string) error
}
