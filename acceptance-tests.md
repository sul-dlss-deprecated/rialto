---
layout: default
---

Note: this is incomplete / WIP.

# Logistics of Acceptance Tests

When there are code changes to be deployed to Stage or Prod, currently that has to be manually triggered by a member of Ops.

Upon any new Terraform build (i.e. any code changes) in Dev or Stage, the following is automatically run to test the end to end completeness of RIALTO Core.

The test suite runs within a lambda written in Go
  * runs the following & outputs success / failure to logs
  * team email (or Honeybadger possibly) notifications via SNS reading said logs

If successful, new build is good; if unsuccessful, debugging and rebuild or reversion needs to occur. This should run at minimum on builds in Stage before pushing said builds to Production.

All data generated in this process should be deleted also by end of process.

Uncertain:
 * what triggers this at end of terraform build


# Creating new resources

## Creating an Organization from CAP

1. Put into SPARQL Proxy:

```
INSERT DATA {
  <http://example.org/orgs/1234> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://xmlns.com/foaf/0.1/Organization> ;
                                 <http://www.w3.org/2004/02/skos/core#prefLabel> "Sample Organization 1 Label" ;
                                 <http://www.w3.org/2004/02/skos/core#altLabel> "Sample Organization 1 Alternate Label" ;
                                 <http://vivoweb.org/ontology/core#abbreviation> "Sample Org. 1 Abbreviation" .
}
```

2. Expect a 200 response from Proxy

3. Upon 200, the following is run:
  * Query Neptune (directly)
    ```
    SELECT ?type ?prefLabel ?altLabel ?abbreviation WHERE {
      <http://example.org/orgs/1234> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> ?type ;
                                     <http://www.w3.org/2004/02/skos/core#prefLabel> ?prefLabel ;
                                     <http://www.w3.org/2004/02/skos/core#altLabel> ?altLabel ;
                                     <http://vivoweb.org/ontology/core#abbreviation> ?abbreviation .
    }
    ```
  * Match Neptune response fields with original query data
  ```
  ?type == <http://xmlns.com/foaf/0.1/Organization> ; type URI
  ?prefLabel == "Sample Organization 1 Label" ; type String
  ?altLabel == "Sample Organization 1 Alternate Label" ; type String
  ?abbreviation == "Sample Org. 1 Abbreviation" ; type String
  ```

  ## *Query SNS*

  ## Query Solr
  ```
  curl https://mysolrendpoint.org/search?id=http://example.org/orgs/1234.json
  ```
  * Parse & match Solr response fields to original query data as expected to be mapped


  ## Query MySQL
  ```
  MySQL client call with the following SQL query:

  ```
  * Parse & match SQL response fields to original query data as expected to be mapped

## Creating a Person (primary) from Profiles


## Creating a Person (secondary, e.g. advisee) from Profiles


## Creating a Grant from SeRA


## Creating an Organization or Person from SeRA


## Creating a Work (Publication) from WoS


## Creating a Grant (secondary) from WoS

# Updating existing resources

## Updating an Organization from CAP

1. Put into SPARQL Proxy:
```
INSERT DATA {
  <http://example.org/orgs/1234> <http://www.w3.org/2004/02/skos/core#prefLabel> "Sample Organization 1 New Label" ;
                                 <http://www.w3.org/2004/02/skos/core#altLabel> "Sample Organization 1 New Alternate Label" ;
                                 <http://vivoweb.org/ontology/core#abbreviation> "Sample Org. 1 New Abbreviation" .
}
DELETE {
  <http://example.org/orgs/1234> <http://www.w3.org/2004/02/skos/core#prefLabel> ?p .
}
WHERE {
  <http://example.org/orgs/1234> <http://www.w3.org/2004/02/skos/core#prefLabel> ?p .
}
```

2. Expect a 200 response from Proxy

3. Upon 200, the following is run:
  * Query Neptune (directly)
    ```
    SELECT ?type ?prefLabel ?altLabel ?abbreviation WHERE {
      <http://example.org/orgs/1234> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> ?type ;
                                     <http://www.w3.org/2004/02/skos/core#prefLabel> ?prefLabel ;
                                     <http://www.w3.org/2004/02/skos/core#altLabel> ?altLabel ;
                                     <http://vivoweb.org/ontology/core#abbreviation> ?abbreviation .
    }
    ```
  * Match Neptune response fields with original query data
  ```
  ?type == <http://xmlns.com/foaf/0.1/Organization> ; type URI
  ?prefLabel == "Sample Organization 1 Label" ; type String
  ?altLabel == "Sample Organization 1 Alternate Label" ; type String
  ?abbreviation == "Sample Org. 1 Abbreviation" ; type String
  ```

  ## *Query SNS*

  ## Query Solr
  ```
  curl https://mysolrendpoint.org/search?id=http://example.org/orgs/1234.json
  ```
  * Parse & match Solr response fields to original query data as expected to be mapped


  ## Query MySQL
  ```
  MySQL client call with the following SQL query:

  ```
  * Parse & match SQL response fields to original query data as expected to be mapped

## Updating a Person (primary) from Profiles


## Updating a Person (secondary, e.g. advisee) from Profiles


## Updating a Grant from SeRA


## Updating an Organization or Person from SeRA


## Updating a Work (Publication) from WoS


## Updating a Grant (secondary) from WoS


# Deleting existing resources

## Deleting test entities

1. Put into SPARQL Proxy:
```
DELETE {
  <http://example.org/orgs/1234> ?p ?o .
}
WHERE {
  <http://example.org/orgs/1234> ?p ?o .
}
```

2. Expect a 200 response from Proxy

3. Upon 200, the following is run:
  * Query Neptune (directly)
    ```
    SELECT ?type ?prefLabel ?altLabel ?abbreviation WHERE {
      <http://example.org/orgs/1234> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> ?type ;
                                     <http://www.w3.org/2004/02/skos/core#prefLabel> ?prefLabel ;
                                     <http://www.w3.org/2004/02/skos/core#altLabel> ?altLabel ;
                                     <http://vivoweb.org/ontology/core#abbreviation> ?abbreviation .
    }
    ```
  * Match Neptune response fields with original query data
  ```
  ?type == <http://xmlns.com/foaf/0.1/Organization> ; type URI
  ?prefLabel == "Sample Organization 1 Label" ; type String
  ?altLabel == "Sample Organization 1 Alternate Label" ; type String
  ?abbreviation == "Sample Org. 1 Abbreviation" ; type String
  ```

  ## *Query SNS*

  ## Query Solr
  ```
  curl https://mysolrendpoint.org/search?id=http://example.org/orgs/1234.json
  ```
  * Parse & match Solr response fields to original query data as expected to be mapped


  ## Query MySQL
  ```
  MySQL client call with the following SQL query:

  ```
  * Parse & match SQL response fields to original query data as expected to be mapped
