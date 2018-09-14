---
layout: default
---

# Data Mappings, Models, and Application Profiles

## Namespaces

* RIALTO namespace base (not dereferenceable): http://sul.stanford.edu/rialto/
  * Organizations: http://sul.stanford.edu/rialto/agents/orgs/
  * People: http://sul.stanford.edu/rialto/agents/people/
  * Concepts: http://sul.stanford.edu/rialto/concepts/
  * Grants: http://sul.stanford.edu/rialto/grants/
* RIALTO Context Nodes:
  * Addresses: http://sul.stanford.edu/rialto/context/addresses/
  * Names: http://sul.stanford.edu/rialto/context/names/
  * Positions: http://sul.stanford.edu/rialto/context/positions/
  * Relationships: http://sul.stanford.edu/rialto/context/relationships/
  * Roles: http://sul.stanford.edu/rialto/context/roles/
* External:
  * dct: http://purl.org/dc/terms/
  * foaf: http://xmlns.com/foaf/0.1/
  * obo: http://purl.obolibrary.org/obo/
  * rdf: http://www.w3.org/1999/02/22-rdf-syntax-ns#
  * rdfs: http://www.w3.org/2000/01/rdf-schema#
  * skos: http://www.w3.org/2004/02/skos/core#
  * vivo: http://vivoweb.org/ontology/core#
  * vcard: http://www.w3.org/2006/vcard/ns#

## Data Models


## Metadata Application Profiles

### Organization

Label | Property | Sub-property Of | Range | Usage | Data Type | Vocab / Syntax Schema | Obligation | Source
----- | -------- | --------------- | ----- | ----- | --------- | --------------------- | ---------- | ------
URI   | n/a      | n/a             | URI   | RIALTO URI for the resource | URI | UUID | {1,1} | CAP API
Identifier | dct:identifier | n/a  | URI   | Other Identifiers for the resource | Literal | n/a | {1,n} | CAP API
Resource Type | rdf:type | n/a     | URI   | Class of the described resource | URI | n/a | {1,n} | CAP API
Alias | skos:altLabel | rdfs:label | Literal | Alternate name or label for the resource. | Literal | n/a | {0,n} | CAP API
Child | obo:BFO_0000051 | ?     | IRI   | Child or contained resource of the given resource | URI | n/a | {0,n} | CAP API
Parent | obo:BFO_0000050 | ?     | IRI   | Parent or container resource of the given resource | URI | n/a | {0,1} | CAP API
Name  | skos:prefLabel | rdfs:label | Literal | Preferred name or label for the resource. | Literal | n/a | {1,1} | CAP API
Admin Role on Grant | obo:RO_0000053 vivo:AdminRole ; vivo:relatedBy Grant-URI | vivo Roles pattern | Org Role URI ; Grant URI | Grant the Organization is administering, associated in tandem with a Admin Role context node | URI, URI | n/a | {0,n} | SERA API

### Person

Label | Property | Sub-property Of | Range | Usage | Data Type | Vocab / Syntax Schema | Obligation | Source
----- | -------- | --------------- | ----- | ----- | --------- | --------------------- | ---------- | ------
URI   | n/a      | n/a             | URI   | RIALTO URI for the resource | URI | UUID  | {1,1}      | SERA API
Identifier | dct:identifier | n/a  | URI   | Other Identifiers for the resource | Literal | n/a | {1,n} | SERA API
Full Name | vcard:fn | n/a         | Literal | Full name for the Person resource. | Literal | n/a | {1,1} | SERA API
PI on Grant | obo:RO_0000053 vivo:PrincipalInvestigatorRole ; vivo:relatedBy Grant-URI | vivo Roles pattern | PI Role URI ; Grant URI | Grant the Person is a principal investigator for, associated in tandem with a PI Role context node | URI, URI | n/a | {0,n} | SERA API
PI Role on Grant | obo:RO_0000053 vivo:PrincipalInvestigatorRole ; vivo:relatedBy Grant-URI | vivo Roles pattern | PI Role URI ; Grant URI | Grant the Person is a principal investigator for, associated in tandem with a PI Role context node | URI, URI | n/a | {0,n} | SERA API

