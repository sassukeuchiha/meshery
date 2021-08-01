package system

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/layer5io/meshery/mesheryctl/pkg/utils"
)

func TestTokenCreateCmd(t *testing.T) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Not able to get current working directory")
	}
	currDir := filepath.Dir(filename)
	tests := []utils.CmdTestInput{
		{
			Name:                 "create the passed token with default location",
			Args:                 []string{"token", "create", "Default2"},
			ExpectedResponse:     "create_default.golden",
			ExpectedResponseYaml: "create_default.yaml",
			ExpectErr:            false,
		},
		{
			Name:                 "create the passed token with passed location",
			Args:                 []string{"token", "create", "Default2", "-f", "~/.meshery/auth.json"},
			ExpectedResponse:     "create.golden",
			ExpectedResponseYaml: "create.yaml",
			ExpectErr:            false,
		},
		{
			Name:                 "creating a token which already exists",
			Args:                 []string{"token", "create", "Default"},
			ExpectedResponse:     "create_err.golden",
			ExpectedResponseYaml: "create_err.yaml",
			ExpectErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			utils.SetupCustomContextEnv(t, currDir+"/testdata/token/"+tt.ExpectedResponseYaml)
			var b *bytes.Buffer
			var e *bytes.Buffer
			if tt.ExpectErr {
				e = utils.SetupLogrusGrabTesting(t)
				SystemCmd.SetErr(e)
			} else {
				b = utils.SetupLogrusGrabTesting(t)
				SystemCmd.SetOut(b)
			}
			SystemCmd.SetArgs(tt.Args)
			var actualResponse string
			err := SystemCmd.Execute()
			if err != nil {
				if !tt.ExpectErr {
					t.Error(err)
					return
				}
			}
			//Check the stdout/stderr against the golden file
			if tt.ExpectErr {
				actualResponse = e.String()
			} else {
				actualResponse = b.String()
			}
			// Expected response
			testdatatokenDir := filepath.Join(currDir, "testdata/token")
			golden := utils.NewGoldenFile(t, tt.ExpectedResponse, testdatatokenDir)
			if *update {
				golden.Write(actualResponse)
			}
			expectedResponse := golden.Load()
			if expectedResponse != actualResponse {
				t.Errorf("expected response [%v] and actual response [%v] don't match", expectedResponse, actualResponse)
			}
			//Skip checking the yamls if we had an error
			if tt.ExpectErr {
				return
			}
			//Check the modified yaml against the golden file
			path, err := os.Getwd()
			if err != nil {
				t.Error("unable to locate meshery directory")
			}
			filepath := path + "/testdata/token/" + tt.ExpectedResponseYaml
			content, err := ioutil.ReadFile(filepath)
			if err != nil {
				t.Error(err)
			}
			actualResponse = string(content)
			golden = utils.NewGoldenFile(t, tt.ExpectedResponseYaml, testdatatokenDir)
			if *update {
				golden.Write(actualResponse)
			}
			createExpected := golden.Load()
			if actualResponse != createExpected {
				t.Errorf("expected response %v and actual response %v don't match", createExpected, actualResponse)
			}
			if err := utils.Copy(path+"/fixtures/.meshery/config.yaml", filepath); err != nil {
				t.Error(err, "Could not complete test. Unable to configure create test file")
			}
			BreakupFunc(t)
		})
	}
}
func TestTokenDeleteCmd(t *testing.T) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Not able to get current working directory")
	}
	currDir := filepath.Dir(filename)
	tests := []utils.CmdTestInput{
		{
			Name:                 "delete the passed token",
			Args:                 []string{"token", "delete", "Default"},
			ExpectedResponse:     "delete_default.golden",
			ExpectedResponseYaml: "delete_default.yaml",
			ExpectErr:            false,
		},
		{
			Name:                 "delete the passed token(with a token name that doesn't exist)",
			Args:                 []string{"token", "create", "Default2"},
			ExpectedResponse:     "delete.golden",
			ExpectedResponseYaml: "delete_err.yaml",
			ExpectErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			utils.SetupCustomContextEnv(t, currDir+"/testdata/token/"+tt.ExpectedResponseYaml)
			var b *bytes.Buffer
			var e *bytes.Buffer
			if tt.ExpectErr {
				e = utils.SetupLogrusGrabTesting(t)
				SystemCmd.SetErr(e)
			} else {
				b = utils.SetupLogrusGrabTesting(t)
				SystemCmd.SetOut(b)
			}
			SystemCmd.SetArgs(tt.Args)
			var actualResponse string
			err := SystemCmd.Execute()
			if err != nil {
				if !tt.ExpectErr {
					t.Error(err)
					return
				}
			}
			//Check the stdout/stderr against the golden file
			if tt.ExpectErr {
				actualResponse = e.String()
			} else {
				actualResponse = b.String()
			}
			// Expected response
			testdatatokenDir := filepath.Join(currDir, "testdata/token")
			golden := utils.NewGoldenFile(t, tt.ExpectedResponse, testdatatokenDir)
			if *update {
				golden.Write(actualResponse)
			}
			expectedResponse := golden.Load()
			if expectedResponse != actualResponse {
				t.Errorf("expected response [%v] and actual response [%v] don't match", expectedResponse, actualResponse)
			}
			//Skip checking the yamls if we had an error
			if tt.ExpectErr {
				return
			}
			//Check the modified yaml against the golden file
			path, err := os.Getwd()
			if err != nil {
				t.Error("unable to locate meshery directory")
			}
			filepath := path + "/testdata/token/" + tt.ExpectedResponseYaml
			content, err := ioutil.ReadFile(filepath)
			if err != nil {
				t.Error(err)
			}
			actualResponse = string(content)
			golden = utils.NewGoldenFile(t, tt.ExpectedResponseYaml, testdatatokenDir)
			if *update {
				golden.Write(actualResponse)
			}
			createExpected := golden.Load()
			if actualResponse != createExpected {
				t.Errorf("expected response %v and actual response %v don't match", createExpected, actualResponse)
			}
			if err := utils.Copy(path+"/fixtures/.meshery/config.yaml", filepath); err != nil {
				t.Error(err, "Could not complete test. Unable to configure create test file")
			}
			BreakupFunc(t)
		})
	}
}

