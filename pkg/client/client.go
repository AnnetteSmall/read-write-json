package client

type FileClient interface {
	Read(name string) (data []map[string]interface{}, err error)
	Write(data []map[string]interface{}, name string) (err error)

}
