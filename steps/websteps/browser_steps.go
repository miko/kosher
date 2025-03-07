package websteps

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/afero"

	"github.com/cbush06/godog/gherkin"
	"github.com/cbush06/kosher/steps/steputils"
	"github.com/sclevine/agouti/api"
)

func iMaximizeTheWindow(s *steputils.StepUtils) func() error {
	return func() error {
		if err := s.Page.Size(s.GetMaxWindowSize()); err != nil {
			return fmt.Errorf("error encountered resizing window: %s", err)
		}
		return nil
	}
}

func iWaitSeconds() func(int) error {
	return func(seconds int) error {
		time.Sleep(time.Duration(seconds) * time.Second)
		return nil
	}
}

func iTakeAScreenshot(s *steputils.StepUtils) func(*gherkin.Embeddings) error {
	return func(e *gherkin.Embeddings) error {
		now := time.Now()

		fileName := "screenshot_" + now.Format("02Jan2006-150405.000.png")
		filePath, _ := s.Settings.FileSystem.ResultsDir.RealPath(fileName)

		if err := s.Page.Screenshot(filePath); err != nil {
			return fmt.Errorf("error encountered while taking screenshot: %s", err)
		}

		imageBytes, _ := afero.ReadFile(s.Settings.FileSystem.ResultsDir, fileName)
		e.AddEmbedding("image/png", base64.StdEncoding.EncodeToString(imageBytes))

		return nil
	}
}

func iSwitchToTheWindow(s *steputils.StepUtils) func(string) error {
	return func(nth string) error {
		var (
			errMsg      = fmt.Sprintf("error encountered while switching to [%s] window: ", nth) + "%s"
			nthNumeric  int
			windowCount int
			err         error
		)

		// Determine how many windows there are
		if windowCount, err = s.Page.WindowCount(); err != nil {
			return fmt.Errorf(errMsg, fmt.Sprintf("error encountered determining number of open windows: %s", err))
		}

		// if something other than an integer was specified
		if nthNumeric, err = strconv.Atoi(nth); err != nil {
			// if "first" or "last" was specified
			if strings.EqualFold(nth, "first") {
				return iSwitchToWindowIndex(s, nth, 0)
			}
			if strings.EqualFold(nth, "last") {
				return iSwitchToWindowIndex(s, nth, windowCount-1)
			}

			re := regexp.MustCompile(`\d+`)

			var match string
			if match = re.FindString(nth); len(match) < 1 {
				return fmt.Errorf(errMsg, "no valid index specified for window")
			}
			nthNumeric, _ = strconv.Atoi(match)
		}

		if nthNumeric > windowCount {
			return fmt.Errorf(errMsg, fmt.Sprintf("requested switch to window [%d], but only [%d] windows found", nthNumeric, windowCount))
		}

		return iSwitchToWindowIndex(s, nth, nthNumeric-1)
	}
}

func iSwitchToWindowIndex(s *steputils.StepUtils, nth string, idx int) error {
	var (
		windows []*api.Window
		errMsg  = fmt.Sprintf("error encountered while switching to [%s] window: ", nth) + "%s"
		err     error
	)

	if windows, err = s.Page.Session().GetWindows(); err != nil {
		return fmt.Errorf(errMsg, err)
	}

	return s.Page.Session().SetWindow(windows[idx])
}

func iReloadThePage(s *steputils.StepUtils) func() error {
	return func() error {
		return s.Page.Refresh()
	}
}

func sendKeysToActiveElement(s *steputils.StepUtils) func(string) error {
	return func(value string) error {
		var (
			activeEl          *api.Element
			interpolatedValue string
			err               error
			errMsg            = "error encountered while sending keys to active element: %s"
		)

		if activeEl, err = s.Page.Session().GetActiveElement(); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		// replace variables in value
		if interpolatedValue, err = s.ReplaceVariables(value); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		if err = activeEl.Value(interpolatedValue); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		return nil
	}
}
