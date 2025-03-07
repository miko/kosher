Feature: Verify Miscellaneous Elements of the Page

    Verify that Kosher's steps for verifying
    various elements on the page work properly.

    Background:
        Given I maximize the window

    Scenario: Confirm I see all the Texts
        Verify that I see various labels and
        excerpts of text on the page.

        Given I go to the "jquery-dropdown" page
        Then I should see all of the texts:
            | Select Country                 |
            | Select State                   |
            | Select US Outlying Territories |

	Scenario: Confirm Flags Work
        Verify that selector flags work.

        Given I go to the "home" page
        Then I should see the "@{INVISIBLE}Bootstrap Date Picker" link
        And I should not see the "Bootstrap Date Picker" link
        And the "GlyphiconHomeIcon" element should exist
        And the "GlyphiconMailIcon" element should not exist
        And the "HomeCenterMessage" element should contain "We have listed most of the components that are used by developers to build web applications."
        And the "HomeCenterMessage" element should not contain "Hello, World."
		

    Scenario: Confirm Buttons/Links are Visible on the Page
        Verify that I see a button and/or link on the page
        and that I do not see a button and/or link on the page.

        Given I go to the "table-search" page
        When I click "Date pickers"
        Then I should see the "Bootstrap Date Picker" link
        And I should see a link that points to "./bootstrap-date-picker-demo.html"
        And I should not see the "Table Pagination" link
        And I should not see a link that points to "./table-pagination-demo.html"
        And I should see the "Filter" button

    Scenario: Confirm Table of Data is Visible on the Page
        Verify that a table of data is on the page.

        Given I go to the "table-search" page
        Then I should see the following:
            | # | Task         | Assignee   | Status      |
            | 1 | Wireframes   | John Smith | in progress |
            | 2 | Landing Page | Mike Trout | completed   |
            | 3 | SEO tags     | Loblab Dan | failed qa   |

    Scenario: Confirm INPUTs with NAME
        Verify that I see an INPUT with a specific NAME.

        Given I go to the "input-form" page
        Then I should see an "input" with "name" of "first_name"
        And I should not see an "input" with "name" of "middle_name"

    Scenario: Confirm DIVs with ID
        Verify that I see a DIV with a specific ID.

        Given I go to the "basicform" page
        Then I should see a "div" with "id" of "user-message"
        And I should not see a "div" with "id" of "user-message2"