func TestTokenSetCmd(t *testing.T) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Not able to get current working directory")
	}
	currDir := filepath.Dir(filename)
	tests := []utils.CmdTestInput{
		{
			Name:                 "set the token for default context",
			Args:                 []string{"token", "set", "DefaultNew"},
			ExpectedResponse:     "set_default.golden",
			ExpectedResponseYaml: "set_default.yaml",
			ExpectErr:            false,
		},
		{
			Name:                 "set the token for custom context",
			Args:                 []string{"token", "set", "DefaultNew", "--context", "local2"},
			ExpectedResponse:     "set.golden",
			ExpectedResponseYaml: "set.yaml",
			ExpectErr:            false,
		},
		{
			Name:                 "set the token(without passing any token name)",
			Args:                 []string{"token", "set"},
			ExpectedResponse:     "set_err.golden",
			ExpectedResponseYaml: "set_err.yaml",
			ExpectErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			utils.SetupCustomContextEnv(t, currDir+"/testdata/token/"+tt.ExpectedResponseYaml)
			var b *bytes.Buffer
			var e *bytes.Buffer
			if tt.ExpectErr {
				e = utils.SetupLogrusGrabTesting(t)
				SystemCmd.SetErr(e)
			} else {
				b = utils.SetupLogrusGrabTesting(t)
				SystemCmd.SetOut(b)
			}
			SystemCmd.SetArgs(tt.Args)
			var actualResponse string
			err := SystemCmd.Execute()
			if err != nil {
				if !tt.ExpectErr {
					t.Error(err)
					return
				}
			}
			//Check the stdout/stderr against the golden file
			if tt.ExpectErr {
				actualResponse = e.String()
			} else {
				actualResponse = b.String()
			}
			// Expected response
			testdatatokenDir := filepath.Join(currDir, "testdata/token")
			golden := utils.NewGoldenFile(t, tt.ExpectedResponse, testdatatokenDir)
			if *update {
				golden.Write(actualResponse)
			}
			expectedResponse := golden.Load()
			if expectedResponse != actualResponse {
				t.Errorf("expected response [%v] and actual response [%v] don't match", expectedResponse, actualResponse)
			}
			//Skip checking the yamls if we had an error
			if tt.ExpectErr {
				return
			}
			//Check the modified yaml against the golden file
			path, err := os.Getwd()
			if err != nil {
				t.Error("unable to locate meshery directory")
			}
			filepath := path + "/testdata/token/" + tt.ExpectedResponseYaml
			content, err := ioutil.ReadFile(filepath)
			if err != nil {
				t.Error(err)
			}
			actualResponse = string(content)
			golden = utils.NewGoldenFile(t, tt.ExpectedResponseYaml, testdatatokenDir)
			if *update {
				golden.Write(actualResponse)
			}
			createExpected := golden.Load()
			if actualResponse != createExpected {
				t.Errorf("expected response %v and actual response %v don't match", createExpected, actualResponse)
			}
			if err := utils.Copy(path+"/testdata/token/set_reset.yaml", filepath); err != nil {
				t.Error(err, "Could not complete test. Unable to configure create test file")
			}
			BreakupFunc(t)
		})
	}
}
func TestTokenViewCmd(t *testing.T) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Not able to get current working directory")
	}
	currDir := filepath.Dir(filename)
	utils.SetupCustomContextEnv(t, currDir+"/testdata/token/view.yaml")
	tests := []utils.CmdTestInput{
		{
			Name:             "view the default2 token",
			Args:             []string{"token", "view", "Default2"},
			ExpectedResponse: "view.golden",
		},
		{
			Name:             "view with token name unspecified",
			Args:             []string{"token", "view"},
			ExpectedResponse: "view_err.golden",
			ExpectErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var b *bytes.Buffer
			var e *bytes.Buffer

			if tt.ExpectErr {
				e = utils.SetupLogrusGrabTesting(t)
				SystemCmd.SetErr(e)
			} else {
				b = utils.SetupLogrusGrabTesting(t)
				SystemCmd.SetOut(b)
			}
			SystemCmd.SetArgs(tt.Args)
			err := SystemCmd.Execute()
			if err != nil {
				if !tt.ExpectErr {
					t.Error(err)
					return
				}
			}
			var actualResponse string
			//Check the stdout against the golden file
			if tt.ExpectErr {
				actualResponse = e.String()
			} else {
				actualResponse = b.String()
			}
			// Expected response
			testdatatokenDir := filepath.Join(currDir, "testdata/token")
			golden := utils.NewGoldenFile(t, tt.ExpectedResponse, testdatatokenDir)
			if *update {
				golden.Write(actualResponse)
			}
			expectedResponse := golden.Load()
			if expectedResponse != actualResponse {
				t.Errorf("expected response [%v] and actual response [%v] don't match", expectedResponse, actualResponse)
			}
			BreakupFunc(t)
		})
	}
}
func TestTokenListCmd(t *testing.T) {
	// get current directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Not able to get current working directory")
	}
	currDir := filepath.Dir(filename)
	utils.SetupCustomContextEnv(t, currDir+"/testdata/token/list.yaml")
	tests := []utils.CmdTestInput{
		{
			Name:             "list all available tokens",
			Args:             []string{"token", "list"},
			ExpectedResponse: "list.golden",
		},
		{
			Name:             "list command with an extra argument passed to token",
			Args:             []string{"token", "list", "random"},
			ExpectedResponse: "list_err.golden",
			ExpectErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var b *bytes.Buffer
			var e *bytes.Buffer
			if tt.ExpectErr {
				e = utils.SetupLogrusGrabTesting(t)
				SystemCmd.SetErr(e)
			} else {
				b = utils.SetupLogrusGrabTesting(t)
				SystemCmd.SetOut(b)
			}
			SystemCmd.SetArgs(tt.Args)
			err := SystemCmd.Execute()
			if err != nil {
				if !tt.ExpectErr {
					t.Error(err)
					return
				}
			}
			var actualResponse string
			//Check the stdout/stderr against the golden file
			if tt.ExpectErr {
				actualResponse = e.String()
			} else {
				actualResponse = b.String()
			}
			// Expected response
			testdatatokenDir := filepath.Join(currDir, "testdata/token")
			golden := utils.NewGoldenFile(t, tt.ExpectedResponse, testdatatokenDir)
			if *update {
				golden.Write(actualResponse)
			}
			expectedResponse := golden.Load()
			if expectedResponse != actualResponse {
				t.Errorf("expected response [%v] and actual response [%v] don't match", expectedResponse, actualResponse)
			}
			BreakupFunc(t)
		})
	}
}
