package nationbuilder

import (
	"fmt"
	"net/http"
	"strconv"
)

// The enormous Person type within Nationbuilder.  Note that many of the references to other Person objects
// do not populate all of the fields in the Person type and instead return an abridged version.
// Only a request for a specific person returns the full Person object.
type Person struct {
	ActiveCustomerExpiresAt     *Date    `json:"active_customer_expires_at,omitempty"`
	ActiveCustomerStartedAt     *Date    `json:"active_customer_started_at,omitempty"`
	AgeCategory                 int      `json:"age_category,omitempty"`
	AuthorID                    int      `json:"author_id,omitempty"`
	Author                      *Person  `json:"author,omitempty"`
	AutoImportID                int      `json:"auto_import_id,omitempty"`
	Availability                string   `json:"availability,omitempty"`
	BannedAt                    *Date    `json:"banned_at,omitempty"`
	BillingAddress              *Address `json:"billing_address,omitempty"`
	Bio                         string   `json:"bio,omitempty"`
	BirthDate                   string   `json:"birthdate,omitempty"`
	CallStatusID                int      `json:"call_status_id,omitempty"`
	CallStatusName              string   `json:"call_status_name,omitempty"`
	CapitalAmountInCents        int      `json:"capital_amount_in_cents,omitempty"`
	ChildrenCount               int      `json:"children_count,omitempty"`
	Church                      string   `json:"church,omitempty"`
	CityDistrict                string   `json:"city_district,omitempty"`
	CitySubDistrict             string   `json:"city_sub_district,omitempty"`
	CiviCRMID                   int      `json:"civicrm_id,omitempty"`
	ClosedInvoicesAmountInCents int      `json:"closed_invoices_amount_in_cents,omitempty"`
	ClosedInvoicesCount         int      `json:"closed_invoices_count,omitempty"`
	ContactStatusID             int      `json:"contact_status_id,omitempty"`
	ContactStatusName           string   `json:"contact_status_name,omitempty"`
	//CouldVoteStatus                       bool     `json:"could_vote_status,omitempty"`
	CountyDistrict                        string   `json:"county_district,omitempty"`
	CountyFileID                          string   `json:"county_file_id,omitempty"`
	CreatedAt                             *Date    `json:"created_at,omitempty"`
	DataTrustID                           string   `json:"datatrust_id,omitempty"`
	Demographic                           string   `json:"demo,omitempty"`
	DoNotCall                             bool     `json:"do_not_call,omitempty"`
	DoNotContact                          bool     `json:"do_not_contact,omitempty"`
	DonationsAmountInCents                int      `json:"donations_amount_in_cents,omitempty"`
	DonationsAmountThisCycleIncCents      int      `json:"donations_amount_this_cycle_in_cents,omitempty"`
	DonationsCountThisCycle               int      `json:"donations_count_this_cycle,omitempty"`
	DonationsCount                        int      `json:"donations_count,omitempty"`
	DonationsPledgedAmountInCents         int      `json:"donations_pledged_amount_in_cents,omitempty"`
	DonationsRaisedAmountInCents          int      `json:"donations_raised_amount_in_cents,omitempty"`
	DonationsRaisedAmountThisCycleInCents int      `json:"donations_raised_amount_this_cycle_in_cents,omitempty"`
	DonationsRaisedCountThisCycle         int      `json:"donations_raised_count_this_cycle,omitempty"`
	DonationsRaisedCount                  int      `json:"donations_raised_count,omitempty"`
	DonationsToRaiseAmount                int      `json:"donations_to_raise_amount_in_cents,omitempty"`
	DwID                                  int      `json:"dw_id,omitempty"`
	EmailOneIsBad                         bool     `json:"email1_is_bad,omitempty"`
	EmailOne                              string   `json:"email1,omitempty"`
	EmailTwoIsBad                         bool     `json:"email2_is_bad,omitempty"`
	EmailTwo                              string   `json:"email2,omitempty"`
	EmailThreeIsBad                       bool     `json:"email3_is_bad,omitempty"`
	EmailThree                            string   `json:"email3,omitempty"`
	EmailFourIsBad                        bool     `json:"email4_is_bad,omitempty"`
	EmailFour                             string   `json:"email4,omitempty"`
	Email                                 string   `json:"email,omitempty"`
	EmailOptIn                            bool     `json:"email_opt_in,omitempty"`
	Employer                              string   `json:"employer,omitempty"`
	Ethnicity                             string   `json:"ethnicity,omitempty"`
	ExternalID                            string   `json:"external_id,omitempty"`
	FacebookAddress                       string   `json:"facebook_address,omitempty"`
	FacebookProfileURL                    string   `json:"facebook_profile_url,omitempty"`
	FacebookUpdatedAt                     string   `json:"facebook_updated_at,omitempty"`
	FacebookUsername                      string   `json:"facebook_username,omitempty"`
	FaxNumber                             string   `json:"fax_number,omitempty"`
	FederalDistrict                       string   `json:"federal_district,omitempty"`
	FederalDoNotCall                      bool     `json:"federal_donotcall,omitempty"`
	FireDistrict                          string   `json:"fire_district,omitempty"`
	FirstDonatedAt                        string   `json:"first_donated_at,omitempty"`
	FirstFundraisedAt                     string   `json:"first_fundraised_at,omitempty"`
	FirstInvoiceAt                        string   `json:"first_invoice_at,omitempty"`
	FirstName                             string   `json:"first_name,omitempty"`
	FirstProspectAt                       string   `json:"first_prospect_at,omitempty"`
	FirstRecruitedAt                      string   `json:"first_recruited_at,omitempty"`
	FirstSupporterAt                      string   `json:"first_supporter_at,omitempty"`
	FirstVolunteerAt                      string   `json:"first_volunteer_at,omitempty"`
	FullName                              string   `json:"full_name,omitempty"`
	HasFacebook                           bool     `json:"has_facebook,omitempty"`
	HomeAddress                           *Address `json:"home_address,omitempty"`
	ID                                    int      `json:"id,omitempty"`
	ImportID                              int      `json:"import_id,omitempty"`
	InferredParty                         string   `json:"inferred_party,omitempty"`
	InferredSupportLevel                  int      `json:"inferred_support_level,omitempty"`
	InvoicePaymentsAmountInCents          int      `json:"invoice_payments_amount_in_cents,omitempty"`
	InvoicePaymentsReferredAmountInCents  int      `json:"invoice_payments_referred_amount_in_cents,omitempty"`
	InvoicesAmountInCents                 int      `json:"invoices_amount_in_cents,omitempty"`
	InvoicesCount                         int      `json:"invoices_count,omitempty"`
	IsDeceased                            bool     `json:"is_deceased,omitempty"`
	IsDonor                               bool     `json:"is_donor,omitempty"`
	IsFundraiser                          bool     `json:"is_fundraiser,omitempty"`
	IsIgnoreDonationLimits                bool     `json:"is_ignore_donation_limits,omitempty"`
	IsLeaderboardable                     bool     `json:"is_leaderboardable,omitempty"`
	IsMobileBad                           bool     `json:"is_mobile_bad,omitempty"`
	IsPossibleDuplicate                   bool     `json:"is_possible_duplicate,omitempty"`
	IsProfilePrivate                      bool     `json:"is_profile_private,omitempty"`
	IsProfileSearchable                   bool     `json:"is_profile_searchable,omitempty"`
	IsProspect                            bool     `json:"is_prospect,omitempty"`
	IsSupporter                           bool     `json:"is_supporter,omitempty"`
	IsSurveyQuestionPrivate               bool     `json:"is_survey_question_private,omitempty"`
	IsTwitterFollower                     bool     `json:"is_twitter_follower,omitempty"`
	IsVolunteer                           bool     `json:"is_volunteer,omitempty"`
	JudicialDistrict                      string   `json:"judicial_district,omitempty"`
	LabourRegion                          string   `json:"labour_region,omitempty"`
	Language                              string   `json:"language,omitempty"`
	LastCallID                            int      `json:"last_call_id,omitempty"`
	LastContactedAt                       *Date    `json:"last_contacted_at,omitempty"`
	LastContactedBy                       *Person  `json:"last_contacted_by,omitempty"`
	LastDonatedAt                         *Date    `json:"last_donated_at,omitempty"`
	LastFundraisedAt                      *Date    `json:"last_fundraised_at,omitempty"`
	LastInvoiceAt                         *Date    `json:"last_invoice_at,omitempty"`
	LastName                              string   `json:"last_name,omitempty"`
	LastRuleViolationAt                   *Date    `json:"last_rule_violation_at,omitempty"`
	LegalName                             string   `json:"legal_name,omitempty"`
	LinkedInID                            string   `json:"linkedin_id,omitempty"`
	Locale                                string   `json:"locale,omitempty"`
	MailingAddress                        *Address `json:"mailing_address,omitempty"`
	MaritalStatus                         string   `json:"marital_status,omitempty"`
	MediaMarketName                       string   `json:"media_market_name,omitempty"`
	MeetupAddress                         *Address `json:"meetup_address,omitempty"`
	MembershipExpiresAt                   *Date    `json:"membership_expires_at,omitempty"`
	MembershipLevelName                   string   `json:"membership_level_name,omitempty"`
	MembershipStartedAt                   *Date    `json:"membership_started_at,omitempty"`
	MiddleName                            string   `json:"middle_name,omitempty"`
	Mobile                                string   `json:"mobile,omitempty"`
	MobileNormalised                      string   `json:"mobile_normalized,omitempty"`
	MobileOptIn                           bool     `json:"mobile_opt_in,omitempty"`
	NbecGUID                              string   `json:"nbec_guid,omitempty"`
	//NbecPrecinctCode                      int      `json:"nbec_precinct_code,omitempty"` //TODO unclear on whether this is an integer or string
	NgpID                            string   `json:"ngp_id,omitempty"`
	NoteUpdatedAt                    string   `json:"note_updated_at,omitempty"`
	Note                             string   `json:"note,omitempty"`
	Occupation                       string   `json:"note,omitempty"`
	OutstandingInvoicesAmountInCents int      `json:"outstanding_invoices_amount_in_cents,omitempty"`
	OutstandingInvoicesCount         int      `json:"outstanding_invoices_count,omitempty"`
	OverdueInvoicesCount             int      `json:"overdue_invoices_count,omitempty"`
	PageSlug                         string   `json:"page_slug,omitempty"`
	ParentID                         int      `json:"parent_id,omitempty"`
	Parent                           *Person  `json:"parent,omitempty"`
	PartyMember                      bool     `json:"party_member,omitempty"`
	Party                            string   `json:"party,omitempty"`
	PfStratID                        int      `json:"pf_strat_id,omitempty"`
	PhoneNormalised                  string   `json:"phone_normalized,omitempty"`
	PhoneTime                        string   `json:"phone_time,omitempty"`
	Phone                            string   `json:"phone,omitempty"`
	PrecinctCode                     string   `json:"precinct_code,omitempty"`
	PrecinctID                       int      `json:"precinct_id,omitempty"`
	PrecintName                      string   `json:"precinct_name,omitempty"`
	Prefix                           string   `json:"prefix,omitempty"`
	PreviousParty                    string   `json:"previous_party,omitempty"`
	PrimaryAddress                   *Address `json:"primary_address,omitempty"`
	PrimaryEmailID                   int      `json:"primary_email_id,omitempty"`
	PriorityLevelChangedAt           string   `json:"priority_level_changed_at,omitempty"`
	PriorityLevel                    int      `json:"priority_level,omitempty"`
	ProfileContentHTML               string   `json:"profile_content_html,omitempty"`
	ProfileContent                   string   `json:"profile_content,omitempty"`
	ProfileHeadline                  string   `json:"profile_headline,omitempty"`
	ProfileImageURLSSL               string   `json:"profile_image_url_ssl,omitempty"`
	ReceivedCapitalAmountInCents     int      `json:"received_capital_amount_in_cents,omitempty"`
	RecruiterID                      int      `json:"recruiter_id,omitempty"`
	Recruiter                        *Person  `json:"recruited,omitempty"`
	RecruitsCount                    int      `json:"recruits_count,omitempty"`
	RegisteredAddress                *Address `json:"registered_address,omitempty"`
	RegisteredAt                     *Date    `json:"registered_at,omitempty"`
	Religion                         string   `json:"religion,omitempty"`
	RncID                            int      `json:"rnc_id,omitempty"`
	RncRegID                         string   `json:"rnc_regid,omitempty"`
	RuleViolationsCount              int      `json:"rule_violations_count,omitempty"`
	SalesforceID                     string   `json:"salesforce_id,omitempty"`
	SchoolDistrict                   string   `json:"school_district,omitempty"`
	SchoolSubDistrict                string   `json:"school_sub_district,omitempty"`
	Sex                              string   `json:"sex,omitempty"`
	SignupType                       int      `json:"signup_type,omitempty"`
	SpentCapitalAmountInCents        int      `json:"spent_capital_amount_in_cents,omitempty"`
	StateFileID                      string   `json:"state_file_id,omitempty"`
	StateLowerDistrict               string   `json:"state_lower_district,omitempty"`
	StateUpperDistrict               string   `json:"state_upper_district,omitempty"`
	SubmittedAddress                 string   `json:"submitted_address,omitempty"`
	SubNations                       []string `json:"subnations,omitempty"`
	Suffix                           string   `json:"suffix,omitempty"`
	SupportLevelChangedAt            string   `json:"support_level_changed_at,omitempty"`
	SupportLevel                     int      `json:"support_level,omitempty"`
	SupportProbabilityScore          float32  `json:"support_probability_score,omitempty"`
	SupraNationalDistrict            string   `json:"supranational_district,omitempty"`
	Tags                             []string `json:"tags,omitempty"`
	Township                         string   `json:"township,omitempty"`
	TurnoutProbabilityScore          float32  `json:"turnout_probability_score,omitempty"`
	TwitterDescription               string   `json:"twitter_description,omitempty"`
	TwitterFollowersCount            int      `json:"twitter_followers_count,omitempty"`
	TwitterFriendsCount              int      `json:"twitter_friends_count,omitempty"`
	TwitterID                        string   `json:"twitter_id,omitempty"`
	TwitterLocation                  string   `json:"twitter_location,omitempty"`
	TwitterLogin                     string   `json:"twitter_login,omitempty"`
	TwitterName                      string   `json:"twitter_name,omitempty"`
	TwitterUpdatedAt                 *Date    `json:"twitter_updated_at,omitempty"`
	TwitterWebsite                   string   `json:"twitter_website,omitempty"`
	UnsubscribedAt                   *Date    `json:"unsubscribed_at,omitempty"`
	UpdatedAt                        *Date    `json:"updated_at,omitempty"`
	UserSubmittedAddress             *Address `json:"user_submitted_address,omitempty"`
	UserName                         string   `json:"username,omitempty"`
	VanID                            string   `json:"van_id,omitempty"`
	VillageDistrict                  string   `json:"village_district,omitempty"`
	Ward                             string   `json:"ward,omitempty"`
	WarningsCount                    int      `json:"warnings_count,omitempty"`
	Website                          string   `json:"website,omitempty"`
	WorkAddress                      *Address `json:"work_address,omitempty"`
	WorkPhoneNumber                  string   `json:"work_phone_number,omitempty"`
}

