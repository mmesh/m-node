package input

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-lib/pkg/cli/status"
	"mmesh.dev/m-lib/pkg/resources"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

const NewResource = "<new>"

func GetInput(inputMsg, helpMsg, defaultValue string, v survey.Validator) string {
	var resp string

	prompt := &survey.Input{Message: inputMsg, Help: helpMsg, Default: defaultValue}

	if v == nil {
		if err := survey.AskOne(prompt, &resp, survey.WithIcons(SurveySetIcons)); err != nil {
			status.Error(err, "Unable to get response")
		}
	} else {
		if err := survey.AskOne(prompt, &resp, survey.WithValidator(v), survey.WithIcons(SurveySetIcons)); err != nil {
			status.Error(err, "Unable to get response")
		}
	}

	return strings.TrimSpace(resp)
}

func GetInputInt(inputMsg, helpMsg, defaultValue string, v survey.Validator) int {
	n := GetInput(inputMsg, helpMsg, defaultValue, v)

	i, err := strconv.Atoi(n)
	if err != nil {
		status.Error(err, "Invalid number")
	}

	return i
}

func GetMultiline(inputMsg, helpMsg, defaultValue string, v survey.Validator) string {
	var resp string

	prompt := &survey.Multiline{Message: inputMsg, Help: helpMsg, Default: defaultValue}

	if v == nil {
		if err := survey.AskOne(prompt, &resp, survey.WithIcons(SurveySetIcons)); err != nil {
			status.Error(err, "Unable to get response")
		}
	} else {
		if err := survey.AskOne(prompt, &resp, survey.WithValidator(v), survey.WithIcons(SurveySetIcons)); err != nil {
			status.Error(err, "Unable to get response")
		}
	}

	return strings.TrimSpace(resp)
}

func GetSelect(inputMsg, helpMsg string, opts []string, v survey.Validator) string {
	var resp string

	promptSelect := &survey.Select{Message: inputMsg, Options: opts, Help: helpMsg}

	if v == nil {
		if err := survey.AskOne(promptSelect, &resp, survey.WithIcons(SurveySetIcons)); err != nil {
			status.Error(err, "Unable to get response")
		}
	} else {
		if err := survey.AskOne(promptSelect, &resp, survey.WithValidator(v), survey.WithIcons(SurveySetIcons)); err != nil {
			status.Error(err, "Unable to get response")
		}
	}

	return resp
}

func GetMultiSelect(inputMsg, helpMsg string, opts []string, v survey.Validator) []string {
	var resp []string

	promptMSelect := &survey.MultiSelect{Message: inputMsg, Options: opts, Help: helpMsg}

	if v == nil {
		if err := survey.AskOne(promptMSelect, &resp, survey.WithIcons(SurveySetIcons)); err != nil {
			status.Error(err, "Unable to get response")
		}
	} else {
		if err := survey.AskOne(promptMSelect, &resp, survey.WithValidator(v), survey.WithIcons(SurveySetIcons)); err != nil {
			status.Error(err, "Unable to get response")
		}
	}

	return resp
}

func GetConfirm(inputMsg string, defaultValue bool) bool {
	var resp bool

	promptConfirm := &survey.Confirm{Message: inputMsg, Default: defaultValue}

	if err := survey.AskOne(promptConfirm, &resp, survey.WithValidator(survey.Required), survey.WithIcons(SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	return resp
}

func GetPassword(inputMsg, helpMsg string) string {
	var pw string

	prompt := &survey.Password{Message: inputMsg, Help: helpMsg}

	if err := survey.AskOne(prompt, &pw, survey.WithValidator(ValidPassword), survey.WithIcons(SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get password")
	}

	return strings.TrimSpace(pw)
}

func PickResource(nxc rpc.CoreAPIClient, kind resource.Kind, realm, name, objSet string, edit bool) string {
	r := &resource.Resource{
		Kind:   kind,
		Realm:  realm,
		Name:   name,
		ObjSet: objSet,
	}

	rl, err := nxc.ListResources(context.TODO(), r)
	status.Error(err, "Unable to list resources")

	var objects []string
	for _, res := range rl.Resources {
		objects = append(objects, res.Name)
	}

	sort.Strings(objects)
	if edit {
		objects = append(objects, NewResource)
	}

	if len(objects) == 0 {
		msg.Alert("No objects found")
		os.Exit(1)
	}

	// if len(objects) == 1 && !edit {
	// 	return objects[0]
	// }

	inputMsg := fmt.Sprintf("%s:", resourceFromKind(kind))
	objID := GetSelect(inputMsg, "", objects, survey.Required)

	return objID
}

func MultiPickResource(nxc rpc.CoreAPIClient, kind resource.Kind, realm, name, objSet string) []string {
	r := &resource.Resource{
		Kind:   kind,
		Realm:  realm,
		Name:   name,
		ObjSet: objSet,
	}

	rl, err := nxc.ListResources(context.TODO(), r)
	status.Error(err, "Unable to list resources")

	var objects []string
	for _, res := range rl.Resources {
		objects = append(objects, res.Name)
	}

	if len(objects) == 0 {
		msg.Alertf("No %s found. Please, set at least one before working with this resource type.", kind)
		os.Exit(1)
	}
	sort.Strings(objects)

	inputMsg := fmt.Sprintf("%s:", resourceFromKind(kind))
	objIDs := GetMultiSelect(inputMsg, "", objects, nil)

	return objIDs
}

func MultiPickStrings(message string, opts []string) []string {
	if len(opts) == 0 {
		msg.Alert("No options found")
		os.Exit(1)
	}
	sort.Strings(opts)

	objIDs := GetMultiSelect(message, "", opts, nil)

	return objIDs
}

func resourceFromKind(kind resource.Kind) string {
	switch kind {
	case resource.Kind_VRF:
		return "Subnet"
	default:
		return resources.ObjectKindMap[kind]
	}
}