### Grant

Label | Property | Sub-property Of | Range | Usage | Data Type | Vocab / Syntax Schema | Obligation | Source
----- | -------- | --------------- | ----- | ----- | --------- | --------------------- | ---------- | ------
URI   | n/a      | n/a             | URI   | RIALTO URI for the resource | URI | UUID | {1,1} | SERA API
Identifier | dct:identifier | n/a  | URI   | Other Identifiers for the resource | Literal | n/a | {1,n} | SERA API
Resource Type | rdf:type | n/a     | URI   | Class of the described resource | URI | n/a | {1,n} | SERA API
Alias | skos:altLabel | rdfs:label | Literal | Alternate name or label for the resource. | Literal | n/a | {0,n} | SERA API
Child | obo:BFO_0000051 | ?     | IRI   | Child or contained resource of the given resource | URI | n/a | {0,n} | SERA API
Parent | obo:BFO_0000050 | ?     | IRI   | Parent or container resource of the given resource | URI | n/a | {0,1} | SERA API
Name  | skos:prefLabel | rdfs:label | Literal | Preferred name or label for the resource. | Literal | n/a | {1,1} | SERA API
Dates | dct:date | dc:date   | Literal | Dates or date range applicable to the resource. | Literal | EDTF | {0,1} | SERA API
PI    | vivo:relates vivo:PrincipalInvestigatorRole . vivo:PrincipalInvestigatorRole obo:RO_0000052 PI-URI | vivo Roles pattern | Role URI ; Person URI | The principal investigator for the Grant, associated through a Role context node | URI, URI | n/a | {0,n} | SERA API
Administering Organization | vivo:relates vivo:AdminRole . vivo:AdminRole obo:RO_0000052 Organization-URI | vivo Roles pattern | Role URI ; Org URI | The administrating organization for the Grant, associated through a Role context node | URI, URI | n/a | {0,n} | SERA API
Funded By | vivo:assignedBy | n/a | URI   | URI for the Organization that funded the Grant. | URI | Organization URI | {0,n} | SERA API

## Data Sources

