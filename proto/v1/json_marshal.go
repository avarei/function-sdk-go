package v1

import (
	"encoding/json"
	"errors"
	"fmt"
)

type jsonResourceSelector struct {
	ApiVersion  string            `json:"apiVersion"`
	Kind        string            `json:"kind"`
	MatchName   *string            `json:"matchName,omitempty"`
	MatchLabels *MatchLabels `json:"matchLabels,omitempty"`
	Namespace   *string           `json:"namespace,omitempty"`
}

func (r *ResourceSelector) UnmarshalJSON(data []byte) error {
    var tmp jsonResourceSelector

    if err := json.Unmarshal(data, &tmp); err != nil {
        return err
    }

    r.ApiVersion = tmp.ApiVersion
    r.Kind = tmp.Kind
    r.Namespace = tmp.Namespace

    switch {
    case tmp.MatchName != nil:
        r.Match = &ResourceSelector_MatchName{
            MatchName: *tmp.MatchName,
        }

    case tmp.MatchLabels != nil:
        r.Match = &ResourceSelector_MatchLabels{
            MatchLabels: tmp.MatchLabels,
        }

    default:
        r.Match = nil
    }

    return nil
}

func (r *ResourceSelector) MarshalJSON() ([]byte, error) {
	var tmp jsonResourceSelector

	tmp.ApiVersion = r.ApiVersion
	tmp.Kind = r.Kind
	tmp.Namespace = r.Namespace

	switch m := r.Match.(type) {
	case *ResourceSelector_MatchName:
		tmp.MatchName = &m.MatchName
	case *ResourceSelector_MatchLabels:
	  tmp.MatchLabels = m.MatchLabels
	}

	return json.Marshal(tmp)
}

func (s *Severity) UnmarshalJSON(data []byte) error {
		var (
				tmpSeverity *int32
				tmpString *string
		)
	  if err := json.Unmarshal(data, &tmpSeverity); err == nil {
				*s = Severity(*tmpSeverity)
				return nil
		}

		errString := json.Unmarshal(data, &tmpString)
		switch {
		case errString == nil && tmpString == nil:
				*s = Severity_SEVERITY_UNSPECIFIED
	  case errString == nil && *tmpString == "fatal":
				*s = Severity_SEVERITY_FATAL
		case errString == nil && *tmpString == "warning":
				*s = Severity_SEVERITY_WARNING
		case errString == nil && *tmpString == "normal":
				*s = Severity_SEVERITY_NORMAL
		default:
				return errors.New(fmt.Sprintf("unable to unmarshal unknown Severity: \"%s\"", data))
		}

    return nil
}

func (t *Target) UnmarshalJSON(data []byte) error {

		var (
				tmpTarget *int32
				tmpString *string
		)
	  if err := json.Unmarshal(data, &tmpTarget); err == nil {
				*t = Target(*tmpTarget)
				return nil
		}

		errString := json.Unmarshal(data, &tmpString)
		switch {
		case errString == nil && tmpString == nil:
				*t = Target_TARGET_UNSPECIFIED
	  case errString == nil && *tmpString == "composite":
				*t = Target_TARGET_COMPOSITE
		case errString == nil && *tmpString == "compositeAndClaim":
				*t = Target_TARGET_COMPOSITE_AND_CLAIM
		default:
				return errors.New(fmt.Sprintf("unable to unmarshal unknown Target \"%s\"", data))
		}

    return nil
}

func (r *Ready) UnmarshalJSON(data []byte) error {
		var (
				tmpReady *int32
				tmpBool *bool
		)
	  if err := json.Unmarshal(data, &tmpReady); err == nil {
				*r = Ready(*tmpReady)
				return nil
		}

		errBool := json.Unmarshal(data, &tmpBool)

		switch {
		case errBool == nil && tmpBool == nil:
				*r = Ready_READY_UNSPECIFIED
	  case errBool == nil && *tmpBool == false:
				*r = Ready_READY_FALSE
		case errBool == nil && *tmpBool == true:
				*r = Ready_READY_TRUE
		default:
				return errors.New(fmt.Sprintf("unable to unmarshal unknown Ready \"%s\"", data))
		}

    return nil
}

func (s *Status) UnmarshalJSON(data []byte) error {
		var (
				tmpBool *bool
				tmpStatus *int32
				tmpString *string
		)
		if err := json.Unmarshal(data, &tmpStatus); err == nil {
				*s = Status(*tmpStatus)
				return nil
		}

		errBool := json.Unmarshal(data, &tmpBool)
		errString := json.Unmarshal(data, &tmpString)
		switch {
		case errBool == nil && tmpBool == nil:
				fallthrough
		case errString == nil && tmpString == nil:
				*s = Status_STATUS_CONDITION_UNSPECIFIED
		case errBool == nil && *tmpBool == true:
				fallthrough
		case errString == nil && *tmpString == "true":
				*s = Status_STATUS_CONDITION_TRUE
		case errBool == nil && *tmpBool == false:
				fallthrough
		case errString == nil && *tmpString == "false":
				*s = Status_STATUS_CONDITION_FALSE
		case errString == nil && *tmpString == "unknown":
				*s = Status_STATUS_CONDITION_UNKNOWN
		default:
				return errors.New(fmt.Sprintf("unable to unmarshal unknown Status: \"%s\"", data))
		}

    return nil
}