//TwitterAddress                        *Address `json:"twitter_address,omitempty"` // TODO determine the actual type returned

func (p *Person) String() string {
	return fmt.Sprintf("Person ID: %d, Name: %s %s", p.ID, p.FirstName, p.LastName)
}

// People is a collection of Person objects with pagination support
type People struct {
	Results []*Person `json:"results"`
	Pagination
}

type personWrap struct {
	Person *Person `json:"person"`
}

type Tag struct {
	PersonID  int    `json:"person_id"`
	Tag       string `json:"tag"`
	CreatedAt *Date  `json:"created_at"`
}

func (t *Tag) String() string {
	return fmt.Sprintf("Tag: %s", t.Tag)
}

type Tags struct {
	Taggings []*Tag `json:"taggings"`
}

type tagWrap struct {
	Tagging *Tag `json:"tagging"`
}

type tagStringWrap struct {
	Tagging tagString `json:"tagging"`
}

type tagString struct {
	Tag []string `json:"tag"`
}

// Retrieve a page of people
func (n *Client) GetPeople(options *Options) (people *People, result *Result) {
	req := n.getRequest("GET", "/people", options)
	result = n.retrieve(req, &people)

	return
}

// Create a person
func (n *Client) CreatePerson(person *Person, options *Options) (newPerson *Person, result *Result) {
	req := n.getRequest("POST", "/people", options)
	pw := &personWrap{}
	result = n.create(&personWrap{person}, req, pw, http.StatusCreated)
	newPerson = pw.Person

	return
}

