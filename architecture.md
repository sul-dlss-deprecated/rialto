---
layout: default
---

# Architecture

Here are diagrams and links to specification, models, and related documentation to the RIALTO Architecture.

## Originally Proposed Conceptual Architecture

This is very early, high level systems architecture Diagram for RIALTO. It is shared here only for historical and motivational purposes (e.g. where we started in our thinking and why).

![Very early, high level systems architecture Diagram for RIALTO.](https://docs.google.com/drawings/d/e/2PACX-1vSQ02m-7tdxzE7UYSbWPsl8DmeWboT952DhosgTLjNCAIUb1f95q71XpijdMQiD60MaWCGBsURLkSmP/pub?w=960&h=720)

<small>[Link to Image](https://docs.google.com/drawings/d/1A_X8krKQcbuonwtzuwMuk1rkEtzLrx-kcSQiuKEoPns/edit?usp=sharing)</small>

## Current Infrastructure Architecture

These are diagrams on our current infrastructure implementations. They are primarily a systems level, and they are broken into multiple sub-systems that include the following:

**RIALTO Combine / RIALTO ETL**: Our data ingestion pipeline. Written in Ruby (primarily), running on premise. This calls various APIs for data to harvest, then maps, reconciles, deduplicates, and normalizes for loading into RIALTO Core.

**RIALTO Core**: Our primary graph store holding RIALTO data following the [RIALTO core data application profiles](/rialto/data-models). This is primarily a proxy API for RIALTO Combine to call upon ingest; a SPARQL query (or set of them) for loading the mapped data into a graph store; and a notification queue of changes for then materializing read-only views into Solr and Postgres databases used by any RIALTO downstream clients.

**RIALTO Web Application**: A Rails and Vue.js application that reads off RIALTO Core's materialized databases and presents data views to end users.

### RIALTO (full system view)

Full RIALTO hybrid infrastructure system architecture diagram:

![Full RIALTO System Architecture on AWS](https://docs.google.com/drawings/d/e/2PACX-1vSArQzCspzUsngeAUs40GCXvVo9j_-o_2eq6S5V7lsNTiV37sPlIrDEXP0wnnlsICFQSs6QIH3DMfpG/pub?w=603&h=736)

<small>[Link to Image](https://docs.google.com/drawings/d/1Xl2cQI-nOfgNi1weFAD8F6kFPhJq57xudFmPbtDZrA4/edit)</small>

### RIALTO Core

Functional system diagram architecture diagram for RIALTO Core alone as on AWS:

![Architectural Diagram for RIALTO Core](https://docs.google.com/drawings/d/e/2PACX-1vTEpwYIDv9e18wGASJxtRktBA68SiafwYKz4RcUd3N5lTmFQEliRisqO4Xx_7YDyE5tZ8DUrdAZrbgc/pub?w=551&h=704)

<small>[Link to Image](https://docs.google.com/drawings/d/1G_w53xeqko0ZEBLvRd7UJ5YTErqLcsuqLbPRDC1NCmI/edit)</small>

### Another view on RIALTO Core AWS Architecture

Earlier AWS implementation diagram for RIALTO Core alone:

![AWS Architecture for RIALTO Core](https://docs.google.com/drawings/d/e/2PACX-1vTHINF7ZJJxPZe_IBzmU63Ev50vnQiLR5D9uWp2GqVO4RfstGu5GiPEwgJenlwTuDrowuG92VVKYV3v/pub?w=863&h=716)

<small>[Link to Image](https://docs.google.com/drawings/d/1bLg28G1qyYBrU6ruLe50cUrZKQuo4CtwrUkmOsbErnc/view)</small>
