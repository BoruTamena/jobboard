package response

type Status string

const (
	Sucess Status = "SUCCESS"
	Fail   Status = "Fail"
)

type Response struct {
	Status string
	Data   interface{}
}