// Retrieve a person
func (n *Client) GetPerson(id int, options *Options) (person *Person, result *Result) {
	u := fmt.Sprintf("/people/%d", id)
	req := n.getRequest("GET", u, options)
	pw := &personWrap{}
	result = n.retrieve(req, pw)
	person = pw.Person

	return
}

// Update a person
func (n *Client) UpdatePerson(id int, person *Person, options *Options) (updatedPerson *Person, result *Result) {
	u := fmt.Sprintf("/people/%d", id)
	req := n.getRequest("PUT", u, options)
	pw := &personWrap{}
	result = n.create(&personWrap{person}, req, pw, http.StatusOK)
	updatedPerson = pw.Person

	return
}

// Delete a person
func (n *Client) DeletePerson(id int, options *Options) (result *Result) {
	u := fmt.Sprintf("/people/%d", id)
	req := n.getRequest("DELETE", u, options)
	result = n.delete(req)

	return
}

// Retrieve people close to the provided latitude/longtitude within the given distance
func (n *Client) NearbyPeople(lattitude float32, longtitude float32, distance int, options *Options) (people *People, result *Result) {
	if options == nil {
		options = NewOptions()
	}
	options.SetQueryOption("location", fmt.Sprintf("%f,%f", lattitude, longtitude))
	options.SetQueryOption("distance", strconv.FormatUint(uint64(distance), 10))
	req := n.getRequest("GET", "/people/nearby", options)
	result = n.retrieve(req, &people)

	return
}