Source                | URL | Data Example (link) | Notes
--------------------- | --- | ------------------- | -----
⭐️ CAP = Profiles API | [GUI](https://profiles.stanford.edu/) ; [CAP API](https://cap.stanford.edu/cap-api/console) Or [CAP API](https://api.stanford.edu) | [Author sample](https://gist.github.com/peetucket/07746abaa2b8d9bce7b38499cbeab9bb) | SUL Pub only calls a subset of CAP, so we need the full CAP API.
⭐️ SERA               | [SERA API](https://aswsuat.stanford.edu/mais/sera/v1/api?scope=sera.stanford-only&sunetId=); [MAIS Docs](https://asconfluence.stanford.edu/confluence/display/MaIS/SeRA+API+-+User+Documentation) | See RIALTO Sample Data repo. | Requires look-up via Agent (Person) SunetId.
⭐️ Web of Science     |     |  |
⭐️ Dimensions         |     |  |
Stanford LDAP | [requesting access to LDAP](https://uit.stanford.edu/service/directory/access/requesting) |  | Subset of HR information? API access leveraging sunets only?
HR PeopleSoft Database |  |  | Likely not needed since we should have access to a more complete CAP API that will provide Agent data. No access? Fully included information (relevant for us) is in CAP already? (Some subset feeds into CAP, thus into SUL Pub)
SUL Pub = SUL CAP API | [GUI](http://sulcap.stanford.edu/) ; API: ? | [sample results for 3 publications based on author search](https://gist.github.com/peetucket/f572c9fe678998cb785accc22d5b5b64) | Pulls in subset of Web of Science, Pubmed, Profiles, CAP, Database access possible.
US Spending (Not done yet) | https://api.usaspending.gov/ |  | Check Github for the API for the status
Grants.gov (Pointed to by Dept of Labor) | Lookup API? |  |
Data.gov |  |  | Same as or includes Grants.gov?
PubMED |  |  | All captured in SUL Pub already?
ORCID lookup API |  |  |
ISNI lookup API |  |  |
DOI lookup API? |  |  |

## Data Mappings

### CAP Organizations

* Harvest endpoint: https://api.stanford.edu/cap/v1/orgs/stanford?includeChildren=True
* Returns Nested JSON, parent (Stanford University) => children
* type = `$.type` (string)
* alias = `$.alias` (string)
* orgCodes = `$.orgCodes` (array of strings)
* name = `$.name` (string)
* CAP identifier == alias
* children = `$.children` (array of hashes)

MAP Field | CAP Field(s) | Normalization Notes | Obligation
--------- | ------------ | ------------------- | ----------
URI       | n/a          | RIALTO organizations namespace + auto-generated UUID | {1,1}
Identifier | $.alias, $.orgCodes     | grab each org code (literal) within the orgCode array, do not assert the array | {1,n}
Resource Type | $.type | always foaf:Agent, foaf:Organization; then normalize $.type (DEPARTMENT == vivo:Department; DIVISION == vivo:Division; ROOT == vivo:university; SCHOOL == vivo:School; SUB_DIVISION == vivo:Division) | {1,n}
Alias     | $.alias      | none                | {1,?}
Children  | $.children.[\*].alias | Use Alias (as identifier) to lookup RIALTO organization URI | {0,n}
Parent    | $.alias for hash one level up in $.children.[\*] | Use Alias (as identifier) to lookup RIALTO organization URI | {0,1}
Name      | $.name    | none                | {1,1}

1. Retrieve API response (monthly?)
2. "Denest" with pointers the nested JSON
3. For each organization JSON:
  i. Look-up
  ii. Map
  

### SERA Grants

* Harvest endpoint: https://aswsuat.stanford.edu/mais/sera/v1/api?scope=sera.stanford-only&sunetId=
* Requires lookup via SUNETId of PI or related Person; returns JSON document where all relevant Grants are in an array ($.SeRARecord.[\*])
* For each hash in $.SeRARecord.[\*]:
  * `$.SeRARecord.[\*].agreementNumber` (string, ex: XL40402214)
  * `$.SeRARecord.[\*].agreementType` (string, ex: Contract | Cooperative Agreement | Grant | Subcontract)
  * `$.SeRARecord.[\*].directSponsorForeignFlag` (string, ex: N | Y)
  * `$.SeRARecord.[\*].directSponsorName` (string, ex: National Institutes of Health)
  * `$.SeRARecord.[\*].edwCreateDateTime` (string, ex: 24-FEB-14)
  * `$.SeRARecord.[\*].edwUpdateDateTime` (string, ex: 24-FEB-14)
  * `$.SeRARecord.[\*].internationalFlag` (string, ex: N | Y)
  * `$.SeRARecord.[\*].piAffiliationType` (string, ex: Former staff,Staff,Faculty,Emeritus faculty,Former faculty)
  * `$.SeRARecord.[\*].piDepartmentName` (string, ex: Anesthesiology&Periop&Pain Med ; *not currently used*)
  * `$.SeRARecord.[\*].piEmployeeId` (string, ex: 07885510)
  * `$.SeRARecord.[\*].piFullName` (string, ex: Pfefferbaum, Adolf)
  * `$.SeRARecord.[\*].piSunetId` (string, ex: dolfp)
  * `$.SeRARecord.[\*].primarySponsorForeignFlag` (string, ex: N | Y)
  * `$.SeRARecord.[\*].primarySponsorName` (string, ex; National Aeronautics and Space Administration)
  * `$.SeRARecord.[\*].projectEndDate` (string, ex: 01-JUL-84)
  * `$.SeRARecord.[\*].projectStartDate` (string, ex: 30-JUN-96)
  * `$.SeRARecord.[\*].projectStatus` (string, should be one value only, ex: Awarded)
  * `$.SeRARecord.[\*].projectTitle` (string, ex: Northern California Perinatal Dispatch Center)
  * `$.SeRARecord.[\*].spoNumber` (string, ex: 1)
  * `$.SeRARecord.[\*].spoProjectDeptName` (string, ex: Anesthesiology&Periop&Pain Med)
  * `$.SeRARecord.[\*].spoProjectSchool` (string, ex: Humanities and Sciences)
  * `$.SeRARecord.[\*].subawardsIncludedFlag` (string, ex: N | Y)

#### Grant Mapping

MAP Field | CAP Field(s) | Normalization Notes | Obligation
--------- | ------------ | ------------------- | ----------
URI       | n/a          | RIALTO grants namespace + auto-generated UUID | {1,1}
Identifier | $.SeRARecord.[\*].spoNumber ; $.SeRARecord.[\*].agreementNumber | none | {1,n}
Resource Type | $.SeRARecord.[\*].agreementType | if value == Grant vivo:Grant | {1,n}
Grant Start Date | $.SeRARecord.[\*].projectStartDate | change local date format to EDTF | {0,1}
Grant End Date | $.SeRARecord.[\*].projectEndDate | change local date format to EDTF | {0,1}
Name      | $.SeRARecord.[\*].projectTitle | none | {1,1}
PI        | $.SeRARecord.[\*].piEmployeeId | Lookup Person URI based on the Employee ID field, mint PI Role URI if needed | {0,n}
Administering Organization | $.SeRARecord.[\*].spoProjectDeptName or $.SeRARecord.[\*].spoProjectSchool | Lookup Organization URI based on the Name fields, mint Admininstration Role URI if needed | {0,n}
Funded By| $.SeRARecord.[\*].directSponsorName, $.SeRARecord.[\*].primarySponsorName | Look up (or mint / add) Organization URI based on field | {0,1}


#### PI Mapping

MAP Field | CAP Field(s) | Normalization Notes | Obligation
--------- | ------------ | ------------------- | ----------
URI       | n/a          | RIALTO persons namespace + auto-generated UUID | {1,1}
Identifier | $.SeRARecord.[\*].piEmployeeId, $.SeRARecord.[\*].piSunetId | none | {1,n}
Resource Type | none | always foaf:Agent, foaf:Person | {1,n}
Full Name | $.SeRARecord.[\*].piFullName | none | {1,1}
PI on Grant | $.SeRARecord.[\*].spoNumber | Lookup Grant URI based on SPO Number. | {0,n}
PI Role on Grant | $.SeRARecord.[\*].spoNumber | Lookup Grant URI based on SPO Number. | {0,n}


#### Administering Organization(s) Mapping

MAP Field | CAP Field(s) | Normalization Notes | Obligation
--------- | ------------ | ------------------- | ----------
URI       | n/a          | RIALTO organizations namespace + auto-generated UUID | {1,1}
Identifier | none? | none | {0,n}
Resource Type | $.SeRARecord.[\*].spoProjectDeptName or $.SeRARecord.[\*].spoProjectSchool | always foaf:Agent, foaf:Organization, Department or School based on field mapped off of | {1,n}
Full Name | $.SeRARecord.[\*].piFullName | none | {1,1}
Administering Grant | $.SeRARecord.[\*].spoProjectDeptName or $.SeRARecord.[\*].spoProjectSchool | Lookup Grant URI based on SPO Number. | {0,n}
Administration Role on Grant | $.SeRARecord.[\*].spoProjectDeptName or $.SeRARecord.[\*].spoProjectSchool | Lookup Grant URI based on SPO Number. | {0,n}
