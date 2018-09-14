---
layout: default
---

# Use Cases, Requirements, Constraints, and Domain

## Use Cases

See more detailed information on our approaches to these use cases [here](https://github.com/sul-dlss/rialto/wiki/Web-App-Specs). These are loosely in prioritization order and functionally grouped, with only a subset being followed for the Phase 1 Work Cycle.

### RIALTO-8: Map Stanford faculty members' co-authors to institution or country:

* Select department or school, select date range:
  * Generate list of countries of institutions of authors who have co-authored within the given date range a work with authors in the selected department or school, and the number of co-authors in an institution in that country
  * Generate list of institutions of authors who have co-authored within the given date range a work with authors in the selected department or school, and the number of co-authors at that institution
* Export via API or CSV the above as tabular data with these extra fields:
  * Author's name (author within the selected department and date range)
  * Author's organization (the selected department or school)
  * Author's co-author's name(s) (aggregated based on institution unique to the publication)
  * Author's co-author's institution name
  * Author's co-author's institution's country
* Visualize a heat map of the world, with color coding on countries to indicate frequency of collaboration (based on report 1 above).

### RIALTO-15: Cross-disciplinary Topics to Publications per Institution:

* Select years range, select topic area(s):
  * row for each organization, number of publications with the selected topic(s) by faculty of that organization
  * column per year
* Export as CSV the above

### RIALTO-13: Faculty Research Interests Trends

* Select years range, select organization:
  * One column per year, row for each subject heading with number of publications with that subject heading by authors in the selected organization & year
  * Ability to sort this by number of publications
* Export as CSV the above

### RIALTO-12: Publications from Grants

* Select organization, select date range:
  * List of Grants, number of connected publications
  * For each Grant, embedded list of titles of associated publications
* Ability to export the above as CSV.

### RIALTO-21: Number of Collaborations with Institutions

* Select date range, enter search string for an external institution (e.g. "Yale"):
   * Number of publications by an author(s) from Stanford and an author(s) from the selected, external institution
   * Above broken down by each year in the range and total for the range
* Interact with publication numbers above by clicking on it then seeing a list of the titles, authors, source, year for each of those publications

### RIALTO-11: Upcoming Grants by Organization

* Select an organization, select a date range:
  * Generate a report with Grant name, Grant Amount (optional), PI Name(s), Awarded Status, Start Date

### RIALTO-10: Number of Publications in a Journal or Publisher

* Select an organization, select a time range (optional):
  * Generate a report with journals and number of publications in the selected timeframe by authors in the department in that journal.
  * Generate a report with publishers and number of publications in the selected timeframe by authors in the department from that publisher.

### RIALTO-16: Publications by Topic by Department over Time

* Select a range of years, select a topic(s):
  * Generate a report that shows one column per month, row for each department, number of publications by faculty in that department in the select topic area.
  * Sort by number of publications.
* Export the above as CSV.

### RIALTO-26: RIALTO <=> ORCID Synchronization

For a faculty member, if we know their ORCID:
* Confirm publications listed for them in RIALTO by checking against ORCID
* Push confirmed publications from RIALTO to ORCID
* Retrieve publications by that faculty member from ORCID

### RIALTO-20: List Faculty Members who've Collaborated with Certain Countries

* Select organization, select date range, select country:
 * Generate report with list of authors in that department where author produced works (papers, projects, grants) associated with institutions in the select country

### RIALTO-25: Notify Faculty Members of Funding Opportunities

For a given user:
* Generate list of topics associated with them through their Works;
* Search funding opportunities matching that topic areas;
* Email that user with the list of funding opportunities.

For RIALTO:
* Aggregate information on funding opportunities from NSF, NIH, DOE, etc.

### RIALTO-17: Collaborations & Grants Networks

* Select organization, select date range:
  * Generate report of all grants or works from that date range associated with faculty members
  * Select one of the resulting grants or work, create visualization:
    * network visualization, institutions as node, lines connecting institutions are co-authors of shared works or grants, and size of nodes is determined by number of authors at the institution.
* Select a grant:
  * Generate number of students funded by that grant;
  * Generate number of articles produced per grant dollar.

### RIALTO-7: Number of Publications & Citations per Faculty Member

* Select organization, select date range, filter / search by author's name:
  * Generate report with Author Name, Total Number of Publications, Total Number of Publications Confirmed by Author (via Profile or other), Total Number of times Author was Cited

### RIALTO-18: Authors Retrieved by MESH Heading or Topic

* Select an organization, select a date range, select topics OR MESH headings (call numbers, expanded out across tree nodes & children):
  * Generate report with authors in that organization who authored works from that date range and with those topics or MESH headings

### RIALTO-22: Suggest Mentors for Faculty Member

* Select organization (optional), select topics:
  * Generate report of faculty members in the selected organization who have authored works in that topic area

### RIALTO-9: Publications Totals by Department & University

* Select a date range:
  * Generate report with total count of publications by organization, including total across the University, within the selected timeframe.

### RIALTO-23: Relevancy Ranking of Researchers Connections via Topics

* Select a time range, select an organization, select a MESH term or topic:
  * Generate list of researchers who produced works in that time range, in that organization, and on that topic;
  * Apply relevancy ranking algorithm to that list, and assign a ranking or score of working on the selected topic;
* From the above, visualize a network graph, with nodes as the authors/researchers on the topic, vertices are collaborations on works, thickness of vertices determined by number of connections, and filter to hide nodes that fall below a certain relevancy ranking threshold.

### RIALTO-14: Determine Students Funded by Grants

* Select Grant:
  * Generate list of authors (faculty, researchers, and students) associated with works associated with the grant

### RIALTO-19: Embed Visualizations

* Support embedding of visualizations into other websites, with ability to apply query filters

### RIALTO-24: Manual Indicate Publications by an Author

Within RIALTO:
* Allow researchers to authenticate, then upload a CV or list of publications as a Word Doc, BibTex file, or pasting a list into a text box.
  * RIALTO upload would parse the citations, extracting relevant information including DOIs and Pubmed IDs;
  * It would check to see if those publications already exist within RIALTO;
  * If not, it would add the publication to RIALTO as a verified publication (from the provided data, WoS, or Dimensions).

# Requirements

* Generate Reports with this information:
* Generate Visualizations of this type and with this information:
* Ingest & Aggregate Data about these entities:
* Ingest & Aggregate Data from these sources:
* Expose data via:


# Constraints

* A SDR (Stanford Digital Repository) researcher / depositor analytics dashboard.
* An SDR authorities management space. This will be a different system influenced by this work.
* At least in the short to mid-term, a faculty research profiles web application.

# Domain