// The Push endpoint lazily creates a person if a matching record is not found or updates one if it is
func (n *Client) PushPerson(person *Person, options *Options) (pushedPerson *Person, result *Result) {
	req := n.getRequest("PUT", "/people/push", options)
	pw := &personWrap{}
	result = n.create(&personWrap{person}, req, pw, http.StatusOK)
	pushedPerson = pw.Person

	return
}

// Retrieve yourself - i.e. the person object associated with the request
func (n *Client) GetYourself(options *Options) (person *Person, result *Result) {
	req := n.getRequest("GET", "/people/me", options)
	pw := &personWrap{}
	result = n.retrieve(req, pw)
	person = pw.Person

	return
}

// Search people by providing a PeopleSearchOptions object together with any additional options (if any)
func (n *Client) SearchPeople(searchOptions *PeopleSearchOptions, options *Options) (people *People, result *Result) {
	if searchOptions != nil {
		if options == nil {
			options = NewOptions()
		}
		searchOptions.setOptions(options)
	}
	req := n.getRequest("GET", "/people/search", options)
	result = n.retrieve(req, &people)

	return
}

// Retrieve tags for the given person
func (n *Client) GetPersonTags(id int, options *Options) (tags []*Tag, result *Result) {
	u := fmt.Sprintf("/people/%d/taggings", id)
	req := n.getRequest("GET", u, options)
	tw := &Tags{}
	result = n.retrieve(req, tw)
	tags = tw.Taggings

	return
}

