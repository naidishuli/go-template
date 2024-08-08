package apierror

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"

    "go-template/config"
    "go-template/internal/app"
    "gopkg.in/yaml.v3"
)

func init() {
    err := loadConfig()
    if err != nil {
        panic(err)
    }
}

func loadConfig() error {
    data, err := config.FS.ReadFile("error_map.yaml")
    if err != nil {
        return err
    }

    return yaml.Unmarshal(data, &configMapInstance)
}

type Nested map[string]ConfigMap

type ConfigMap struct {
    ApiError *Error `yaml:"error"`
    Nested   `yaml:",inline"`
}

func (c *ConfigMap) FetchError(code string, target error) *Error {
    var err *Error

    codeParts := strings.Split(code, ".")
    node := configMapInstance.Nested
    if node != nil {
        for _, part := range codeParts {

            cfgEntry, ok := node[part]
            if !ok {
                break
            }

            err = cfgEntry.ApiError

            if cfgEntry.Nested != nil {
                node = cfgEntry.Nested
                continue
            }
        }
    }

    return err
}

var configMapInstance ConfigMap

type Response struct {
    Message string      `json:"message" yaml:"message"`
    Details interface{} `json:"details,omitempty" yaml:"details"`
}

// Error represent an errors to be returned to the client
type Error struct {
    Status   int `json:"-" yaml:"status"`
    Response `yaml:",inline"`

    err error
}

// New return a new ApiError based on the errors passed on argument
func New(err error) *Error {
    var message string
    wErr := err

    for {
        var appErr *app.Error
        ok := errors.As(wErr, &appErr)
        if ok {
            return apiErrorFromAppError(appErr, message)
        }

        message = err.Error()
        wErr = errors.Unwrap(wErr)
        if wErr == nil {
            break
        }
    }

    return &Error{
        Status: 500,
        Response: Response{
            Message: "Internal server error",
        },
        err: &app.Error{Code: "undefined", Err: err},
    }
}

// Error used to implement errors interface
func (e Error) Error() string {
    return e.Message
}

func (e Error) Log() string {
    var code, message, details, trace, sql string
    var err error

    var appErr *app.Error
    isAppErr := errors.As(e.err, &appErr)

    if isAppErr {
        err = appErr.Unwrap()
        code = string(appErr.Code)
        message = appErr.Message

        data, _ := json.Marshal(err)
        details = string(data)
    } else {
        data, _ := json.Marshal(e.Response)
        message = string(data)

        dataDetails, _ := json.Marshal(e.Details)
        details = string(dataDetails)
    }

    return fmt.Sprintf(
        "\n\nError: %s\n\nCode: %s\n\nMessage: %s\n\nDetails: %s\n\nStack: \t%s\n\nSql: \t%s\n"+
            "--------------------------------------------------------------------------------------------------------------\n",
        err,
        code,
        message,
        details,
        trace,
        sql,
    )
}

// apiErrorFromAppError transform internal error to an api error
func apiErrorFromAppError(appErr *app.Error, msg string) *Error {
    apiError := &Error{
        Status: 500,
        Response: Response{
            Message: "internal server error",
        },
        err: appErr,
    }

    // todo implement checks with errors.Is() method for specific scenarios

    cfgErr := configMapInstance.FetchError(appErr.Code, appErr)
    if cfgErr != nil {
        apiError = cfgErr
    }

    //switch err.Code {
    //
    //default:
    //	apiError.Status = 500
    //	apiError.Message = "Internal server error"
    //}

    return apiError
}
