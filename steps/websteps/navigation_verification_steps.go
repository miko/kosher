package websteps

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/cbush06/kosher/steps/steputils"
)

func iShouldBeRedirectedTo(s *steputils.StepUtils) func(string) error {
	return iShouldBeOn(s)
}

func iShouldBeOn(s *steputils.StepUtils) func(string) error {
	return func(pageName string) error {
		var (
			expectedURL string
			err         error
		)

		if expectedURL, err = s.ResolvePage(pageName); err != nil {
			return fmt.Errorf("error encountered while verifying page URL: %s", err)
		}

		// get this page's URL
		currentURL, _ := s.Page.URL()

		// trim trailing forward slashes
		trailingForwardSlashRegex := regexp.MustCompile("/$")
		currentURL = trailingForwardSlashRegex.ReplaceAllString(currentURL, "")

		expectedUnescapedURL, _ := url.PathUnescape(expectedURL)
		// assert their equality
		if expectedUnescapedURL != currentURL {
			return fmt.Errorf("expected URL to be [%s] but was [%s]", expectedUnescapedURL, currentURL)
		}
		return nil
	}
}

// I should see the popup text
func iShouldSeeThePopupText(s *steputils.StepUtils) func(string) error {
	return iSeeThePopupText(s, true)
}

// I should not see the popup text
func iShouldNotSeeThePopupText(s *steputils.StepUtils) func(string) error {
	return iSeeThePopupText(s, false)
}

// I should see / should not see the the popup text
func iSeeThePopupText(s *steputils.StepUtils, shouldSee bool) func(string) error {
	return func(text string) error {
		var (
			actualText       string
			interpolatedText string
			err              error
			errMsg           = `error encountered verifying popup text: %s`
		)

		// replace variables
		if interpolatedText, err = s.ReplaceVariables(text); err != nil {
			return fmt.Errorf(errMsg, err)
		}

		if actualText, err = s.Page.PopupText(); (err != nil || len(actualText) < 1) && shouldSee {
			return fmt.Errorf(errMsg, fmt.Sprintf("expected popup text [%s], but found none", interpolatedText))
		}

		doesMatch := strings.EqualFold(interpolatedText, actualText)
		if !doesMatch && shouldSee {
			return fmt.Errorf(errMsg, fmt.Sprintf("expected popup text [%s], but found [%s]", interpolatedText, actualText))
		} else if doesMatch && !shouldSee {
			return fmt.Errorf(errMsg, fmt.Sprintf("expected not to see popup text [%s], but did", interpolatedText))
		}
		return nil
	}
}
