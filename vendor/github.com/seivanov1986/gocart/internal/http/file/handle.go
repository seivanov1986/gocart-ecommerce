package file

const (
	serviceBasePathKey       = "service_base_path"
	adminPrefix              = "/admin/static"
	adminPostfix             = "/schemes/admin/index.html"
	serviceBasePathPostfix   = "/schemes/admin"
	dynamicPrefix            = "/tmp/project/images/"
	stripPrefix              = "/static"
	fileServerHandlerPostfix = "/schemes/public"
)

type handle struct {
}

func New() *handle {
	return &handle{}
}