// Set a tag on a person
func (n *Client) CreatePersonTag(id int, tags []string, options *Options) (newTags *Tags, result *Result) {
	u := fmt.Sprintf("/people/%d/taggings", id)
	r := n.getRequest("PUT", u, options)
	tw := &tagStringWrap{
		Tagging: tagString{tags},
	}

	if len(tags) == 1 {
		newTag := &tagWrap{}
		result = n.create(tw, r, newTag, http.StatusOK)
		newTags = &Tags{
			Taggings: []*Tag{
				newTag.Tagging,
			},
		}
	} else {
		result = n.create(tw, r, &newTags, http.StatusOK)
	}

	return
}

// Delete a tag on a person
func (n *Client) DeletePersonTag(id int, tagName string, options *Options) (result *Result) {
	u := fmt.Sprintf("/people/%d/taggings/%s", id, tagName)
	r := n.getRequest("DELETE", u, options)
	result = n.delete(r)

	return
}

// MatchPerson attempts to match the provided match options to a person in the people database
func (n *Client) MatchPerson(matchOptions *PersonMatchOptions, options *Options) (person *Person, result *Result) {
	if matchOptions != nil {
		if options == nil {
			options = NewOptions()
		}
		matchOptions.setOptions(options)
	}
	req := n.getRequest("GET", "/people/match", options)
	pw := &personWrap{}
	result = n.retrieve(req, pw)
	person = pw.Person

	return
}

// Represents the various options that can be passed to the people search endpoint
type PeopleSearchOptions struct {
	FirstName         string
	LastName          string
	City              string
	State             string
	Sex               string
	BirthDate         string
	UpdatedSince      string
	WithMobile        string
	CiviCRMID         string
	CountyFileID      string
	StateFileID       string
	DataTrustID       string
	DwID              string
	MediaMarketID     string
	MembershipLevelID string
	NgpID             string
	PfStratID         string
	VanID             string
	SalesforceID      string
	RncID             string
	RncRegID          string
	ExternalID        string
}

func (p *PeopleSearchOptions) setOptions(o *Options) {
	p.setOptionString("first_name", p.FirstName, o)
	p.setOptionString("last_name", p.LastName, o)
	p.setOptionString("city", p.City, o)
	p.setOptionString("state", p.State, o)
	p.setOptionString("sex", p.Sex, o)
	p.setOptionString("birthdate", p.BirthDate, o)
	p.setOptionString("updated_since", p.UpdatedSince, o)
	p.setOptionString("with_mobile", p.WithMobile, o)
	p.setOptionString("civicrm_id", p.CiviCRMID, o)
	p.setOptionString("county_file_id", p.CountyFileID, o)
	p.setOptionString("state_file_id", p.StateFileID, o)
	p.setOptionString("datatrust_id", p.DataTrustID, o)
	p.setOptionString("dw_id", p.DwID, o)
	p.setOptionString("media_market_id", p.MediaMarketID, o)
	p.setOptionString("membership_level_id", p.MembershipLevelID, o)
	p.setOptionString("ngp_id", p.NgpID, o)
	p.setOptionString("pf_strat_id", p.PfStratID, o)
	p.setOptionString("van_id", p.VanID, o)
	p.setOptionString("salesforce_id", p.SalesforceID, o)
	p.setOptionString("rnc_id", p.RncID, o)
	p.setOptionString("rnc_regid", p.RncRegID, o)
	p.setOptionString("external_id", p.ExternalID, o)
}

func (p *PeopleSearchOptions) setOptionString(key string, value string, o *Options) {
	if value != "" {
		o.SetQueryOption(key, value)
	}
}

// TODO convert people search into using this func now there's the match endpoint also
// look at easier way of handling this
func setOptionString(key string, value string, o *Options) {
	if value != "" {
		o.SetQueryOption(key, value)
	}
}

// PeopleMatchOptions represents the different query params that can be set for the people match endpoint
type PersonMatchOptions struct {
	Email     string
	FirstName string
	LastName  string
	Phone     string
	Mobile    string
}

func (pmo *PersonMatchOptions) setOptions(o *Options) {
	setOptionString("email", pmo.Email, o)
	setOptionString("first_name", pmo.FirstName, o)
	setOptionString("last_name", pmo.LastName, o)
	setOptionString("phone", pmo.Phone, o)
	setOptionString("mobile", pmo.Mobile, o)
}
